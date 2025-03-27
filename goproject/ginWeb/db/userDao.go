package db

import (
	"fmt"
	"ginWeb/config"
	"ginWeb/model"

	"gorm.io/gorm"
)

func (handler *dbHandler) CreateUser(user *model.User) error {
	// 自动事务
	return handler.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(user)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}

func (handler *dbHandler) CreateUsers(users []*model.User) error {
	return handler.db.Transaction(func(tx *gorm.DB) error {
		result := tx.CreateInBatches(users, config.Cfg.DB.BatchSize)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}

func (handler *dbHandler) DeleteUsers(conditions []Condition) error {
	return handler.db.Transaction(func(tx *gorm.DB) error {
		query := tx.Model(&model.User{})
		for _, condition := range conditions {
			query = query.Where(fmt.Sprintf("%s %s ?", condition.Field, condition.Operator), condition.Value)
		}
		result := query.Delete(&model.User{})
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}

func (handler *dbHandler) UpdateUsers(updates map[string]interface{}, conditions []Condition) error {
	return handler.db.Transaction(func(tx *gorm.DB) error {
		query := tx.Model(&model.User{})
		for _, condition := range conditions {
			query = query.Where(fmt.Sprintf("%s %s ?", condition.Field, condition.Operator), condition.Value)
		}
		result := query.Updates(updates)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}

func (handler *dbHandler) GetUsers(conditions []Condition) ([]*model.User, error) {
	var users []*model.User = []*model.User{}
	err := handler.db.Transaction(func(tx *gorm.DB) error {
		query := tx.Model(&model.User{})
		for _, condition := range conditions {
			query = query.Where(fmt.Sprintf("%s %s ?", condition.Field, condition.Operator), condition.Value)
		}
		result := query.Find(&users)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return users, nil
}
