package main

import (
	"context"
	"fmt"
	"os"

	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/joho/godotenv"
)

type Movie struct {
	ID          int64
	Title       string
	ReleaseDate string
	Overview    string
	PosterPath  string
	Rating      float32
}

type Show struct {
	ID          int64
	Name        string
	ReleaseDate string
	Overview    string
	PosterPath  string
	Rating      float32
	Seasons     []tmdb.Season
}

type Episode struct {
	AirDate        string
	EpisodeNumber  int
	ID             int64
	Name           string
	Overview       string
	ProductionCode string
	Runtime        int
	SeasonNumber   int
	ShowID         int64
	StillPath      string
}

var tmdbClient *tmdb.Client
var err = error(nil)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	a.ctx = ctx
	tmdbClient, err = tmdb.Init(os.Getenv("TMDB_API"))
	if err != nil {
		fmt.Println(err)
	}
}

func (a *App) GetTrendingMovies() ([]Movie, error) {
	trending, err := tmdbClient.GetTrending("movie", "week", nil)
	if err != nil {
		fmt.Println(err)
		return []Movie{}, err
	}
	trendingResults := []Movie{}
	for _, result := range trending.Results {
		poster := tmdb.GetImageURL(result.PosterPath, tmdb.Original)
		if result.PosterPath != "" {
			movie := Movie{result.ID, result.Title, result.ReleaseDate, result.Overview, poster, result.VoteAverage}
			trendingResults = append(trendingResults, movie)
		}
	}
	return trendingResults, nil
}

func (a *App) GetTrendingShows() ([]Show, error) {
	trending, err := tmdbClient.GetTrending("tv", "week", nil)
	if err != nil {
		fmt.Println(err)
		return []Show{}, err
	}
	trendingResults := []Show{}
	for _, result := range trending.Results {
		poster := tmdb.GetImageURL(result.PosterPath, tmdb.Original)
		if result.PosterPath != "" {
			show := Show{result.ID, result.Name, result.FirstAirDate, result.Overview, poster, result.VoteAverage, nil}
			trendingResults = append(trendingResults, show)
		}
	}
	return trendingResults, nil
}

func (a *App) SearchMovie(searchTerm string) ([]Movie, error) {
	options := make(map[string]string)
	options["language"] = "en-US"
	searchResult, err := tmdbClient.GetSearchMovies(searchTerm, options)
	if err != nil {
		fmt.Println(err)
		return []Movie{}, err
	}

	movieResults := []Movie{}
	for _, result := range searchResult.Results {
		poster := tmdb.GetImageURL(result.PosterPath, tmdb.Original)
		if result.PosterPath != "" {
			movie := Movie{result.ID, result.Title, result.ReleaseDate, result.Overview, poster, result.VoteAverage}
			movieResults = append(movieResults, movie)
		}
	}
	return movieResults, nil
}

func (a *App) SearchShow(searchTerm string) ([]Show, error) {
	options := make(map[string]string)
	options["language"] = "en-US"
	searchResult, err := tmdbClient.GetSearchTVShow(searchTerm, options)
	if err != nil {
		fmt.Println(err)
		return []Show{}, err
	}

	showResults := []Show{}
	for _, result := range searchResult.Results {
		poster := tmdb.GetImageURL(result.PosterPath, tmdb.Original)
		if result.PosterPath != "" {
			show := Show{result.ID, result.Name, result.FirstAirDate, result.Overview, poster, result.VoteAverage, nil}
			showResults = append(showResults, show)
		}
	}
	return showResults, nil
}

func (a *App) GetMovieDetails(id int) (Movie, error) {
	options := make(map[string]string)
	options["language"] = "en-US"
	details, err := tmdbClient.GetMovieDetails(id, options)
	if err != nil {
		fmt.Println(err)
		return Movie{}, err
	}
	poster := tmdb.GetImageURL(details.PosterPath, tmdb.Original)
	return Movie{details.ID, details.Title, details.ReleaseDate, details.Overview, poster, details.VoteAverage}, nil
}

func (a *App) GetShowDetails(id int) (Show, error) {
	options := make(map[string]string)
	options["language"] = "en-US"
	details, err := tmdbClient.GetTVDetails(id, options)
	if err != nil {
		fmt.Println(err)
		return Show{}, err
	}
	poster := tmdb.GetImageURL(details.PosterPath, tmdb.Original)
	return Show{details.ID, details.Name, details.FirstAirDate, details.Overview, poster, details.VoteAverage, details.Seasons}, nil
}

func (a *App) GetSeasonEpisodes(id int, seasonNr int) ([]Episode, error) {
	options := make(map[string]string)
	options["language"] = "en-US"
	details, err := tmdbClient.GetTVSeasonDetails(id, seasonNr, options)
	if err != nil {
		fmt.Println(err)
		return []Episode{}, err
	}
	showEpisodes := []Episode{}
	for _, episode := range details.Episodes {
		stillPath := tmdb.GetImageURL(episode.StillPath, tmdb.Original)
		showEpisode := Episode{
			AirDate:        episode.AirDate,
			EpisodeNumber:  episode.EpisodeNumber,
			ID:             episode.ID,
			Name:           episode.Name,
			Overview:       episode.Overview,
			ProductionCode: episode.ProductionCode,
			Runtime:        episode.Runtime,
			StillPath:      stillPath,
		}
		showEpisodes = append(showEpisodes, showEpisode)
	}
	return showEpisodes, nil
}
