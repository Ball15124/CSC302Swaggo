package di

import (
	"backend/internal/routes"
	"github.com/google/wire"
)

var routesSet = wire.NewSet(
	routes.ProvideAuthRouter,
	routes.ProvideRouter,
)
