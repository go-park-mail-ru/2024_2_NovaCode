package usecase

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/album"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/album/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/artist"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
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
		return nil, fmt.Errorf("Album wasn't found")
	}
	usecase.logger.Info("Album found")

	dtoAlbum, err := usecase.convertAlbumToDTO(ctx, foundAlbum)
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s album: %v", foundAlbum.Name, err))
		return nil, fmt.Errorf("Can't create DTO")
	}

	return dtoAlbum, nil
}

func (usecase *albumUsecase) Search(ctx context.Context, name string) ([]*dto.AlbumDTO, error) {
	foundAlbums, err := usecase.albumRepo.FindByName(ctx, name)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Albums with name '%s' were not found: %v", name, err))
		return nil, fmt.Errorf("Can't find albums")
	}
	usecase.logger.Info("Albums found")

	var dtoAlbums []*dto.AlbumDTO
	for _, album := range foundAlbums {
		dtoAlbum, err := usecase.convertAlbumToDTO(ctx, album)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s album: %v", album.Name, err))
			return nil, fmt.Errorf("Can't create DTO")
		}
		dtoAlbums = append(dtoAlbums, dtoAlbum)
	}

	return dtoAlbums, nil
}

func (usecase *albumUsecase) GetAll(ctx context.Context) ([]*dto.AlbumDTO, error) {
	albums, err := usecase.albumRepo.GetAll(ctx)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load albums: %v", err))
		return nil, fmt.Errorf("Can't find albums")
	}
	usecase.logger.Info("Albums found")

	var dtoAlbums []*dto.AlbumDTO
	for _, album := range albums {
		dtoAlbum, err := usecase.convertAlbumToDTO(ctx, album)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s album: %v", album.Name, err))
			return nil, fmt.Errorf("Can't create DTO")
		}
		dtoAlbums = append(dtoAlbums, dtoAlbum)
	}

	return dtoAlbums, nil
}

func (usecase *albumUsecase) GetAllByArtistID(ctx context.Context, artistID uint64) ([]*dto.AlbumDTO, error) {
	albums, err := usecase.albumRepo.GetAllByArtistID(ctx, artistID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load albums by artist ID %d: %v", artistID, err))
		return nil, fmt.Errorf("Can't load albums by artist ID %d", artistID)
	}
	usecase.logger.Infof("Found %d albums for artist ID %d", len(albums), artistID)

	var dtoAlbums []*dto.AlbumDTO
	for _, album := range albums {
		dtoAlbum, err := usecase.convertAlbumToDTO(ctx, album)
		if err != nil {
			usecase.logger.Errorf("Can't create DTO for %s album: %v", album.Name, err)
			return nil, fmt.Errorf("Can't create DTO")
		}
		dtoAlbums = append(dtoAlbums, dtoAlbum)
	}

	return dtoAlbums, nil
}

func (usecase *albumUsecase) convertAlbumToDTO(ctx context.Context, album *models.Album) (*dto.AlbumDTO, error) {
	artist, err := usecase.artistRepo.FindById(ctx, album.ArtistID)
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("Can't find artist for album %s: %v", album.Name, err))
		return nil, fmt.Errorf("Can't find artist for album")
	}

	return dto.NewAlbumDTO(album, artist), nil
}
