package utitliy

import (
	"Parking_Slot_in_GO/base/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func SuccessResponse(ctx *gin.Context, data interface{}, msg string) {
	response := models.Response{
		Code:      200,
		Message:   msg,
		Data:      data,
		IsSuccess: true,
	}
	ctx.JSON(http.StatusOK, response)
}

func CreatedResponse(ctx *gin.Context, data interface{}, msg string) {
	response := models.Response{
		Code:      201,
		Message:   msg,
		Data:      data,
		IsSuccess: true,
	}
	ctx.JSON(http.StatusCreated, response)
}

func ClientErrResponse(ctx *gin.Context, msg string) {
	response := models.Response{
		Code:      400,
		Message:   msg,
		Data:      nil,
		IsSuccess: false,
	}
	ctx.JSON(http.StatusBadRequest, response)
}

func ServerErrResponse(ctx *gin.Context, msg string) {
	response := models.Response{
		Code:      500,
		Message:   msg,
		Data:      nil,
		IsSuccess: false,
	}
	ctx.JSON(http.StatusInternalServerError, response)
}

func UnauthorizedErrResponse(ctx *gin.Context, msg string) {
	response := models.Response{
		Code:      401,
		Message:   msg,
		Data:      nil,
		IsSuccess: false,
	}
	ctx.JSON(http.StatusUnauthorized, response)
}

func GetJsonString(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

func GetCurTimeStamp() time.Time {
	return time.Now()
}
func GetUUID() gocql.UUID {
	return gocql.UUID(uuid.New())
}

// ErrorCheck() is a helper function that takes a value of any type and an error.
// If the error is nil, it returns the value; if the error is non-nil, it panics.
func ErrorCheck[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
