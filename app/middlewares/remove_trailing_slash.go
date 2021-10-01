package middlewares

import (
	"net/http"
	"strings"
)

// 除首页以外，移除所有请求路径后面的斜杆
func RemoveTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		next.ServeHTTP(w, r)
	})
}
