package db

import (
	"fmt"
	"github.com/nightlord189/gormery"
	"github.com/nightlord189/so5hw/internal/model"
	"math"
)

func (d *Manager) CreateProduct(req *model.CreateProductRequest) (model.ProductDB, error) {
	product, images := req.ToDbModels()
	tx := d.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return product, fmt.Errorf("error tx: %v", err)
	}
	err := tx.Create(&product).Error
	if err != nil {
		tx.Rollback()
		return product, fmt.Errorf("err create product: %w", err)
	}
	for i := range images {
		images[i].ProductID = product.ID
	}
	err = tx.Create(&images).Error
	if err != nil {
		tx.Rollback()
		return product, fmt.Errorf("err create images: %w", err)
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return product, fmt.Errorf("err commit transaction: %w", err)
	}
	product.Images = req.Images
	return product, nil
}

func (d *Manager) GetCategories() ([]string, error) {
	entities := make([]string, 0)
	err := d.DB.Raw("SELECT distinct category FROM product WHERE status='active' ORDER BY category").Scan(&entities).Error
	return entities, err
}

func (d *Manager) GetProducts(req *model.GetProductsRequest) (model.GetProductsResponse, error) {
	entities := make([]model.ProductDB, 0)
	queryElems := make([]gormery.ConditionElement, 0)
	result := model.GetProductsResponse{CurrentPage: req.Page}
	if req.ID != "" {
		queryElems = append(queryElems, gormery.Equal("id", req.ID))
	}
	if req.Articul != "" {
		queryElems = append(queryElems, gormery.Equal("articul", req.Articul))
	}
	if req.Vendor != "" {
		queryElems = append(queryElems, gormery.Equal("vendor", req.Vendor))
	}
	if req.Status != "" {
		queryElems = append(queryElems, gormery.Equal("status", req.Status))
	}
	if req.Category != "" {
		queryElems = append(queryElems, gormery.Equal("category", req.Category))
	}

	sql, elems := gormery.CombineSimpleQuery(queryElems, "AND")
	query := d.DB.Where(sql, elems...).Order("id")

	// pagination
	if req.Page >= 1 && req.Limit > 0 {
		// counting...
		var count int64
		err := query.Model(&model.ProductDB{}).Count(&count).Error
		if err != nil {
			return result, err
		}
		result.RecordsCount = int(count)
		result.PagesCount = int(math.Ceil(float64(count) / float64(req.Limit)))

		// select by pages
		query = query.Limit(req.Limit)
		offset := req.Limit * (req.Page - 1)
		query = query.Offset(offset)
	}
	err := query.Find(&entities).Error
	result.Records = entities
	return result, err
}
