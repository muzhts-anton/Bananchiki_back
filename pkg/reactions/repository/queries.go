package reacrep

const (
	queryGetAllPres = `
	SELECT id, code FROM presentation;
	`

	queryUpdEmotions = `
	UPDATE presentation
	SET
		emotion_like = $1,
		emotion_love = $2,
		emotion_laughter = $3,
		emotion_surprise = $4,
		emotion_sad = $5
	WHERE id = $6;
	`
)
