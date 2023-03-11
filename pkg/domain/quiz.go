package domain

const (
	QuizType1 = iota
	QuizType2
	QuizType3
)

type Quiz struct {
	Id        uint64            `json:"quizId,omitempty"`
	CreatorId uint64            `json:"creatorId,omitempty"`
	Type      int               `json:"type,omitempty"`
	Question  string            `json:"question,omitempty"`
	Vote      map[string]uint64 `json:"vote,omitempty"`
}

type QuizRepository interface {
	CreateQuiz(q *Quiz) error
	VoteQuiz(q *Quiz) error
	ShowQuiz(q *Quiz) error
}

type QuizUsecase interface {
	CreateQuiz(q *Quiz) error
	VoteQuiz(q *Quiz) error
	ShowQuiz(q *Quiz) error
}
