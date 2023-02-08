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
			router.DELETE("/patient/delet/:id", controller.DeletePatient)
			//ข้อมูล path ใบสั่งยา//
			router.GET("/prescription/get/:id", controller.GetPrescription)
			router.GET("/prescriptions/list", controller.ListPrescription)
			router.POST("/prescription/create", controller.CreatePrescription)
			router.PATCH("/prescription/edit", controller.UpdatePrescription)
			router.DELETE("/prescription/delet/:id", controller.DeletePrescription)
			router.GET("/employee/doctor/list", controller.ListEmployeeDoctor)
			//ข้อมูล path patient_type//
			router.GET("/patient/type/:id", controller.GetPatientType)
			router.GET("/patient/types/list", controller.ListPatientType)
			//ข้อมูล path patient_right//
			router.GET("/patient/right/:id", controller.GetPatientRight)
			router.GET("/patient/rights/list", controller.ListPatientRight)
			//ข้อมูล path ข้อมูลยา//
			router.GET("/medicine/:id", controller.GetMedicine)
			router.GET("/medicine/list", controller.ListMedicine)

			router.GET("/employee/get/:id", controller.GetEmployee)
			// List
			router.GET("/employees/list", controller.ListEmployee)
			// Create
			router.POST("/employee/create", controller.CreateEmployee)
			// UPDATE
			router.PATCH("/employee/update", controller.UpdateEmployee)
			// DELETE
			router.DELETE("/employees/delete/:id", controller.DeleteEmployee)
			// ----------------- Employee ----------------------------
			// Role Routes
			//r.GET("/roles/list", controller.ListRole)
			//r.GET("/role/get/:id", controller.GetRole)
			// router.POST("/roles", controller.CreateRole)
			// router.PATCH("/roles", controller.UpdateRole)
			// router.DELETE("/roles/:id", controller.DeleteRole)
			// Gender
			router.GET("/genders/list", controller.ListGender)
			router.GET("/gender/get/:id", controller.GetGender)
			// Department
			//r.GET("/departments/list", controller.ListDepartment)
			//r.GET("/department/get/:id", controller.GetDepartment)
			//r.GET("/departmentbyrole/get/:id", controller.ListDepartmentByRole)

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
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
