package discounthandler

import (
	"app/param/discountparam"
	"app/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) applyDiscount(c echo.Context) error {

	var req discountparam.ApplyDiscountRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	//if filedErrors, err := h.validation.ValidateApplyDiscountRequest(req); err != nil {
	//	msg, code := httpmsg.Error(err)
	//	return c.JSON(code, echo.Map{
	//		"message": msg,
	//		"errors":  filedErrors,
	//	})
	//	return echo.NewHTTPError(code, msg, filedErrors)
	//}

	resp, err := h.svc.ApplyDiscount(c.Request().Context(), req)
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}
	return c.JSON(http.StatusOK, resp)
}
