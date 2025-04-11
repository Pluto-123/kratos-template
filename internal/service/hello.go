package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kratos-template/api"
	"kratos-template/api/gcode"
	"kratos-template/api/gerror"
	hello_api "kratos-template/api/hello"
	"kratos-template/internal/biz"
	"kratos-template/internal/conf"
)

type HelloAPIService struct {
	c   *conf.Data
	hUc *biz.HelloUsecase
}

func NewHelloAPIService(c *conf.Data, hUc *biz.HelloUsecase) *HelloAPIService {
	return &HelloAPIService{
		c:   c,
		hUc: hUc,
	}
}

func (s *HelloAPIService) Hello(ctx *gin.Context) (api.HttpResponse, error) {
	resp := &hello_api.HelloReply{
		Code:    int32(gcode.CodeOK.Code()),
		Message: gcode.CodeOK.Message(),
	}
	req := &hello_api.HelloRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		resp.Code = int32(gcode.CodeInvalidRequest.Code())
		resp.Message = fmt.Sprintf("%v: %v", gcode.CodeInvalidRequest.Message(), err)
		return resp, gerror.NewCode(gcode.CodeInvalidRequest)
	}
	_ = s.hUc.Hello(ctx, req)
	return resp, nil
}
