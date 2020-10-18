package app

import "github.com/balsuvendukumar/bookstore_users-api/controller/bookuser_controller"

func MapUrls() {
	router.GET("/ping", bookuser_controller.Pong)
	router.POST("/Users", bookuser_controller.CreateUser)
	router.GET("/users/:user_id", bookuser_controller.GetUser)
}
