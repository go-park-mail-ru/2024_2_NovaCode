package usecase

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	uuid "github.com/google/uuid"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	mockAlbum "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/mock"
	mockArtist "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	artistService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/artist"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUsecase_View_FoundAlbum(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	albumUsecase := NewAlbumUsecase(albumRepoMock, artistClientMock, logger)

	findByIDResponseArtist := &artistService.FindByIDResponse{
		Artist: &artistService.Artist{
			Id:      1,
			Name:    "quinn",
			Bio:     "Some random bio",
			Country: "USA",
			Image:   "/imgs/artists/artist_1.jpg",
		},
	}

	album := &models.Album{
		ID:          1,
		Name:        "Attempted Lover",
		ReleaseDate: time.Date(2024, 07, 19, 0, 0, 0, 0, time.UTC),
		Image:       "/imgs/albums/album_1.jpg",
		ArtistID:    1,
	}

	ctx := context.Background()
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: album.ArtistID}).Return(findByIDResponseArtist, nil)
	albumRepoMock.EXPECT().FindById(ctx, album.ID).Return(album, nil)

	dtoAlbum, err := albumUsecase.View(ctx, album.ID)

	require.NoError(t, err)
	require.NotNil(t, dtoAlbum)
	require.Equal(t, findByIDResponseArtist.Artist.Name, dtoAlbum.ArtistName)
	require.Equal(t, album.Name, dtoAlbum.Name)
}

func TestUsecase_View_NotFoundAlbums(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	ctx := context.Background()

	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	albumUsecase := NewAlbumUsecase(albumRepoMock, artistClientMock, logger)

	albumRepoMock.EXPECT().FindById(ctx, uint64(1)).Return(nil, errors.New("Album wasn't found"))
	dtoAlbum, err := albumUsecase.View(ctx, uint64(1))

	require.Error(t, err)
	require.Nil(t, dtoAlbum)
	require.EqualError(t, err, "Album wasn't found")
}

func TestUsecase_Search_FoundAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	albumUsecase := NewAlbumUsecase(albumRepoMock, artistClientMock, logger)

	now := time.Now()
	findByIDResponseArtists := []*artistService.FindByIDResponse{
		{
			Artist: &artistService.Artist{
				Id:      uint64(1),
				Name:    "artist1",
				Bio:     "1",
				Country: "1",
				Image:   "1",
			},
		},
		{
			Artist: &artistService.Artist{
				Id:      uint64(2),
				Name:    "artist2",
				Bio:     "2",
				Country: "2",
				Image:   "2",
			},
		},
	}

	albums := []*models.Album{
		{
			ID: uint64(1), Name: "test", ReleaseDate: now, Image: "1",
			ArtistID: uint64(1), CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(2), Name: "test", ReleaseDate: now, Image: "2",
			ArtistID: uint64(2), CreatedAt: now, UpdatedAt: now,
		},
	}

	ctx := context.Background()
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: albums[0].ArtistID}).Return(findByIDResponseArtists[0], nil)
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: albums[1].ArtistID}).Return(findByIDResponseArtists[1], nil)
	albumRepoMock.EXPECT().FindByQuery(ctx, "test").Return([]*models.Album{albums[0], albums[1]}, nil)

	dtoAlbums, err := albumUsecase.Search(ctx, "test")

	require.NoError(t, err)
	require.NotNil(t, dtoAlbums)
	require.Equal(t, 2, len(dtoAlbums))
	require.Equal(t, albums[0].Name, dtoAlbums[0].Name)
	require.Equal(t, albums[1].Name, dtoAlbums[1].Name)
	require.Equal(t, findByIDResponseArtists[0].Artist.Name, dtoAlbums[0].ArtistName)
	require.Equal(t, findByIDResponseArtists[1].Artist.Name, dtoAlbums[1].ArtistName)
}

func TestUsecase_Search_NotFoundAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	albumUsecase := NewAlbumUsecase(albumRepoMock, artistClientMock, logger)

	ctx := context.Background()
	albumRepoMock.EXPECT().FindByQuery(ctx, "album").Return(nil, errors.New("Can't find albums"))

	dtoAlbums, err := albumUsecase.Search(ctx, "album")

	require.Error(t, err)
	require.Nil(t, dtoAlbums)
	require.EqualError(t, err, "Can't find albums")
}

func TestUsecase_GetAll_FoundAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	albumUsecase := NewAlbumUsecase(albumRepoMock, artistClientMock, logger)

	now := time.Now()
	findByIDResponseArtists := []*artistService.FindByIDResponse{
		{
			Artist: &artistService.Artist{
				Id:      uint64(1),
				Name:    "artist1",
				Bio:     "1",
				Country: "1",
				Image:   "1",
			},
		},
		{
			Artist: &artistService.Artist{
				Id:      uint64(2),
				Name:    "artist2",
				Bio:     "2",
				Country: "2",
				Image:   "2",
			},
		},
		{
			Artist: &artistService.Artist{
				Id:      uint64(3),
				Name:    "artist3",
				Bio:     "3",
				Country: "3",
				Image:   "3",
			},
		},
	}

	albums := []*models.Album{
		{
			ID: uint64(1), Name: "album1", ReleaseDate: now, Image: "1",
			ArtistID: uint64(1), CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(2), Name: "album2", ReleaseDate: now, Image: "2",
			ArtistID: uint64(2), CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(3), Name: "album3", ReleaseDate: now, Image: "3",
			ArtistID: uint64(3), CreatedAt: now, UpdatedAt: now,
		},
	}

	ctx := context.Background()
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: albums[0].ArtistID}).Return(findByIDResponseArtists[0], nil)
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: albums[1].ArtistID}).Return(findByIDResponseArtists[1], nil)
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: albums[2].ArtistID}).Return(findByIDResponseArtists[2], nil)
	albumRepoMock.EXPECT().GetAll(ctx).Return(albums, nil)

	dtoAlbums, err := albumUsecase.GetAll(ctx)

	require.NoError(t, err)
	require.NotNil(t, dtoAlbums)
	require.Equal(t, len(albums), len(dtoAlbums))

	for i := 0; i < len(albums); i++ {
		require.Equal(t, findByIDResponseArtists[i].Artist.Name, dtoAlbums[i].ArtistName)
		require.Equal(t, albums[i].Name, dtoAlbums[i].Name)
	}
}

func TestUsecase_GetAll_NotFoundAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	albumUsecase := NewAlbumUsecase(albumRepoMock, artistClientMock, logger)

	ctx := context.Background()
	albumRepoMock.EXPECT().GetAll(ctx).Return(nil, errors.New("Can't find albums"))
	dtoAlbums, err := albumUsecase.GetAll(ctx)

	require.Error(t, err)
	require.Nil(t, dtoAlbums)
	require.EqualError(t, err, "Can't find albums")
}

func TestUsecase_GetAllByArtistID_FoundAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	albumUsecase := NewAlbumUsecase(albumRepoMock, artistClientMock, logger)

	now := time.Now()
	findByIDResponseArtist := &artistService.FindByIDResponse{
		Artist: &artistService.Artist{
			Id:      uint64(1),
			Name:    "artist1",
			Bio:     "1",
			Country: "1",
			Image:   "1",
		},
	}

	albums := []*models.Album{
		{
			ID: uint64(1), Name: "album1", ReleaseDate: now, Image: "1",
			ArtistID: uint64(1), CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(2), Name: "album2", ReleaseDate: now, Image: "2",
			ArtistID: uint64(1), CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(3), Name: "album3", ReleaseDate: now, Image: "3",
			ArtistID: uint64(1), CreatedAt: now, UpdatedAt: now,
		},
	}

	ctx := context.Background()
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: uint64(1)}).Return(findByIDResponseArtist, nil).Times(3)
	albumRepoMock.EXPECT().GetAllByArtistID(ctx, uint64(1)).Return(albums, nil)

	dtoAlbums, err := albumUsecase.GetAllByArtistID(ctx, uint64(1))

	require.NoError(t, err)
	require.NotNil(t, dtoAlbums)
	require.Equal(t, len(albums), len(dtoAlbums))

	for i := 0; i < len(albums); i++ {
		require.Equal(t, findByIDResponseArtist.Artist.Name, dtoAlbums[i].ArtistName)
		require.Equal(t, albums[i].Name, dtoAlbums[i].Name)
	}
}

func TestUsecase_GetAllByArtistID_NotFoundAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	albumUsecase := NewAlbumUsecase(albumRepoMock, artistClientMock, logger)

	ctx := context.Background()
	albumRepoMock.EXPECT().GetAllByArtistID(ctx, uint64(1)).Return(nil, errors.New("Can't load albums by artist ID 1"))

	dtoAlbums, err := albumUsecase.GetAllByArtistID(ctx, uint64(1))

	require.Error(t, err)
	require.Nil(t, dtoAlbums)
	require.EqualError(t, err, "Can't load albums by artist ID 1")
}

func TestAlbumUsecase_AddFavoriteAlbum(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	mockAlbumRepo := mockAlbum.NewMockRepo(ctrl)
	logger := logger.New(&cfg.Service.Logger)

	albumUsecase := &albumUsecase{
		albumRepo: mockAlbumRepo,
		logger:    logger,
	}

	userID := uuid.New()
	albumID := uint64(12345)
	requestID := "request-id"
	ctx := context.WithValue(context.Background(), utils.RequestIDKey{}, requestID)

	t.Run("success", func(t *testing.T) {
		mockAlbumRepo.EXPECT().AddFavoriteAlbum(ctx, userID, albumID).Return(nil)
		err := albumUsecase.AddFavoriteAlbum(ctx, userID, albumID)
		require.NoError(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		mockError := fmt.Errorf("repository error")
		mockAlbumRepo.EXPECT().AddFavoriteAlbum(ctx, userID, albumID).Return(mockError)

		err := albumUsecase.AddFavoriteAlbum(ctx, userID, albumID)
		require.Error(t, err)
		require.Contains(t, err.Error(), "repository error")
	})
}

func TestAlbumUsecase_DeleteFavoriteAlbum(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	mockAlbumRepo := mockAlbum.NewMockRepo(ctrl)
	logger := logger.New(&cfg.Service.Logger)

	albumUsecase := &albumUsecase{
		albumRepo: mockAlbumRepo,
		logger:    logger,
	}

	userID := uuid.New()
	albumID := uint64(12345)
	requestID := "request-id"
	ctx := context.WithValue(context.Background(), utils.RequestIDKey{}, requestID)

	t.Run("success", func(t *testing.T) {
		mockAlbumRepo.EXPECT().DeleteFavoriteAlbum(ctx, userID, albumID).Return(nil)
		err := albumUsecase.DeleteFavoriteAlbum(ctx, userID, albumID)
		require.NoError(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		mockError := fmt.Errorf("repository error")
		mockAlbumRepo.EXPECT().DeleteFavoriteAlbum(ctx, userID, albumID).Return(mockError)

		err := albumUsecase.DeleteFavoriteAlbum(ctx, userID, albumID)
		require.Error(t, err)
		require.Contains(t, err.Error(), "repository error")
	})
}

func TestAlbumUsecase_IsFavoriteAlbum(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	mockAlbumRepo := mockAlbum.NewMockRepo(ctrl)
	logger := logger.New(&cfg.Service.Logger)

	albumUsecase := &albumUsecase{
		albumRepo: mockAlbumRepo,
		logger:    logger,
	}

	userID := uuid.New()
	albumID := uint64(12345)
	requestID := "request-id"
	ctx := context.WithValue(context.Background(), utils.RequestIDKey{}, requestID)

	t.Run("success", func(t *testing.T) {
		mockAlbumRepo.EXPECT().IsFavoriteAlbum(ctx, userID, albumID).Return(true, nil)
		exists, err := albumUsecase.IsFavoriteAlbum(ctx, userID, albumID)
		require.NoError(t, err)
		require.True(t, exists)
	})

	t.Run("album not found", func(t *testing.T) {
		mockAlbumRepo.EXPECT().IsFavoriteAlbum(ctx, userID, albumID).Return(false, nil)
		exists, err := albumUsecase.IsFavoriteAlbum(ctx, userID, albumID)
		require.NoError(t, err)
		require.False(t, exists)
	})

	t.Run("repository error", func(t *testing.T) {
		mockError := fmt.Errorf("repository error")
		mockAlbumRepo.EXPECT().IsFavoriteAlbum(ctx, userID, albumID).Return(false, mockError)

		exists, err := albumUsecase.IsFavoriteAlbum(ctx, userID, albumID)
		require.Error(t, err)
		require.Contains(t, err.Error(), "repository error")
		require.False(t, exists)
	})
}

func TestUsecase_GetFavoriteAlbums_FoundAlbums(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumUsecase := NewAlbumUsecase(albumRepoMock, artistClientMock, logger)

	now := time.Now()
	albums := []*models.Album{
		{
			ID: 1, Name: "album1", Image: "image1", ReleaseDate: now, ArtistID: 1,
		},
		{
			ID: 2, Name: "album2", Image: "image2", ReleaseDate: now, ArtistID: 1,
		},
		{
			ID: 3, Name: "album3", Image: "image3", ReleaseDate: now, ArtistID: 2,
		},
	}

	findByIDResponseArtists := []*artistService.FindByIDResponse{
		{
			Artist: &artistService.Artist{
				Id:      1,
				Name:    "artist1",
				Bio:     "bio1",
				Country: "country1",
				Image:   "image1",
			},
		},
		{
			Artist: &artistService.Artist{
				Id:      2,
				Name:    "artist2",
				Bio:     "bio2",
				Country: "country2",
				Image:   "image2",
			},
		},
	}

	userID := uuid.New()
	ctx := context.Background()
	albumRepoMock.EXPECT().GetFavoriteAlbums(ctx, userID).Return(albums, nil)
	for _, album := range albums {
		artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: album.ArtistID}).Return(findByIDResponseArtists[album.ArtistID-1], nil)
	}

	dtoAlbums, err := albumUsecase.GetFavoriteAlbums(ctx, userID)

	require.NoError(t, err)
	require.NotNil(t, dtoAlbums)
	require.Equal(t, len(albums), len(dtoAlbums))

	for i := 0; i < len(albums); i++ {
		require.Equal(t, albums[i].Name, dtoAlbums[i].Name)
		require.Equal(t, findByIDResponseArtists[albums[i].ArtistID-1].Artist.Name, dtoAlbums[i].ArtistName)
	}
}

func TestUsecase_GetFavoriteAlbums_NotFoundAlbums(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumUsecase := NewAlbumUsecase(albumRepoMock, artistClientMock, logger)

	userID := uuid.New()
	ctx := context.Background()
	albumRepoMock.EXPECT().GetFavoriteAlbums(ctx, userID).Return(nil, errors.New(fmt.Sprintf("Can't load albums by user ID %v", userID)))

	dtoAlbums, err := albumUsecase.GetFavoriteAlbums(ctx, userID)

	require.Error(t, err)
	require.Nil(t, dtoAlbums)
	require.EqualError(t, err, fmt.Sprintf("Can't load albums by user ID %v", userID))
}

func TestUsecase_GetFavoriteAlbumsCount_FoundAlbums(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumUsecase := NewAlbumUsecase(albumRepoMock, artistClientMock, logger)

	userID := uuid.New()
	ctx := context.Background()
	expectedCount := uint64(3)
	albumRepoMock.EXPECT().GetFavoriteAlbumsCount(ctx, userID).Return(expectedCount, nil)

	count, err := albumUsecase.GetFavoriteAlbumsCount(ctx, userID)

	require.NoError(t, err)
	require.Equal(t, expectedCount, count)
}

func TestUsecase_GetFavoriteAlbumsCount_ErrorGettingCount(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumUsecase := NewAlbumUsecase(albumRepoMock, artistClientMock, logger)

	userID := uuid.New()
	ctx := context.Background()
	expectedError := fmt.Errorf("Can't load albums by user ID %v", userID)
	albumRepoMock.EXPECT().GetFavoriteAlbumsCount(ctx, userID).Return(uint64(0), expectedError)

	count, err := albumUsecase.GetFavoriteAlbumsCount(ctx, userID)

	require.Error(t, err)
	require.EqualError(t, err, expectedError.Error())
	require.Equal(t, uint64(0), count)
}

func TestUsecase_GetAlbumLikesCount_Success(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumUsecase := NewAlbumUsecase(albumRepoMock, artistClientMock, logger)

	ctx := context.Background()
	ctx = context.WithValue(ctx, utils.RequestIDKey{}, "test-request-id")

	albumID := uint64(123)
	expectedLikesCount := uint64(10)

	albumRepoMock.EXPECT().GetAlbumLikesCount(ctx, albumID).Return(expectedLikesCount, nil)

	likesCount, err := albumUsecase.GetAlbumLikesCount(ctx, albumID)
	require.NoError(t, err)
	require.Equal(t, expectedLikesCount, likesCount)
}

func TestUsecase_GetAlbumLikesCount_ErrorGettingLikesCount(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Service: config.ServiceConfig{
			Logger: config.LoggerConfig{
				Level:  "info",
				Format: "json",
			},
		},
	}

	logger := logger.New(&cfg.Service.Logger)
	albumRepoMock := mockAlbum.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumUsecase := NewAlbumUsecase(albumRepoMock, artistClientMock, logger)

	ctx := context.Background()
	ctx = context.WithValue(ctx, utils.RequestIDKey{}, "test-request-id")

	albumID := uint64(123)
	expectedError := fmt.Errorf("Can't load albums by user ID %v", albumID)

	albumRepoMock.EXPECT().GetAlbumLikesCount(ctx, albumID).Return(uint64(0), expectedError)

	likesCount, err := albumUsecase.GetAlbumLikesCount(ctx, albumID)
	require.Error(t, err)
	require.EqualError(t, err, expectedError.Error())
	require.Equal(t, uint64(0), likesCount)
}
