package usecase

import "github.com/taku10101/internal/domain"


type CaffeineRepository interface {
    Create(caffeine *domain.Caffeine) error
    Update(caffeine *domain.Caffeine) error
    Delete(caffeineID int) error
    FindByID(caffeineID int) (*domain.Caffeine, error)
    FindAll(userID int) ([]domain.Caffeine, error)
    GetScore(userID int) (int, error)
    FindLast7Days(userID int) ([]domain.Caffeine, error)
}

type CaffeineUseCase struct {
    caffeineRepository CaffeineRepository
}

func NewCaffeineUseCase(caffeineRepository CaffeineRepository) *CaffeineUseCase {
    return &CaffeineUseCase{caffeineRepository: caffeineRepository}
}

func (u *CaffeineUseCase) CreateCaffeine(caffeine *domain.Caffeine) error {
    return u.caffeineRepository.Create(caffeine)
}

func (u *CaffeineUseCase) UpdateCaffeine(caffeine *domain.Caffeine) error {
    return u.caffeineRepository.Update(caffeine)
}

func (u *CaffeineUseCase) DeleteCaffeine(caffeineID int) error {
    return u.caffeineRepository.Delete(caffeineID)
}

func (u *CaffeineUseCase) GetCaffeineByID(caffeineID int) (*domain.Caffeine, error) {
    return u.caffeineRepository.FindByID(caffeineID)
}

func (u *CaffeineUseCase) GetCaffeinesByUserID(userID int) ([]domain.Caffeine, error) {
    return u.caffeineRepository.FindAll(userID)
}

func (u *CaffeineUseCase) GetCaffeineScore(userID int) (int, error) {
    return u.caffeineRepository.GetScore(userID)
}

func (u *CaffeineUseCase) GetCaffeinesLast7Days(userID int) ([]domain.Caffeine, error) {
    return u.caffeineRepository.FindLast7Days(userID)
}