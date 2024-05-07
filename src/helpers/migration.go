package helpers

import (
	"gofiber-batch2/src/configs"
	"gofiber-batch2/src/models"
)

func Migration() {
	configs.DB.AutoMigrate(&models.Product{})
}
