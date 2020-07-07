package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DefaultDB *gorm.DB
)

func InitMySQL(dbConnStr string) error {
	db, err := gorm.Open("mysql", dbConnStr)
	if err != nil {
		return fmt.Errorf("init db failed: %w", err)
	}

	DefaultDB = db
	return nil
}

func CloseMySQL() error {
	if err := DefaultDB.Close(); err != nil {
		return fmt.Errorf("close db failed: %w", err)
	}
	return nil
}
