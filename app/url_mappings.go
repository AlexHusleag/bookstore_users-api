package app

import (
	"github.com/AlexHusleag/bookstore_users-api/controllers/ping"
	"github.com/AlexHusleag/bookstore_users-api/controllers/users"
)

func mapURLs() {
	router.GET("/ping", controllers.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.DELETE("/users/:user_id", users.DeleteUser)
	router.GET("/internal/users/search/", users.Search)
}
