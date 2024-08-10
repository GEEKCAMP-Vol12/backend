package usecase

import (
	"github.com/google/uuid"
	"github.com/taku10101/internal/domain"
	"github.com/taku10101/internal/repository"
)

type CaffeineRepository interface {
    Create(caffeine *repository.Caffeine) error
    Update(caffeine *repository.Caffeine) error
    Delete(caffeineID string) error
    FindByID(caffeineID string) (*repository.Caffeine, error)
    FindAll(userID string) ([]repository.Caffeine, error)
    GetScore(userID string) (int, error)
    FindLast7Days(userID string) ([]repository.Caffeine, error)
}

type CaffeineUseCase struct {
    caffeineRepository CaffeineRepository
}

func NewCaffeineUseCase(caffeineRepository CaffeineRepository) *CaffeineUseCase {
    return &CaffeineUseCase{caffeineRepository: caffeineRepository}
}

func (u *CaffeineUseCase) CreateCaffeine(caffeine *domain.CreateCaffeineRequest) error {
    dbCaffeine := &repository.Caffeine{
        UserID: caffeine.UserID,
        Score: caffeine.Score,
    }
    return u.caffeineRepository.Create(dbCaffeine)
}

func (u *CaffeineUseCase) UpdateCaffeine(caffeine *domain.UpdateCaffeineRequest) error {
    
    dbCaffeine := &repository.Caffeine{
        ID:     uuid.Must(uuid.NewRandom()),
        UserID: caffeine.UserID,
        Score:  caffeine.Score,
    }
    return u.caffeineRepository.Update(dbCaffeine)
}

func (u *CaffeineUseCase) DeleteCaffeine(caffeineID string) error {
    return u.caffeineRepository.Delete(caffeineID)
}

func (u *CaffeineUseCase) GetCaffeineByID(caffeineID string) (*repository.Caffeine, error) {
    return u.caffeineRepository.FindByID(caffeineID)
}

func (u *CaffeineUseCase) GetCaffeinesByUserID(userID string) ([]repository.Caffeine, error) {
    return u.caffeineRepository.FindAll(userID)
}

func (u *CaffeineUseCase) GetCaffeineScore(userID string) (int, error) {
    return u.caffeineRepository.GetScore(userID)
}

func (u *CaffeineUseCase) GetCaffeinesLast7Days(userID string) ([]repository.Caffeine, error) {
    return u.caffeineRepository.FindLast7Days(userID)
}