package rest

import (
	"encoding/json"
	"net/http"

	"github.com/ShunyaNagashige/go-ddd-layered-sample/usecase"
	"github.com/julienschmidt/httprouter"
)

// userにおけるHandlerのインタフェース
type UserHandler interface {
	Index(http.ResponseWriter, *http.Request, httprouter.Params)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

// Userデータに関するHandlerを生成
func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{uu}
}

// UserIndex : GET /users -> 検索結果を返す
func (uh userHandler) Index(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	// GETパラメータ
	name := r.FormValue("name")

	user, err := uh.userUseCase.Search(name)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// クライアントにレスポンスを返却
	if err = json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
