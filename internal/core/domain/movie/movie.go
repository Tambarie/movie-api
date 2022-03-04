package domain

type Movie struct {
	EpisodeId    int    `json:"episode_id"`
	Title        string `json:"title"`
	OpeningCrawl string `json:"opening_crawl"`
	CommentCount int64  `json:"comment_count"`
	ReleaseDate  string `json:"release_date"`
}

type Movies struct {
	Results []Movie `json:"results"`
}
