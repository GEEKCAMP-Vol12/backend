package usecase

import "github.com/taku10101/internal/domain"

type SleepRepository interface {
    Create(sleep *domain.Sleep) error
    Update(sleep *domain.Sleep) error
    Delete(sleepID int) error
    FindByID(sleepID int) (*domain.Sleep, error)
    FindAll() ([]domain.Sleep, error)
    GetScore(userID int) (int, error)
}

type SleepUseCase struct {
    sleepRepository SleepRepository
}

func NewSleepUseCase(sleepRepository SleepRepository) *SleepUseCase {
    return &SleepUseCase{sleepRepository: sleepRepository}
}

func (u *SleepUseCase) CreateSleep(sleep *domain.Sleep) error {
    return u.sleepRepository.Create(sleep)
}

func (u *SleepUseCase) UpdateSleep(sleep *domain.Sleep) error {
    return u.sleepRepository.Update(sleep)
}

func (u *SleepUseCase) DeleteSleep(sleepID int) error {
    return u.sleepRepository.Delete(sleepID)
}

func (u *SleepUseCase) GetSleepByID(sleepID int) (*domain.Sleep, error) {
    return u.sleepRepository.FindByID(sleepID)
}

func (u *SleepUseCase) GetAllSleeps() ([]domain.Sleep, error) {
    return u.sleepRepository.FindAll()
}

func (u *SleepUseCase) GetSleepScore(userID int) (int, error) {
    return u.sleepRepository.GetScore(userID)
}
