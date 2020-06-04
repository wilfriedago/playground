package server

import (
	"github.com/techx/playground/controllers"
	"github.com/techx/playground/socket"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func newRouter(hub *socket.Hub) *echo.Echo {
	e := echo.New()

	// Define middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Jukebox controller
	jukebox := new(controllers.JukeboxController).Init(hub)
	e.POST("/jukebox/songs", jukebox.QueueSong)

	// Rooms controller
	room := new(controllers.RoomController)
	e.GET("/rooms", room.GetRooms)
	e.POST("/rooms", room.CreateRoom)

	return e
}
