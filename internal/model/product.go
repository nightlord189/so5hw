package model

const ProductStatusActive ProductStatus = "active"
const ProductStatusInactive ProductStatus = "inactive"

type ProductStatus string

type ProductDB struct {
	ID                      int           `json:"id" gorm:"column:id;primaryKey;autoIncrement;"`
	Articul                 string        `json:"articul" gorm:"column:articul;unique"`
	Price                   float64       `json:"price" gorm:"column:price"`
	DeliveryTimeDescription string        `json:"deliveryTimeDescription" gorm:"column:delivery_time_description"`
	Status                  ProductStatus `json:"status" gorm:"column:status"`
	Inventory               int           `json:"inventory" gorm:"column:inventory"`
	Vendor                  string        `json:"vendor" gorm:"column:vendor"`
	Category                string        `json:"category" gorm:"column:category"`
	Images                  [][]byte      `json:"images" gorm:"-"`
}

func (ProductDB) TableName() string {
	return "product"
}

type CreateProductRequest struct {
	Articul                 string  `binding:"required"`
	Price                   float64 `binding:"required"`
	DeliveryTimeDescription string
	Status                  ProductStatus `binding:"required"`
	Category                string        `binding:"required"`
	Inventory               int           `binding:"required"`
	Vendor                  string        `binding:"required"`
	Images                  [][]byte
}

func (c *CreateProductRequest) ToDbModels() (ProductDB, []ImageDB) {
	product := ProductDB{
		Articul:                 c.Articul,
		Price:                   c.Price,
		DeliveryTimeDescription: c.DeliveryTimeDescription,
		Status:                  c.Status,
		Category:                c.Category,
		Inventory:               c.Inventory,
		Vendor:                  c.Vendor,
	}
	images := make([]ImageDB, len(c.Images))
	for i := range c.Images {
		images[i] = ImageDB{
			Data: c.Images[i],
		}
	}
	return product, images
}

type GetProductsRequest struct {
	ID       string        `json:"id" form:"id"`
	Articul  string        `json:"articul" form:"articul"`
	Category string        `json:"category" form:"category"`
	Status   ProductStatus `json:"status" form:"status"`
	Vendor   string        `json:"vendor" form:"vendor"`
	Limit    int           `json:"limit" form:"limit"`
	Page     int           `json:"page" form:"page"`
}

type GetProductsResponse struct {
	Records      []ProductDB `json:"records"`
	RecordsCount int         `json:"recordsCount"`
	CurrentPage  int         `json:"currentPage"`
	PagesCount   int         `json:"pagesCount"`
}
