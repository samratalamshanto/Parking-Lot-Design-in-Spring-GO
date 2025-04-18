package global_exception_handler

import (
	"Parking_Slot_in_GO/pkg/utitliy"
	"github.com/gin-gonic/gin"
	"log"
)

func GlobalExceptionHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		errs := ctx.Errors
		if len(errs) > 0 {
			err := errs.Last()
			utitliy.ServerErrResponse(ctx, "Something went wrong. Please try again later.")
			log.Printf("GlobalExceptionHandler() Error Occured=%v\n", err)
			ctx.Abort()
		}
	}
}
