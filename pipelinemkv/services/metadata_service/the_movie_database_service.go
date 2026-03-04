package metadataservice

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const (
	detailsBaseUrl  = "https://api.themoviedb.org/3/"
	imageBaseUrl    = "https://image.tmdb.org/"
	searchMoviePath = "/search/movie"
	searchTvPath    = "/search/tv"
	movieDetailsUrl = "https://api.themoviedb.org/3/movie/"
)

type MetadataService struct {
	// the endpoint for when you already
	authToken string
}

func New(authToken string) *MetadataService {
	return &MetadataService{
		authToken: authToken,
	}
}

func (s MetadataService) SearchMovie(title, release_year, language string, ctx context.Context) {
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

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.authToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	fmt.Println(buf.String())

}

func (s MetadataService) GetMovieDetails(movieId string) {

}
