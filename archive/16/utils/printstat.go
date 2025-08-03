package utils

import (
	"fmt"
	"modular-app/models"
)

func PrintStats(t *models.Tank) {
	fmt.Println("Модель:", t.Model)
	fmt.Println("Уровень:", t.Level)
	fmt.Println("Макс. HP:", t.MaxHP)
}
