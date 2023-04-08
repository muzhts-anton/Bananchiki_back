package domain

type PresEmotions struct {
	Like     uint64 `json:"like"`
	Love     uint64 `json:"love"`
	Laughter uint64 `json:"laughter"`
	Surprise uint64 `json:"surprise"`
	Sad      uint64 `json:"sad"`
}

type ReactionsApi struct {
	PresHash string       `json:"hash"`
	Emotions PresEmotions `json:"emotions"`
}

type Question struct {
	Idx    uint64 `json:"idx"`
	Option string `json:"question"`
	Likes  uint64 `json:"likes"`
}

type QuestionApi struct {
	Hash   string   `json:"hash"`
	Option Question `json:"question"`
}

type ReacUsecase interface {
	ReactionsUpd(emo ReactionsApi) error
	QuestionAsk(h string, q Question) error
}

type ReacRepository interface {
	GetPresIdByHash(hash string) (uint64, error)
	ReactionsUpd(pid uint64, emo PresEmotions) error
	QuestionAsk(pid uint64, q Question) error
}
