package controllers

import (
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/mashup/mashup"
)

type RootController struct {
	generator mashup.Generator
}

func NewRootController(g mashup.Generator) *RootController {
	return &RootController{
		generator: g,
	}
}

func (r *RootController) Routes() []*fireball.Route {
	routes := []*fireball.Route{
		{
			Path: "/",
			Handlers: map[string]fireball.Handler{
				"GET": r.getVerse,
			},
		},
	}

	return routes
}

func (r *RootController) getVerse(c *fireball.Context) (fireball.Response, error) {
	return c.HTML(200, "index.html", r.generator())
}
