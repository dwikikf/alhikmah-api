package main

import (
	"log"

	"github.com/dwikikf/alhikmah-api/internal/handler"
	"github.com/dwikikf/alhikmah-api/internal/infrastucture/database"
	infraRepo "github.com/dwikikf/alhikmah-api/internal/infrastucture/repository"
	"github.com/dwikikf/alhikmah-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	classRepo := infraRepo.NewClassRepository(db)
	studentRepo := infraRepo.NewStudentRepository(db)
	uow := infraRepo.NewUnitOfWork(db)

	studentUsecase := usecase.NewStudentUseCase(classRepo, studentRepo, uow)
	studentHandler := handler.NewStudentHandler(studentUsecase)
	classUsecase := usecase.NewClassUseCase(classRepo, uow)
	classHandler := handler.NewClassHandler(classUsecase)

	r := gin.Default()

	r.GET("/students", studentHandler.GetAll)
	r.GET("/students/:id", studentHandler.GetByID)
	r.POST("/students", studentHandler.Create)
	r.PUT("/students/:id", studentHandler.Update)
	r.DELETE("/students/:id", studentHandler.Delete)

	r.GET("/classes", classHandler.GetAll)
	r.GET("/classes/:id", classHandler.GetClassByID)
	r.POST("/classes", classHandler.CreateClass)
	r.PUT("/classes/:id", classHandler.UpdateClass)
	r.DELETE("/classes/:id", classHandler.DeleteClass)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong Bro!",
		})
	})

	r.Run(":8080")
}
