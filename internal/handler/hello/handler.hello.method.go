package hello

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) GetHelloMessageHandler(e echo.Context) error {
	msg, err := h.helloUsecase.GetHelloMessageUsecase()
	if err != nil {
		return e.JSON(err.StatusCode, err.Error.Error())
	}
	return e.JSON(http.StatusOK, map[string]any{
		"message": msg,
	})
}
