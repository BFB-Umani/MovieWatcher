type Movie = {
    ID: number;
    Title: string;
    ReleaseDate: string;
    Overview: string;
    PosterPath: string;
    Rating: number
}

type Show = {
    ID: number;
    Name: string;
    ReleaseDate: string;
    Overview: string;
    PosterPath: string;
    Rating: number
    NumberOfSeasons?: number
    Seasons?: Season[]
}

type CardID = {
    ID: string;
    index: number;
}

type Season = {
    air_date: string;
    episode_count: number;
    id: number;
    name: string;
    overview: string;
    poster_path: string;
    season_number: number;
    vote_average?: number;
    show_id?: number;
}

type Episode = {
    AirDate: string;
    EpisodeNumber: number;
    ID: number;
    Name: string;
    Overview: string;
    ProductionCode: string;
    Runtime: number;
    SeasonNumber: number;
    ShowID: number;
    StillPath: string;
}
