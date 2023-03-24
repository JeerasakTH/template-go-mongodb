package controller

import (
	"log"
	"net/http"
	model "template-go-mongodb/model"
	repo "template-go-mongodb/repository"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCoupon() func(c *gin.Context) {
	return func(c *gin.Context) {
		coupon := model.CouponDetail{}
		if err := c.BindJSON(&coupon); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			log.Fatal(err)
			return
		}

		payloadDetail := model.CouponDetail{
			ID:          primitive.NewObjectID(),
			Name:        coupon.Name,
			CouponCount: 3,
		}

		objCode := []interface{}{}

		for i := 1; i <= payloadDetail.CouponCount; i++ {
			couponCode := model.CouponCode{
				ID:     primitive.NewObjectID(),
				Status: 1,
				Total:  1,
				Count:  1,
			}
			objCode = append(objCode, couponCode)
		}

		if _, err := repo.CreateOne("coupon_detail", payloadDetail); err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		if _, err := repo.CreateMany("coupon_code", objCode); err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Posted successfully", "Data": payloadDetail})
	}
}

func GetCoupon() func(c *gin.Context) {
	return func(c *gin.Context) {
		couponID := c.Query("id")
		objId, _ := primitive.ObjectIDFromHex(couponID)

		var coupon model.CouponDetail
		filter := bson.M{"_id": objId}
		filterOp := bson.M{"coupon_count": -1}
		if err := repo.GetOne("coupon_detail", filter, filterOp, &coupon); err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "success!", "Data": coupon})
	}
}

func GetmanyCoupon() func(c *gin.Context) {
	return func(c *gin.Context) {
		var coupon []model.CouponDetail

		filter := bson.M{"status": 1}
		if err := repo.GetManyLimitOne("coupon_detail", filter, nil, &coupon); err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "success!", "Data": coupon})
	}
}

func UpdateCoupon() func(c *gin.Context) {
	return func(c *gin.Context) {
		couponID := c.Query("id")
		objId, _ := primitive.ObjectIDFromHex(couponID)

		var coupon model.CouponDetail

		if err := c.BindJSON(&coupon); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		filter := bson.M{"_id": objId}
		set := bson.M{
			"$set": bson.M{"status": coupon.Status, "reward": coupon.Reward},
		}
		if _, err := repo.UpdateOne("coupon_detail", filter, set); err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "data updated successfully!", "Data": set})
	}
}

func UpdateManyCoupon() func(c *gin.Context) {
	return func(c *gin.Context) {
		var coupon model.CouponDetail

		if err := c.BindJSON(&coupon); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		filter := bson.M{"status": 1}
		set := bson.M{
			"$set": bson.M{"status": 0},
		}

		if err := repo.UpdateMany("coupon_detail", filter, set); err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "data updated successfully!", "Data": set})
	}
}

func DeleteCoupon() func(c *gin.Context) {
	return func(c *gin.Context) {
		couponID := c.Query("id")
		objId, _ := primitive.ObjectIDFromHex(couponID)

		filter := bson.M{"_id": objId}
		if _, err := repo.DeleteOne("coupon_detail", filter); err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Article deleted successfully", "Data": nil})
	}
}
