package config

import (
	"os"

	"gorm.io/gorm"
)

var(
	db *gorm.DB
	err error
)

// DB接続
func Connect()*gorm.DB{
	// 環境変数取得
	env:=os.Getenv("ENV")

	if env!="development"&&env!="production"{
		return nil
	}

	// 環境変数取得
	
}