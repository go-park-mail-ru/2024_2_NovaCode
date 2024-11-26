package server

// import (
// 	"context"
// 	"fmt"
// 	"strconv"

// 	"/internal/microservices/pkg/album/usecase"
// 	"/internal/microservices/album/proto"
// 	converters "/internal/microservices/models/proto_converters"
// )

// type AlbumServer struct {
// 	proto.UnimplementedAlbumServiceServer
// 	AlbumUseCase usecase.AlbumUseCase
// }

// func NewAlbumServer(albumUseCase usecase.AlbumUseCase) *AlbumServer {
// 	return &AlbumServer{AlbumUseCase: albumUseCase}
// }

// func (as *AlbumServer) GetAlbumByID(ctx context.Context, input *proto.AlbumIDAndLogin) (*proto.Album, error) {
// 	if input.Id <= 0 {
// 		return nil, fmt.Errorf("invalid album id: %s", strconv.Itoa(int(input.Id)))
// 	}

// 	album, err := as.AlbumUseCase.GetAlbumByID(input.Id, input.Login, ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("album not found")
// 	}

// 	return converters.AlbumConvertCoreInProto(*album), nil
// }

// func (as *AlbumServer) SearchAlbum(ctx context.Context, input *proto.SearchQuery) (*proto.Albums, error) {
// 	if input.Query == "" {
// 		return nil, fmt.Errorf("missing query parameter 'query'")
// 	}

// 	foundAlbums, err := as.AlbumUseCase.Search(ctx, input.Query)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to find albums")
// 	} else if len(foundAlbums) == 0 {
// 		return nil, fmt.Errorf("no albums with %s were found", input.Query)
// 	}

// 	albumsProto := make([]*proto.Album, len(foundAlbums))
// 	for i, a := range foundAlbums {
// 		albumsProto[i] = converters.AlbumConvertCoreInProto(*a)
// 	}

// 	albumProto := new(proto.Albums)
// 	albumProto.Albums = albumsProto
// 	return albumProto, nil
// }

// func (as *AlbumServer) GetAllAlbums(ctx context.Context, input *proto.LoginOffsetLimit) (*proto.Albums, error) {
// 	if input.Login == "" {
// 		return nil, fmt.Errorf("invalid album login: %s", input.Login)
// 	}

// 	albumsCore, err := as.AlbumUseCase.GetAllAlbums(input.Login, input.Offset, input.Limit, ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("albums not found")
// 	}

// 	albumsProto := make([]*proto.Album, len(albumsCore))
// 	for i, a := range albumsCore {
// 		albumsProto[i] = converters.AlbumConvertCoreInProto(*a)
// 	}

// 	albumProto := new(proto.Albums)
// 	albumProto.Albums = albumsProto
// 	return albumProto, nil
// }

// func (as *AlbumServer) GetAllByArtistID(ctx context.Context, input *proto.ArtistID) (*proto.Albums, error) {
// 	if input.ArtistID <= 0 {
// 		return nil, fmt.Errorf("invalid artist ID: %s", strconv.Itoa(int(input.ArtistID)))
// 	}

// 	albums, err := as.AlbumUseCase.GetAllByArtistID(ctx, input.ArtistID)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to load albums by artist ID")
// 	} else if len(albums) == 0 {
// 		return nil, fmt.Errorf("no albums found for the artist")
// 	}

// 	albumsProto := make([]*proto.Album, len(albums))
// 	for i, a := range albums {
// 		albumsProto[i] = converters.AlbumConvertCoreInProto(*a)
// 	}

// 	albumProto := new(proto.Albums)
// 	albumProto.Albums = albumsProto
// 	return albumProto, nil
// }