package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Level{},
		&Informer{},
		&Patient{},
		&Characteristic{},
		&Case{},

		&User{},
		&Register{},

		&Assess{},
		&Symptom{},
		&State{},
		&Case{},
		&AssessmentSheet{},

		&CheckList{},
		&Car_path{},
		&Path_status{},

		&Ambulance{},
		&AmbulanceType{},
		&Status{},
		&Brand{},
	)

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	db.Model(&Informer{}).Create(&Informer{
		Name:     "Butsakorn",
		Email:    "butsakorn@gmail.com",
		Tel:      "065-280-3444",
		Password: string(password),
	})

	db.Model(&Informer{}).Create(&Informer{
		Name:     "Yotsaphon",
		Email:    "yotsaphon@gmail.com",
		Tel:      "082-875-3990",
		Password: string(password),
	})

	db.Model(&Informer{}).Create(&Informer{
		Name:     "Name",
		Email:    "name@gmail.com",
		Tel:      "089-998-8877",
		Password: string(password),
	})

	var butsakorn Informer
	var yotsaphon Informer
	var name Informer
	db.Raw("SELECT * FROM informers WHERE email = ?", "butsakorn@gmail.com").Scan(&butsakorn)
	db.Raw("SELECT * FROM informers WHERE email = ?", "yotsaphon@gmail.com").Scan(&yotsaphon)
	db.Raw("SELECT * FROM informers WHERE email = ?", "name@gmail.com").Scan(&name)

	// --- Characteristic

	congenitaldisease := Characteristic{
		Category: "โรคประจำตัว",
	}
	db.Model(&Characteristic{}).Create(&congenitaldisease)

	accident := Characteristic{
		Category: "เกิดอุบัติเหตุ",
	}
	db.Model(&Characteristic{}).Create(&accident)

	// --- Level
	little := Level{
		Rating: "เล็กน้อย",
	}
	db.Model(&Level{}).Create(&little)

	moderate := Level{
		Rating: "ปานกลาง",
	}
	db.Model(&Level{}).Create(&moderate)

	serious := Level{
		Rating: "หนัก",
	}
	db.Model(&Level{}).Create(&serious)

	// --- Patient
	PatientOfTonphaii := Patient{
		Name:   "Tonphaii",
		Age:    "20",
		Gender: "หญิง",
	}
	db.Model(&Patient{}).Create(&PatientOfTonphaii)

	PatientOfAwesome := Patient{
		Name:   "Awesome",
		Age:    "21",
		Gender: "ชาย",
	}
	db.Model(&Patient{}).Create(&PatientOfAwesome)

	PatientOfHunki := Patient{
		Name:   "Hunki",
		Age:    "26",
		Gender: "ชาย",
	}
	db.Model(&Patient{}).Create(&PatientOfHunki)

	// Case 1
	Case1 := Case{
		Characteristic: accident,
		Level:          serious,
		CaseTime:       time.Now(),
		Address:        "หอ หลังมอ มข. ต.ในเมือง อ.เมืองขอนแก่น จ.ขอนแก่น",
		Patient:        PatientOfHunki,
		Informer:       butsakorn,
	}
	db.Model(&Case{}).Create(&Case1)

	// Case 2
	Case2 := Case{
		Characteristic: congenitaldisease,
		Level:          little,
		CaseTime:       time.Now(),
		Address:        "หน้าวัดพระธาตุพนม ต.ธาตุพนม อ.ธาตุพนม จ.นครพนม",
		Patient:        PatientOfTonphaii,
		Informer:       yotsaphon,
	}
	db.Model(&Case{}).Create(&Case2)

	// Case 3
	Case3 := Case{
		Characteristic: accident,
		Level:          moderate,
		CaseTime:       time.Now(),
		Address:        "ถนนหน้าเซเว่นประตู 4 มทส. ต.สุรนารี อ.เมือง จ.นครราชสีมา",
		Patient:        PatientOfAwesome,
		Informer:       name,
	}
	db.Model(&Case{}).Create(&Case3)

	db.Model(&User{}).Create(&User{
		Name:     "Chanon",
		Email:    "chanon@gmail.com",
		Password: string(password),
	})

	db.Model(&User{}).Create(&User{
		Name:     "Apple",
		Email:    "Apple@gmail.com",
		Password: string(password),
	})

	var chanon User
	var apple User
	db.Raw("SELECT * FROM users WHERE email = ?", "chanon@gmail.com").Scan(&chanon)
	db.Raw("SELECT * FROM users WHERE email = ?", "Apple@gmail.com").Scan(&apple)

	// ระบบบันทึกข้อมูลรถพยาบาล
	// --- Type Data
	basicLifeSupport := AmbulanceType{
		TypeName: "รถพยาบาลปฏิบัติการพื้นฐาน",
	}
	db.Model(&AmbulanceType{}).Create(&basicLifeSupport)

	advancedLifeSupport := AmbulanceType{
		TypeName: "รถพยาบาลปฏิบัติการฉุกเฉิน",
	}
	db.Model(&AmbulanceType{}).Create(&advancedLifeSupport)

	//----Brand Data
	brand1 := Brand{
		BrandName: "Honda Stepwgn Spada",
	}
	db.Model(&Brand{}).Create(&brand1)

	brand2 := Brand{
		BrandName: "Toyota Hilux Vigo",
	}
	db.Model(&Brand{}).Create(&brand2)

	brand3 := Brand{
		BrandName: "Nilsson Volvo",
	}
	db.Model(&Brand{}).Create(&brand3)

	//----Status Data
	status1 := Status{
		StatusName: "พร้อมใช้งาน",
	}
	db.Model(&Status{}).Create(&status1)

	status2 := Status{
		StatusName: "ไม่พร้อมใช้งาน",
	}
	db.Model(&Status{}).Create(&status2)

	// ambulance 1
	Ambulance1 := Ambulance{
		AmbulanceType: basicLifeSupport,
		Registration:  "กข1234",
		Brand:         brand1,
		RecordingTime: time.Now(),
		Status:        status1,
		Owner:         chanon,
	}
	db.Model(&Ambulance{}).Create(&Ambulance1)

	// ambulance 2
	Ambulance2 := Ambulance{
		AmbulanceType: advancedLifeSupport,
		Registration:  "งม12345",
		Brand:         brand2,
		RecordingTime: time.Now(),
		Status:        status2,
		Owner:         chanon,
	}
	db.Model(&Ambulance{}).Create(&Ambulance2)

	// ambulance 3
	Ambulance3 := Ambulance{
		AmbulanceType: advancedLifeSupport,
		Registration:  "พส8888",
		Brand:         brand3,
		RecordingTime: time.Now(),
		Status:        status1,
		Owner:         apple,
	}
	db.Model(&Ambulance{}).Create(&Ambulance3)

	// ระบบประเมินสถานะผู้ป่วย
	// --- Assess Data
	Assess1 := Assess{
		AssessData: "ปกติ",
	}
	db.Model(&Assess{}).Create(&Assess1)

	Assess2 := Assess{
		AssessData: "ฉุกเฉิน",
	}
	db.Model(&Assess{}).Create(&Assess2)

	// State Data
	State1 := State{
		StateData: "ไม่มีสติ",
	}
	db.Model(&State{}).Create(&State1)

	State2 := State{
		StateData: "มีสติ",
	}
	db.Model(&State{}).Create(&State2)

	// Symptom Data
	Symptom1 := Symptom{
		SymptomData: "กระดูกหัก",
	}
	db.Model(&Symptom{}).Create(&Symptom1)

	Symptom2 := Symptom{
		SymptomData: "เลือดคั่งในสมอง",
	}
	db.Model(&Symptom{}).Create(&Symptom2)

	Symptom3 := Symptom{
		SymptomData: "ลมบ้าหมู",
	}
	db.Model(&Symptom{}).Create(&Symptom3)

	Symptom4 := Symptom{
		SymptomData: "เลือดจาง",
	}
	db.Model(&Symptom{}).Create(&Symptom4)

	// ams 1
	AssessmentSheet1 := AssessmentSheet{
		Case:       Case1,
		Symptom:    Symptom1,
		State:      State1,
		Assess:     Assess1,
		AssessTime: time.Now(),
		Owner:      chanon,
	}
	db.Model(&AssessmentSheet{}).Create(&AssessmentSheet1)

	// ams 2
	AssessmentSheet2 := AssessmentSheet{
		Case:       Case2,
		Symptom:    Symptom3,
		State:      State2,
		Assess:     Assess2,
		AssessTime: time.Now(),
		Owner:      chanon,
	}
	db.Model(&AssessmentSheet{}).Create(&AssessmentSheet2)

	// ams 3
	AssessmentSheet3 := AssessmentSheet{
		Case:       Case3,
		Symptom:    Symptom2,
		State:      State2,
		Assess:     Assess2,
		AssessTime: time.Now(),
		Owner:      apple,
	}
	db.Model(&AssessmentSheet{}).Create(&AssessmentSheet3)

	// ระบบลงทะเบียนรถ

	// Register1
	db.Model(&Register{}).Create(&Register{
		Ambulance:       Ambulance1,
		Case:            Case1,
		RegisterTime:    time.Now(),
		AssessmentSheet: AssessmentSheet1,
		Owner:           chanon,
	})
	// Register2
	db.Model(&Register{}).Create(&Register{
		Ambulance:       Ambulance2,
		Case:            Case2,
		RegisterTime:    time.Now(),
		AssessmentSheet: AssessmentSheet2,
		Owner:           chanon,
	})
	// Register3
	db.Model(&Register{}).Create(&Register{
		Ambulance:       Ambulance3,
		Case:            Case3,
		RegisterTime:    time.Now(),
		AssessmentSheet: AssessmentSheet3,
		Owner:           apple,
	})

	// ระบบตรวจเช็คสภาพรถ

	pathstatus1 := Path_status{
		Status: "ใช้งานได้ปกติ",
	}
	db.Model(&Path_status{}).Create(&pathstatus1)

	pathstatus2 := Path_status{
		Status: "เกินระยะที่กำหนด",
	}
	db.Model(&Path_status{}).Create(&pathstatus2)

	pathstatus3 := Path_status{
		Status: "ชำรุดเสียหาย",
	}
	db.Model(&Path_status{}).Create(&pathstatus3)
	pathstatus4 := Path_status{
		Status: "น้อยกว่าระยะที่กำหนด",
	}
	db.Model(&Path_status{}).Create(&pathstatus4)

	path1 := Car_path{
		Path_titel: "ตรวจเช็คแบตเตอรี่",
	}
	db.Model(&Car_path{}).Create(&path1)
	path2 := Car_path{
		Path_titel: "ล้อและยางรถยนต์",
	}
	db.Model(&Car_path{}).Create(&path2)
	path3 := Car_path{
		Path_titel: "เช็คช่วงล่าง",
	}
	db.Model(&Car_path{}).Create(&path3)
	path4 := Car_path{
		Path_titel: "เช็คระดับน้ำมันเบรก และ ระบบเบรก",
	}
	db.Model(&Car_path{}).Create(&path4)
	path5 := Car_path{
		Path_titel: "เช็คระบบไฟส่องสว่าง",
	}
	db.Model(&Car_path{}).Create(&path5)
	path6 := Car_path{
		Path_titel: "น้ำมันเครื่อง",
	}
	db.Model(&Car_path{}).Create(&path6)
	path7 := Car_path{
		Path_titel: "เช็คหม้อน้ำ ท่อยาง และ ระบบหล่อเย็น",
	}
	db.Model(&Car_path{}).Create(&path7)
	path8 := Car_path{
		Path_titel: "ชุดเครื่องมือประจำรถ",
	}
	db.Model(&Car_path{}).Create(&path8)

	// CheckList1
	db.Model(&CheckList{}).Create(&CheckList{
		Checked_time: time.Now(),
		Ambulance:    Ambulance1,
		Car_path:     path1,
		Path_status:  pathstatus1,
		Owner:        chanon,
	})

	// CheckList1
	db.Model(&CheckList{}).Create(&CheckList{
		Checked_time: time.Now(),
		Ambulance:    Ambulance2,
		Car_path:     path2,
		Path_status:  pathstatus2,
		Owner:        apple,
	})

}
