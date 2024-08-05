package router

import (
	"net/http"

	"github.com/RayhanAsauqi/blog-app/internal/http/handler"
	"github.com/labstack/echo/v4"
)

type Route struct {
	Method      string
	Path        string
	Handler     echo.HandlerFunc
	Middlewares []echo.MiddlewareFunc
}

func PublicRoutes(authHandler *handler.AuthHandler, articleHandler *handler.ArticleHandler) []*Route {
	return []*Route{
		{
			Method:  http.MethodPost,
			Path:    "/register",
			Handler: authHandler.Register,
		},
		{
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: authHandler.Login,
		},
		{
			Method:  http.MethodGet,
			Path:    "/articles",
			Handler: articleHandler.GetArticles,
		},
	}
}

func PrivateRoutes(articleHandler *handler.ArticleHandler) []*Route {
	return []*Route{
		{
			Method:  http.MethodPost,
			Path:    "/articles",
			Handler: articleHandler.CreateArticle,
		},
	}
}

func RegisterRoutes(e *echo.Echo, routes []*Route) {
	for _, route := range routes {
		e.Add(route.Method, route.Path, route.Handler, route.Middlewares...)
	}
}