package di

import "github.com/google/wire"

var appSet = wire.NewSet(
	ProvideGinEngine,
)
