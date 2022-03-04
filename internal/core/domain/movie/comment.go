package domain

type Comment struct {
	Model
	tableName struct{} `pg:"movie_comments"`
	ID        uint     `gorm:"primaryKey;autoIncrement;"`
	MovieId   int      `json:"movie_id" pg:"movie_id"`
	IP        string   `json:"ip" pg:"ip"`
	Content   string   `json:"content" pg:"content"`
}
