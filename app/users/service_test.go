package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService(t *testing.T) {
	mockRepo := &MockUserRepository{}
	service := &UserService{
		userRepository: mockRepo,
	}

	mockRepo.On("Create", mock.Anything).Return(User{}, nil)

	user, err := service.CreateUser(CreateUserDto{})

	assert.Nil(t, err)
	assert.NotNil(t, user)
}
