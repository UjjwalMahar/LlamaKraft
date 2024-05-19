package initializers

import "github.com/UjjwalMahar/llamakraft/models"



func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}