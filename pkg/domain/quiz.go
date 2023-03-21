package domain

const (
	QuizTypeNone       = ""
	QuizTypeHorizontal = "horizontal"
	QuizTypeVertical   = "vertical"
	QuizTypePie        = "pie"
	QuizTypeCloud      = "cloud"
	QuizTypeDoughnut   = "doughnut"
)

type Quiz struct {
	Id         uint64 `json:"quizId"`
	Idx        uint32 `json:"idx"`
	Type       string `json:"type"`
	Question   string `json:"question"`
	Votes      []Vote `json:"votes"`
	Background string `json:"background"`
	FontColor  string `json:"fontColor"`
	FontSize   uint32 `json:"fontSize"`
	GraphColor string `json:"graphColor"`
}

type QuizHTTP struct {
	CreatorId uint64 `json:"creatoreId"`
	PresId    uint64 `json:"presId"`
	Quiz
}

type Vote struct {
	Idx    uint32 `json:"idx"`
	Option string `json:"option"`
	Votes  uint64 `json:"votes"`
	Color  string `json:"color"`
}

type VoteHTTP struct {
	CreatorId uint64 `json:"creatoreId"`
	QuizId    uint64 `json:"quizId"`
	Vote
}

type QuizRepository interface {
	CreateQuiz(q Quiz, pid uint64) (uint64, error)
	DeleteQuiz(qid, pid uint64) error
	UpdateQuiz(q Quiz, pid uint64) error
	CreateQuizVote(q Vote, qid uint64) error
	UpdateQuizVote(q Vote, qid uint64) error
	DeleteQuizVote(idx uint32, qid uint64) error
}

type QuizUsecase interface {
	CreateQuiz(q Quiz, cid, pid uint64) (uint64, error)
	DeleteQuiz(qid, cid, pid uint64) error
	UpdateQuiz(q Quiz, cid, pid uint64) error
	CreateQuizVote(q Vote, qid, cid uint64) error
	UpdateQuizVote(q []Vote, qid, cid uint64) error
	DeleteQuizVote(idx uint32, qid, cid uint64) error
}
