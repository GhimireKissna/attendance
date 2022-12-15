package service
import(
	"fmt"
	"gorm.io/gorm"
	"attendance/model"
)
func caseUser(db*gorm.DB ){
	var number int
	fmt.Println("SELECT YOUR OPTION")
	fmt.Println("0:Go Back To Main Menu")
	fmt.Println("1:Create User ")
	fmt.Println("2:Get User By Id")
	fmt.Println("3:update user")
	fmt.Println("4:Delete User")
	fmt.Println("5:Get All Users")
	fmt.Print("Your Choice:- ")
	fmt.Scan(&number)
		switch number {
		case 0:
		MainMenu(db)
		case 1:
		    CreateUser(db)
		case 2:
			GetUserById(db)
		case 3:
			Updateuser(db)
		case 4:
             Deleteuser(db)
	    case 5:
			GetAllUser(db)
		default:
		fmt.Println("Invalid option")
		}
}

func CreateUser(db*gorm.DB ){
	// var user model.User
	var firstName string
	var middleName string
	var lastName string
	var designation string
	fmt.Println("Enter Your First Name")
	fmt.Scan(&firstName)
	fmt.Println("Enter Your Middle Name")
	fmt.Scan(&middleName)
	fmt.Println("Enter Your Last Name")
	fmt.Scan(&lastName)
	fmt.Println("Enter Your designation ")
	fmt.Scan(&designation)
	user:= model.User{FirstName: firstName,MiddleName :middleName,LastName : lastName,Designation:designation }
	db.Create(&user)

}
func GetUserById (db*gorm.DB){
	var id int
	fmt.Println("Enter your id:")
	fmt.Scan(&id)
	var user = []model.User{}
    db.Where("id = ?", id).Find(&user)
	fmt.Println("id\t", "firstname\t","middlename\t","LastName\t","designation\t") 
	for _, user:= range user { 
    		fmt.Println(user.ID,"\t",user.FirstName,"\t",user.MiddleName,"\t\t",user.LastName,"\t",user.Designation) 
			fmt.Println("---------------------------------------------------------------------------------------")
}
}
func Updateuser(db*gorm.DB){
var id int
var firstName string
var middleName string
var lastName string
var designation string
fmt.Println("Enter an Id")
fmt.Scan(&id)
fmt.Println("Enter Updated First Name")
fmt.Scan(&firstName)
fmt.Println("Enter Updatd Middle Name")
fmt.Scan(&middleName)
fmt.Println("Enter Updated Last Name")
fmt.Scan(&lastName)
fmt.Println("Enter New Designation")
fmt.Scan(&designation)
updateQuery :=fmt.Sprintf("UPDATE users SET first_name ='%s',middle_name = '%s',last_name ='%s',designation ='%s'where id = %d",firstName,middleName,lastName,designation,id)
db.Exec(updateQuery)
}

func Deleteuser(db*gorm.DB){
	var id int
	fmt.Println("Enter an Id")
	fmt.Scan(&id)
	query := fmt.Sprintf("DELETE from users Where id = %d",id)
	db.Exec(query)
}
func GetAllUser(db*gorm.DB){
  var id int
  var user[]model.User
  db.Where("id = ?",id).Find(&user)
  fmt.Println("id\t", "firstname\t","middlename\t","LastName\t","designation\t") 
  query := "SELECT * FROm users";
  db.Raw(query).Scan(&user)
	for _, user:= range user { 
		fmt.Println(user.ID,"\t",user.FirstName,"\t",user.MiddleName,"\t\t",user.LastName,"\t\t",user.Designation) 
  }

}