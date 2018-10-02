package controllers

import (
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/mashup/engine"
)

type RootController struct {
	generator engine.Generator
}

func NewRootController(g engine.Generator) *RootController {
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
