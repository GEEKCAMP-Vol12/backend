package usecase

import "github.com/taku10101/internal/domain"


type CaffeineRepository interface {
    Create(caffeine *domain.Caffeine) error
    Update(caffeine *domain.Caffeine) error
    Delete(caffeineID string) error
    FindByID(caffeineID string) (*domain.Caffeine, error)
    FindAll(userID string) ([]domain.Caffeine, error)
    GetScore(userID string) (int, error)
    FindLast7Days(userID string) ([]domain.Caffeine, error)
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

func (u *CaffeineUseCase) DeleteCaffeine(caffeineID string) error {
    return u.caffeineRepository.Delete(caffeineID)
}

func (u *CaffeineUseCase) GetCaffeineByID(caffeineID string) (*domain.Caffeine, error) {
    return u.caffeineRepository.FindByID(caffeineID)
}

func (u *CaffeineUseCase) GetCaffeinesByUserID(userID string) ([]domain.Caffeine, error) {
    return u.caffeineRepository.FindAll(userID)
}

func (u *CaffeineUseCase) GetCaffeineScore(userID string ) (int, error) {
    return u.caffeineRepository.GetScore(userID)
}

func (u *CaffeineUseCase) GetCaffeinesLast7Days(userID string) ([]domain.Caffeine, error) {
    return u.caffeineRepository.FindLast7Days(userID)
}