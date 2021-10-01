package route

import (
	"goblog/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func SetupRoute(r *mux.Router) {
	Router = r
}

func Name2URL(routeName string, pairs ...string) string {
	url, err := Router.Get(routeName).URL(pairs...)
	if err != nil {
		logger.LogError(err)
		return ""
	}
	return url.String()
}

// 获取URI路由参数
func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
