package domain

const (
	SildeTypeConvertedSlide = "slide"
	SlideTypeQuiz           = "question"
)

type Presentation struct {
	Id        uint64
	CreatorId uint64
	Url       string
	SlideNum  uint32
	QuizNum   uint32
	Slides    []Slide
	Quizzes   []Quiz
}

type Slide struct {
	Name   string
	Idx    uint32
	Width  uint32
	Height uint32
}

type PresHTTP struct {
	Url      string `json:"url"`
	SlideNum uint32 `json:"slideNum"`
	QuizNum  uint32 `json:"quizNum"`
	Slides   []struct {
		Idx        uint32 `json:"idx"`
		Kind       string `json:"kind"`
		QuizId     uint64 `json:"quizId"`
		Name       string `json:"name"`
		Width      uint32 `json:"width"`
		Height     uint32 `json:"height"`
		Type       string `json:"type"`
		Question   string `json:"question"`
		Vote       []Vote `json:"vote"`
		Background string `json:"background"`
		FontColor  string `json:"fontColor"`
		FontSize   string `json:"fontSize"`
		GraphColor string `json:"graphColor"`
	} `json:"slides"`
}

type PresRepository interface {
	GetPres(p *Presentation) error
	CreatePres(p *Presentation) error
}

type PresUsecase interface {
	GetPres(p *Presentation) error
	CreatePres(p *Presentation) error
}
