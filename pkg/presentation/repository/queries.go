package presrep

const (
	queryCreatePres = `
	INSERT INTO presentation (creator_id, url, converted_slide_num, quiz_num)
	VALUES ($1, $2, $3, $4)
	RETURNING id;
	`

	queryCreateConvertedSlide = `
	INSERT INTO convertedslide (name, width, height)
	VALUES ($1, $2, $3)
	RETURNING id;
	`

	queryInsertConvertedSlide = `
	INSERT INTO slideorder (presentation_id, type, item_id, idx)
	VALUES ($1, $2, $3, $4);
	`

	queryUpdateConvertedSlideNum = `
	UPDATE presentation SET converted_slide_num = $1 WHERE id = $2;
	`

	queryUpdatePresUrl = `
	UPDATE presentation SET url = $1 WHERE id = $2;
	`
)

const (
	queryGetPres = `
	SELECT id, creator_id, url, name, code, converted_slide_num, quiz_num
	FROM presentation
	WHERE id = $1;
	`

	queryGetConvertedSlides = `
	SELECT slideorder.idx, convertedslide.name, convertedslide.width, convertedslide.height
	FROM convertedslide
	JOIN slideorder ON convertedslide.id = slideorder.item_id
	WHERE slideorder.type = $1 AND slideorder.presentation_id = $2
	ORDER BY slideorder.idx;
	`

	queryGetQuizzes = `
	SELECT
		quiz.id, slideorder.idx, quiz.type, quiz.question,
		quiz.background, quiz.font_color, quiz.font_size, quiz.graph_color
	FROM quiz
	JOIN slideorder ON quiz.id = slideorder.item_id
	WHERE slideorder.type = $1 AND slideorder.presentation_id = $2
	ORDER BY slideorder.idx;
	`

	queryGetVotes = `
	SELECT idx, option, votes_num, color
	FROM vote
	WHERE quiz_id = $1
	ORDER BY idx;
	`

	queryGetPresOwner = `
	SELECT creator_id FROM presentation WHERE id = $1;
	`

	queryChangePresName = `
	UPDATE presentation SET name = $1 WHERE id = $2;
	`
)

const (
	queryGetSlidesIds = `
	SELECT item_id, type FROM slideorder WHERE presentation_id = $1;
	`

	queryDeleteVotes = `
	DELETE FROM vote WHERE quiz_id = $1;
	`

	queryDeleteQuiz = `
	DELETE FROM quiz WHERE id = $1;
	`

	queryDeleteConvSlides = `
	DELETE FROM convertedslide WHERE id = $1;
	`

	queryDeleteSlideorder = `
	DELETE FROM slideorder WHERE presentation_id = $1;
	`

	queryDeleteQuestions = `
	DELETE FROM question WHERE presentation_id = $1;
	`

	queryDeletePresentation = `
	DELETE FROM presentation WHERE id = $1;
	`
)