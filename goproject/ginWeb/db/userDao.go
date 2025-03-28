package db

import (
	"context"
	"fmt"
	"ginWeb/config"
	"ginWeb/model"
	"ginWeb/utils"

	"gorm.io/gorm"
)

func (handler *dbHandler) CreateUser(ctx context.Context, user *model.User) error {
	log := utils.LoggerFromContext(ctx, "logger")
	// 自动事务
	return handler.db.Transaction(func(tx *gorm.DB) error {
		log.Info("start create user with db")
		result := tx.Create(user)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}

func (handler *dbHandler) CreateUsers(ctx context.Context, users []*model.User) error {
	return handler.db.Transaction(func(tx *gorm.DB) error {
		result := tx.CreateInBatches(users, config.Cfg.DB.BatchSize)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}

func (handler *dbHandler) DeleteUsers(ctx context.Context, conditions []Condition) error {
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

func (handler *dbHandler) UpdateUsers(ctx context.Context, updates map[string]interface{}, conditions []Condition) error {
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

func (handler *dbHandler) GetUsers(ctx context.Context, conditions []Condition) ([]*model.User, error) {
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
