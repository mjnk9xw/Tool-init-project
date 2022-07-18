package routers

import (
	"{{PROJECT_NAME}}/api/controllers"
	"{{PROJECT_NAME}}/services"

	{{FRAMEWORK_API_LINK}}
)

type Router struct {
}

func New(s services.IService) {
	"{{FRAMEWORK_API_NEW}}"

	controllers.NewControllers(g, s)

	"{{FRAMEWORK_API_RUN}}"

}
