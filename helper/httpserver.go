package helper

import (
	"log"
	"log/slog"
	"os"

	slogecho "github.com/samber/slog-echo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteConfigurable interface {
	Add(method, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route
	Any(path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) []*echo.Route
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	Group(prefix string, middleware ...echo.MiddlewareFunc) (sg *echo.Group)
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	Match(methods []string, path string, handler echo.HandlerFunc, m ...echo.MiddlewareFunc) []*echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	RouteNotFound(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	Use(middleware ...echo.MiddlewareFunc)
}

// PATTERN Factory Options
type SrvOption func(e RouteConfigurable)

// type GroupOption func(g *echo.Group)

type Srv struct {
	*echo.Echo
	Addr string
}

func RunServer(opt ...SrvOption) error {
	e := echo.New()
	e.HideBanner = true
	e.Pre(middleware.RemoveTrailingSlash())

	srv := &Srv{Echo: e, Addr: ":8080"}

	for _, o := range opt {
		o(srv)
	}
	return e.Start(srv.Addr)
}

func WithLog() SrvOption {
	return func(e RouteConfigurable) {
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

		config := slogecho.Config{
			DefaultLevel:       slog.LevelInfo,
			ClientErrorLevel:   slog.LevelWarn,
			ServerErrorLevel:   slog.LevelWarn,
			WithRequestHeader:  false,
			WithRequestBody:    true,
			WithResponseHeader: false,
			WithResponseBody:   false,
			Filters:            []slogecho.Filter{slogecho.IgnorePathPrefix("/ping")},
		}

		e.Use(slogecho.NewWithConfig(logger, config))
	}
}

func WithGroup(grouppath string, opt ...SrvOption) SrvOption {
	return func(e RouteConfigurable) {
		g := e.Group(grouppath)
		for _, o := range opt {
			o(g)
		}
	}
}

func WithMiddleware(m ...echo.MiddlewareFunc) SrvOption {
	return func(e RouteConfigurable) {
		e.Use(m...)
	}
}

func WithAdd(method string, path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) SrvOption {
	return func(e RouteConfigurable) {
		e.Add(method, path, h, m...)
	}
}

// m ... = Variadic Functions
func WithGet(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) SrvOption {
	return func(e RouteConfigurable) {
		e.GET(path, h, m...)
	}
}

func WithPost(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) SrvOption {
	return func(e RouteConfigurable) {
		e.POST(path, h, m...)
	}
}

func WithPut(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) SrvOption {
	return func(e RouteConfigurable) {
		e.PUT(path, h, m...)
	}
}

func WithDelete(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) SrvOption {
	return func(e RouteConfigurable) {
		e.DELETE(path, h, m...)
	}
}

// type casting
func WithAddr(addr string) SrvOption {
	return func(e RouteConfigurable) {
		if srv, isSrv := e.(*Srv); isSrv {
			srv.Addr = addr
		} else {
			log.Fatal("cannot set Addr on non *Serv type")
		}
	}
}

func WithAddrProd() SrvOption {
	return func(e RouteConfigurable) {
		WithAddr(os.Getenv("ADDR"))
	}
}
