package route

import (
	"github.com/google/wire"
	"kratos-template/internal/conf"
	"kratos-template/internal/service"
)

var ProviderSet = wire.NewSet(RegisterHttpService)

func RegisterHttpService(
	cd *conf.Data,
	helloService *service.HelloAPIService,
) []GroupUrl {

	var routes []GroupUrl
	//routes = append(routes, RegisterWebsocketService(wsService, cd))
	routes = append(routes, RegisterHelloService(helloService, cd))
	return routes
}
