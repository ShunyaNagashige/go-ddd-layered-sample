package usecase

import (
	"github.com/ShunyaNagashige/go-ddd-layered-sample/domain/model"
	"github.com/ShunyaNagashige/go-ddd-layered-sample/domain/repository"
)

// UserにおけるUseCaseのインタフェース
type UserUseCase interface {
	Search(name string) ([]*model.User, error)
}

// repository.UserRepositoryを実装
// DBにアクセスするときもdomain層のrepositoryを介してアクセスすることにより
// usecase層を，domain層のみに依存させている
type userUseCase struct {
	userRepository repository.UserRepository
}

// Userデータに関するUseCaseを生成
func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

// 検索
func (uu userUseCase) Search(name string) (user []*model.User, err error) {
	user, err = uu.userRepository.Search(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}
