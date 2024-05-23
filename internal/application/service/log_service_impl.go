package service

import (
	"errors"
	"product-app-go/internal/domain/model"
	"product-app-go/internal/domain/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogServiceImpl struct {
	LogRepository repository.LogRepository
}

func NewLogServiceImpl(logRepository repository.LogRepository) LogService {
	return &LogServiceImpl{
		LogRepository: logRepository,
	}
}

func (s *LogServiceImpl) GetLogByLogId(logId primitive.ObjectID) (model.LoginLog, error) {
	log, err := s.LogRepository.GetLogByLogId(logId)
	if err != nil {
		return model.LoginLog{}, err
	}

	return log, nil
}

func (s *LogServiceImpl) GetAllLogs() ([]model.LoginLog, error) {
	logs, err := s.LogRepository.GetAllLogs()
	if err != nil {
		return nil, err
	}

	if len(logs) == 0 {
		return nil, errors.New("no logs found")
	}

	return logs, nil
}

func (s *LogServiceImpl) GetLogsByUserId(userId int) ([]model.LoginLog, error) {
	if userId <= 0 {
		return nil, errors.New("invalid user ID")
	}

	logs, err := s.LogRepository.GetLogsByUserId(userId)
	if err != nil {
		return nil, err
	}

	if len(logs) == 0 {
		return nil, errors.New("no logs found for the given user ID")
	}

	return logs, nil
}

func (s *LogServiceImpl) DeleteLogById(logId primitive.ObjectID) error {
	if logId.IsZero() {
		return errors.New("invalid log ID")
	}

	err := s.LogRepository.DeleteLogByLogId(logId)
	if err != nil {
		return err
	}

	return nil
}

func (s *LogServiceImpl) DeleteAllLogs() error {
	err := s.LogRepository.DeleteAllLogs()
	if err != nil {
		return err
	}

	return nil
}
