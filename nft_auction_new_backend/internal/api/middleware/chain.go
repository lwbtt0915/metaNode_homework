package middleware

import (
	"backend/internal/config"
	"github.com/gin-gonic/gin"
)

func InjectChainContext(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("chain_id", cfg.ChainID)
		c.Set("rpc_url", cfg.RpcURL)
		c.Set("contract", cfg.AuctionContract)
		c.Next()
	}
}
