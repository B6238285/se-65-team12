package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		&PatientType{},
		&PatientRight{},
		&Gender{},
		&Employee{},
		&Patient{},
	)

	db = database

	password1, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	password2, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	password3, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	password4, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	password5, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	patienttype1 := PatientType{
		Type: "ปกติ",
	}
	db.Model(&PatientType{}).Create(&patienttype1)

	patienttype2 := PatientType{
		Type: "อุบัติเหตุ/ฉุกเฉิน",
	}
	db.Model(&PatientType{}).Create(&patienttype2)

	patienttype3 := PatientType{
		Type: "เด็กแรกเกิด",
	}
	db.Model(&PatientType{}).Create(&patienttype3)

	patienttype4 := PatientType{
		Type: "คลอดบุตร",
	}
	db.Model(&PatientType{}).Create(&patienttype4)

	patientright1 := PatientRight{
		Type: "เจ็บป่วยปกติ",
	}
	db.Model(&PatientRight{}).Create(&patientright1)

	patientright2 := PatientRight{
		Type: "เจ็บป่วยฉุกเฉิน/อุบัติเหตุ",
	}
	db.Model(&PatientRight{}).Create(&patientright2)

	patientright3 := PatientRight{
		Type: "คลอดบุตร",
	}
	db.Model(&PatientRight{}).Create(&patientright3)

	male := Gender{
		Identity: "ชาย",
	}
	db.Model(&Gender{}).Create(&male)

	female := Gender{
		Identity: "หญิง",
	}
	db.Model(&Gender{}).Create(&female)

	doctor := Employee{
		FirstName:  "preechapat",
		LastName:   "anpanit",
		Civ:        "1250008896345",
		Phone:      "0811111111",
		Email:      "preechapat@mail.com",
		Password:   string(password1),
		Address:    "45 บ้านฟ้าปิยรมณ์ ต.บึงคำพร้อย อ.ลำลูกกา จ.ปทุมธานี 11350",
		Role:       "Doctor",
		Department: "Surgical Department",
		Gender:     "Women",
	}
	db.Model(&Employee{}).Create(&doctor)

	humanresourse := Employee{
		FirstName:  "aam",
		LastName:   "love",
		Civ:        "1234567890124",
		Phone:      "0899999999",
		Email:      "kawin@mail.com",
		Password:   string(password2),
		Address:    "37/123 บ้านหนองพิลุม ต.บ้านท่า อ.เมือง จ.ปราจีนบุรี 12150",
		Role:       "Human Resourse",
		Department: "Radiology Department",
		Gender:     "Women",
	}
	db.Model(&Employee{}).Create(&humanresourse)

	nurse := Employee{
		FirstName:  "sirinya",
		LastName:   "kotpanya",
		Civ:        "1258896675256",
		Phone:      "0633333333",
		Email:      "sirinya@mail.com",
		Password:   string(password3),
		Address:    "23/777 บ้านหนองบึง ต.ท่าช้าง อ.เมือง จ.ลพบุรี 13000",
		Role:       "Nurse",
		Department: "Human and Resourse",
		Gender:     "Women",
	}
	db.Model(&Employee{}).Create(&nurse)

	pharmacist := Employee{
		FirstName:  "poramate",
		LastName:   "jitlamom",
		Civ:        "1234445678055",
		Phone:      "0432536678",
		Email:      "poramate@mail.com",
		Password:   string(password4),
		Address:    "56/77 บ้านตาก ต.หนองคุ้ม อ.ระงัน จ.ระยอง 13500",
		Role:       "Pharmacist",
		Department: "Phamarceutical Department",
		Gender:     "Women",
	}
	db.Model(&Employee{}).Create(&pharmacist)

	accounting := Employee{
		FirstName:  "siwana",
		LastName:   "julaiwarnsutti",
		Civ:        "1274563346856",
		Phone:      "0456673256",
		Email:      "siwa@mail.com",
		Password:   string(password5),
		Address:    "324 ฟาร์มโชคชัย ต.โชคชัย อ.เมือง จ.นครราชศรีมา 12300",
		Role:       "Accounting",
		Department: "Finance Department",
		Gender:     "Women",
	}
	db.Model(&Employee{}).Create(&accounting)

	patient1 := Patient{
		Civ:          "1309902756650",
		FirstName:    "paramet",
		LastName:     "chitlamom",
		PatientType:  patienttype1,
		Employee:     nurse,
		PatientRight: patientright1,
		Gender:       male,
		Age:          22,
		Weight:       56.23,
		Brithdate:    time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		Underlying:   "ภูมิแพ้",
		PatientTime:  time.Date(2002, 12, 3, 0, 0, 0, 0, time.UTC),
	}
	db.Model(&Patient{}).Create(&patient1)

	patient2 := Patient{
		Civ:          "1309956926351",
		FirstName:    "parama",
		LastName:     "chitlamom",
		PatientType:  patienttype2,
		Employee:     nurse,
		PatientRight: patientright2,
		Gender:       female,
		Age:          26,
		Weight:       56.20,
		Underlying:   "หอบหืด",
		Brithdate:    time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		PatientTime:  time.Date(2002, 12, 12, 0, 0, 0, 0, time.UTC),
	}
	db.Model(&Patient{}).Create(&patient2)

}
