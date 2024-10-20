package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/album"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/album/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/artist"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/httpErrors"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
)

type albumUsecase struct {
	albumRepo  album.Repo
	artistRepo artist.Repo
	logger     logger.Logger
}

func NewAlbumUsecase(albumRepo album.Repo, artistRepo artist.Repo, logger logger.Logger) album.Usecase {
	return &albumUsecase{albumRepo, artistRepo, logger}
}

func (usecase *albumUsecase) View(ctx context.Context, albumID uint64) (*dto.AlbumDTO, error) {
	foundAlbum, err := usecase.albumRepo.FindById(ctx, albumID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Album wasn't found: %v", err))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, "Album wasn't found", err)
	}
	usecase.logger.Info("Album found")

	dtoAlbum, err := usecase.convertAlbumToDTO(ctx, foundAlbum)
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s album: %v", foundAlbum.Name, err))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrCreateDTOFailed, err)
	}

	return dtoAlbum, nil
}

func (usecase *albumUsecase) Search(ctx context.Context, name string) ([]*dto.AlbumDTO, error) {
	foundAlbums, err := usecase.albumRepo.FindByName(ctx, name)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Albums with name '%s' were not found: %v", name, err))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, "Can't find albums", err)
	}
	usecase.logger.Info("Albums found")

	var dtoAlbums []*dto.AlbumDTO
	for _, album := range foundAlbums {
		dtoAlbum, err := usecase.convertAlbumToDTO(ctx, album)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s album: %v", album.Name, err))
			return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrCreateDTOFailed, err)
		}
		dtoAlbums = append(dtoAlbums, dtoAlbum)
	}

	return dtoAlbums, nil
}

func (usecase *albumUsecase) GetAll(ctx context.Context) ([]*dto.AlbumDTO, error) {
	albums, err := usecase.albumRepo.GetAll(ctx)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load albums: %v", err))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, "Can't find albums", err)
	}
	usecase.logger.Info("Albums found")

	var dtoAlbums []*dto.AlbumDTO
	for _, album := range albums {
		dtoAlbum, err := usecase.convertAlbumToDTO(ctx, album)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s album: %v", album.Name, err))
			return nil, httpErrors.NewRestError(http.StatusBadRequest, httpErrors.StrCreateDTOFailed, err)
		}
		dtoAlbums = append(dtoAlbums, dtoAlbum)
	}

	return dtoAlbums, nil
}

func (usecase *albumUsecase) convertAlbumToDTO(ctx context.Context, album *models.Album) (*dto.AlbumDTO, error) {
	artist, err := usecase.artistRepo.FindById(ctx, album.ArtistID)
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("Can't find artist for album %s: %v", album.Name, err))
		return nil, httpErrors.NewRestError(http.StatusBadRequest, "Can't find artist for album", err)
	}

	return dto.NewAlbumDTO(album, artist), nil
}
