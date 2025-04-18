package ticket

import "github.com/gin-gonic/gin"

func SetupRouter(ginRouter *gin.Engine) {
	ticketRouter := ginRouter.Group("ticket")
	{
		ticketRouter.POST("/create-ticket", GetTicket)
		ticketRouter.POST("/checkout-ticket", CheckoutTicket)
	}
}
