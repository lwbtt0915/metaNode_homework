package db

import (
	"errors"
	"gorm.io/gorm"
)

var ErrCursorNotFound = errors.New("cursor not found")

// 查询游标
func (m *MySQL) GetCursor(chainID uint64,
	contractAddr string,
	eventName string) (uint64, error) {

	var c SyncCursor
	err := m.DB.Where(
		"chain_id = ? AND contract_address = ? AND event_name = ?",
		chainID,
		contractAddr,
		eventName,
	).First(&c).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, ErrCursorNotFound
	}

	return c.LastBlock, err
}

// 更新游标
func (m *MySQL) UpdateCursor(chainID uint64,
	contractAddr string,
	eventName string,
	lastBlock uint64) error {
	var c SyncCursor
	err := m.DB.Where("chain_id = ? AND contract_address = ? AND event_name = ?",
		chainID, contractAddr, eventName).First(&c).Error

	if err != nil {
		return err
	}

	c.LastBlock = lastBlock
	return m.DB.Save(&c).Error
}
