package usecase

import (
	"context"

	"github.com/wasuwa/night-view-api/domain/model"
	"github.com/wasuwa/night-view-api/domain/repository"
)

// NightViewUsecase 夜景のユースケース
type NightViewUsecase interface {
	FetchNightViewByID(ctx context.Context, id string) (*model.NightView, error)
}

type nightViewUsecase struct {
	nightViewRepository repository.NightViewRepository
}

// NewNightViewUsecase NightViewUsecaseの実装を初期化する
func NewNightViewUsecase(rp repository.NightViewRepository) NightViewUsecase {
	return &nightViewUsecase{nightViewRepository: rp}
}

// FetchNightViewByID IDに紐づく夜景を取得する
func (nu *nightViewUsecase) FetchNightViewByID(ctx context.Context, id string) (*model.NightView, error) {
	nightView, err := nu.nightViewRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return nightView, nil
}
