package domain

import (
	"github.com/google/wire"
	"github.com/takkee/go-sample-api/domain/service"
)

var WireSet = wire.NewSet(
	service.NewProfileService,
	service.NewUserService,
)
