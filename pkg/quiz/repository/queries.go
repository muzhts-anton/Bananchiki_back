package quizrep

const (
	queryCreateQuiz = `
	INSERT INTO
		quiz (
			type, question, seconds_num, result_after, price,
			extra_points, background, font_color, font_size, graph_color
		)
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
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
		type = $6,
		seconds_num = $7,
		result_after = $8,
		price = $9,
		extra_points = $10
	WHERE id = $11;
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
	INSERT INTO vote (quiz_id, idx, option, correct, votes_num, color)
	VALUES ($1, $2, $3, $4, $5, $6);
	`

	queryUpdateQuizVote = `
	UPDATE vote
	SET
		option = $1,
		votes_num = $2,
		color = $3,
		correct = $4
	WHERE idx = $5 AND quiz_id = $6;
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

	queryGetVoterCount = `
	SELECT COUNT(*)
	FROM voter_quiz
	JOIN quiz ON voter_quiz.quiz_id = quiz.id
	JOIN voters ON voter_quiz.voter_id = voters.id
	WHERE quiz.id = $1 AND voters.id = $2;
	`

	queryAppendQuizVoter = `
	INSERT INTO voter_quiz (quiz_id, voter_id) VALUES ($1, $2);
	`

	queryGetQuizInfo = `
	SELECT
		extra_points, price, seconds_num, runout,
		EXTRACT(EPOCH FROM AGE(current_timestamp, start_time))
	FROM quiz
	WHERE id = $1;
	`

	queryPutPts = `
	UPDATE voters SET points = points + $1 WHERE id = $2;
	`

	queryIsVoteCorrect = `
	SELECT correct FROM vote WHERE idx = $1 AND quiz_id = $2;
	`
)

const (
	queryCompetitionStart = `
	UPDATE quiz SET runout = FALSE, start_time = current_timestamp WHERE id = $1;
	`

	queryCompetitionStop = `
	UPDATE quiz SET runout = TRUE WHERE id = $1;
	`

	queryCompetitionVoterRegister = `
	INSERT INTO voters (presentation_id, name) VALUES ($1, $2) RETURNING id;
	`

	queryGetAllPres = `
	SELECT id, code FROM presentation;
	`
	
	queryClearAllTopPlaces = `
	UPDATE voter SET points = 0 WHERE presentation_id = $1;
	`

	queryGetTopByPts = `
	SELECT id FROM voters WHERE presentation_id = $1 ORDER BY points DESC LIMIT 5;
	`

	querySetTopVoter = `
	UPDATE voters SET top_place = $1 WHERE presentation_id = $2 AND id = $3;
	`

	queryGetTop = `
	SELECT id, name, points
	FROM voters
	WHERE presentation_id = $1 AND top_place > 0
	ORDER BY id
	LIMIT 5;
	`
)
