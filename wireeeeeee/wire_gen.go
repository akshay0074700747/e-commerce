// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wireeeeeee

import (
	"ecommerce/internal/adapters"
	"ecommerce/internal/usecases"
	"ecommerce/web/api/Routes"
	"ecommerce/web/api/handlers"
	"ecommerce/web/config"
	"ecommerce/web/database"
)

// Injectors from wire.go:

func InitializeAPI(config2 config.Config) (*routes.GinEngine, error) {
	db := database.Connect_to(config2)
	userRepo := adapters.NewUserRepository(db)
	userUsecaseInterface := usecases.NewUserUseCase(userRepo)
	userHandler := handlers.NewUserHandler(config2, userUsecaseInterface)
	adminRepo := adapters.NewAdminRepository(db)
	adminUsecaseInterface := usecases.NewAdminUsecase(adminRepo)
	adminHandler := handlers.NewAdminHandler(adminUsecaseInterface, config2)
	suAdminRepo := adapters.NewSuAdminRepository(db)
	suAdminUsecaseInterface := usecases.NewSuAdminUsecase(suAdminRepo)
	suAdminHandler := handlers.NewSuAdminHandler(suAdminUsecaseInterface, config2)
	ginEngine := routes.NewGinEngine(userHandler, adminHandler, suAdminHandler)
	return ginEngine, nil
}