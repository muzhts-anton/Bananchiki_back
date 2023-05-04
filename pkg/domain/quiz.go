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
	Id          uint64 `json:"quizId"`
	Idx         uint32 `json:"idx"`
	Type        string `json:"type"`
	Question    string `json:"question"`
	AnswerTime  uint64 `json:"answerTime"`
	ResultAfter bool   `json:"answerAfter"`
	Cost        uint64 `json:"cost"`
	ExtraPts    bool   `json:"extrapts"`
	Runout      bool   `json:"runout"`
	Votes       []Vote `json:"votes"`
	Background  string `json:"background"`
	FontColor   string `json:"fontColor"`
	FontSize    string `json:"fontSize"`
	GraphColor  string `json:"graphColor"`
}

type QuizHTTP struct {
	CreatorId uint64 `json:"creatoreId"`
	PresId    uint64 `json:"presId"`
	Quiz
}

type Vote struct {
	Idx     uint32 `json:"idx"`
	Option  string `json:"option"`
	Correct bool   `json:"correct"`
	Votes   uint64 `json:"votes"`
	Color   string `json:"color"`
}

type VoteHTTP struct {
	CreatorId uint64 `json:"creatoreId"`
	QuizId    uint64 `json:"quizId"`
	VoterName string `json:"user"`
	VoterId   uint64 `json:"userId"`
	Vote
}

type CompetitionHttp struct {
	QuizId uint64 `json:"quizId"`
	PresId uint64 `json:"presId"`
}

type VoteRegisterHttp struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
}

type ResultItem struct {
	Id     uint64 `json:"id"`
	Name   string `json:"name"`
	Points uint64 `json:"points"`
}

type QuizRepository interface {
	CreateQuiz(q Quiz, pid uint64) (uint64, error)
	DeleteQuiz(qid, pid uint64) error
	UpdateQuiz(q Quiz, pid uint64) error
	CreateQuizVote(q Vote, qid uint64) error
	UpdateQuizVote(q Vote, qid uint64) error
	DeleteQuizVote(idx uint32, qid uint64) error
	PollQuizVote(idx uint32, qid uint64) error
	PollQuizVoteTracked(idx uint32, qid uint64, vid uint64) error
	CalculatePoints(idx uint32, qid, vid uint64) error
	FinishCompetition(qid uint64) error

	CompetitionStart(quizId uint64, presId uint64) error
	CompetitionStop(quizId uint64, presId uint64) error
	CompetitionVoterRegister(name string, presId uint64) (uint64, error)
	GetPresIdByHash(hash string) (uint64, error)
	GetCompetitionResult(presId uint64) ([]ResultItem, error)
	SetCompetitionResult(presId uint64) error
}

type QuizUsecase interface {
	CreateQuiz(q Quiz, cid, pid uint64) (uint64, error)
	DeleteQuiz(qid, cid, pid uint64) error
	UpdateQuiz(q Quiz, cid, pid uint64) error
	CreateQuizVote(q Vote, qid, cid uint64) error
	UpdateQuizVote(q []Vote, qid, cid uint64) error
	DeleteQuizVote(idx uint32, qid, cid uint64) error
	PollQuizVote(idx uint32, qid uint64, votername string, vid uint64) error

	CompetitionStart(quizId uint64, presId uint64) error
	CompetitionStop(quizId uint64, presId uint64) error
	CompetitionVoterRegister(name string, hash string) (uint64, error)
	GetCompetitionResult(presId uint64) ([]ResultItem, error)
}
