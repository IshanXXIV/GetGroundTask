package routes

import (
	"github.com/IshanXXIV/GetGroundTask/models/controllers"
	"github.com/go-siris/siris/core/router"
)

//endpoints to make all adjustments in the tables accoding to the guests

func MapUrls() {

	//preparation for the party
	router.POST("/guest_list/:name", controllers.AddGuest())
	// router.GET("/guest_list", TableController.getGuestList)

	// //after guests arrive
	// router.PUT("/guests/:name", TableController.editGuestList)
	// router.DELETE("/guests/:name", TableController.deleteGuests)
	// router.GET("/seats_empty", TableController.GetUser)
	// router.GET("/users/:user_id", TableController.GetUser)

}
