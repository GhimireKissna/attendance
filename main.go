package main
import (
	"fmt"
	"attendance/service"
	"attendance/model"
	
)
func main() {
	db, err := connect()
	if err != nil {
		fmt.Println("error validating sql.Open arguments")
		panic(err.Error())
	}
	db.AutoMigrate(&model.User{})
	service.MainMenu(db)

	}

