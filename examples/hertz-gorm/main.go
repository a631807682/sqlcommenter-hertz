package main

import (
	"context"
	"time"

	sqlcommentergorm "github.com/a631807682/sqlcommenter-gorm"
	sqlcommenterhertz "github.com/a631807682/sqlcommenter-hertz"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	dbInit()

	h := server.Default()
	h.Use(sqlcommenterhertz.SQLCommenterMiddleware())
	register(h)
	go h.Spin()

	time.Sleep(time.Microsecond * 100)
	clientRun()
	h.Shutdown(context.Background())
}

func register(h *server.Hertz) {
	h.GET("/foo", func(c context.Context, ctx *app.RequestContext) {
		var t Test
		DB.Scopes(sqlcommentergorm.ContextInject(c)).First(&t)
		// inject the following comments
		/* action='main.register.func1',application='example-hertz-gorm',db_driver='mysql',framework='cloudwego%2Fhertz',route='GET--%2Ffoo' */
		ctx.JSON(consts.StatusOK, t)
	})

	h.POSTEX("/bar", func(c context.Context, ctx *app.RequestContext) {
		t := &Test{Name: "foo"}
		DB.Scopes(sqlcommentergorm.ContextInject(c)).Create(&t)
		// inject the following comments
		/* action='BarHandler',application='example-hertz-gorm',db_driver='mysql',framework='cloudwego%2Fhertz',route='POST--%2Fbar' */
		ctx.JSON(consts.StatusOK, t)
	}, "BarHandler")
}

func clientRun() {
	cli, err := client.NewClient()
	if err != nil {
		panic(err)
	}
	cli.Post(context.Background(), nil, "http://127.0.0.1:8888/bar", nil)
	cli.Get(context.Background(), nil, "http://127.0.0.1:8888/foo")
}
