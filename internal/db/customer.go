package db

import (
	"encoding/json"
	"github.com/nightlord189/so5hw/internal/model"
)

func (d *Manager) GetCustomer(id string) (model.CustomerDB, error) {
	var entity model.CustomerDB
	err := d.GetEntityByField("id", id, &entity)
	if err != nil {
		return entity, err
	}

	if entity.CreditCardRaw != "" {
		var creditCard model.CreditCardInfo
		err = json.Unmarshal([]byte(entity.CreditCardRaw), &creditCard)
		if err != nil {
			return entity, err
		}
		entity.CreditCard = creditCard
	}

	return entity, nil
}
