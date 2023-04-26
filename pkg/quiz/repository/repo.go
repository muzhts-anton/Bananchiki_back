package quizrep

import (
	"banana/pkg/domain"
	"banana/pkg/utils/cast"
	"banana/pkg/utils/database"
	"banana/pkg/utils/hash"
	"banana/pkg/utils/log"
	"banana/pkg/utils/points"
)

type dbQuizRepository struct {
	dbm *database.DBManager
}

func InitQuizRep(manager *database.DBManager) domain.QuizRepository {
	return &dbQuizRepository{
		dbm: manager,
	}
}

// quizzes
func (r *dbQuizRepository) CreateQuiz(q domain.Quiz, pid uint64) (uint64, error) {
	resp, err := r.dbm.Query(queryCreateQuiz,
		q.Type, q.Question, q.AnswerTime, q.ResultAfter, q.Cost,
		q.ExtraPts, q.Background, q.FontColor, q.FontSize, q.GraphColor,
	)
	if err != nil {
		log.Warn("{CreateQuiz} in query: " + queryCreateQuiz)
		log.Error(err)
		return 0, err
	}
	if len(resp) == 0 {
		log.Warn("{CreateQuiz}")
		log.Error(domain.ErrDatabaseRange)
		return 0, domain.ErrDatabaseRange
	}

	q.Id = cast.ToUint64(resp[0][0])

	err = r.dbm.Execute(queryShiftUpIdxs, q.Idx, pid)
	if err != nil {
		log.Warn("{CreateQuiz} in query: " + queryShiftUpIdxs)
		log.Error(err)
		return 0, err
	}

	err = r.dbm.Execute(queryInsertQuiz, pid, domain.SlideTypeQuiz, q.Id, q.Idx)
	if err != nil {
		log.Warn("{CreateQuiz} in query: " + queryInsertQuiz)
		log.Error(err)
		return 0, err
	}

	err = r.dbm.Execute(queryIncrementQuizNum, pid)
	if err != nil {
		log.Warn("{CreateQuiz} in query: " + queryIncrementQuizNum)
		log.Error(err)
		return 0, err
	}

	return q.Id, nil
}

func (r *dbQuizRepository) DeleteQuiz(qid, pid uint64) error {
	resp, err := r.dbm.Query(queryGetQuizIdx, domain.SlideTypeQuiz, qid)
	if err != nil {
		log.Warn("{DeleteQuiz} in query: " + queryDeleteQuiz)
		log.Error(err)
		return err
	}
	if len(resp) == 0 {
		log.Error(domain.ErrDatabaseRequest)
		return domain.ErrDatabaseRequest
	}

	err = r.dbm.Execute(queryDeleteVotes, qid)
	if err != nil {
		log.Warn("{DeleteQuiz} in query: " + queryDeleteVotes)
		log.Error(err)
		return err
	}

	err = r.dbm.Execute(queryDeleteQuiz, qid)
	if err != nil {
		log.Warn("{DeleteQuiz} in query: " + queryDeleteQuiz)
		log.Error(err)
		return err
	}

	err = r.dbm.Execute(queryCutQuiz, domain.SlideTypeQuiz, qid)
	if err != nil {
		log.Warn("{DeleteQuiz} in query: " + queryCutQuiz)
		log.Error(err)
		return err
	}

	err = r.dbm.Execute(queryShiftDownIdxs, cast.ToUint16(resp[0][0]), pid)
	if err != nil {
		log.Warn("{DeleteQuiz} in query: " + queryShiftDownIdxs)
		log.Error(err)
		return err
	}

	err = r.dbm.Execute(queryDecrementQuizNum, pid)
	if err != nil {
		log.Warn("{CreateQuiz} in query: " + queryDecrementQuizNum)
		log.Error(err)
		return err
	}

	return nil
}

func (r *dbQuizRepository) UpdateQuiz(q domain.Quiz, pid uint64) error {
	err := r.dbm.Execute(queryUpdateQuiz,
		q.Question, q.Background, q.FontColor, q.FontSize, q.GraphColor,
		q.Type, q.AnswerTime, q.ResultAfter, q.Cost, q.ExtraPts, q.Id,
	)
	if err != nil {
		log.Warn("{UpdateQuiz} in query: " + queryUpdateQuiz)
		log.Error(err)
		return err
	}

	return nil
}

// votes
func (r *dbQuizRepository) CreateQuizVote(q domain.Vote, qid uint64) error {
	err := r.dbm.Execute(queryShiftUpVote, q.Idx, qid)
	if err != nil {
		log.Warn("{CreateQuizVote} in query: " + queryShiftUpVote)
		log.Error(err)
		return err
	}

	err = r.dbm.Execute(queryCreateQuizVote, qid, q.Idx, q.Option, q.Correct, q.Votes, q.Color)
	if err != nil {
		log.Warn("{CreateQuizVote} in query: " + queryCreateQuizVote)
		log.Error(err)
		return err
	}

	return nil
}

func (r *dbQuizRepository) UpdateQuizVote(q domain.Vote, qid uint64) error {
	err := r.dbm.Execute(queryUpdateQuizVote, q.Option, q.Votes, q.Color, q.Correct, q.Idx, qid)
	if err != nil {
		log.Warn("{UpdateQuizVote} in query: " + queryUpdateQuizVote)
		log.Error(err)
		return err
	}

	return nil
}

func (r *dbQuizRepository) DeleteQuizVote(idx uint32, qid uint64) error {
	err := r.dbm.Execute(queryDeleteQuizVote, idx, qid)
	if err != nil {
		log.Warn("{DeleteQuizVote} in query: " + queryDeleteQuizVote)
		log.Error(err)
		return err
	}

	err = r.dbm.Execute(queryShiftDownVote, idx, qid)
	if err != nil {
		log.Warn("{DeleteQuizVote} in query: " + queryShiftDownVote)
		log.Error(err)
		return err
	}

	return nil
}

func (r *dbQuizRepository) PollQuizVote(idx uint32, qid uint64) error {
	err := r.dbm.Execute(queryPollQuizVote, qid, idx)
	if err != nil {
		log.Warn("{PollQuizVote} in query: " + queryPollQuizVote)
		log.Error(err)
		return err
	}

	return nil
}

func (r *dbQuizRepository) PollQuizVoteTracked(idx uint32, qid uint64, vid uint64) error {
	resp, err := r.dbm.Query(queryGetVoterCount, qid, vid)
	if err != nil {
		log.Warn("{PollQuizVoteTracked} in query: " + queryGetVoterCount)
		log.Error(err)
		return err
	}

	if cast.ToUint64(resp[0][0]) != 0 {
		return domain.ErrSecondVote
	}

	err = r.dbm.Execute(queryAppendQuizVoter, qid, vid)
	if err != nil {
		log.Warn("{PollQuizVoteTracked} in query: " + queryAppendQuizVoter)
		log.Error(err)
	}

	return err
}

func (r *dbQuizRepository) CalculatePoints(idx uint32, qid uint64, vid uint64) error {
	resp, err := r.dbm.Query(queryGetQuizInfo, qid)
	if err != nil {
		log.Warn("{CalculatePoints} in query: " + queryGetQuizInfo)
		log.Error(err)
		return err
	}

	if cast.ToBool(resp[0][3]) {
		return domain.ErrRunout
	}

	tmpresp, err := r.dbm.Query(queryIsVoteCorrect, idx, qid)
	if err != nil {
		log.Warn("{CalculatePoints} in query: " + queryPutPts)
		log.Error(err)
		return err
	}

	if !cast.ToBool(tmpresp[0][0]) {
		return nil
	}

	pts, err := points.CalculateQuizPoints(
		cast.ToBool(resp[0][0]), cast.ToUint64(resp[0][1]),
		cast.ToUint64(resp[0][2]), cast.ToFloat64(resp[0][4]),
	)
	if err != nil {
		log.Error(err)
	}

	err = r.dbm.Execute(queryPutPts, pts, vid)
	if err != nil {
		log.Warn("{CalculatePoints} in query: " + queryPutPts)
		log.Error(err)
	}

	return err
}

func (r *dbQuizRepository) CompetitionStart(quizId uint64, presId uint64) error {
	err := r.dbm.Execute(queryCompetitionStart, quizId)
	if err != nil {
		log.Warn("{queryCompetitionStart} in query: " + queryCompetitionStart)
		log.Error(err)
		return err
	}

	return nil
}

func (r *dbQuizRepository) CompetitionStop(quizId uint64, presId uint64) error {
	err := r.dbm.Execute(queryCompetitionStop, quizId)
	if err != nil {
		log.Warn("{queryCompetitionStop} in query: " + queryCompetitionStop)
		log.Error(err)
		return err
	}

	return nil
}

func (r *dbQuizRepository) CompetitionVoterRegister(name string, presId uint64) (uint64, error) {
	resp, err := r.dbm.Query(queryCompetitionVoterRegister, presId, name)
	if err != nil {
		log.Warn("{queryCompetitionStop} in query: " + queryCompetitionStop)
		log.Error(err)
		return 0, err
	}

	return cast.ToUint64(resp[0][0]), nil
}

func (r *dbQuizRepository) GetPresIdByHash(h string) (uint64, error) {
	resp, err := r.dbm.Query(queryGetAllPres)
	if err != nil {
		log.Warn("{GetPresIdByHash} in query: " + queryGetAllPres)
		log.Error(err)
		return 0, err
	}

	for _, pres := range resp {
		if code := cast.ToString(pres[1]); hash.EncodeToHash(code) == h {
			return cast.ToUint64(pres[0]), nil
		}
	}

	return 0, domain.ErrCodeNotFound
}

func (r *dbQuizRepository) GetCompetitionResult(pid uint64) ([]domain.ResultItem, error) {
	resp, err := r.dbm.Query(queryGetTop, pid)
	if err != nil {
		log.Warn("{GetCompetitionResult} in query: " + queryGetTop)
		log.Error(err)
		return nil, err
	}

	out := make([]domain.ResultItem, 0)
	for _, voter := range resp {
		out = append(out, domain.ResultItem{
			Id:     cast.ToUint64(voter[0]),
			Name:   cast.ToString(voter[1]),
			Points: cast.ToUint64(voter[2]),
		})
	}

	return out, nil
}

func (r *dbQuizRepository) SetCompetitionResult(pid uint64) error {
	err := r.dbm.Execute(queryClearAllTopPlaces, pid)
	if err != nil {
		log.Warn("{SetCompetitionResult} in query: " + queryClearAllTopPlaces)
		log.Error(err)
		return err
	}

	resp, err := r.dbm.Query(queryGetTopByPts, pid)
	if err != nil {
		log.Warn("{SetCompetitionResult} in query: " + queryGetTopByPts)
		log.Error(err)
		return err
	}

	for i, voter := range resp {
		err = r.dbm.Execute(querySetTopVoter, uint16(i+1), pid, cast.ToUint64(voter[0]))
		if err != nil {
			log.Warn("{SetCompetitionResult} in query: " + querySetTopVoter)
			log.Error(err)
			return err
		}
	}

	return nil
}
