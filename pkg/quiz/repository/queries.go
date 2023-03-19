package quizrep

const (
	queryCreateQuiz = `
	INSERT INTO
		quiz (type, question, background, font_color, font_size, graph_color)
	VALUES
		($1, $2, $3, $4, $5, $6)
	RETURNING id;
	`

	queryInsertQuiz = `
	INSERT INTO slideorder (presentation_id, type, item_id, idx) values ($1, $2, $3, $4)
	`

	queryShiftUpIdxs = `
	UPDATE slideorder SET idx = idx + 1 WHERE idx >= $1 AND pres_id = $2;
	`

	queryDeleteQuiz = `
	DELETE FROM quiz WHERE id = $1;
	`

	queryShiftDownIdxs = `
	UPDATE slideorder SET idx = idx - 1 WHERE idx >= $1 AND presentation_id = $2;
	`

	queryDeleteVotes = `
	DELETE FROM vote WHERE quiz_id = $1;
	`

	queryUpdateQuiz = `
	UPDATE quiz
	SET
		type = $1,
		question = $2,
		background = $3,
		font_color = $4,
		font_size = $5,
		graph_color
	WHERE id = $6;
	`
)

const (
	queryCreateQuizVote = `
	INSERT INTO vote (quiz_id, idx, option, votes_num, color)
	VALUES ($1, $2, $3, $4, $5);
	`

	queryUpdateQuizVote = `
	UPDATE vote
	SET
		option = $1,
		votes_num = $2,
		color = $3
	WHERE idx = $1 AND quiz_id = $5;
	`

	queryDeleteQuizVote = `
	DELETE FROM vote WHERE quiz_id = $1 AND idx = $2;
	`

	queryShiftDownVote = `
	UPDATE vote SET idx = idx - 1 WHERE idx >= $1 AND quiz_id = $2;
	`

	queryShiftUpVote = `
	UPDATE vote SET idx = idx + 1 WHERE idx >= $1 AND quiz_id = $2;
	`
)
