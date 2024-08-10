package usecase

import (
	"github.com/taku10101/internal/domain"
	"github.com/taku10101/internal/repository"
)

type UserRepository interface {
    Create(user *repository.User) error
    Update(user *repository.User) error
    Delete(userID string) error
    FindByID(userID string) (*repository.User, error)
    FindAll() ([]repository.User, error)
}

type UserUseCase struct {
    userRepository UserRepository
}

func NewUserUseCase(userRepository UserRepository) *UserUseCase {
    return &UserUseCase{userRepository: userRepository}
}


func (u *UserUseCase) CreateUser(user *domain.CreateUserRequest) error {
    dbUser := &repository.User{
        ID:     user.ID,
        Name:   user.Name,
        Icon:   user.Icon,
        Age:    user.Age,
        Gender: user.Gender,
    }
    return u.userRepository.Create(dbUser)
}

func (u *UserUseCase) UpdateUser(user *domain.UpdateUserRequest) error {
    dbUser := &repository.User{
        ID:     user.ID,
        Name:   user.Name,
        Icon:   user.Icon,
        Age:    user.Age,
    }
    return u.userRepository.Update(dbUser)
}
func (u *UserUseCase) DeleteUser(userID string) error {
    return u.userRepository.Delete(userID)
}

func (u *UserUseCase) GetUserByID(userID string) (*repository.User, error) {
    return u.userRepository.FindByID(userID)
}

func (u *UserUseCase) GetAllUsers() ([]repository.User, error) {
    return u.userRepository.FindAll()
}
