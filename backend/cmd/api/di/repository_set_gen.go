package di

import (
	"backend/internal/repository"
	"github.com/google/wire"
)

var repositorySet = wire.NewSet(
	repository.ProvideUserRepository,
)
