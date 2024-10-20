package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/artist"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/artist/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/httpErrors"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
)

type artistUsecase struct {
	artistRepo artist.Repo
	logger     logger.Logger
}

func NewArtistUsecase(artistRepo artist.Repo, logger logger.Logger) artist.Usecase {
	return &artistUsecase{artistRepo, logger}
}

func (usecase *artistUsecase) View(ctx context.Context, artistID uint64) (*dto.ArtistDTO, error) {
	foundArtist, err := usecase.artistRepo.FindById(ctx, artistID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Artist wasn't found: %v", err))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, "Artist wasn't found", err)
	}
	usecase.logger.Info("Artist found")

	dtoArtist, err := usecase.convertArtistToDTO(foundArtist)
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s artist: %v", foundArtist.Name, err))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrCreateDTOFailed, err)
	}

	return dtoArtist, nil
}

func (usecase *artistUsecase) Search(ctx context.Context, name string) ([]*dto.ArtistDTO, error) {
	foundArtists, err := usecase.artistRepo.FindByName(ctx, name)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Artist '%s' wasn't found: %v", name, err))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, "Can't find artist", err)
	}
	usecase.logger.Info("Artists found")

	var dtoArtists []*dto.ArtistDTO
	for _, artist := range foundArtists {
		dtoArtist, err := usecase.convertArtistToDTO(artist)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s artist: %v", artist.Name, err))
			return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrCreateDTOFailed, err)
		}
		dtoArtists = append(dtoArtists, dtoArtist)
	}

	return dtoArtists, nil
}

func (usecase *artistUsecase) GetAll(ctx context.Context) ([]*dto.ArtistDTO, error) {
	artists, err := usecase.artistRepo.GetAll(ctx)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load artists: %v", err))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, "Can't load artists", err)
	}
	usecase.logger.Info("Artists found")

	var dtoArtists []*dto.ArtistDTO
	for _, artist := range artists {
		dtoArtist, err := usecase.convertArtistToDTO(artist)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s artist: %v", artist.Name, err))
			return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrCreateDTOFailed, err)
		}
		dtoArtists = append(dtoArtists, dtoArtist)
	}

	return dtoArtists, nil
}

func (usecase *artistUsecase) convertArtistToDTO(artist *models.Artist) (*dto.ArtistDTO, error) {
	return dto.NewArtistDTO(artist), nil
}
