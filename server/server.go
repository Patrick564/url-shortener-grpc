package server

import (
	"context"

	pb "github.com/Patrick564/url-shortener-grpc/url_shortener"
)

type UrlShortenerServer struct {
	pb.UnimplementedUrlSortenerServer
}

func (u *UrlShortenerServer) AddUrl(ctx context.Context, in *pb.Url) (*pb.ShortUrl, error) {
	return &pb.ShortUrl{ShortUrl: "example short url response"}, nil
}

func (u *UrlShortenerServer) ListUrls(in *pb.Empty, url pb.UrlSortener_ListUrlsServer) error {
	return nil
}

func (u *UrlShortenerServer) GetUrl(ctx context.Context, in *pb.ShortUrl) (*pb.Url, error) {
	return &pb.Url{Url: "example url original"}, nil
}

func NewServer() *UrlShortenerServer {
	return &UrlShortenerServer{}
}
