// persistenceはrepositoryと同義
// Domain層とInfra層でパッケージが被ってしまうため，
// このようにしている
package persistence

import (
	"github.com/ShunyaNagashige/go-ddd-layered-sample/domain/model"
	"github.com/ShunyaNagashige/go-ddd-layered-sample/domain/repository"
	"gorm.io/gorm"
)

// UserにおけるPersistenceのインタフェース
// この構造体にSearchメソッドを持たせることによって，
// repository packageのUserRepositoryを実装させる
type userPersistence struct {
	Conn *gorm.DB
}

// Userデータに関するPersistenceを生成
func NewUserPersistence(conn *gorm.DB) repository.UserRepository {
	return &userPersistence{Conn: conn}
}

// 検索
func (up *userPersistence) Search(name string) ([]*model.User, error) {
	var user []*model.User

	// DB接続確認
	if err := up.Conn.Take(&user).Error; err != nil {
		return nil, err
	}

	db := up.Conn.Find(&user)

	// 名前検索
	if name != "" {
		db = db.Where("name = ?", name).Find(&user)
	}

	return user, nil
}
