package controllers

import (
	"{{PROJECT_NAME}}/services"

	{{FRAMEWORK_API_LINK}}
)

type Controller struct {
	g *{{PACKAGE_FRAMEWORK_API}}.{{PACKAGE_FRAMEWORK_ENGINE}}
	s services.IService
}

func NewControllers(g *{{PACKAGE_FRAMEWORK_API}}.{{PACKAGE_FRAMEWORK_ENGINE}}, s services.IService)  *Controller {
	return &Controller{
		g: g,
		s: s,
	}
}
