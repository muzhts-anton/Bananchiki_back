package domain

type CurrentDemoSlide struct {
	ViewMode bool             `json:"viewMode"`
	Width    uint32           `json:"width"`
	Height   uint32           `json:"height"`
	Url      string           `json:"url"`
	Emotions PresEmotions     `json:"emotions"`
	Slide    SlideApiResponse `json:"slide"`
}

type DemoUsecase interface {
	GetCurrentDemoSlide(hash string) (CurrentDemoSlide, error)
	ShowDemoGo(presId, userId uint64, idx uint32) error
	ShowDemoSop(presId, userId uint64) error
}

type DemoRepository interface {
	GetPresIdByHash(hash string) (uint64, error)
	GetPresViewMode(pid uint64) (bool, error)
	GetCurrentDemoSlide(pid uint64) (SlideApiResponse, error)
	GetPresCreatorId(pid uint64) (uint64, error)
	DemoGo(pid uint64, idx uint32) error
	DemoStop(pid uint64) error
	GetPresEmotions(pid uint64) (PresEmotions, error)
	ZeroingReactions(pid uint64) error // обнуление презентаций
	SetAllVotes(pid uint64, value int) error // выставить все votes равному value
}
