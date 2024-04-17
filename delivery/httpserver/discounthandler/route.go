package discounthandler

import (
	"github.com/labstack/echo/v4"
)

func (h Handler) SetUserRoutes(e *echo.Echo) {

	e.POST("apply_discount", h.applyDiscount)

}
