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
	nightViewRepository      repository.NightViewRepository
	nearestStationRepository repository.NearestStationRepository
}

// NewNightViewUsecase NightViewUsecaseの実装を初期化する
func NewNightViewUsecase(nvRepo repository.NightViewRepository, nsRepo repository.NearestStationRepository) NightViewUsecase {
	return &nightViewUsecase{
		nightViewRepository:      nvRepo,
		nearestStationRepository: nsRepo,
	}
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

	nearestStations, err := nu.nearestStationRepository.FindByNightViewID(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "最寄り駅一覧の取得ができません", slog.String("Error", err.Error()), slog.String("NightViewID", id))
		return nil, apierrors.InternalServerError
	}
	nightView.NearestStations = nearestStations
	return nightView, nil
}
