package main

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"

	"apx103.com/super-mid/controllors"
	"apx103.com/super-mid/controllors/common"
	// "apx103.com/super-mid/events"
	// "apx103.com/super-mid/middleware"
	// "apx103.com/super-mid/services"
	// "apx103.com/super-mid/utils"
)

// Set logrus format

type Formatter struct{}

func (m *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string

	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		newLog = fmt.Sprintf("[logrus] [%s] [%s] [%s:%d %s] %s\n",
			timestamp, entry.Level, fName, entry.Caller.Line, entry.Caller.Function, entry.Message)
	} else {
		newLog = fmt.Sprintf("[logrus] [%s] [%s] %s\n", timestamp, entry.Level, entry.Message)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}

// HTTP server

type RouterFx struct {
	router *gin.Engine
}

func NewRouterFx(routers []common.Controller) RouterFx {
	r := RouterFx{}
	router := gin.Default()
	router.Use(cors.Default())
	// router.Use(lm.RateLimitMiddleware())
	// router.Use(am.AuthenticationMiddleware())
	// router.Static("/puyu_larkbot_web_config", "./web/dist")
	// router.Static("/public", "./public")
	router.Use(gin.Recovery())
	for _, r := range routers {
		switch r.GetMethod() {
		case "POST":
			router.POST(r.GetPath(), r.GetController())
		case "GET":
			router.GET(r.GetPath(), r.GetController())
		default:
			logrus.Infof("%s method is invalided", r.GetMethod())
		}
	}
	// router.NoRoute(func(c *gin.Context) {
	// 	prefix := "./web/dist"
	// 	c.File(prefix + "/index.html")
	// })
	r.router = router
	return r
}

func Server(lc fx.Lifecycle, routerFx RouterFx) *gin.Engine {
	srv := &http.Server{Addr: ":8080", Handler: routerFx.router}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				logrus.Errorf("[LarkBot] Failed to start HTTP Server at %s", srv.Addr)
				return err
			}
			go func() {
				err := srv.Serve(ln)
				if err != nil {
					logrus.Infof("Server stoped: %s", err)
				}
			}()
			logrus.Infof("[LarkBot]Succeeded to start HTTP Server at %s", srv.Addr)
			return nil

		},
		OnStop: func(ctx context.Context) error {
			err := srv.Shutdown(ctx)
			if err != nil {
				logrus.Errorf("Got error when shutting the Server: %s", err)
			}
			logrus.Info("[LarkBot] HTTP Server is stopped")
			return err
		},
	})

	return routerFx.router
}

// Init log setting

func init() {
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		lvl = "debug"
	}
	ll, err := logrus.ParseLevel(lvl)
	if err != nil {
		ll = logrus.DebugLevel
	}
	logrus.SetLevel(ll)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&Formatter{})
}

// APP Entry

func main() {
	app := fx.New(
		// utils.Module,
		// events.Module,
		// services.Module,
		controllors.Module,
		// middleware.Module,
		fx.Provide(
			fx.Annotate(
				NewRouterFx,
				fx.ParamTags(`group:"controller"`),
			),
			Server,
		),
		fx.Invoke(func(*gin.Engine) {}),
	)
	app.Run()
}
