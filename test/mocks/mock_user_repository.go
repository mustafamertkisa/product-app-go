package mocks

import (
	"errors"
	"product-app-go/internal/domain/model"
)

type MockUserRepository struct {
	users  map[int]model.User
	nextID int
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users:  make(map[int]model.User),
		nextID: 1,
	}
}

func (m *MockUserRepository) Save(user model.User) error {
	if user.Email == "error@example.com" {
		return errors.New("forced save error")
	}
	user.Id = m.nextID
	m.users[m.nextID] = user
	m.nextID++
	return nil
}

func (m *MockUserRepository) FindByEmail(email string) (model.User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}
	return model.User{}, nil
}

func (m *MockUserRepository) FindById(id int) (model.User, error) {
	user, exists := m.users[id]
	if !exists {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}

func (m *MockUserRepository) FindAll() ([]model.User, error) {
	users := make([]model.User, 0, len(m.users))
	for _, user := range m.users {
		users = append(users, user)
	}
	return users, nil
}

func (m *MockUserRepository) Update(user model.User) error {
	if _, exists := m.users[user.Id]; !exists {
		return errors.New("user not found")
	}
	m.users[user.Id] = user
	return nil
}

func (m *MockUserRepository) Delete(id int) error {
	if _, exists := m.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(m.users, id)
	return nil
}

func (m *MockUserRepository) AddLogToMongo(log model.LoginLog) error {
	// Dummy implementation, as it's not essential for testing UserServiceImpl
	return nil
}
