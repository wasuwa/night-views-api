package usecase

import (
	"context"
	"log/slog"

	apierrors "github.com/wasuwa/night-view-api/api-errors"
	"github.com/wasuwa/night-view-api/domain/model"
	"github.com/wasuwa/night-view-api/domain/repository"
)

// NightViewUsecase 夜景のユースケース
type NightViewUsecase interface {
	FetchNightViewByID(ctx context.Context, id string) (*model.NightView, *apierrors.APIError)
}

type nightViewUsecase struct {
	nightViewRepository repository.NightViewRepository
}

// NewNightViewUsecase NightViewUsecaseの実装を初期化する
func NewNightViewUsecase(rp repository.NightViewRepository) NightViewUsecase {
	return &nightViewUsecase{nightViewRepository: rp}
}

// FetchNightViewByID IDに紐づく夜景を取得する
func (nu *nightViewUsecase) FetchNightViewByID(ctx context.Context, id string) (*model.NightView, *apierrors.APIError) {
	nightView, err := nu.nightViewRepository.FindByID(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "夜景情報の取得ができません", slog.String("Error", err.Error()), slog.String("ID", id))
		return nil, apierrors.InternalServerError
	}
	if nightView == nil {
		slog.WarnContext(ctx, "夜景情報が見つかりません", slog.String("ID", id))
		return nil, apierrors.NightViewNotFound
	}
	return nightView, nil
}
