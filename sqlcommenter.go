package sqlcommenterhertz

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/sqlcommenter/go/core"
	"github.com/google/sqlcommenter/go/net/http"
)

const framework = "cloudwego/hertz"

func SQLCommenterMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		path := c.FullPath()
		method := string(c.Method())
		route := fmt.Sprintf("%s--%s", method, path)
		handlerName := app.GetHandlerName(c.Handler())
		if handlerName == "" { // handler name not set, use func name.
			handlerName = c.HandlerName()
		}

		ctx = core.ContextInject(ctx, http.NewHTTPRequestTags(framework, route, handlerName))
		c.Next(ctx)
	}
}
