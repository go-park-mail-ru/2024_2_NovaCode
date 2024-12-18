package usecase

import (
	"context"
	"fmt"

	uuid "github.com/google/uuid"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album"
	"github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	artistService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/artist"
)

type albumUsecase struct {
	albumRepo    album.Repo
	artistClient artistService.ArtistServiceClient
	logger       logger.Logger
}

func NewAlbumUsecase(
	albumRepo album.Repo,
	artistClient artistService.ArtistServiceClient,
	logger logger.Logger,
) album.Usecase {
	return &albumUsecase{albumRepo, artistClient, logger}
}

func (usecase *albumUsecase) View(ctx context.Context, albumID uint64) (*dto.AlbumDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	foundAlbum, err := usecase.albumRepo.FindById(ctx, albumID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Album wasn't found: %v", err), requestID)
		return nil, fmt.Errorf("Album wasn't found")
	}
	usecase.logger.Info("Album found", requestID)

	dtoAlbum, err := usecase.convertAlbumToDTO(ctx, foundAlbum)
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s album: %v", foundAlbum.Name, err), requestID)
		return nil, fmt.Errorf("Can't create DTO")
	}

	return dtoAlbum, nil
}

func (usecase *albumUsecase) Search(ctx context.Context, name string) ([]*dto.AlbumDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	foundAlbums, err := usecase.albumRepo.FindByQuery(ctx, name)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Albums with name '%s' were not found: %v", name, err), requestID)
		return nil, fmt.Errorf("Can't find albums")
	}
	usecase.logger.Info("Albums found", requestID)

	var dtoAlbums []*dto.AlbumDTO
	for _, album := range foundAlbums {
		dtoAlbum, err := usecase.convertAlbumToDTO(ctx, album)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s album: %v", album.Name, err), requestID)
			return nil, fmt.Errorf("Can't create DTO")
		}
		dtoAlbums = append(dtoAlbums, dtoAlbum)
	}

	return dtoAlbums, nil
}

func (usecase *albumUsecase) GetAll(ctx context.Context) ([]*dto.AlbumDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	albums, err := usecase.albumRepo.GetAll(ctx)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load albums: %v", err), requestID)
		return nil, fmt.Errorf("Can't find albums")
	}
	usecase.logger.Info("Albums found", requestID)

	var dtoAlbums []*dto.AlbumDTO
	for _, album := range albums {
		dtoAlbum, err := usecase.convertAlbumToDTO(ctx, album)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s album: %v", album.Name, err), requestID)
			return nil, fmt.Errorf("Can't create DTO")
		}
		dtoAlbums = append(dtoAlbums, dtoAlbum)
	}

	return dtoAlbums, nil
}

func (usecase *albumUsecase) GetAllByArtistID(ctx context.Context, artistID uint64) ([]*dto.AlbumDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	albums, err := usecase.albumRepo.GetAllByArtistID(ctx, artistID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load albums by artist ID %d: %v", artistID, err), requestID)
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

func (usecase *albumUsecase) AddFavoriteAlbum(ctx context.Context, userID uuid.UUID, albumID uint64) error {
	requestID := ctx.Value(utils.RequestIDKey{})
	if err := usecase.albumRepo.AddFavoriteAlbum(ctx, userID, albumID); err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't add album %d to favorite for user %v: %v", albumID, userID, err), requestID)
		return fmt.Errorf("Can't add album %d to favorite for user %v: %v", albumID, userID, err)
	}

	return nil
}

func (usecase *albumUsecase) DeleteFavoriteAlbum(ctx context.Context, userID uuid.UUID, albumID uint64) error {
	requestID := ctx.Value(utils.RequestIDKey{})
	if err := usecase.albumRepo.DeleteFavoriteAlbum(ctx, userID, albumID); err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't delete album %d from favorite for user %v: %v", albumID, userID, err), requestID)
		return fmt.Errorf("Can't delete album %d from favorite for user %v: %v", albumID, userID, err)
	}

	return nil
}

func (usecase *albumUsecase) IsFavoriteAlbum(ctx context.Context, userID uuid.UUID, albumID uint64) (bool, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	exists, err := usecase.albumRepo.IsFavoriteAlbum(ctx, userID, albumID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't find album %d in favorite for user %v: %v", albumID, userID, err), requestID)
		return false, fmt.Errorf("Can't find album %d in favorite for user %v: %v", albumID, userID, err)
	}

	return exists, nil
}

func (usecase *albumUsecase) GetFavoriteAlbums(ctx context.Context, userID uuid.UUID) ([]*dto.AlbumDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	albums, err := usecase.albumRepo.GetFavoriteAlbums(ctx, userID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load albums by user ID %v: %v", userID, err), requestID)
		return nil, fmt.Errorf("Can't load albums by user ID %v", userID)
	}
	usecase.logger.Infof("Found %d albums for user ID %v", len(albums), userID)

	var dtoAlbums []*dto.AlbumDTO
	for _, album := range albums {
		dtoAlbum, err := usecase.convertAlbumToDTO(ctx, album)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for %s album: %v", album.Name, err), requestID)
			return nil, fmt.Errorf("Can't create DTO for album")
		}
		dtoAlbums = append(dtoAlbums, dtoAlbum)
	}

	return dtoAlbums, nil
}

func (usecase *albumUsecase) convertAlbumToDTO(ctx context.Context, album *models.Album) (*dto.AlbumDTO, error) {
	requestID := ctx.Value(utils.RequestIDKey{})
	artist, err := usecase.artistClient.FindByID(ctx, &artistService.FindByIDRequest{Id: album.ArtistID})
	if err != nil {
		usecase.logger.Error(fmt.Sprintf("Can't find artist for album %s: %v", album.Name, err), requestID)
		return nil, fmt.Errorf("Can't find artist for album")
	}

	albumDTO := dto.NewAlbumDTO(album)
	albumDTO.ArtistName = artist.Artist.Name
	albumDTO.ArtistID = artist.Artist.Id

	return albumDTO, nil
}
