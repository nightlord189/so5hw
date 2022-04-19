package model

type CreditCardInfo struct {
	Number string
	Holder string
	Date   string
	CVV    string
}

func (c CreditCardInfo) IsFilled() bool {
	return c.Date != "" && c.Holder != "" && c.CVV != "" && c.Number != ""
}

type CustomerDB struct {
	ID              int            `json:"id" gorm:"column:id;primaryKey;autoIncrement;"`
	Email           string         `json:"email" gorm:"column:email"`
	PasswordHash    string         `json:"-" gorm:"column:password_hash"`
	BillingAddress  string         `json:"billingAddress" gorm:"column:billing_address"`
	CreditCardRaw   string         `json:"-" gorm:"column:credit_card"`
	CreditCard      CreditCardInfo `json:"creditCard" gorm:"-"`
	ShippingAddress string         `json:"shippingAddress" gorm:"column:shipping_address"`
}

func (c CustomerDB) IsFilled() bool {
	return c.Email != "" && c.BillingAddress != "" && c.ShippingAddress != "" && c.CreditCard.IsFilled()
}

func (CustomerDB) TableName() string {
	return "customer"
}
