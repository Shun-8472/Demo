package inject

import (
	"demo/internal/applied/cache"
	"github.com/google/wire"

	"demo/external/receiver/demo"
	"demo/internal/applied/database"
	"demo/internal/applied/engine"
)

var InjectsSet = wire.NewSet(wire.Struct(new(Injects), "*"))

type Injects struct {
	Receiver   demo.Receiver
	GrpcServer engine.Engine
	Database   database.Database
	Redis      cache.Cache
}
