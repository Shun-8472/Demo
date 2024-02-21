// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package inject

import (
	"demo/external/receiver/demo"
	"demo/internal/inject/applied"
	"demo/internal/processor/version/implement"
)

// Injectors from wire.go:

func BuildInjector() (*Injects, error) {
	procedure := implement.NewDemoProcedure()
	receiver := demo.ProvideReceiver(procedure)
	engine := applied.InitGrpcServer()
	database := applied.InitDatabaseConnect()
	injects := &Injects{
		Receiver:   receiver,
		GrpcServer: engine,
		Database:   database,
	}
	return injects, nil
}