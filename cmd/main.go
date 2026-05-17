package main

import (
	"log"

	"github.com/dwikikf/alhikmah-api/internal/di"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app, err := di.InitializeApp()
	if err != nil {
		log.Fatal(err)
	}

	studentHandler := app.Handlers.Student
	classHandler := app.Handlers.Class
	attendanceHandler := app.Handlers.Attendance

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		// Izinkan origin dari aplikasi
		AllowOrigins: []string{"http://localhost:5173"},

		// Method HTTP yang diizinkan untuk diakses frontend
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},

		// Header yang diizinkan dalam request
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},

		// Izinkan browser mengirimkan cookie atau header autentikasi
		AllowCredentials: true,
	}))

	r.GET("/students", studentHandler.GetAll)
	r.GET("/students/:id", studentHandler.GetByID)
	r.POST("/students", studentHandler.Create)
	r.PUT("/students/:id", studentHandler.Update)
	r.DELETE("/students/:id", studentHandler.Delete)

	r.GET("/classes", classHandler.GetAll)
	r.GET("/classes/:id", classHandler.GetByID)
	r.POST("/classes", classHandler.Create)
	r.PUT("/classes/:id", classHandler.Update)
	r.DELETE("/classes/:id", classHandler.Delete)

	r.GET("/attendances", attendanceHandler.GetAll)
	r.GET("/attendances/:id", attendanceHandler.GetByID)
	r.POST("/checkin", attendanceHandler.CheckedIn)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong Bro!",
		})
	})

	r.Run(":8080")
}
