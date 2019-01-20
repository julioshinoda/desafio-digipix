package app

import(
	"github.com/gin-gonic/gin"
	"github.com/desafio-digipix/app/routes"
	"github.com/joho/godotenv"
	"log"
)

func Start(){
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	routes.Shipments(r)

	r.Run() 
}