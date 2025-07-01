package main

import (
	"modular-app/models"
	"modular-app/utils"
)

func main() {
	t := models.Tank{
		Model: "Pz-4",
		Level: 6,
		MaxHP: 500,
	}
	utils.PrintStats(&t)
}
