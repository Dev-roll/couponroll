package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	UserCoupon struct {
		ID        uuid.UUID `json:"id,omitempty" db:"id"`
		UserID    uuid.UUID `json:"user_id,omitempty" db:"user_id"`
		CouponID  uuid.UUID `json:"coupon_id,omitempty" db:"coupon_id"`
		IsUsed    bool      `json:"is_used,omitempty" db:"is_used"`
		CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
		UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
	}
)

func CreateUserCoupon(uc UserCoupon) {
	v, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	if _, err := db.Exec("INSERT INTO user_coupons (id, user_id, coupon_id, is_used) VALUES (?, ?, ?, ?)", v, uc.UserID, uc.CouponID, uc.IsUsed); err != nil {
		panic(err)
	}
}

func GetUserCouponFromID(uc UserCoupon) UserCoupon {
	var userCoupon UserCoupon
	if err := db.Get(&userCoupon, "SELECT * FROM user_coupons WHERE id=?", uc.ID); err != nil {
		panic(err)
	}
	return userCoupon
}

func UpdateUserCoupon(userCouponID uuid.UUID, uc UserCoupon) {
	if _, err := db.Exec("UPDATE user_coupons SET is_used=? WHERE id=?", uc.IsUsed, userCouponID); err != nil {
		panic(err)
	}
}

func DeleteUserCoupon(uc UserCoupon) {
	if _, err := db.Exec("DELETE FROM user_coupons WHERE id=?", uc.ID); err != nil {
		panic(err)
	}
}

func GetUserCouponsFromUserID(userID uuid.UUID) []Coupon {
	var coupons []Coupon
	// TODO@notchcoder: Get coupons of the authenticated user

	return coupons
}

func GetUserCouponsFromUserIDByStore(userID uuid.UUID, storeID uuid.UUID) []Coupon {
	var coupons []Coupon
	// TODO@notchcoder: Get coupons for store determined by store_id of the authenticated user

	return coupons
}
