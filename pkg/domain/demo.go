package domain

type PresEmotions struct {
	Like     uint64 `json:"like"`
	Love     uint64 `json:"love"`
	Laughter uint64 `json:"laughter"`
	Surprise uint64 `json:"surprise"`
	Sad      uint64 `json:"sad"`
}

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
}
