package db

import "github.com/nightlord189/so5hw/internal/model"

func (d *Manager) GetUserEntity(username string, userType model.UserType) (model.UserEntity, error) {
	result := model.UserEntity{
		Username: username,
	}
	if userType == model.UserTypeMerchandiser {
		var entity model.MerchandiserDB
		err := d.GetEntityByField("username", username, &entity)
		if err != nil {
			return result, err
		}
		result.ID = entity.ID
		result.PasswordHash = entity.PasswordHash
	} else {
		var entity model.CustomerDB
		err := d.GetEntityByField("email", username, &entity)
		if err != nil {
			return result, err
		}
		result.ID = entity.ID
		result.PasswordHash = entity.PasswordHash
	}
	return result, nil
}
