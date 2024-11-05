package usecase

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/genre"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/genre/dto"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
)

type genreUsecase struct {
	genreRepo genre.Repo
	logger    logger.Logger
}

func NewGenreUsecase(genreRepo genre.Repo, logger logger.Logger) genre.Usecase {
	return &genreUsecase{genreRepo, logger}
}

func (usecase *genreUsecase) GetAll(ctx context.Context) ([]*dto.GenreDTO, error) {
	genres, err := usecase.genreRepo.GetAll(ctx)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load genres: %v", err))
		return nil, fmt.Errorf("Can't find genres")
	}
	usecase.logger.Info("Genres found")

	var dtoGenres []*dto.GenreDTO
	for _, genre := range genres {
		dtoGenre, err := usecase.convertGenreToDTO(ctx, genre)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for genre: %v", err))
			return nil, fmt.Errorf("Can't create DTO")
		}
		dtoGenres = append(dtoGenres, dtoGenre)
	}

	return dtoGenres, nil
}

func (usecase *genreUsecase) GetAllByArtistID(ctx context.Context, artistID uint64) ([]*dto.GenreDTO, error) {
	genres, err := usecase.genreRepo.GetAllByArtistID(ctx, artistID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load genres by artist ID %d: %v", artistID, err))
		return nil, fmt.Errorf("Can't load genres by artist ID %d", artistID)
	}
	usecase.logger.Infof("Genres found for artist ID %d", artistID)

	var dtoGenres []*dto.GenreDTO
	for _, genre := range genres {
		dtoGenre, err := usecase.convertGenreToDTO(ctx, genre)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for genre: %v", err))
			return nil, fmt.Errorf("Can't create DTO")
		}
		dtoGenres = append(dtoGenres, dtoGenre)
	}

	return dtoGenres, nil
}

func (usecase *genreUsecase) GetAllByTrackID(ctx context.Context, trackID uint64) ([]*dto.GenreDTO, error) {
	genres, err := usecase.genreRepo.GetAllByTrackID(ctx, trackID)
	if err != nil {
		usecase.logger.Warn(fmt.Sprintf("Can't load genres by track ID %d: %v", trackID, err))
		return nil, fmt.Errorf("Can't load genres by track ID %d", trackID)
	}
	usecase.logger.Infof("Genres found for track ID %d", trackID)

	var dtoGenres []*dto.GenreDTO
	for _, genre := range genres {
		dtoGenre, err := usecase.convertGenreToDTO(ctx, genre)
		if err != nil {
			usecase.logger.Error(fmt.Sprintf("Can't create DTO for genre: %v", err))
			return nil, fmt.Errorf("Can't create DTO")
		}
		dtoGenres = append(dtoGenres, dtoGenre)
	}

	return dtoGenres, nil
}

func (usecase *genreUsecase) convertGenreToDTO(ctx context.Context, genre *models.Genre) (*dto.GenreDTO, error) {
	return dto.NewGenreDTO(genre), nil
}
