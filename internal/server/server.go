package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"kratos-template/internal/conf"
	"kratos-template/internal/route"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewAllHttpServer, NewCronWorker)

func NewAllHttpServer(c *conf.Server, d *conf.Data, logger log.Logger,
	urls []route.GroupUrl) []*http.Server {
	return []*http.Server{
		NewHTTPServer(c, logger),
		NewGinServer(c, d, logger, urls),
	}
}
