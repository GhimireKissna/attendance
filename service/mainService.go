package service
import (
	"fmt"
	"gorm.io/gorm"
)

func MainMenu(db * gorm.DB){
	willContinue:
	for {
	var number int
	fmt.Println("SELECT YOUR OPTION")
	fmt.Println("0:Exit")
	fmt.Println("1:User:")
	fmt.Println("2:Contact:")
	fmt.Println("3:Attendance:")
	fmt.Scan(&number)
		switch number {
		case 0:
		break willContinue
		case 1:
		    caseUser(db)
		default:
		fmt.Println("Invalid option")
		}
	}
}