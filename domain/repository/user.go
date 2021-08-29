// DBとのやりとりを定義する．
// 技術的関心ごとはinfrastructure層に書くため，
// ここではインタフェースとしてメソッドを定義する．
package repository

import (
	"github.com/ShunyaNagashige/go-ddd-layered-sample/domain/model"
)

type UserRepository interface {
	// UserRepositoryインタフェースでは，Searchを定義している
	Search(name string) ([]*model.User, error)
}

// リポジトリは，ドメインオブジェクトの
// 永続化に関わる操作を隠蔽している