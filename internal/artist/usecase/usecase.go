package usecase

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/artist"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/artist/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
)

type artistUsecase struct {
	cfg        *config.Config
	artistRepo artist.Repo
	logger     logger.Logger
}

func NewArtistUsecase(cfg *config.Config, artistRepo artist.Repo, logger logger.Logger) artist.Usecase {
	return &artistUsecase{cfg, artistRepo, logger}
}

func (usecase *artistUsecase) View(ctx context.Context, artistID uint64) (*dto.ArtistDTO, error) {
	foundArtist, err := usecase.artistRepo.FindById(ctx, artistID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Artist wasn't found: %v", err))
		return nil, fmt.Errorf("Artist wasn't found")
	}
	usecase.logger.Info("Artist found")

	dtoArtist, err := usecase.convertArtistToDTO(foundArtist)
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s artist: %v", foundArtist.Name, err))
		return nil, fmt.Errorf("Can't create DTO")
	}

	return dtoArtist, nil
}

func (usecase *artistUsecase) Search(ctx context.Context, name string) ([]*dto.ArtistDTO, error) {
	foundArtists, err := usecase.artistRepo.FindByName(ctx, name)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Artist '%s' wasn't found: %v", name, err))
		return nil, fmt.Errorf("Can't find artist")
	}
	usecase.logger.Info("Artists found")

	var dtoArtists []*dto.ArtistDTO
	for _, artist := range foundArtists {
		dtoArtist, err := usecase.convertArtistToDTO(artist)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s artist: %v", artist.Name, err))
			return nil, fmt.Errorf("Can't create DTO")
		}
		dtoArtists = append(dtoArtists, dtoArtist)
	}

	return dtoArtists, nil
}

func (usecase *artistUsecase) GetAll(ctx context.Context) ([]*dto.ArtistDTO, error) {
	artists, err := usecase.artistRepo.GetAll(ctx)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load artists: %v", err))
		return nil, fmt.Errorf("Can't load artists")
	}
	usecase.logger.Info("Artists found")

	var dtoArtists []*dto.ArtistDTO
	for _, artist := range artists {
		dtoArtist, err := usecase.convertArtistToDTO(artist)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s artist: %v", artist.Name, err))
			return nil, fmt.Errorf("Can't create DTO")
		}
		dtoArtists = append(dtoArtists, dtoArtist)
	}

	return dtoArtists, nil
}

func (usecase *artistUsecase) convertArtistToDTO(artist *models.Artist) (*dto.ArtistDTO, error) {
	return dto.NewArtistDTO(artist), nil
}
