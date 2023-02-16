package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
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

type Employee struct {
	gorm.Model
	FirstName string
	LastName  string
	Civ       string `gorm:"uniqueIndex"`
	Phone     string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Address   string
	//save Role ID in FK
	RoleID *uint
	//to easier for adding FK
	Role Role `gorm:"references:id"`
	//save Gender ID in FK
	GenderID *uint
	//to easier for adding FK
	Gender Gender `gorm:"references:id"`
	//save Department ID in FK
	DepartmentID *uint
	//to easier for adding FK
	Department        Department     `gorm:"references:id"`
	Patient           []Patient      `gorm:"foreignKey:EmployeeID"`
	Prescription      []Prescription `gorm:"foreignKey:EmployeeID"`
	PrescriptionOrder []Prescription `gorm:"foreignKey:OrderID"`
}

type Role struct {
	gorm.Model
	Name      string
	Employees []Employee `gorm:"foreignKey:RoleID"`
}
type Gender struct {
	gorm.Model
	Name      string
	Employees []Employee `gorm:"foreignKey:GenderID"`
	Patient   []Patient  `gorm:"foreignKey:GenderID"`
}
type Department struct {
	gorm.Model
	Type      string
	Employees []Employee `gorm:"foreignKey:DepartmentID"`
}

type Patient struct {
	gorm.Model
	Civ         string    `gorm:"uniqueIndex" valid:"matches(^([0-9]{13})$)~Identification Number must have only number and lenght is 13,required~Identification Number cannot be blank"`
	FirstName   string    `valid:"required~FirstName cannot be blank,thai_eng_char_vowel~FirstName must have only character"`
	LastName    string    `valid:"required~LastName cannot be blank,thai_eng_char_vowel~LastName must have only character"`
	Age         int       `valid:"range(0|122)~Age not in range 0-122"`
	Weight      float32   `valid:"range(0|595)~Weight not in range 0-595"`
	Underlying  string    `valid:"required~Underlying cannot be blank"`
	Brithdate   time.Time `valid:"past~Brithdate must be in the past"`
	PatientTime time.Time

	//FK
	PatientTypeID  *uint
	EmployeeID     *uint
	PatientRightID *uint
	GenderID       *uint

	//JOIN
	PatientType  PatientType  `gorm:"references:id" valid:"-"`
	Employee     Employee     `gorm:"references:id" valid:"-"`
	PatientRight PatientRight `gorm:"references:id" valid:"-"`
	Gender       Gender       `gorm:"references:id" valid:"-"`

	Prescription []Prescription `gorm:"foreignKey:PatientID"`
}

// -------------------------------------------------------------------------------//
type Medicine struct {
	gorm.Model
	Drug         string
	Cost         float32
	Prescription []Prescription `gorm:"foreignKey:MedicineID"`
}

type Prescription struct {
	gorm.Model
	Annotation string `valid:"required~Annotation cannot be blank,maxstringlength(300)~Annotation length is too long,thai_eng_char_vowel_number~Annotation must have only character and number"`
	ScriptTime time.Time

	//FK
	PatientID  *uint
	MedicineID *uint
	EmployeeID *uint
	OrderID    *uint

	//JOIN
	Patient  Patient  `gorm:"references:id" valid:"-"`
	Medicine Medicine `gorm:"references:id" valid:"-"`
	Employee Employee `gorm:"references:id" valid:"-"`
	Order    Employee `gorm:"references:id" valid:"-"`
}

func init() {
	govalidator.CustomTypeTagMap.Set("past", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.Before(time.Now())
	})

	govalidator.CustomTypeTagMap.Set("future", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.After(time.Now())
	})

	govalidator.CustomTypeTagMap.Set("Now", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.Equal(time.Now())
	})

	govalidator.CustomTypeTagMap.Set("positive", func(i interface{}, context interface{}) bool {
		num := i
		return num.(int) >= 0
	})
	govalidator.CustomTypeTagMap.Set("positivenum", func(i interface{}, context interface{}) bool {
		num := i
		return num.(int) > 0
	})

	govalidator.CustomTypeTagMap.Set("thai_eng_char_vowel", func(i interface{}, context interface{}) bool {
		s, ok := i.(string)
		if !ok {
			return false
		}

		for _, c := range s {
			if !(('ก' <= c && c <= 'ฮ') || ('ะ' <= c && c <= 'ู') || ('เ' <= c && c <= '์') || ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')) {
				return false
			}
		}
		return true
	})

	govalidator.CustomTypeTagMap.Set("thai_eng_char_vowel_number", func(i interface{}, context interface{}) bool {
		s, ok := i.(string)
		if !ok {
			return false
		}

		for _, c := range s {
			if !(('ก' <= c && c <= 'ฮ') || ('ะ' <= c && c <= 'ู') || ('เ' <= c && c <= '์') || ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9') || ('๐' <= c && c <= '๙') || (' ' == c)) {
				return false
			}
		}
		return true
	})

}
