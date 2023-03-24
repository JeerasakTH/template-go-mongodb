package router

import (
	"template-go-mongodb/controller"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.POST("/coupon", controller.CreateCoupon())
	r.GET("/coupon", controller.GetCoupon())
	r.GET("/many_coupon", controller.GetmanyCoupon())
	r.POST("/update_coupon", controller.UpdateCoupon())
	r.POST("/updatemany_coupon", controller.UpdateManyCoupon())
	r.POST("/delete_coupon", controller.DeleteCoupon())
}
