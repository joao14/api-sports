package main

import (
	"github.com/joao14/api-sports.git/routes"
)

var err error

func main() {
	/*config.DB, err = gorm.Open("postgres", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer config.DB.Close()

	config.DB.AutoMigrate(&model.Player{})*/

	r := routes.SetupRouter()

	r.Run()
}
