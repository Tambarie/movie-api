package domain

type Character struct {
	Model
	Name      string `json:"name"`
	Height    string `json:"height"`
	Mass      string `json:"mass"`
	HairColor string `json:"hair_color"`
	SkinColor string `json:"skin_color"`
	EyeColor  string `json:"eye_color"`
	BirthYear string `json:"birth_year"`
	Gender    string `json:"gender"`
}

type CharacterLinks struct {
	Characters []string `json:"characters"`
}
