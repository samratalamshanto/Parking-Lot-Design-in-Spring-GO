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

//func SuccessResponse(data interface{}, msg string) *models.Response {
//	response := models.Response{
//		Code:      200,
//		Message:   msg,
//		Data:      data,
//		IsSuccess: true,
//	}
//	return &response
//}
//
//func CreatedResponse(data interface{}, msg string) *models.Response {
//	response := models.Response{
//		Code:      201,
//		Message:   msg,
//		Data:      data,
//		IsSuccess: true,
//	}
//	return &response
//}
//
//func ClientErrResponse(data interface{}, msg string) *models.Response {
//	response := models.Response{
//		Code:      400,
//		Message:   msg,
//		Data:      data,
//		IsSuccess: false,
//	}
//	return &response
//}
//
//func ServerErrResponse(data interface{}, msg string) *models.Response {
//	response := models.Response{
//		Code:      500,
//		Message:   msg,
//		Data:      data,
//		IsSuccess: false,
//	}
//	return &response
//}
//
//func UnauthorizedErrResponse(msg string) *models.Response {
//	response := models.Response{
//		Code:      401,
//		Message:   msg,
//		Data:      nil,
//		IsSuccess: false,
//	}
//	return &response
//}
