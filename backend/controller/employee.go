package controller

import (
	"fmt"
	"net/http"

	"github.com/B6238285/se-65-team12/entity"
	"github.com/gin-gonic/gin"
)

// ---------------- Employee ---------------

// List all Employee
// GET /employees
func ListEmployee(c *gin.Context) {
	var employees []entity.Employee
	if err := entity.DB().Preload("Gender").Preload("Role").Preload("Department").Raw("SELECT * FROM employees").Find(&employees).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": employees,
	})
}

// Get Once Employee
// GET /employee/:id
func GetEmployee(c *gin.Context) {
	var employee entity.Employee
	id := c.Param("id")
	if err := entity.DB().Preload("Gender").Preload("Role").Preload("Department").Raw("SELECT * FROM employees WHERE id = ?", id).Find(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": employee,
	})
}

// Create Employee
// POST /employee
func CreateEmployee(c *gin.Context) {
	//main
	var employee entity.Employee
	//relation
	var role entity.Role
	var gender entity.Gender
	var department entity.Department

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	// get role from database
	if tx := entity.DB().Where("id = ?", employee.RoleID).First(&role); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "role is not found",
		})
		return
	}

	// get gender from database
	if tx := entity.DB().Where("id = ?", employee.GenderID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "gender is not found",
		})
		return
	}

	// get department from database
	if tx := entity.DB().Where("id = ?", employee.Department).First(&department); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Status is not found",
		})
		return
	}

	emp := entity.Employee{
		FirstName:  employee.FirstName,
		LastName:   employee.LastName,
		Civ:        employee.Civ,
		Phone:      employee.Phone,
		Email:      employee.Email,
		Password:   employee.Password,
		Address:    employee.Address,
		Role:       role,
		Department: department,
		Gender:     gender,
	}

	if err := entity.DB().Create(&emp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "create Employee Success",
		"data":   emp,
	})
}

// Update Employee
// PATCH /employee
func UpdateEmployee(c *gin.Context) {

	//main
	var employee entity.Employee
	var oldemployee entity.Employee
	//relation
	var role entity.Role
	var gender entity.Gender
	var department entity.Department

	// Bind Json to var emp
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	// Check emp is haved ?
	if tx := entity.DB().Where("id = ?", employee.ID).First(&oldemployee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Employee id = %d not found", employee.ID)})
		c.Abort()
		return
	}

	if employee.FirstName == "" {
		employee.FirstName = oldemployee.FirstName
	}

	if employee.LastName == "" {
		employee.LastName = oldemployee.LastName
	}

	if employee.Civ == "" {
		employee.Civ = oldemployee.Civ
	}

	if employee.Phone == "" {
		employee.Phone = oldemployee.Phone
	}

	if employee.Email == "" {
		employee.Email = oldemployee.Email
	}

	if employee.Password == "" {
		employee.Password = oldemployee.Password
	}
	// if new have role id
	if employee.RoleID != nil {
		if tx := entity.DB().Where("id = ?", employee.RoleID).First(&role); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not found role"})
			return
		}
		fmt.Print("NOT NULL")
		employee.Role = role
	} else {
		if tx := entity.DB().Where("id = ?", oldemployee.Role).First(&role); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not found user"})
			return
		}
		fmt.Print("NULL")
		employee.Role = role
	}

	// if new have department id
	if employee.DepartmentID != nil {
		if tx := entity.DB().Where("id = ?", employee.DepartmentID).First(&department); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not found status"})
			return
		}
		employee.Department = department
	} else {
		if tx := entity.DB().Where("id = ?", oldemployee.Department).First(&department); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not found status"})
			return
		}
		employee.Department = department
	}

	//if new have gender id
	if employee.GenderID != nil {
		if tx := entity.DB().Where("id = ?", employee.GenderID).First(&gender); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not found workingarea"})
			return
		}
		employee.Gender = gender
	} else {
		if tx := entity.DB().Where("id = ?", oldemployee.GenderID).First(&gender); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not found workingarea"})
			return
		}
		employee.Gender = gender
	}

	// Update emp in database
	if err := entity.DB().Save(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "Update Success",
		"data":   employee,
	})

}

// Delete Employee
// DELETE /employee/:id
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM employees WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee is not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})

}
