package entity

import (
	"time"

	"gorm.io/gorm"
)

type PatientType struct {
	gorm.Model
	Type    string
	Patient []Patient `gorm:"foreignKey:PatientTypeID"`
}

type PatientRight struct {
	gorm.Model
	Type    string
	Patient []Patient `gorm:"foreignKey:PatientRightID"`
}

type Gender struct {
	gorm.Model
	Identity string
	Patient  []Patient `gorm:"foreignKey:GenderID"`
}

type Employee struct {
	gorm.Model
	FirstName  string
	LastName   string
	Civ        string `gorm:"uniqueIndex"`
	Phone      string
	Email      string `gorm:"uniqueIndex"`
	Password   string
	Address    string
	Role       string
	Department string
	Gender     string
	Patient    []Patient `gorm:"foreignKey:EmployeeID"`
}

type Patient struct {
	gorm.Model
	Civ         string `gorm:"uniqueIndex"`
	FirstName   string
	LastName    string
	Age         int
	Weight      float32
	Underlying  string
	Brithdate   time.Time
	PatientTime time.Time

	//FK
	PatientTypeID  *uint
	EmployeeID     *uint
	PatientRightID *uint
	GenderID       *uint

	//JOIN
	PatientType  PatientType  `gorm:"references:id"`
	Employee     Employee     `gorm:"references:id"`
	PatientRight PatientRight `gorm:"references:id"`
	Gender       Gender       `gorm:"references:id"`
}
