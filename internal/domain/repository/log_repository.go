package repository

import (
	"product-app-go/internal/domain/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogRepository interface {
	AddLogToMongo(log model.LoginLog) error
	GetLogByLogId(logId primitive.ObjectID) (model.LoginLog, error)
	GetAllLogs() ([]model.LoginLog, error)
	GetLogsByUserId(userId int) ([]model.LoginLog, error)
	DeleteLogByLogId(logId primitive.ObjectID) error
	DeleteAllLogs() error
}
