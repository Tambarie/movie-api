package domain

type Comment struct {
	Model
	tableName struct{} `pg:"movie_comments"`
	MovieID   int      `json:"movie_id" pg:"movie_id"`
	Content   string   `json:"content" pg:"content"`
	IP        string   `json:"ip" pg:"ip"`
}
