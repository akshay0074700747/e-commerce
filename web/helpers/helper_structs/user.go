package helperstructs

type UserReq struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Otp      string `json:"otp"`
}

type AddressReq struct {
	ID            uint   `json:"id"`
	Email         string `json:"email"`
	HouseName     string `json:"house_name"`
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	District      string `json:"district"`
	PO            string `json:"po"`
	State         string `json:"state"`
}
