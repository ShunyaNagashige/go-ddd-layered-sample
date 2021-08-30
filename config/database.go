package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var (
// 	db  *gorm.DB
// 	err error
// )

func createDsn() string {
	return os.Getenv("DB_USER") +
		":" + os.Getenv("DB_PASS") +
		"@" + os.Getenv("DB_PROTOCOL") +
		"(" + os.Getenv("DB_ADDRESS") +
		":" + os.Getenv("DB_PORT") + ")" +
		"/" + os.Getenv("DB_NAME")
}

// DB接続
func Connect() (*gorm.DB, error) {
	// 実行環境取得
	env := os.Getenv("ENV")
	fmt.Println(env)
	if env != "dev" && env != "stg" && env != "prod" {
		return nil, errors.New("環境変数ENVの値が正しく設定されていません．")
	}

	// 環境変数取得
	if err := godotenv.Load("./config/env/" + env + ".env"); err != nil {
		return nil, err
	}

	// DB接続
	dsn := createDsn()
	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
