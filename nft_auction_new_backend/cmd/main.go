package main

import (
	"backend/internal/api"
	"backend/internal/api/middleware"
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/indexer"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.Load()
	log.Printf("config loaded: env=%s port=%s chain_id=%d rpc=%s contract=%s",
		cfg.Env, cfg.Port, cfg.ChainID, cfg.RpcURL, cfg.AuctionContract)

	mysqlDB, err := db.NewMySQL(cfg)

	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	log.Println("mysql db connected.")

	// 初始化索引器：传入配置和 MySQL 客户端（索引器需要写数据到数据库）
	indx, err := indexer.NewIndexer(cfg, mysqlDB)
	if err != nil {
		log.Fatal("failed to create indexer: ", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	indx.Start(ctx)

	r := gin.Default()
	r.Use(middleware.InjectChainContext(cfg))

	// 注册健康检查接口：/health，用于监控服务是否正常运行（比如 k8s 探针、运维监控）
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "env": cfg.Env})
	})

	// 注册业务 API 路由：api.RegisterRoutes 是自定义函数，注册拍卖相关的接口（如查询拍卖列表、出价记录等）
	// 传入 Gin 实例和 MySQL 客户端（API 需要从数据库查数据）
	api.RegisterRoutes(r, mysqlDB)

	// 创建 HTTP 服务器实例：基于 Gin 路由创建标准的 http.Server
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	// 开启协程启动 HTTP 服务：ListenAndServe 是阻塞调用，所以用 goroutine 避免阻塞主协程
	go func() {
		log.Println("server listening on port: ", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 创建信号通道：用于接收系统信号（缓冲区大小 1，避免信号丢失）
	quit := make(chan os.Signal, 1)
	// 注册信号监听：监听 SIGINT（Ctrl+C）和 SIGTERM（kill 命令）信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 阻塞等待信号：主协程停在这里，直到收到终止信号
	<-quit

	log.Println("shutdown signal received")

	cancel()

	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelTimeout()

	if err := srv.Shutdown(ctxTimeout); err != nil {
		log.Fatal("server shutdown failed: ", err)
	}

	log.Println("server exited gracefully")
}
