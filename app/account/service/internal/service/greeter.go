package service

import (
	v1 "github.com/fringelin/fgo/api/account/v1"
	"github.com/fringelin/fgo/app/account/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

// AccountService is a account service.
type AccountService struct {
	v1.UnimplementedAccountServer

	uc  *biz.AccountUsecase
	log *log.Helper
}

// NewAccountService new a greeter service.
func NewAccountService(uc *biz.AccountUsecase, logger log.Logger) *AccountService {
	return &AccountService{uc: uc, log: log.NewHelper("service/account", logger)}
}
