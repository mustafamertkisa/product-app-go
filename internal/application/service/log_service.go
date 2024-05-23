package service

import (
	"product-app-go/internal/domain/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogService interface {
	GetLogByLogId(logId primitive.ObjectID) (model.LoginLog, error)
	GetAllLogs() ([]model.LoginLog, error)
	GetLogsByUserId(userId int) ([]model.LoginLog, error)
	DeleteLogById(logId primitive.ObjectID) error
	DeleteAllLogs() error
}
