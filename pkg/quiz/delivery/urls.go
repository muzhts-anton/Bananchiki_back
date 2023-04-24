package quizdel

const (
	urlCreateQuiz     = "/quiz/create"
	urlUpdateQuiz     = "/quiz/update"
	urlDeleteQuiz     = "/quiz/delete"
	urlCreateQuizVote = "/quiz/vote/create"
	urlUpdateQuizVote = "/quiz/vote/update"
	urlDeleteQuizVote = "/quiz/vote/delete"
	urlPollQuizVote   = "/quiz/vote/poll"

	urlCompetitionStart = "/competition/start"
	urlCompetitionStop = "/competition/stop"
	urlCompetitionVoterRegister = "/competition/voter/register"
	urlCompetitionResult = "/competition/{uint64 presId}/result"

// GET /competition/{uint64 presId}/result
// <- { “top”: [{name, points}, …] }
)
