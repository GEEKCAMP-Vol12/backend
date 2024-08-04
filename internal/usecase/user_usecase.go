package usecase

import "github.com/taku10101/internal/domain"

type UserRepository interface {
    Create(user *domain.User) error
    Update(user *domain.User) error
    Delete(userID int) error
    FindByID(userID int) (*domain.User, error)
    FindAll() ([]domain.User, error)
}

type UserUseCase struct {
    userRepository UserRepository
}

func NewUserUseCase(userRepository UserRepository) *UserUseCase {
    return &UserUseCase{userRepository: userRepository}
}

func (u *UserUseCase) CreateUser(user *domain.User) error {
    return u.userRepository.Create(user)
}

func (u *UserUseCase) UpdateUser(user *domain.User) error {
    return u.userRepository.Update(user)
}

func (u *UserUseCase) DeleteUser(userID int) error {
    return u.userRepository.Delete(userID)
}

func (u *UserUseCase) GetUserByID(userID int) (*domain.User, error) {
    return u.userRepository.FindByID(userID)
}

func (u *UserUseCase) GetAllUsers() ([]domain.User, error) {
    return u.userRepository.FindAll()
}
