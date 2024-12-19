package usecase

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/go-park-mail-ru/2024_2_NovaCode/config"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/models"
	"github.com/go-park-mail-ru/2024_2_NovaCode/internal/utils"
	mockAlbum "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/album/mock"
	mockArtist "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/artist/mock"
	mockTrack "github.com/go-park-mail-ru/2024_2_NovaCode/microservices/track/mock"
	"github.com/go-park-mail-ru/2024_2_NovaCode/pkg/logger"
	albumService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/album"
	artistService "github.com/go-park-mail-ru/2024_2_NovaCode/proto/artist"
	"github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestUsecase_View_FoundTrack(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumClientMock := mockAlbum.NewMockAlbumServiceClient(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, artistClientMock, albumClientMock, logger)

	track := &models.Track{
		ID:          1,
		Name:        "ok im cool",
		Duration:    167,
		FilePath:    "/songs/track_1.mp4",
		Image:       "/imgs/tracks/track_1.jpg",
		ArtistID:    1,
		AlbumID:     1,
		ReleaseDate: time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
	}

	findByIDResponseArtist := &artistService.FindByIDResponse{
		Artist: &artistService.Artist{
			Id:      1,
			Name:    "quinn",
			Bio:     "Some random bio",
			Country: "USA",
			Image:   "/imgs/artists/artist_1.jpg",
		},
	}

	findByIDResponseAlbum := &albumService.FindByIDResponse{
		Album: &albumService.Album{
			Id:          1,
			Name:        "Attempted Lover",
			ReleaseDate: timestamppb.New(time.Date(2024, 07, 19, 0, 0, 0, 0, time.UTC)),
			Image:       "/imgs/albums/album_1.jpg",
			ArtistID:    1,
		},
	}

	ctx := context.Background()
	trackRepoMock.EXPECT().FindById(ctx, track.ID).Return(track, nil)
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: track.ArtistID}).Return(findByIDResponseArtist, nil)
	albumClientMock.EXPECT().FindByID(ctx, &albumService.FindByIDRequest{Id: track.AlbumID}).Return(findByIDResponseAlbum, nil)

	dtoTrack, err := trackUsecase.View(ctx, track.ID)

	require.NoError(t, err)
	require.NotNil(t, dtoTrack)
	require.Equal(t, track.Name, dtoTrack.Name)
	require.Equal(t, findByIDResponseArtist.Artist.Name, dtoTrack.ArtistName)
	require.Equal(t, findByIDResponseAlbum.Album.Name, dtoTrack.AlbumName)
}

func TestUsecase_View_NotFoundTrack(t *testing.T) {
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

	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumClientMock := mockAlbum.NewMockAlbumServiceClient(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, artistClientMock, albumClientMock, logger)

	trackRepoMock.EXPECT().FindById(ctx, uint64(1)).Return(nil, errors.New("Track wasn't found"))
	dtoTrack, err := trackUsecase.View(ctx, uint64(1))

	require.Error(t, err)
	require.Nil(t, dtoTrack)
	require.EqualError(t, err, "Track wasn't found")
}

func TestUsecase_Search_FoundTracks(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumClientMock := mockAlbum.NewMockAlbumServiceClient(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, artistClientMock, albumClientMock, logger)

	now := time.Now()
	tracks := []*models.Track{
		{
			ID: uint64(1), Name: "test", Duration: uint64(1), FilePath: "1", Image: "1",
			ArtistID: uint64(1), AlbumID: uint64(1), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(2), Name: "none", Duration: uint64(2), FilePath: "2", Image: "2",
			ArtistID: uint64(2), AlbumID: uint64(2), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(3), Name: "test", Duration: uint64(3), FilePath: "3", Image: "3",
			ArtistID: uint64(3), AlbumID: uint64(3), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
	}

	findByIDResponseArtist1 := &artistService.FindByIDResponse{
		Artist: &artistService.Artist{
			Id:      uint64(1),
			Name:    "artist1",
			Bio:     "1",
			Country: "1",
			Image:   "1",
		},
	}

	findByIDResponseArtist3 := &artistService.FindByIDResponse{
		Artist: &artistService.Artist{
			Id:      uint64(3),
			Name:    "artist3",
			Bio:     "3",
			Country: "3",
			Image:   "3",
		},
	}

	findByIDResponseAlbum1 := &albumService.FindByIDResponse{
		Album: &albumService.Album{
			Id:          uint64(1),
			Name:        "album1",
			ReleaseDate: timestamppb.New(now),
			Image:       "1",
			ArtistID:    uint64(1),
		},
	}

	findByIDResponseAlbum3 := &albumService.FindByIDResponse{
		Album: &albumService.Album{
			Id:          uint64(3),
			Name:        "album3",
			ReleaseDate: timestamppb.New(now),
			Image:       "3",
			ArtistID:    uint64(3),
		},
	}

	ctx := context.Background()
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: tracks[0].ArtistID}).Return(findByIDResponseArtist1, nil)
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: tracks[2].ArtistID}).Return(findByIDResponseArtist3, nil)
	albumClientMock.EXPECT().FindByID(ctx, &albumService.FindByIDRequest{Id: tracks[0].AlbumID}).Return(findByIDResponseAlbum1, nil)
	albumClientMock.EXPECT().FindByID(ctx, &albumService.FindByIDRequest{Id: tracks[2].AlbumID}).Return(findByIDResponseAlbum3, nil)
	trackRepoMock.EXPECT().FindByQuery(ctx, "test").Return([]*models.Track{tracks[0], tracks[2]}, nil)

	dtoTracks, err := trackUsecase.Search(ctx, "test")

	require.NoError(t, err)
	require.NotNil(t, dtoTracks)
	require.Equal(t, 2, len(dtoTracks))
	require.Equal(t, findByIDResponseArtist1.Artist.Name, dtoTracks[0].ArtistName)
	require.Equal(t, findByIDResponseAlbum1.Album.Name, dtoTracks[0].AlbumName)
	require.Equal(t, findByIDResponseArtist3.Artist.Name, dtoTracks[1].ArtistName)
	require.Equal(t, findByIDResponseAlbum3.Album.Name, dtoTracks[1].AlbumName)
}

func TestUsecase_Search_NotFoundTracks(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumClientMock := mockAlbum.NewMockAlbumServiceClient(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, artistClientMock, albumClientMock, logger)

	ctx := context.Background()
	trackRepoMock.EXPECT().FindByQuery(ctx, "song").Return(nil, errors.New("Can't find tracks"))

	dtoTracks, err := trackUsecase.Search(ctx, "song")

	require.Error(t, err)
	require.Nil(t, dtoTracks)
	require.EqualError(t, err, "Can't find tracks")
}

func TestUsecase_GetAll_FoundTracks(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumClientMock := mockAlbum.NewMockAlbumServiceClient(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, artistClientMock, albumClientMock, logger)

	now := time.Now()
	tracks := []*models.Track{
		{
			ID: uint64(1), Name: "test1", Duration: uint64(1), FilePath: "1", Image: "1",
			ArtistID: uint64(1), AlbumID: uint64(1), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(2), Name: "test2", Duration: uint64(2), FilePath: "2", Image: "2",
			ArtistID: uint64(2), AlbumID: uint64(2), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(3), Name: "test3", Duration: uint64(3), FilePath: "3", Image: "3",
			ArtistID: uint64(3), AlbumID: uint64(3), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
	}

	findByIDResponseArtist1 := &artistService.FindByIDResponse{
		Artist: &artistService.Artist{
			Id:      uint64(1),
			Name:    "artist1",
			Bio:     "1",
			Country: "1",
			Image:   "1",
		},
	}

	findByIDResponseArtist2 := &artistService.FindByIDResponse{
		Artist: &artistService.Artist{
			Id:      uint64(2),
			Name:    "artist2",
			Bio:     "2",
			Country: "2",
			Image:   "2",
		},
	}

	findByIDResponseArtist3 := &artistService.FindByIDResponse{
		Artist: &artistService.Artist{
			Id:      uint64(3),
			Name:    "artist3",
			Bio:     "3",
			Country: "3",
			Image:   "3",
		},
	}

	findByIDResponseAlbum1 := &albumService.FindByIDResponse{
		Album: &albumService.Album{
			Id:          uint64(1),
			Name:        "album1",
			ReleaseDate: timestamppb.New(now),
			Image:       "1",
			ArtistID:    uint64(1),
		},
	}

	findByIDResponseAlbum2 := &albumService.FindByIDResponse{
		Album: &albumService.Album{
			Id:          uint64(2),
			Name:        "album2",
			ReleaseDate: timestamppb.New(now),
			Image:       "2",
			ArtistID:    uint64(2),
		},
	}

	findByIDResponseAlbum3 := &albumService.FindByIDResponse{
		Album: &albumService.Album{
			Id:          uint64(3),
			Name:        "album3",
			ReleaseDate: timestamppb.New(now),
			Image:       "3",
			ArtistID:    uint64(3),
		},
	}

	ctx := context.Background()
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: tracks[0].ArtistID}).Return(findByIDResponseArtist1, nil)
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: tracks[1].ArtistID}).Return(findByIDResponseArtist2, nil)
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: tracks[2].ArtistID}).Return(findByIDResponseArtist3, nil)
	albumClientMock.EXPECT().FindByID(ctx, &albumService.FindByIDRequest{Id: tracks[0].AlbumID}).Return(findByIDResponseAlbum1, nil)
	albumClientMock.EXPECT().FindByID(ctx, &albumService.FindByIDRequest{Id: tracks[1].AlbumID}).Return(findByIDResponseAlbum2, nil)
	albumClientMock.EXPECT().FindByID(ctx, &albumService.FindByIDRequest{Id: tracks[2].AlbumID}).Return(findByIDResponseAlbum3, nil)
	trackRepoMock.EXPECT().GetAll(ctx).Return(tracks, nil)

	dtoTracks, err := trackUsecase.GetAll(ctx)

	require.NoError(t, err)
	require.NotNil(t, dtoTracks)
	require.Equal(t, len(tracks), len(dtoTracks))

	for i := 0; i < len(tracks); i++ {
		require.Equal(t, tracks[i].Name, dtoTracks[i].Name)
		require.Equal(t, findByIDResponseArtist1.Artist.Name, dtoTracks[0].ArtistName)
		require.Equal(t, findByIDResponseArtist2.Artist.Name, dtoTracks[1].ArtistName)
		require.Equal(t, findByIDResponseArtist3.Artist.Name, dtoTracks[2].ArtistName)
		require.Equal(t, findByIDResponseAlbum1.Album.Name, dtoTracks[0].AlbumName)
		require.Equal(t, findByIDResponseAlbum2.Album.Name, dtoTracks[1].AlbumName)
		require.Equal(t, findByIDResponseAlbum3.Album.Name, dtoTracks[2].AlbumName)
	}
}

func TestUsecase_GetAll_NotFoundTracks(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumClientMock := mockAlbum.NewMockAlbumServiceClient(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, artistClientMock, albumClientMock, logger)

	ctx := context.Background()
	trackRepoMock.EXPECT().GetAll(ctx).Return(nil, errors.New("Can't load tracks"))
	dtoTracks, err := trackUsecase.GetAll(ctx)

	require.Error(t, err)
	require.Nil(t, dtoTracks)
	require.EqualError(t, err, "Can't load tracks")
}

func TestUsecase_GetAllByArtistID_FoundTracks(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumClientMock := mockAlbum.NewMockAlbumServiceClient(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, artistClientMock, albumClientMock, logger)

	now := time.Now()
	tracks := []*models.Track{
		{
			ID: uint64(1), Name: "test1", Duration: uint64(1), FilePath: "1", Image: "1",
			ArtistID: uint64(1), AlbumID: uint64(1), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(2), Name: "test2", Duration: uint64(2), FilePath: "2", Image: "2",
			ArtistID: uint64(1), AlbumID: uint64(2), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(3), Name: "test3", Duration: uint64(3), FilePath: "3", Image: "3",
			ArtistID: uint64(1), AlbumID: uint64(3), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
	}

	findByIDResponseArtist := &artistService.FindByIDResponse{
		Artist: &artistService.Artist{
			Id:      uint64(1),
			Name:    "artist1",
			Bio:     "1",
			Country: "1",
			Image:   "1",
		},
	}

	findByIDResponseAlbums := []*albumService.FindByIDResponse{
		{
			Album: &albumService.Album{
				Id:          uint64(1),
				Name:        "album1",
				ReleaseDate: timestamppb.New(now),
				Image:       "1",
				ArtistID:    uint64(1),
			},
		},
		{
			Album: &albumService.Album{
				Id:          uint64(2),
				Name:        "album2",
				ReleaseDate: timestamppb.New(now),
				Image:       "2",
				ArtistID:    uint64(1),
			},
		},
		{
			Album: &albumService.Album{
				Id:          uint64(3),
				Name:        "album3",
				ReleaseDate: timestamppb.New(now),
				Image:       "3",
				ArtistID:    uint64(1),
			},
		},
	}

	ctx := context.Background()
	for i := 0; i < len(findByIDResponseAlbums); i++ {
		artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: uint64(1)}).Return(findByIDResponseArtist, nil)
		albumClientMock.EXPECT().FindByID(ctx, &albumService.FindByIDRequest{Id: findByIDResponseAlbums[i].Album.Id}).Return(findByIDResponseAlbums[i], nil)
	}
	trackRepoMock.EXPECT().GetAllByArtistID(ctx, uint64(1)).Return(tracks, nil)

	dtoTracks, err := trackUsecase.GetAllByArtistID(ctx, uint64(1))

	require.NoError(t, err)
	require.NotNil(t, dtoTracks)
	require.Equal(t, len(tracks), len(dtoTracks))

	for i := 0; i < len(tracks); i++ {
		require.Equal(t, tracks[i].Name, dtoTracks[i].Name)
		require.Equal(t, findByIDResponseArtist.Artist.Name, dtoTracks[i].ArtistName)
		require.Equal(t, findByIDResponseAlbums[i].Album.Name, dtoTracks[i].AlbumName)
	}
}

func TestUsecase_GetAllByArtistID_NotFoundTracks(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumClientMock := mockAlbum.NewMockAlbumServiceClient(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, artistClientMock, albumClientMock, logger)

	ctx := context.Background()
	trackRepoMock.EXPECT().GetAllByArtistID(ctx, uint64(1)).Return(nil, errors.New("Can't load tracks by artist ID 1"))

	dtoTracks, err := trackUsecase.GetAllByArtistID(ctx, uint64(1))

	require.Error(t, err)
	require.Nil(t, dtoTracks)
	require.EqualError(t, err, "Can't load tracks by artist ID 1")
}

func TestUsecase_GetAllByAlbumID_FoundTracks(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumClientMock := mockAlbum.NewMockAlbumServiceClient(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, artistClientMock, albumClientMock, logger)

	now := time.Now()
	tracks := []*models.Track{
		{
			ID: uint64(1), Name: "test1", Duration: uint64(1), FilePath: "1", Image: "1",
			ArtistID: uint64(1), AlbumID: uint64(1), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(2), Name: "test2", Duration: uint64(2), FilePath: "2", Image: "2",
			ArtistID: uint64(1), AlbumID: uint64(1), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(3), Name: "test3", Duration: uint64(3), FilePath: "3", Image: "3",
			ArtistID: uint64(1), AlbumID: uint64(1), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
	}

	findByIDResponseArtist := &artistService.FindByIDResponse{
		Artist: &artistService.Artist{
			Id:      1,
			Name:    "artist1",
			Bio:     "1",
			Country: "1",
			Image:   "1",
		},
	}

	findByIDResponseAlbum := &albumService.FindByIDResponse{
		Album: &albumService.Album{
			Id:          1,
			Name:        "album1",
			ReleaseDate: timestamppb.New(now),
			Image:       "1",
			ArtistID:    1,
		},
	}

	ctx := context.Background()
	trackRepoMock.EXPECT().GetAllByAlbumID(ctx, uint64(1)).Return(tracks, nil)
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: tracks[0].ArtistID}).Return(findByIDResponseArtist, nil)
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: tracks[1].ArtistID}).Return(findByIDResponseArtist, nil)
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: tracks[2].ArtistID}).Return(findByIDResponseArtist, nil)
	albumClientMock.EXPECT().FindByID(ctx, &albumService.FindByIDRequest{Id: tracks[0].AlbumID}).Return(findByIDResponseAlbum, nil)
	albumClientMock.EXPECT().FindByID(ctx, &albumService.FindByIDRequest{Id: tracks[1].AlbumID}).Return(findByIDResponseAlbum, nil)
	albumClientMock.EXPECT().FindByID(ctx, &albumService.FindByIDRequest{Id: tracks[2].AlbumID}).Return(findByIDResponseAlbum, nil)

	dtoTracks, err := trackUsecase.GetAllByAlbumID(ctx, uint64(1))

	require.NoError(t, err)
	require.NotNil(t, dtoTracks)
	require.Equal(t, len(tracks), len(dtoTracks))

	for i := 0; i < len(tracks); i++ {
		require.Equal(t, tracks[i].Name, dtoTracks[i].Name)
		require.Equal(t, findByIDResponseArtist.Artist.Name, dtoTracks[i].ArtistName)
		require.Equal(t, findByIDResponseAlbum.Album.Name, dtoTracks[i].AlbumName)
	}
}

func TestUsecase_GetAllByAlbumID_NotFoundTracks(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumClientMock := mockAlbum.NewMockAlbumServiceClient(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, artistClientMock, albumClientMock, logger)

	ctx := context.Background()
	trackRepoMock.EXPECT().GetAllByAlbumID(ctx, uint64(1)).Return(nil, errors.New("Can't load tracks by album ID 1"))

	dtoTracks, err := trackUsecase.GetAllByAlbumID(ctx, uint64(1))

	require.Error(t, err)
	require.Nil(t, dtoTracks)
	require.EqualError(t, err, "Can't load tracks by album ID 1")
}

func TestUsecase_AddFavoriteTrack(t *testing.T) {
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

	mockTrackRepo := mockTrack.NewMockRepo(ctrl)
	logger := logger.New(&cfg.Service.Logger)

	trackUsecase := &trackUsecase{
		trackRepo: mockTrackRepo,
		logger:    logger,
	}

	userID := uuid.New()
	trackID := uint64(12345)
	requestID := "request-id"
	ctx := context.WithValue(context.Background(), utils.RequestIDKey{}, requestID)

	t.Run("success", func(t *testing.T) {
		mockTrackRepo.EXPECT().AddFavoriteTrack(ctx, userID, trackID).Return(nil)
		err := trackUsecase.AddFavoriteTrack(ctx, userID, trackID)
		require.NoError(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		mockError := fmt.Errorf("repository error")
		mockTrackRepo.EXPECT().AddFavoriteTrack(ctx, userID, trackID).Return(mockError)

		err := trackUsecase.AddFavoriteTrack(ctx, userID, trackID)

		require.Error(t, err)
		require.Contains(t, err.Error(), "repository error")
	})
}

func TestUsecase_DeleteFavoriteTrack(t *testing.T) {
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

	mockTrackRepo := mockTrack.NewMockRepo(ctrl)
	logger := logger.New(&cfg.Service.Logger)

	trackUsecase := &trackUsecase{
		trackRepo: mockTrackRepo,
		logger:    logger,
	}

	userID := uuid.New()
	trackID := uint64(12345)
	requestID := "request-id"
	ctx := context.WithValue(context.Background(), utils.RequestIDKey{}, requestID)

	t.Run("success", func(t *testing.T) {
		mockTrackRepo.EXPECT().DeleteFavoriteTrack(ctx, userID, trackID).Return(nil)
		err := trackUsecase.DeleteFavoriteTrack(ctx, userID, trackID)
		require.NoError(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		mockError := fmt.Errorf("repository error")
		mockTrackRepo.EXPECT().DeleteFavoriteTrack(ctx, userID, trackID).Return(mockError)
		err := trackUsecase.DeleteFavoriteTrack(ctx, userID, trackID)
		require.Error(t, err)
		require.Contains(t, err.Error(), "repository error")
	})
}

func TestUsecase_IsFavoriteTrack(t *testing.T) {
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

	mockTrackRepo := mockTrack.NewMockRepo(ctrl)
	logger := logger.New(&cfg.Service.Logger)

	trackUsecase := &trackUsecase{
		trackRepo: mockTrackRepo,
		logger:    logger,
	}

	userID := uuid.New()
	trackID := uint64(12345)
	requestID := "request-id"
	ctx := context.WithValue(context.Background(), utils.RequestIDKey{}, requestID)

	t.Run("success", func(t *testing.T) {
		mockTrackRepo.EXPECT().IsFavoriteTrack(ctx, userID, trackID).Return(true, nil)
		exists, err := trackUsecase.IsFavoriteTrack(ctx, userID, trackID)
		require.NoError(t, err)
		require.True(t, exists)
	})

	t.Run("track not found", func(t *testing.T) {
		mockTrackRepo.EXPECT().IsFavoriteTrack(ctx, userID, trackID).Return(false, nil)
		exists, err := trackUsecase.IsFavoriteTrack(ctx, userID, trackID)
		require.NoError(t, err)
		require.False(t, exists)
	})

	t.Run("repository error", func(t *testing.T) {
		mockError := fmt.Errorf("repository error")
		mockTrackRepo.EXPECT().IsFavoriteTrack(ctx, userID, trackID).Return(false, mockError)
		exists, err := trackUsecase.IsFavoriteTrack(ctx, userID, trackID)
		require.Error(t, err)
		require.Contains(t, err.Error(), "repository error")
		require.False(t, exists)
	})
}

func TestUsecase_GetFavoriteTracks_FoundTracks(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumClientMock := mockAlbum.NewMockAlbumServiceClient(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, artistClientMock, albumClientMock, logger)

	now := time.Now()
	tracks := []*models.Track{
		{
			ID: uint64(1), Name: "test1", Duration: uint64(1), FilePath: "1", Image: "1",
			ArtistID: uint64(1), AlbumID: uint64(1), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(2), Name: "test2", Duration: uint64(2), FilePath: "2", Image: "2",
			ArtistID: uint64(1), AlbumID: uint64(2), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(3), Name: "test3", Duration: uint64(3), FilePath: "3", Image: "3",
			ArtistID: uint64(1), AlbumID: uint64(3), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
	}

	findByIDResponseArtist := &artistService.FindByIDResponse{
		Artist: &artistService.Artist{
			Id:      1,
			Name:    "artist1",
			Bio:     "1",
			Country: "1",
			Image:   "1",
		},
	}

	findByIDResponseAlbums := []*albumService.FindByIDResponse{
		{
			Album: &albumService.Album{
				Id:          1,
				Name:        "album1",
				ReleaseDate: timestamppb.New(now),
				Image:       "1",
				ArtistID:    1,
			},
		},
		{
			Album: &albumService.Album{
				Id:          2,
				Name:        "album2",
				ReleaseDate: timestamppb.New(now),
				Image:       "2",
				ArtistID:    1,
			},
		},
		{
			Album: &albumService.Album{
				Id:          3,
				Name:        "album3",
				ReleaseDate: timestamppb.New(now),
				Image:       "3",
				ArtistID:    1,
			},
		},
	}

	userID := uuid.New()
	ctx := context.Background()
	trackRepoMock.EXPECT().GetFavoriteTracks(ctx, userID).Return(tracks, nil)
	for i := 0; i < len(tracks); i++ {
		artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: tracks[i].ArtistID}).Return(findByIDResponseArtist, nil)
		albumClientMock.EXPECT().FindByID(ctx, &albumService.FindByIDRequest{Id: tracks[i].AlbumID}).Return(findByIDResponseAlbums[i], nil)
	}

	dtoTracks, err := trackUsecase.GetFavoriteTracks(ctx, userID)

	require.NoError(t, err)
	require.NotNil(t, dtoTracks)
	require.Equal(t, len(tracks), len(dtoTracks))

	for i := 0; i < len(tracks); i++ {
		require.Equal(t, tracks[i].Name, dtoTracks[i].Name)
		require.Equal(t, findByIDResponseArtist.Artist.Name, dtoTracks[i].ArtistName)
		require.Equal(t, findByIDResponseAlbums[i].Album.Name, dtoTracks[i].AlbumName)
	}
}

func TestUsecase_GetFavoriteTracks_NotFoundTracks(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumClientMock := mockAlbum.NewMockAlbumServiceClient(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, artistClientMock, albumClientMock, logger)

	userID := uuid.New()
	ctx := context.Background()
	trackRepoMock.EXPECT().GetFavoriteTracks(ctx, userID).Return(nil, errors.New(fmt.Sprintf("Can't load tracks by user ID %v", userID)))

	dtoTracks, err := trackUsecase.GetFavoriteTracks(ctx, userID)

	require.Error(t, err)
	require.Nil(t, dtoTracks)
	require.EqualError(t, err, fmt.Sprintf("Can't load tracks by user ID %v", userID))
}

func TestUsecase_GetFavoriteTracksCount_FoundTracks(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumClientMock := mockAlbum.NewMockAlbumServiceClient(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, artistClientMock, albumClientMock, logger)

	userID := uuid.New()
	ctx := context.Background()
	expectedCount := uint64(10)
	trackRepoMock.EXPECT().GetFavoriteTracksCount(ctx, userID).Return(expectedCount, nil)

	count, err := trackUsecase.GetFavoriteTracksCount(ctx, userID)

	require.NoError(t, err)
	require.Equal(t, expectedCount, count)
}

func TestUsecase_GetFavoriteTracksCount_ErrorGettingCount(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumClientMock := mockAlbum.NewMockAlbumServiceClient(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, artistClientMock, albumClientMock, logger)

	userID := uuid.New()
	ctx := context.Background()
	expectedError := fmt.Errorf("Can't load tracks by user ID %v", userID)
	trackRepoMock.EXPECT().GetFavoriteTracksCount(ctx, userID).Return(uint64(0), expectedError)

	count, err := trackUsecase.GetFavoriteTracksCount(ctx, userID)

	require.Error(t, err)
	require.EqualError(t, err, expectedError.Error())
	require.Equal(t, uint64(0), count)
}

func TestUsecase_GetTracksFromPlaylist_FoundTracks(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumClientMock := mockAlbum.NewMockAlbumServiceClient(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, artistClientMock, albumClientMock, logger)

	now := time.Now()
	playlistTracks := []*models.PlaylistTrack{
		{
			ID:         1,
			PlaylistID: 1,
			TrackID:    1,
		},
		{
			ID:         2,
			PlaylistID: 1,
			TrackID:    2,
		},
		{
			ID:         3,
			PlaylistID: 1,
			TrackID:    3,
		},
	}

	tracks := []*models.Track{
		{
			ID: uint64(1), Name: "test1", Duration: uint64(1), FilePath: "1", Image: "1",
			ArtistID: uint64(1), AlbumID: uint64(1), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(2), Name: "test2", Duration: uint64(2), FilePath: "2", Image: "2",
			ArtistID: uint64(1), AlbumID: uint64(1), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(3), Name: "test3", Duration: uint64(3), FilePath: "3", Image: "3",
			ArtistID: uint64(1), AlbumID: uint64(1), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
	}

	findByIDResponseArtist := &artistService.FindByIDResponse{
		Artist: &artistService.Artist{
			Id:      1,
			Name:    "artist1",
			Bio:     "1",
			Country: "1",
			Image:   "1",
		},
	}

	findByIDResponseAlbum := &albumService.FindByIDResponse{
		Album: &albumService.Album{
			Id:          1,
			Name:        "album1",
			ReleaseDate: timestamppb.New(now),
			Image:       "1",
			ArtistID:    1,
		},
	}

	ctx := context.Background()
	trackRepoMock.EXPECT().GetTracksFromPlaylist(ctx, uint64(1)).Return(playlistTracks, nil)
	trackRepoMock.EXPECT().FindById(ctx, uint64(1)).Return(tracks[0], nil)
	trackRepoMock.EXPECT().FindById(ctx, uint64(2)).Return(tracks[1], nil)
	trackRepoMock.EXPECT().FindById(ctx, uint64(3)).Return(tracks[2], nil)
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: tracks[0].ArtistID}).Return(findByIDResponseArtist, nil)
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: tracks[1].ArtistID}).Return(findByIDResponseArtist, nil)
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: tracks[2].ArtistID}).Return(findByIDResponseArtist, nil)
	albumClientMock.EXPECT().FindByID(ctx, &albumService.FindByIDRequest{Id: tracks[0].AlbumID}).Return(findByIDResponseAlbum, nil)
	albumClientMock.EXPECT().FindByID(ctx, &albumService.FindByIDRequest{Id: tracks[1].AlbumID}).Return(findByIDResponseAlbum, nil)
	albumClientMock.EXPECT().FindByID(ctx, &albumService.FindByIDRequest{Id: tracks[2].AlbumID}).Return(findByIDResponseAlbum, nil)

	dtoTracks, err := trackUsecase.GetTracksFromPlaylist(ctx, uint64(1))

	require.NoError(t, err)
	require.NotNil(t, dtoTracks)
	require.Equal(t, len(tracks), len(dtoTracks))

	for i := 0; i < len(tracks); i++ {
		require.Equal(t, tracks[i].Name, dtoTracks[i].Name)
		require.Equal(t, findByIDResponseArtist.Artist.Name, dtoTracks[i].ArtistName)
		require.Equal(t, findByIDResponseAlbum.Album.Name, dtoTracks[i].AlbumName)
	}
}

func TestUsecase_GetTracksFromPlaylist_NotFoundTracks(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumClientMock := mockAlbum.NewMockAlbumServiceClient(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, artistClientMock, albumClientMock, logger)

	ctx := context.Background()
	trackRepoMock.EXPECT().GetTracksFromPlaylist(ctx, uint64(1)).Return(nil, errors.New("Can't load tracks from playlist 1"))

	dtoTracks, err := trackUsecase.GetTracksFromPlaylist(ctx, uint64(1))

	require.Error(t, err)
	require.Nil(t, dtoTracks)
	require.EqualError(t, err, "Can't load tracks from playlist 1")
}

func TestTrackUsecaseGetPopular(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumClientMock := mockAlbum.NewMockAlbumServiceClient(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, artistClientMock, albumClientMock, logger)

	now := time.Now()
	tracks := []*models.Track{
		{
			ID: uint64(1), Name: "test1", Duration: uint64(1), FilePath: "1", Image: "1",
			ArtistID: uint64(1), AlbumID: uint64(1), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
		{
			ID: uint64(2), Name: "test2", Duration: uint64(2), FilePath: "2", Image: "2",
			ArtistID: uint64(1), AlbumID: uint64(1), ReleaseDate: now, CreatedAt: now, UpdatedAt: now,
		},
	}

	findByIDResponseArtist := &artistService.FindByIDResponse{
		Artist: &artistService.Artist{
			Id:      1,
			Name:    "artist1",
			Bio:     "1",
			Country: "1",
			Image:   "1",
		},
	}

	findByIDResponseAlbum := &albumService.FindByIDResponse{
		Album: &albumService.Album{
			Id:          1,
			Name:        "album1",
			ReleaseDate: timestamppb.New(now),
			Image:       "1",
			ArtistID:    1,
		},
	}

	ctx := context.Background()
	trackRepoMock.EXPECT().GetPopular(ctx).Return(tracks, nil)
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: tracks[0].ArtistID}).Return(findByIDResponseArtist, nil)
	artistClientMock.EXPECT().FindByID(ctx, &artistService.FindByIDRequest{Id: tracks[1].ArtistID}).Return(findByIDResponseArtist, nil)
	albumClientMock.EXPECT().FindByID(ctx, &albumService.FindByIDRequest{Id: tracks[0].AlbumID}).Return(findByIDResponseAlbum, nil)
	albumClientMock.EXPECT().FindByID(ctx, &albumService.FindByIDRequest{Id: tracks[1].AlbumID}).Return(findByIDResponseAlbum, nil)

	dtoTracks, err := trackUsecase.GetPopular(ctx)

	require.NoError(t, err)
	require.NotNil(t, dtoTracks)
	require.Equal(t, len(tracks), len(dtoTracks))

	for i := 0; i < len(tracks); i++ {
		require.Equal(t, tracks[i].Name, dtoTracks[i].Name)
		require.Equal(t, findByIDResponseArtist.Artist.Name, dtoTracks[i].ArtistName)
		require.Equal(t, findByIDResponseAlbum.Album.Name, dtoTracks[i].AlbumName)
	}
}

func TestTrackUsecaseGetPopular_Error(t *testing.T) {
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
	trackRepoMock := mockTrack.NewMockRepo(ctrl)
	artistClientMock := mockArtist.NewMockArtistServiceClient(ctrl)
	albumClientMock := mockAlbum.NewMockAlbumServiceClient(ctrl)
	trackUsecase := NewTrackUsecase(trackRepoMock, artistClientMock, albumClientMock, logger)

	ctx := context.Background()
	trackRepoMock.EXPECT().GetPopular(ctx).Return(nil, errors.New("Can't load tracks"))

	dtoTracks, err := trackUsecase.GetPopular(ctx)

	require.Error(t, err)
	require.Nil(t, dtoTracks)
	require.EqualError(t, err, "Can't load tracks")
}
