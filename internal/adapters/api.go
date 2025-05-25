package adapters

import (
	"go-training-crud/data_source"
	"go-training-crud/internal/adapters/controller"
	"go-training-crud/internal/adapters/repository"
	"go-training-crud/internal/application/cmd"
	"go-training-crud/internal/application/qry"
	"log"
)

func BuildApp() error {
	// Data Source
	userDataSource := data_source.NewUserDataSource()

	// Repository
	userRepository := repository.NewUserRepository(userDataSource)

	// Service
	userCommandService := cmd.NewUserCommandService(userRepository)
	userQueryService := qry.NewUserQueryService(userRepository)

	// Controller
	err := controller.MapRoutes(userCommandService, userQueryService)
	if err != nil {
		log.Panic("startup failed")
	}

	return nil
}
