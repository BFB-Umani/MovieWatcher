export namespace main {
	
	export class Episode {
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
	
	    static createFrom(source: any = {}) {
	        return new Episode(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.AirDate = source["AirDate"];
	        this.EpisodeNumber = source["EpisodeNumber"];
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.Overview = source["Overview"];
	        this.ProductionCode = source["ProductionCode"];
	        this.Runtime = source["Runtime"];
	        this.SeasonNumber = source["SeasonNumber"];
	        this.ShowID = source["ShowID"];
	        this.StillPath = source["StillPath"];
	    }
	}
	export class Movie {
	    ID: number;
	    Title: string;
	    ReleaseDate: string;
	    Overview: string;
	    PosterPath: string;
	    Rating: number;
	
	    static createFrom(source: any = {}) {
	        return new Movie(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Title = source["Title"];
	        this.ReleaseDate = source["ReleaseDate"];
	        this.Overview = source["Overview"];
	        this.PosterPath = source["PosterPath"];
	        this.Rating = source["Rating"];
	    }
	}
	export class Show {
	    ID: number;
	    Name: string;
	    ReleaseDate: string;
	    Overview: string;
	    PosterPath: string;
	    Rating: number;
	    Seasons: tmdb.Season[];
	
	    static createFrom(source: any = {}) {
	        return new Show(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Name = source["Name"];
	        this.ReleaseDate = source["ReleaseDate"];
	        this.Overview = source["Overview"];
	        this.PosterPath = source["PosterPath"];
	        this.Rating = source["Rating"];
	        this.Seasons = this.convertValues(source["Seasons"], tmdb.Season);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace tmdb {
	
	export class Season {
	    air_date: string;
	    episode_count: number;
	    id: number;
	    name: string;
	    overview: string;
	    poster_path: string;
	    season_number: number;
	    vote_average?: number;
	    show_id?: number;
	
	    static createFrom(source: any = {}) {
	        return new Season(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.air_date = source["air_date"];
	        this.episode_count = source["episode_count"];
	        this.id = source["id"];
	        this.name = source["name"];
	        this.overview = source["overview"];
	        this.poster_path = source["poster_path"];
	        this.season_number = source["season_number"];
	        this.vote_average = source["vote_average"];
	        this.show_id = source["show_id"];
	    }
	}

}

