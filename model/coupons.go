package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	Coupon struct {
		ID          uuid.UUID `json:"id,omitempty" db:"id"`
		StoreID     uuid.UUID `json:"store_id,omitempty" db:"store_id"`
		Name        string    `json:"name,omitempty" db:"name"`
		Description string    `json:"description,omitempty" db:"description"`
		ImgUrl      string    `json:"img_url,omitempty" db:"img_url"`
		Discount    int       `json:"discount,omitempty" db:"discount"`
		CreatorID   uuid.UUID `json:"creator_id,omitempty" db:"creator_id"`
		CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
		UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at"`
		ExpiresAt   time.Time `json:"expires_at,omitempty" db:"expires_at"`
		// DeletedAt   time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	}
)

func CreateCoupon(c Coupon) {
	v, err := uuid.NewRandom()
	if err != nil {
		// TODO: handle error
		panic(err)
	}

	if _, err := db.Exec("INSERT INTO coupons (id, store_id, name, description, img_url, discount, creator_id, expires_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", v, c.StoreID, c.Name, c.Description, c.ImgUrl, c.Discount, c.CreatorID, c.ExpiresAt); err != nil {
		panic(err)
	}
}

func GetCoupons() []Coupon {
	var coupons []Coupon
	if err := db.Select(&coupons, "SELECT * FROM coupons"); err != nil {
		panic(err)
	}
	return coupons
}

func GetCouponFromID(c Coupon) Coupon {
	var coupon Coupon
	if err := db.Get(&coupon, "SELECT * FROM coupons WHERE id=?", c.ID); err != nil {
		panic(err)
	}
	return coupon
}

func UpdateCoupon(couponID uuid.UUID, c Coupon) {
	if _, err := db.Exec("UPDATE coupons SET name=?, description=?, img_url=?, discount=?, expires_at=? WHERE id=?", c.Name, c.Description, c.ImgUrl, c.Discount, c.ExpiresAt, couponID); err != nil {
		panic(err)
	}
}

func DeleteCoupon(c Coupon) {
	if _, err := db.Exec("DELETE FROM coupons WHERE id=?", c.ID); err != nil {
		panic(err)
	}
}
