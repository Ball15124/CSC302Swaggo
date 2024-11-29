package di

import (
	"backend/internal/controller"
	"github.com/google/wire"
)

var controllerSet = wire.NewSet(
	controller.ProvideAuthController,
)
