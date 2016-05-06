package controllers

import (
	_ "github.com/hieubq-vng/revel-swagger/modules/swaggerapi"
	"github.com/revel/revel"
)

type SwaggerAPI struct {
	*revel.Controller
}

// ServeUI renders the template for your swagger-ui
func (c SwaggerAPI) ServeUI(spec string) revel.Result {
	c.RenderArgs["SwaggerSpecURL"] = c.Request.URL.Path + spec
	return c.Render()
}
