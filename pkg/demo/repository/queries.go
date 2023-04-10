package demorep

const (
	queryGetAllPres = `
	SELECT id, code FROM presentation;
	`

	queryGetPresVm = `
	SELECT viewmode FROM presentation WHERE id = $1;
	`

	queryGetCurrentDemoSlideType = `
	SELECT slideorder.type, slideorder.item_id, presentation.demo_idx
	FROM slideorder
	JOIN presentation ON slideorder.presentation_id = presentation.id
	WHERE presentation.id = $1 AND slideorder.idx = presentation.demo_idx;
	`

	queryGetPresEmotions = `
	SELECT emotion_like, emotion_love, emotion_laughter, emotion_surprise, emotion_sad
	FROM presentation
	WHERE id = $1;
	`

	queryGetConvertedSlide = `
	SELECT name, width, height FROM convertedslide WHERE id = $1;
	`

	queryGetQuiz = `
	SELECT id, type, question, background, font_color, font_size, graph_color
	FROM quiz
	WHERE id = $1;
	`

	queryGetCreatorId = `
	SELECT creator_id FROM presentation WHERE id = $1;
	`

	queryGetVotes = `
	SELECT idx, option, votes_num, color
	FROM vote
	WHERE quiz_id = $1
	ORDER BY idx;
	`

	queryDemoGo = `
	UPDATE presentation SET viewmode = true, demo_idx = $1 WHERE id = $2;
	`

	queryDemoStop = `
	UPDATE presentation SET viewmode = false WHERE id = $1;
	`
	queryZeroingReations = `
	UPDATE presentation
	SET
		emotion_like = 0,
		emotion_love = 0,
		emotion_laughter = 0,
		emotion_surprise = 0,
		emotion_sad = 0
	WHERE id = $1;
	`

	queryGetAllQuizzes = `
	SELECT item_id
	FROM slideorder
	WHERE presentation_id = $1 AND type = 'question';
	`

	querySetAllVotes = `
	UPDATE vote
	SET votes_num = $1
	WHERE quiz_id = $2;
	`

	queryGetPresQuestions = `
	SELECT idx, option, likes FROM question WHERE presentation_id = $1 ORDER BY idx;
	`

	queryDeletePresQuestions = `
	DELETE FROM questions WHERE presentation_id = $1;
	`
)
