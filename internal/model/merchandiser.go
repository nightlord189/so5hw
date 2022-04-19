package model

type MerchandiserDB struct {
	ID           int    `json:"id" gorm:"column:id;primaryKey;autoIncrement;"`
	Username     string `json:"username" gorm:"column:username"`
	PasswordHash string `json:"-" gorm:"column:password_hash"`
}

func (MerchandiserDB) TableName() string {
	return "merchandiser"
}
