package model

type Game struct {
	Forms []QAForm
}

type Questions struct {
	Who      string `json:"who"`
	What     string `json:"what"`
	When     string `json:"when"`
	Where    string `json:"where"`
	Why      string `json:"why"`
	PlayerID string `json:"playerId"`
}

type Answers struct {
	Who      string `json:"who"`
	What     string `json:"what"`
	When     string `json:"when"`
	Where    string `json:"where"`
	Why      string `json:"why"`
	PlayerID string `json:"playerId"`
}

type QAForm struct {
	Questions Questions
	Answers   Answers
}
