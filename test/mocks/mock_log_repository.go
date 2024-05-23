package mocks

import (
	"errors"
	"product-app-go/internal/domain/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockLogRepository struct {
	Logs       []model.LoginLog
	DeletedIds []primitive.ObjectID
}

func NewMockLogRepository() *MockLogRepository {
	return &MockLogRepository{
		Logs:       make([]model.LoginLog, 0),
		DeletedIds: make([]primitive.ObjectID, 0),
	}
}

func (m *MockLogRepository) AddLogToMongo(log model.LoginLog) error {
	m.Logs = append(m.Logs, log)
	return nil
}

func (m *MockLogRepository) GetLogByLogId(logId primitive.ObjectID) (model.LoginLog, error) {
	for _, log := range m.Logs {
		if log.Id == logId {
			return log, nil
		}
	}
	return model.LoginLog{}, errors.New("log not found")
}

func (m *MockLogRepository) GetAllLogs() ([]model.LoginLog, error) {
	return m.Logs, nil
}

func (m *MockLogRepository) GetLogsByUserId(userId int) ([]model.LoginLog, error) {
	var logs []model.LoginLog
	for _, log := range m.Logs {
		if log.UserId == userId {
			logs = append(logs, log)
		}
	}
	return logs, nil
}

func (m *MockLogRepository) DeleteLogByLogId(logId primitive.ObjectID) error {
	for i, log := range m.Logs {
		if log.Id == logId {
			m.Logs = append(m.Logs[:i], m.Logs[i+1:]...)
			m.DeletedIds = append(m.DeletedIds, logId)
			return nil
		}
	}
	return errors.New("log not found")
}

func (m *MockLogRepository) DeleteAllLogs() error {
	m.Logs = make([]model.LoginLog, 0)
	return nil
}
