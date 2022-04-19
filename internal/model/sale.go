package model

type SaleRequest struct {
	ProductID  int `binding:"required"`
	CustomerID int `binding:"required"`
	Quantity   int `binding:"required"`
}
