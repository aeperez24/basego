package handler

import (
	"aeperez24/basego/dto"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(dto.ResponseDto{Data: payload})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithJSONGin(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, dto.ResponseDto{Data: payload})
}

type HandlerConfigGin struct {
	AuthenticationHandler GinAutenticationHandlerImpl
	UserHandler           GinUserHandlerImpl
}
