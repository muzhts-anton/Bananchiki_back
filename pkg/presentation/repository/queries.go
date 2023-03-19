package presrep

const (
	queryGetPres = `
	SELECT (id, creator_id, url, converted_slide_num, quiz_num)
	FROM presentation
	WHERE id = $1;
	`

	queryGetConvertedSlides = `
	SELECT (slideorder.idx, convertedslide.name, convertedslide.width, convertedslide.height)
	FROM convertedslide
	JOIN slideorder ON convertedslide.id = slideorder.item_id
	WHERE slideorder.type = $1 AND slideorder.presentation_id = $2
	ORDER BY slideorder.idx;
	`

	queryGetQuizzes = `
	SELECT
		(quiz.id, slideorder.idx, quiz.type, quiz.question,
		quiz.background, quiz.font_color, quiz.font_size, quiz.graph_color)
	FROM quiz
	JOIN slideorder ON quiz.id = slideorder.item_id
	WHERE slideorder.type = $1 AND slideorder.presentation_id = $2
	ORDER BY slideorder.idx;
	`

	queryGetVotes = `
	SELECT (idx, option, votes_num, color)
	FROM vote
	WHERE quiz_id = $1
	ORDER BY idx;
	`
)
