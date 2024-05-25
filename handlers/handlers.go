package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/baselrabia/echo-task/pkg/route"
)

func ReconstructRoute(c echo.Context) error {
	var tickets []route.Ticket
	if err := c.Bind(&tickets); err != nil {
		return err
	}

	reconstructedRoute := route.ReconstructRoute(tickets)
	return c.JSON(http.StatusOK, reconstructedRoute)
}
