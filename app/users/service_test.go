package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	service *UserService
}

func (suite *UserServiceTestSuite) SetupTest() {
	repo := new(MockUserRepository)
	suite.service = new(UserService)
	suite.service.userRepository = repo
}

func (suite *UserServiceTestSuite) TestCreateUser() {
	mockRepo := &MockUserRepository{}
	service := &UserService{
		userRepository: mockRepo,
	}

	mockRepo.On("Create", mock.Anything).Return(User{}, nil)

	user, err := service.CreateUser(CreateUserDto{})

	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), user)
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
