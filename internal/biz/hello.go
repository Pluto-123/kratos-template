package biz

import (
	"context"
	hello_api "kratos-template/api/hello"
	"kratos-template/internal/conf"
)

type HelloRepo interface {
}
type HelloUsecase struct {
	cd        *conf.Data
	helloRepo HelloRepo
}

func NewHelloUsecase(cd *conf.Data, helloRepo HelloRepo) *HelloUsecase {
	return &HelloUsecase{
		cd:        cd,
		helloRepo: helloRepo,
	}
}

func (uc *HelloUsecase) Hello(ctx context.Context, req *hello_api.HelloRequest) string {
	return ""
}
