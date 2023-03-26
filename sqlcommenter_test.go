package sqlcommenterhertz

import (
	"context"
	"net/http"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/google/sqlcommenter/go/core"
)

func TestSQLCommenterMiddleware(t *testing.T) {
	r := route.NewEngine(config.NewOptions([]config.Option{}))
	r.Use(SQLCommenterMiddleware())
	r.GET("/test/:id", func(c context.Context, _ *app.RequestContext) {
		expectedFramework := framework
		expectedRoute := "GET--/test/:id"
		expectedAction := "github.com/a631807682/sqlcommenter-hertz.TestSQLCommenterMiddleware.func1"

		framework := c.Value(core.Framework)
		route := c.Value(core.Route)
		action := c.Value(core.Action)

		assert.DeepEqual(t, expectedFramework, framework)
		assert.DeepEqual(t, expectedRoute, route)
		assert.DeepEqual(t, action, expectedAction)
	})

	r.POSTEX("/test2/:id", func(c context.Context, ctx *app.RequestContext) {
		expectedFramework := framework
		expectedRoute := "POST--/test2/:id"
		expectedAction := "HandleTest2"

		framework := c.Value(core.Framework)
		route := c.Value(core.Route)
		action := c.Value(core.Action)

		assert.DeepEqual(t, expectedFramework, framework)
		assert.DeepEqual(t, expectedRoute, route)
		assert.DeepEqual(t, action, expectedAction)
	}, "HandleTest2")

	ut.PerformRequest(r, http.MethodGet, "/test/1", nil)
	ut.PerformRequest(r, http.MethodPost, "/test2/1", nil)
}
