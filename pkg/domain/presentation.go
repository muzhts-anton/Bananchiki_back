package domain

const (
	SildeTypeConvertedSlide = "slide"
	SlideTypeQuiz           = "question"
	SlideTypeTimedQuiz      = "quiz"
)

const (
	PresentationFilePath   = "/static/presentation/files/"
	PresentationSlidesPath = "/static/presentation/slides/"
)

type Presentation struct {
	Id        uint64
	CreatorId uint64
	Url       string
	Name      string
	Code      string
	Hash      string
	Width     uint32
	Height    uint32
	SlideNum  uint32
	QuizNum   uint32
	Slides    []ConvertedSlide
	Quizzes   []Quiz
}

type ConvertedSlide struct {
	Name   string
	Idx    uint32
	Width  uint32
	Height uint32
}

type SlideApiResponse struct {
	Idx         uint32 `json:"idx"`
	Kind        string `json:"kind"`
	QuizId      uint64 `json:"quizId"`
	Name        string `json:"name"`
	Width       uint32 `json:"width"`
	Height      uint32 `json:"height"`
	Type        string `json:"type"`
	Question    string `json:"question"`
	AnswerTime  uint64 `json:"answerTime"`
	ResultAfter bool   `json:"answerAfter"`
	Cost        uint64 `json:"cost"`
	ExtraPts    bool   `json:"extrapts"`
	Runout      bool   `json:"runout"`
	Vote        []Vote `json:"votes"`
	Background  string `json:"background"`
	FontColor   string `json:"fontColor"`
	FontSize    string `json:"fontSize"`
	GraphColor  string `json:"graphColor"`
}

type PresApiResponse struct {
	Name     string             `json:"name"`
	Code     string             `json:"code"`
	Hash     string             `json:"hash"`
	Width    uint32             `json:"width"`
	Height   uint32             `json:"height"`
	Url      string             `json:"url"`
	SlideNum uint32             `json:"slideNum"`
	QuizNum  uint32             `json:"quizNum"`
	Slides   []SlideApiResponse `json:"slides"`
}

type PresRepository interface {
	GetPres(pid uint64) (Presentation, error)
	GetConvertedSlides(t string, pid uint64) ([]ConvertedSlide, error)
	GetQuizzes(t string, pid uint64) ([]Quiz, error)
	CreatePres(cid uint64) (uint64, error)
	CreateCovertedSlides(pid uint64, slides []ConvertedSlide) error
	UpdatePresUrl(pid uint64, url string) error
	GetPresOwner(pid uint64) (uint64, error)
	ChangePresName(pid uint64, name string) error
	DeletePres(pid uint64) error
}

type PresUsecase interface {
	GetPres(cid, pid uint64) (PresApiResponse, error)
	CreatePres(url string, cid uint64) (uint64, error)
	ChangePresName(uid, pid uint64, name string) error
	DeletePres(uid, pid uint64) error
}
