package main 

import(
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/qor/qor"
	"github.com/qor/admin"
)

type User struct {
	gorm.Model
	Name string
}

type Product struct{
	gorm.Model
	Name string
	Description string
}


func main(){
	DB, _ := gorm.Open("mysql", "root:gauva@123@tcp(localhost:3306)/gowebapp?parseTime=true")

	DB.AutoMigrate(&User{}, &Product{})

	//Initalize
	Admin := admin.New(&admin.AdminConfig{DB: DB})

	//Allow to use Admin to manage User, Product 
	Admin.AddResource(&User{})
	Admin.AddResource(&Product{})

	//Initalize an HTTP request multiplexer
	mux := http.NewServeMux()

	//Mount Admin interface to mux
	Admin.MountTo("/admin", mux)

	fmt.Println("Listening on : 8080")
	http.ListenAndServe(":8080", mux)
}