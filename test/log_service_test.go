package test

import (
	"testing"
	"time"

	"product-app-go/internal/application/service"
	"product-app-go/internal/domain/model"
	"product-app-go/test/mocks"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func initLogTestServices() (service.LogService, *mocks.MockLogRepository) {
	mockRepo := mocks.NewMockLogRepository()
	logService := service.NewLogServiceImpl(mockRepo)

	return logService, mockRepo
}

func TestLogServiceImpl_GetLogByLogId(t *testing.T) {
	logService, mockRepo := initLogTestServices()

	testLogID := primitive.NewObjectID()
	expectedLog := model.LoginLog{
		Id:        testLogID,
		Success:   true,
		Message:   "Login successful",
		UserId:    123,
		CreatedAt: time.Now(),
	}

	mockRepo.Logs = append(mockRepo.Logs, expectedLog)

	resultLog, err := logService.GetLogByLogId(testLogID)

	assert.Nil(t, err)
	assert.Equal(t, expectedLog, resultLog)
}

func TestLogServiceImpl_GetAllLogs(t *testing.T) {
	logService, mockRepo := initLogTestServices()

	expectedLogs := []model.LoginLog{
		{
			Id:        primitive.NewObjectID(),
			Success:   true,
			Message:   "Login successful",
			UserId:    123,
			CreatedAt: time.Now(),
		},
	}

	mockRepo.Logs = expectedLogs

	resultLogs, err := logService.GetAllLogs()

	assert.Nil(t, err)
	assert.Equal(t, expectedLogs, resultLogs)
}

func TestLogServiceImpl_GetLogsByUserId(t *testing.T) {
	logService, mockRepo := initLogTestServices()

	userID := 123
	expectedLogs := []model.LoginLog{
		{
			Id:        primitive.NewObjectID(),
			Success:   true,
			Message:   "Login successful",
			UserId:    userID,
			CreatedAt: time.Now(),
		},
	}

	mockRepo.Logs = expectedLogs

	resultLogs, err := logService.GetLogsByUserId(userID)

	assert.Nil(t, err)
	assert.Equal(t, expectedLogs, resultLogs)
}

func TestLogServiceImpl_DeleteLogById(t *testing.T) {
	logService, mockRepo := initLogTestServices()

	testLogID := primitive.NewObjectID()
	mockRepo.Logs = []model.LoginLog{
		{Id: testLogID},
	}

	err := logService.DeleteLogById(testLogID)

	assert.Nil(t, err)
	assert.Len(t, mockRepo.Logs, 0)
}

func TestLogServiceImpl_DeleteAllLogs(t *testing.T) {
	logService, mockRepo := initLogTestServices()

	mockRepo.Logs = []model.LoginLog{}

	err := logService.DeleteAllLogs()

	assert.Nil(t, err)
	assert.Empty(t, mockRepo.Logs)
}
