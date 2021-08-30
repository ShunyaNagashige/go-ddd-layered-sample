package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ShunyaNagashige/go-ddd-layered-sample/config"
	"github.com/ShunyaNagashige/go-ddd-layered-sample/infrastructure/persistence"
	handler "github.com/ShunyaNagashige/go-ddd-layered-sample/interface/handler/rest"
	"github.com/ShunyaNagashige/go-ddd-layered-sample/usecase"
	"github.com/ShunyaNagashige/go-ddd-layered-sample/utils"
	"github.com/julienschmidt/httprouter"
)

func main() {
	utils.LoggingSettings(os.Getenv("LOG_FILE"))

	dbConn, err := config.Connect()
	if err != nil {
		log.Fatalf("DBに接続できません．: err=%v", err)
	}

	// 依存関係を定義
	userPersistence := persistence.NewUserPersistence(dbConn)
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	// ルーティングの設定
	router := httprouter.New()
	router.GET("/api/users", userHandler.Index)

	// サーバ起動
	http.ListenAndServe(":8080", router)
}

/*
type Handler interface{
	ServeHTTP(ResponseWriter,*Request)
}
である．
つまり，上記のServeHTTPさえ持っていれば
Handlerを実装したことになる．
*/
