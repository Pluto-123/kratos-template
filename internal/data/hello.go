package data

import "kratos-template/internal/biz"

type helloRepo struct {
	data *Data
}

func NewHelloRepo(data *Data) biz.HelloRepo {
	return &helloRepo{data: data}
}
