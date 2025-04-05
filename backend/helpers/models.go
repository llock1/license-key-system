package helpers

import (
	"fmt"
	"license/database"
)

func DeleteModel(model interface{}, id string) error {
	var err error
	err = database.Client.First(&model, id).Error
	if err != nil {
		return err
	}

	err = database.Client.Delete(&model).Error

	if err != nil {
		return err
	}

	return nil
}

func CreateModel() {
	fmt.Println("Create MODEL")
}
