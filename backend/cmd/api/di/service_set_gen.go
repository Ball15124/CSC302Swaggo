package di

import (
	"backend/internal/service"
	"github.com/google/wire"
)

var serviceSet = wire.NewSet(
	service.ProvideAuthService,
	service.ProvideHashService,
	service.ProvideTokenService,
)
