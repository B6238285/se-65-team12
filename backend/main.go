package main

import (
	"github.com/B6238285/se-65-team12/controller"
	"github.com/B6238285/se-65-team12/entity"
	"github.com/B6238285/se-65-team12/middlewares"
	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")

	{
		router.Use(middlewares.Authorizes())
		{
			//ข้อมูล path ผู้ป่วย//
			router.GET("/patient/get/:id", controller.GetPatient)
			router.GET("/patients/list", controller.ListPatient)
			router.POST("/patient/create", controller.CreatePatient)
			router.PATCH("/patient/edit", controller.UpdatePatient)
			router.DELETE("/patient/delet", controller.DeletePatient)
			//ข้อมูล path ใบสั่งยา//
			router.GET("/prescription/get/:id", controller.GetPrescription)
			router.GET("/prescriptions/list", controller.ListPrescription)
			router.POST("/prescription/create", controller.CreatePrescription)
			router.PATCH("/prescription/edit", controller.UpdatePrescription)
			router.DELETE("/prescription/delet", controller.DeletePrescription)
			//ข้อมูล path patient_type//
			router.GET("/patient/type/:id", controller.GetPatientType)
			router.GET("/patient/types", controller.ListPatientType)
			//ข้อมูล path patient_right//
			router.GET("/patient/right/:id", controller.GetPatientRight)
			router.GET("/patient/rights", controller.ListPatientRight)

		}
	}
	// // Signup User Route
	r.POST("/signup", controller.CreateLoginUser)
	// // login User Route
	r.POST("/login", controller.Login)

	// Run the server go run main.go
	//r.Run("localhost: " + PORT)
	r.Run("0.0.0.0:8080")

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
