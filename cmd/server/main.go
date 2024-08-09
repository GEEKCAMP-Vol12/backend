package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/taku10101/config"
	"github.com/taku10101/internal/delivery/http"
	"github.com/taku10101/internal/repository"
	"github.com/taku10101/internal/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
    config := config.LoadConfig()
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
        config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&repository.User{}, &repository.Caffeine{}, &repository.Sleep{})

    r := gin.Default()

    userRepo := repository.NewUserRepository(db)
    caffeineRepo := repository.NewCaffeineRepository(db)
    sleepRepo := repository.NewSleepRepository(db)

    userUseCase := usecase.NewUserUseCase(userRepo)
    caffeineUseCase := usecase.NewCaffeineUseCase(caffeineRepo)
    sleepUseCase := usecase.NewSleepUseCase(sleepRepo)

    
    http.NewUserHandler(r, userUseCase)
    http.NewCaffeineHandler(r, caffeineUseCase)
    http.NewSleepHandler(r, sleepUseCase)

    r.Run(config.ServerAddress)
}
