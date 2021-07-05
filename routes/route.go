package routes

import (
	"github.com/IshanXXIV/GetGroundTask/controllers"

	"github.com/labstack/echo"
)

//endpoints to make all adjustments in the tables accoding to the guests

func MapUrls() {

	router := echo.New()
	// preparation for the party
	router.POST("/guest_list/:name", controllers.AddGuest)
	router.GET("/guest_list", controllers.getGuestList)

	// //after guests arrive
	router.PUT("/guests/:name", controllers.editGuestList)
	router.DELETE("/guests/:name", controllers.deleteGuests)
	router.GET("/seats_empty", controllers.GetUser)
	router.GET("/users/:user_id", controllers.GetUser)

}
