package usecase

import "github.com/taku10101/internal/domain"

type UserRepository interface {
    Create(user *domain.CreateUserRequest) error
    Update(user *domain.UpdateUserRequest) error
    Delete(userID string) error
    FindByID(userID string) (*domain.User, error)
    FindAll() ([]domain.User, error)
}

type UserUseCase struct {
    userRepository UserRepository
}

func NewUserUseCase(userRepository UserRepository) *UserUseCase {
    return &UserUseCase{userRepository: userRepository}
}

func (u *UserUseCase) CreateUser(user *domain.CreateUserRequest) error {
    return u.userRepository.Create(user)
}

func (u *UserUseCase) UpdateUser(user *domain.UpdateUserRequest) error {
    return u.userRepository.Update(user)
}

func (u *UserUseCase) DeleteUser(userID string) error {
    return u.userRepository.Delete(userID)
}

func (u *UserUseCase) GetUserByID(userID string) (*domain.User, error) {
    return u.userRepository.FindByID(userID)
}

func (u *UserUseCase) GetAllUsers() ([]domain.User, error) {
    return u.userRepository.FindAll()
}
