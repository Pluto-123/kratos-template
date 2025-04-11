package route

import (
	"github.com/gin-gonic/gin"
	"kratos-template/internal/conf"
	"kratos-template/internal/service"
)

func RegisterHelloService(s *service.HelloAPIService, cd *conf.Data) GroupUrl {
	return GroupUrl{
		GroupAddr: "/api",
		Urls: []Url{{
			JsonHandlerFunc: s.Hello,
			Path:            "hello",
			Method:          GET,
			Middleware:      []gin.HandlerFunc{},
		}},
		Middleware: nil,
	}
}
