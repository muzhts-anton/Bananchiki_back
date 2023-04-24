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
	UPDATE slideorder SET idx = idx + 1 WHERE idx >= $1 AND presentation_id = $2;
	`

	queryGetQuizIdx = `
	SELECT idx from slideorder WHERE type = $1 AND item_id = $2;
	`

	queryDeleteQuiz = `
	DELETE FROM quiz WHERE id = $1;
	`

	queryCutQuiz = `
	DELETE FROM slideorder WHERE type = $1 AND item_id = $2;
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
		question = $1,
		background = $2,
		font_color = $3,
		font_size = $4,
		graph_color = $5,
		type = $6
	WHERE id = $7;
	`

	queryIncrementQuizNum = `
	UPDATE presentation SET quiz_num = quiz_num + 1 WHERE id = $1;
	`

	queryDecrementQuizNum = `
	UPDATE presentation SET quiz_num = quiz_num - 1 WHERE id = $1;
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
	WHERE idx = $4 AND quiz_id = $5;
	`

	queryDeleteQuizVote = `
	DELETE FROM vote WHERE idx = $1 AND quiz_id = $2;
	`

	queryShiftDownVote = `
	UPDATE vote SET idx = idx - 1 WHERE idx >= $1 AND quiz_id = $2;
	`

	queryShiftUpVote = `
	UPDATE vote SET idx = idx + 1 WHERE idx >= $1 AND quiz_id = $2;
	`

	queryPollQuizVote = `
	UPDATE vote SET votes_num = votes_num + 1 WHERE quiz_id = $1 AND idx = $2;
	`
)

const (
	queryCompetitionStart = `
	UPDATE quiz SET runout = False, start_time = current_timestamp WHERE id = $1
	`

	queryCompetitionStop = `
	UPDATE quiz SET runout = True WHERE id = $1
	`

	queryCompetitionVoterRegister = `
	INSERT INTO voters (presentation_id, name)
	VALUES ($1, $2);
	RETURNING id;
	`

	queryGetAllPres = `
	SELECT id, code FROM presentation;
	`

	queryGetPrevCompetitionResult = `
	SELECT id, name, points
	FROM voters
	ORDER BY top_place DESC
	`

	queryGetCurrentCompetitionResult = `
	SELECT id, name, points
	FROM voters
	ORDER BY points DESC
	`
)
