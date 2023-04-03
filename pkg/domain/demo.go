package domain

type CurrentDemoSlide struct {
	ViewMode bool `json:"viewMode"`
	Width uint32 `json:"width"`
	Height uint32 `json:"height"`
	Slide SlideApiResponse `json:"slide"`
}

type DemoUsecase interface {
	GetCurrentDemoSlide(hash string) (CurrentDemoSlide, error)
}

type DemoRepository interface {
	GetPresIdByHash(hash string) (uint64, error)
	GetPresViewMode(pid uint64) (bool, error)
	GetCurrentDemoSlide(pid uint64) (SlideApiResponse, error)
}
