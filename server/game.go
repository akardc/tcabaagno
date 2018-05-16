package tcabaagno

type Game struct {
	Forms []QAForm
}

type Questions struct {
	Who   string
	What  string
	When  string
	Where string
	Why   string
}

type Answers struct {
	Who   string
	What  string
	When  string
	Where string
	Why   string
}

type QAForm struct {
	Questions Questions
	Answers   Answers
	Player    string
}
