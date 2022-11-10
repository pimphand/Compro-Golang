package main

import (
	"compro/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "sammy:password@tcp(127.0.0.1:3306)/madjou?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterInput{}
	userInput.Name = "faisal"
	userInput.Email = "fssf@gmail.com"
	userInput.Password = "password"

	userService.RegisterUser(userInput)

	// user := user.User{
	// 	Name:            "Faisal",
	// 	Email:           "Faisal@gmail.com",
	// 	Password:        "Faisal@gmail.com",
	// 	Status:          true,
	// 	EmailVerifiedAt: time.Now(),
	// }
	// userRepository.Save(user)

}
