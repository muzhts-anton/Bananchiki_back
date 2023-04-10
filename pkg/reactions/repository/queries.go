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

	queryQuestionAsk = `
	INSERT INTO question (presentation_id, idx, option, likes) VALUES ($1, $2, $3, $4);
	`

	queryQuestionLike = `
	UPDATE question SET likes = likes + 1 WHERE presentation_id = $1 AND idx = $2;
	`
)
