package helperstructs

import "time"

type ProductReq struct {
	ID              uint                   `json:"id"`
	Category        uint                   `json:"category_id"`
	Brand           uint                   `json:"brand"`
	Name            string                 `json:"name"`
	Description     string                 `json:"description"`
	Price           int                    `json:"price"`
	Stock           int                    `json:"stock"`
	Specifications  map[string]interface{} `json:"specifications"`
	IsActive        bool                   `json:"is_active"`
	RelatedProducts []uint                 `json:"related_products"`
	UpdatedBy       string                 `json:"updated_by"`
}

type BrandReq struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type DiscountReq struct {
	ID       uint      `json:"id"`
	Category uint      `json:"category_id"`
	Discount float32   `json:"discount"`
	EndDate  time.Time `json:"end_date"`
}

type StockReq struct {
	ID    uint `json:"id"`
	Stock int  `json:"stock"`
}

type CouponReq struct {
	ID                   uint `json:"id"`
	Code                 int
	OFF                  int    `json:"off"`
	GiveOnPurchaseAbove  int    `json:"give_onpurchase_above"`
	ApplyOnPurchaseAbove int    `json:"apply_onpurchase_above"`
	IsWelcome            bool   `json:"is_welcome"`
	Description          string `json:"description"`
}

type AddCouponforUser struct {
	Email  string `json:"email"`
	Coupon uint   `json:"coupon"`
}
