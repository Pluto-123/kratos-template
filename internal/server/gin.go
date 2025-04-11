package server

import (
	_ "embed"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"kratos-template/api/gcode"
	"kratos-template/internal/conf"
	"kratos-template/internal/route"
)

// NewGinServer new a Gin server.
func NewGinServer(c *conf.Server, d *conf.Data, logger log.Logger,
	urls []route.GroupUrl) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
		),
	}
	if c.Gin.Network != "" {
		opts = append(opts, http.Network(c.Gin.Network))
	}
	if c.Gin.Addr != "" {
		opts = append(opts, http.Address(c.Gin.Addr))
	}
	if c.Gin.Timeout != nil {
		opts = append(opts, http.Timeout(c.Gin.Timeout.AsDuration()))
	}

	gEngine := gin.Default()
	gin.SetMode(gin.DebugMode)

	DecoratorRouter(gEngine, d, urls)

	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", gEngine)
	return srv
}

func DecoratorRouter(gEngine *gin.Engine, cd *conf.Data, urls []route.GroupUrl) {
	//解决跨域问题
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AddAllowHeaders("*")
	gEngine.Use(cors.New(corsConfig))

	//Swagger
	gEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//设置健康检查
	gEngine.GET("/healthy", func(c *gin.Context) {
		c.JSON(200,
			struct {
				IsAlive bool `json:"is_alive"`
			}{IsAlive: true})
	})

	for _, gurl := range urls {
		route.MKHandler(gEngine, gurl)
	}

	gEngine.NoRoute(func(c *gin.Context) {
		c.JSON(int(gcode.CodeNotFound.HttpCode()), nil)
	})
}
