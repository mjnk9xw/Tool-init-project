package routers

import (
	"{{PROJECT_NAME}}/api/controllers"
	"{{PROJECT_NAME}}/services"

	{{FRAMEWORK_API_LINK}}
)

func New(s services.IService) {
	"{{FRAMEWORK_API_NEW}}"

	con := controllers.NewControllers(g, s)

	loadRouter(g, con)

	"{{FRAMEWORK_API_RUN}}"

}
func loadRouter(g *{{PACKAGE_FRAMEWORK_API}}.{{PACKAGE_FRAMEWORK_ENGINE}}, con *controllers.Controller) {
	api := g.Group("/api")
	{
		{{GROUP_API_CONTROLLER}}
	}
}
