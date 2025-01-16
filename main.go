package main

import (
	"primeskill/config"
	"primeskill/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitPsqlDB()
	config.MigratePsqlDB()

	router := gin.Default()
	api := router.Group("/api/v0")
	/*
		API Endpoints
		For Testing on Postman: it could simplify by using the following:
		1. POST: http://localhost:8080/api/v0/comment/create
		2. POST: http://localhost:8080/api/v0/comment/create/bulk
		3. GET: http://localhost:8080/api/v0/comment
		4. GET: http://localhost:8080/api/v0/comment/1
		5. GET: http://localhost:8080/api/v0/comment/user/1
		6. DELETE: http://localhost:8080/api/v0/comment/delete/1
		7. DELETE: http://localhost:8080/api/v0/comment/delete/user/1
		8. DELETE: http://localhost:8080/api/v0/comment/delete/all
		9. DELETE: http://localhost:8080/api/v0/comment/delete/bulk
		
	*/
	routes.CommentRoute(api)

	router.Run(":8080")
}
