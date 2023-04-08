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

type ReacUsecase interface {
	ReactionsUpd(emo ReactionsApi) error
}

type ReacRepository interface {
	GetPresIdByHash(hash string) (uint64, error)
	ReactionsUpd(pid uint64, emo PresEmotions) error
}
