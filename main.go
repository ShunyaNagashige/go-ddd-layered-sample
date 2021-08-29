package main

import "github.com/ShunyaNagashige/go-ddd-layered-sample/infrastructure/persistence"

func main() {
	// utils.LoggingSettings(config.Config.LogFile)

	// 依存関係を定義
	userPersistence := persistence.NewUserPersistence()
}
