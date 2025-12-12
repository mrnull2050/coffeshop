package api

import (
	repo "coffeshop/Repo"
	"coffeshop/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	r1 := gin.New()
	r1.Use(gin.Logger(), gin.Recovery())
	userRepo := repo.UserRepo()
	v1 := r1.Group("/api/v1")
	{
		
	}
	cfg, err := config.LoadConfig()
	if err != nil {
		return
	}

	r1.Run(fmt.Sprintf("%d", cfg.Service.Port))

}
