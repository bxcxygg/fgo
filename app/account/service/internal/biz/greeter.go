package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type AccountRepo interface {
}

type AccountUsecase struct {
	repo AccountRepo
	log  *log.Helper
}

func NewGreeterUsecase(repo AccountRepo, logger log.Logger) *AccountUsecase {
	return &AccountUsecase{repo: repo, log: log.NewHelper("usecase/account", logger)}
}
