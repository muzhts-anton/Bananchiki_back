package domain

type Presentation struct {
	Id        uint64
	CreatorId uint64
	Url       string
}

type Slides struct {
	
	Url string
	Num uint64
}

type PresRepository interface {
	GetPres(p *Presentation) error
	UploadPres(p *Presentation) error
}

type PresUsecase interface {
	GetPres(p *Presentation) error
	UploadPres(p *Presentation) error
}
