package initializers

import "github.com/UjjwalMahar/llamakraft/auth/models"



func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}