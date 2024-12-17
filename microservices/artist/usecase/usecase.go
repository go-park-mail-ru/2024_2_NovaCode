package usecase

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	uuid "github.com/google/uuid"
)

type artistUsecase struct {
	artistRepo artist.Repo
	logger     logger.Logger
}

func NewArtistUsecase(artistRepo artist.Repo, logger logger.Logger) artist.Usecase {
	return &artistUsecase{artistRepo, logger}
}

func (usecase *artistUsecase) View(ctx context.Context, artistID uint64) (*dto.ArtistDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	foundArtist, err := usecase.artistRepo.FindById(ctx, artistID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Artist wasn't found: %v", err), requestID)
		return nil, fmt.Errorf("Artist wasn't found")
	}
	usecase.logger.Info("Artist found", requestID)

	dtoArtist, err := usecase.convertArtistToDTO(foundArtist)
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s artist: %v", foundArtist.Name, err), requestID)
		return nil, fmt.Errorf("Can't create DTO")
	}

	return dtoArtist, nil
}

func (usecase *artistUsecase) Search(ctx context.Context, query string) ([]*dto.ArtistDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	foundArtists, err := usecase.artistRepo.FindByQuery(ctx, query)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Artist '%s' wasn't found: %v", query, err), requestID)
		return nil, fmt.Errorf("Can't find artist")
	}
	usecase.logger.Info("Artists found", requestID)

	var dtoArtists []*dto.ArtistDTO
	for _, artist := range foundArtists {
		dtoArtist, err := usecase.convertArtistToDTO(artist)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s artist: %v", artist.Name, err), requestID)
			return nil, fmt.Errorf("Can't create DTO")
		}
		dtoArtists = append(dtoArtists, dtoArtist)
	}

	return dtoArtists, nil
}

func (usecase *artistUsecase) GetAll(ctx context.Context) ([]*dto.ArtistDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	artists, err := usecase.artistRepo.GetAll(ctx)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load artists: %v", err), requestID)
		return nil, fmt.Errorf("Can't load artists")
	}
	usecase.logger.Info("Artists found", requestID)

	var dtoArtists []*dto.ArtistDTO
	for _, artist := range artists {
		dtoArtist, err := usecase.convertArtistToDTO(artist)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s artist: %v", artist.Name, err), requestID)
			return nil, fmt.Errorf("Can't create DTO")
		}
		dtoArtists = append(dtoArtists, dtoArtist)
	}

	return dtoArtists, nil
}

func (usecase *artistUsecase) AddFavoriteArtist(ctx context.Context, userID uuid.UUID, artistID uint64) error {
	requestID := ctx.Value(utils.RequestIDKey{})
	if err := usecase.artistRepo.AddFavoriteArtist(ctx, userID, artistID); err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't add artist %d to favorite for user %v: %v", artistID, userID, err), requestID)
		return fmt.Errorf("Can't add artist %d to favorite for user %v: %v", artistID, userID, err)
	}

	return nil
}

func (usecase *artistUsecase) DeleteFavoriteArtist(ctx context.Context, userID uuid.UUID, artistID uint64) error {
	requestID := ctx.Value(utils.RequestIDKey{})
	if err := usecase.artistRepo.DeleteFavoriteArtist(ctx, userID, artistID); err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't delete artist %d from favorite for user %v: %v", artistID, userID, err), requestID)
		return fmt.Errorf("Can't delete artist %d from favorite for user %v: %v", artistID, userID, err)
	}

	return nil
}

func (usecase *artistUsecase) IsFavoriteArtist(ctx context.Context, userID uuid.UUID, artistID uint64) (bool, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	exists, err := usecase.artistRepo.IsFavoriteArtist(ctx, userID, artistID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't find artist %d in favorite for user %v: %v", artistID, userID, err), requestID)
		return false, fmt.Errorf("Can't find artist %d in favorite for user %v: %v", artistID, userID, err)
	}

	return exists, nil
}

func (usecase *artistUsecase) GetFavoriteArtists(ctx context.Context, userID uuid.UUID) ([]*dto.ArtistDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	artists, err := usecase.artistRepo.GetFavoriteArtists(ctx, userID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load artists by user ID %v: %v", userID, err), requestID)
		return nil, fmt.Errorf("Can't load artists by user ID %v", userID)
	}
	usecase.logger.Infof("Found %d artists for user ID %v", len(artists), userID)

	var dtoArtists []*dto.ArtistDTO
	for _, artist := range artists {
		dtoArtist, err := usecase.convertArtistToDTO(artist)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s artist: %v", artist.Name, err), requestID)
			return nil, fmt.Errorf("Can't create DTO for artist")
		}
		dtoArtists = append(dtoArtists, dtoArtist)
	}

	return dtoArtists, nil
}

func (usecase *artistUsecase) convertArtistToDTO(artist *models.Artist) (*dto.ArtistDTO, error) {
	return dto.NewArtistDTO(artist), nil
}
