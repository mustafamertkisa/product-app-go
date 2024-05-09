package repository

import (
	"context"
	"errors"
	"product-app-go/internal/application/command"
	"product-app-go/internal/domain/model"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	PostgresDb *gorm.DB
	MongoDb    *mongo.Client
}

func NewUserRepositoryImpl(PostgresDb *gorm.DB, MongoDb *mongo.Client) UserRepository {
	return &UserRepositoryImpl{PostgresDb: PostgresDb, MongoDb: MongoDb}
}

func (r *UserRepositoryImpl) Save(user model.User) error {
	result := r.PostgresDb.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepositoryImpl) Update(user model.User) error {
	var updateUser = command.UpdateUserRequest{Id: int(user.Id), Name: user.Name, Email: user.Email}
	result := r.PostgresDb.Model(&model.User{}).Where("id = ?", user.Id).Updates(updateUser)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepositoryImpl) Delete(userId int) error {
	var user model.User
	result := r.PostgresDb.Where("id = ?", userId).Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepositoryImpl) FindAll() ([]model.User, error) {
	var user []model.User
	result := r.PostgresDb.Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *UserRepositoryImpl) FindById(userId int) (model.User, error) {
	var user model.User
	result := r.PostgresDb.Find(&user, userId)
	if result == nil {
		return user, errors.New("user is not found")
	}

	return user, nil
}

func (r *UserRepositoryImpl) FindByEmail(userEmail string) (model.User, error) {
	var user model.User
	result := r.PostgresDb.Where("email = ?", userEmail).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (r *UserRepositoryImpl) AddLogToMongo(log model.LoginLog) error {
	collection := r.MongoDb.Database("logs").Collection("login_logs")

	_, err := collection.InsertOne(context.TODO(), log)
	if err != nil {
		return err
	}

	return nil
}
