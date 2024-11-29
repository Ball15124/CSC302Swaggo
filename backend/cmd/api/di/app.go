package di

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"backend/cmd/api/docs"
	"backend/internal/routes"
)

type IPillar interface {
	RunApp() error
}

type Pillar struct {
	app   *gin.Engine
	route routes.IRouter
}

func (p *Pillar) RunApp() error {
	p.app.Use(gin.Recovery())
	p.app.Use(gin.Logger())
	p.route.InitializedRouter(p.app)

	docs.SwaggerInfo.BasePath = "/"
	p.app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return p.app.Run(":8000")
}

func ProvideGinEngine(route routes.IRouter) IPillar {
	app := gin.New()

	return &Pillar{
		app:   app,
		route: route,
	}
}
