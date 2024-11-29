package routes

import "github.com/gin-gonic/gin"

type IRouter interface {
	InitializedRouter(app *gin.Engine)
}

type Router struct {
	authRouter AuthRouter
}

func ProvideRouter(router AuthRouter) IRouter {
	return &Router{
		authRouter: router,
	}
}

func (r *Router) InitializedRouter(app *gin.Engine) {
	r.authRouter.InitializedGroup(app)
}
