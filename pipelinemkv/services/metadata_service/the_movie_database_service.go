package metadataservice

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	detailsBaseUrl   = "https://api.themoviedb.org/3/"
	imageBaseUrl     = "https://image.tmdb.org/t/p/"
	searchMoviePath  = "/search/movie"
	searchTvPath     = "/search/tv"
	movieDetailsPath = "/movie/"
)

type MetadataService struct {
	authToken string
}

func New(authToken string) *MetadataService {
	return &MetadataService{
		authToken: authToken,
	}
}

func (s MetadataService) SearchMovie(ctx context.Context, title, release_year, language string) {
	u, err := url.Parse(detailsBaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	u = u.JoinPath(searchMoviePath)

	queryParams := u.Query()

	queryParams.Add("query", title)
	if release_year != "" {
		queryParams.Add("primary_release_year", release_year)
	}
	if language != "" {
		queryParams.Add("language", language)

	}
	u.RawQuery = queryParams.Encode()

	sendRequest(ctx, u, s.authToken)
}

func (s MetadataService) GetMovieDetails(ctx context.Context, movieId string, appendToResponse []string) {
	u, err := url.Parse(detailsBaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	u = u.JoinPath(fmt.Sprintf("/%s/%s", movieDetailsPath, movieId))
	appendAppendToResponseToUrl(u, appendToResponse)

	sendRequest(ctx, u, s.authToken)
}

// Documentation: https://developer.themoviedb.org/docs/image-basics
// TODO: consider implementing straight in the frontend
func (s MetadataService) GetPoster(ctx context.Context, imagePath string, imageSize string) {
	u, err := url.Parse(imageBaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	if imageSize == "original" || imageSize == "" {
		u = u.JoinPath("/original/")
	} else {
		var imgPath string
		if imageSize[0] == 'w' {
			imgPath = fmt.Sprintf("/%s/", imageSize)
		} else {
			imgPath = fmt.Sprintf("/w%s/", imageSize)
		}
		u = u.JoinPath(imgPath)
	}
	u = u.JoinPath(fmt.Sprintf("/%s", imagePath))

	sendRequest(ctx, u, s.authToken)

}

func appendAppendToResponseToUrl(u *url.URL, appendToResponse []string) {
	if len(appendToResponse) > 0 {
		queryParams := u.Query()
		queryParams.Add("append_to_response", strings.Join(appendToResponse, ","))
		u.RawQuery = queryParams.Encode()
	}
}

func sendRequest(ctx context.Context, u *url.URL, authToken string) {
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	fmt.Println(buf.String())
}
