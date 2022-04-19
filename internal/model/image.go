package model

type ImageDB struct {
	ID        string `json:"id" gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4();"`
	ProductID int    `json:"productId" gorm:"column:product_id"`
	Data      []byte `json:"data" gorm:"-"`
}

func (ImageDB) TableName() string {
	return "image"
}
