package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CouponDetail struct {
	ID             primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name           string             `bson:"name" json:"name"`
	StartDate      time.Time          `bson:"start_date" json:"start_date"`
	EndDate        time.Time          `bson:"end_date" json:"end_date"`
	Status         int                `bson:"status" json:"status"` //open-close
	CouponCount    int                `bson:"coupon_count" json:"coupon_count"`
	CouponType     int                `bson:"coupon_type" json:"coupon_type"` //1 one 2 many
	Reward         float64            `bson:"reward" json:"reward"`
	WithdrawFix    float64            `bson:"withdraw_fix" json:"withdraw_fix"`
	FixWithdraw    float64            `bson:"fix_withdraw" json:"fix_withdraw"`
	CreatedSuccess int                `bson:"created_success" json:"created_success"`
}

type CouponCode struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Status   int                `bson:"status" json:"status"` // use-unused
	Total    int                `bson:"total" json:"total"`   // ทั้งหมด
	Count    int                `bson:"count" json:"count"`   // คงเหลือ
	Code     string             `bson:"code" json:"code"`
	ConfigID primitive.ObjectID `bson:"config_id" json:"config_id"`
	UserUsed []string           `bson:"username_used" json:"username_used"`
}
