package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sajjad3k/ginbankapiex/controllers"
)

func Setroutes() *gin.Engine {

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/customers", controllers.Getallcustomers)
		api.GET("/customers/:id", controllers.Getcustomerbyid)
		api.POST("/customers", controllers.Createcustomer)
		api.PUT("/customers/:id", controllers.Updatecustomer)
		api.DELETE("/customers/:id", controllers.Deletecustomer)
		api.PATCH("/customers/:id/:amount/", controllers.Updatebalance)
		api.PUT("/customers/:id/:toid/:amount/", controllers.Transfermoney)
		api.GET("/customers/:id/checkid", controllers.Checkidavailable)
	}

	r.NoRoute(func(c *gin.Context) { c.AbortWithStatus(http.StatusNotFound) })

	return r
}
