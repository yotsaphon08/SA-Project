package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yotsaphon08/sa-64-project/controller"
	"github.com/yotsaphon08/sa-64-project/entity"
	"github.com/yotsaphon08/sa-64-project/middlewares"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// Informer Routes
			protected.GET("/informers", controller.ListInformer)
			protected.GET("/informer/:id", controller.GetInformer)
			//protected.POST("/informers", controller.CreateInformer)
			protected.PATCH("/informers", controller.UpdateInformer)
			protected.DELETE("/informers/:id", controller.DeleteInformer)

			// User Routes
			protected.GET("/users", controller.ListUsers)
			protected.GET("/user/:id", controller.GetUser)
			protected.PATCH("/users", controller.UpdateUser)
			protected.DELETE("/users/:id", controller.DeleteUser)

			// Characteristic Routes
			protected.GET("/characteristics", controller.ListCharacteristics)
			protected.GET("/characteristic/:id", controller.GetCharacteristic)
			protected.POST("/characteristics", controller.CreateCharacteristic)
			protected.PATCH("/characteristics", controller.UpdateCharacteristic)
			protected.DELETE("/characteristics/:id", controller.DeleteCharacteristic)

			// Patient Routes
			protected.GET("/patients", controller.ListPatient)
			protected.GET("/patient/:id", controller.GetPatient)
			//protected.POST("/patients", controller.CreatePatient)
			protected.PATCH("/patients", controller.UpdatePatient)
			protected.DELETE("/patients/:id", controller.DeletePatient)

			// Level Routes
			protected.GET("/levels", controller.ListLevels)
			protected.GET("/level/:id", controller.GetLevel)
			protected.POST("/levels", controller.CreateLevel)
			protected.PATCH("/levels", controller.UpdateLevel)
			protected.DELETE("/levels/:id", controller.DeleteLevel)

			// Case Routes
			protected.GET("/cases", controller.ListCase)
			protected.GET("/case/:id", controller.GetCase)
			protected.POST("/cases", controller.CreateCase)
			protected.PATCH("/cases", controller.UpdateCase)
			protected.DELETE("/cases/:id", controller.DeleteCase)

			// Register Routes
			protected.GET("/registers", controller.ListRegisters)
			protected.GET("/register/:id", controller.GetRegister)
			protected.POST("/registers", controller.CreateRegister)
			protected.PATCH("/registers", controller.UpdateRegister)
			protected.DELETE("/registers/:id", controller.DeleteRegister)

			// Assesses Routes

			protected.GET("/assesses", controller.ListAssess)
			protected.GET("/assess/:id", controller.GetAssess)
			protected.POST("/assesses", controller.CreateAssess)
			protected.PATCH("/assesses", controller.UpdateAssess)
			protected.DELETE("/assesses/:id", controller.DeleteAssess)

			// State Routes

			protected.GET("/states", controller.ListState)
			protected.GET("/state/:id", controller.GetState)
			protected.POST("/states", controller.CreateState)
			protected.PATCH("/states", controller.UpdateState)
			protected.DELETE("/states/:id", controller.DeleteState)

			// Symptom Routes

			protected.GET("/symptoms", controller.ListSymptom)
			protected.GET("/symptom/:id", controller.GetSymptom)
			protected.POST("/symptoms", controller.CreateSymptom)
			protected.PATCH("/symptoms", controller.UpdateSymptom)
			protected.DELETE("/symptoms/:id", controller.DeleteSymptom)

			// AssessmentSheet Routes

			protected.GET("/assessment_sheets", controller.ListAssessmentSheet)
			protected.GET("/assessment_sheet/:id", controller.GetAssessmentSheet)
			protected.POST("/assessment_sheets", controller.CreateAssessmentSheet)
			protected.PATCH("/assessment_sheets", controller.UpdateAssessmentSheet)
			protected.DELETE("/assessment_sheets/:id", controller.DeleteAssessmentSheet)

			// AmbulanceType Routes
			protected.GET("/ambulanceTypes", controller.ListAmbulanceType)
			protected.GET("/ambulanceType/:id", controller.GetAmbulanceType)
			protected.POST("/ambulanceTypes", controller.CreateAmbulanceType)
			protected.PATCH("/ambulanceTypes", controller.UpdateAmbulanceType)
			protected.DELETE("/ambulanceTypes/:id", controller.DeleteAmbulanceType)

			// Brand Routes
			protected.GET("/brands", controller.ListBrand)
			protected.GET("/brand/:id", controller.GetBrand)
			protected.POST("/brands", controller.CreateBrand)
			protected.PATCH("/brands", controller.UpdateBrand)
			protected.DELETE("/brands/:id", controller.DeleteBrand)

			// Status Routes
			protected.GET("/statuses", controller.ListStatuses)
			protected.GET("/status/:id", controller.GetStatus)
			protected.POST("/statuses", controller.CreateStatus)
			protected.PATCH("/statuses", controller.UpdateStatus)
			protected.DELETE("/statuses/:id", controller.DeleteStatus)

			protected.GET("/path_statuses", controller.ListPathStatus)
			protected.GET("/path_status/:id", controller.GetPathStatus)
			protected.POST("/path_statuses", controller.CreatePathStatus)
			protected.PATCH("/path_statuses", controller.UpdatePathStatus)
			protected.DELETE("/path_statuses/:id", controller.DeletePathStatus)

			protected.GET("/check_lists", controller.ListCheckList)
			protected.GET("/check_list/:id", controller.GetCheckList)
			protected.POST("/check_lists", controller.CreateCheckList)
			protected.PATCH("/check_lists", controller.UpdateCheckList)
			protected.DELETE("/check_lists/:id", controller.DeleteCheckList)

			protected.GET("/car_paths", controller.ListCarPath)
			protected.GET("/car_path/:id", controller.GetCarPath)
			protected.POST("/car_paths", controller.CreateCarPath)
			protected.PATCH("/car_paths", controller.UpdateCarPath)
			protected.DELETE("/car_paths/:id", controller.DeleteCarPath)

		}
	}

	// Informer Routes
	r.POST("/informers", controller.CreateInformer)

	// Patient Routes
	r.POST("/patients", controller.CreatePatient)

	// User Routes
	r.POST("/users", controller.CreateUser)

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Authentication1 Routes
	r.POST("/login1", controller.Login1)

	// Run the server
	r.Run()
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
