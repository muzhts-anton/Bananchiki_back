package quizrep

import (
	"banana/pkg/database"
	"banana/pkg/domain"
	"banana/pkg/utils/cast"
	"banana/pkg/utils/log"
)

type dbQuizRepository struct {
	dbm *database.DBManager
}

func InitQuizRep(manager *database.DBManager) domain.QuizRepository {
	return &dbQuizRepository{
		dbm: manager,
	}
}

func (r *dbQuizRepository) CreateQuiz(q domain.Quiz, pid uint64) (uint64, error) {
	resp, err := r.dbm.Query(queryCreateQuiz, q.Type, q.Question, q.Background, q.FontColor, q.FontSize, q.GraphColor)
	if err != nil {
		log.Warn("{CreateQuiz} in query: " + queryCreateQuiz)
		log.Error(err)
		return 0, domain.ErrDatabaseRequest
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
		return 0, domain.ErrDatabaseRequest
	}

	err = r.dbm.Execute(queryInsertQuiz, pid, domain.SlideTypeQuiz, q.Id, q.Idx)
	if err != nil {
		log.Warn("{CreateQuiz} in query: " + queryInsertQuiz)
		log.Error(err)
		return 0, domain.ErrDatabaseRequest
	}

	err = r.dbm.Execute(queryIncrementQuizNum, pid)
	if err != nil {
		log.Warn("{CreateQuiz} in query: " + queryIncrementQuizNum)
		log.Error(err)
		return 0, domain.ErrDatabaseRequest
	}

	return q.Id, nil
}

func (r *dbQuizRepository) DeleteQuiz(qid, pid uint64) error {
	resp, err := r.dbm.Query(queryGetQuizIdx, domain.SlideTypeQuiz, qid)
	if err != nil {
		log.Warn("{DeleteQuiz} in query: " + queryDeleteQuiz)
		log.Error(err)
		return domain.ErrDatabaseRequest
	}
	if len(resp) == 0 {
		log.Error(domain.ErrDatabaseRequest)
		return domain.ErrDatabaseRequest
	}

	err = r.dbm.Execute(queryDeleteQuiz, qid)
	if err != nil {
		log.Warn("{DeleteQuiz} in query: " + queryDeleteQuiz)
		log.Error(err)
		return domain.ErrDatabaseRequest
	}

	err = r.dbm.Execute(queryCutQuiz, domain.SlideTypeQuiz, qid)
	if err != nil {
		log.Warn("{DeleteQuiz} in query: " + queryCutQuiz)
		log.Error(err)
		return domain.ErrDatabaseRequest
	}

	err = r.dbm.Execute(queryShiftDownIdxs, cast.ToUint16(resp[0][0]), pid)
	if err != nil {
		log.Warn("{DeleteQuiz} in query: " + queryShiftDownIdxs)
		log.Error(err)
		return domain.ErrDatabaseRequest
	}

	err = r.dbm.Execute(queryDeleteVotes, qid)
	if err != nil {
		log.Warn("{DeleteQuiz} in query: " + queryDeleteVotes)
		log.Error(err)
		return domain.ErrDatabaseRequest
	}

	err = r.dbm.Execute(queryDecrementQuizNum, pid)
	if err != nil {
		log.Warn("{CreateQuiz} in query: " + queryDecrementQuizNum)
		log.Error(err)
		return domain.ErrDatabaseRequest
	}

	return nil
}

func (r *dbQuizRepository) UpdateQuiz(q domain.Quiz, pid uint64) error {
	err := r.dbm.Execute(queryUpdateQuiz, q.Question, q.Background, q.FontColor, q.FontSize, q.GraphColor, q.Id)
	if err != nil {
		log.Warn("{UpdateQuiz} in query: " + queryUpdateQuiz)
		log.Error(err)
		return domain.ErrDatabaseRequest
	}

	return nil
}

//////////////////////////////////
func (r *dbQuizRepository) CreateQuizVote(q domain.Vote, qid uint64) error {
	err := r.dbm.Execute(queryShiftUpVote, q.Idx, qid)
	if err != nil {
		log.Warn("{CreateQuizVote} in query: " + queryShiftUpVote)
		log.Error(err)
		return domain.ErrDatabaseRequest
	}

	err = r.dbm.Execute(queryCreateQuizVote, qid, q.Idx, q.Option, q.Votes, q.Color)
	if err != nil {
		log.Warn("{CreateQuizVote} in query: " + queryCreateQuizVote)
		log.Error(err)
		return domain.ErrDatabaseRequest
	}

	return nil
}

func (r *dbQuizRepository) UpdateQuizVote(q domain.Vote, qid uint64) error {
	err := r.dbm.Execute(queryUpdateQuizVote, q.Option, q.Votes, q.Color, q.Idx, qid)
	if err != nil {
		log.Warn("{UpdateQuizVote} in query: " + queryUpdateQuizVote)
		log.Error(err)
		return domain.ErrDatabaseRequest
	}

	return nil
}

func (r *dbQuizRepository) DeleteQuizVote(idx uint32, qid uint64) error {
	err := r.dbm.Execute(queryDeleteQuizVote, idx, qid)
	if err != nil {
		log.Warn("{DeleteQuizVote} in query: " + queryDeleteQuizVote)
		log.Error(err)
		return domain.ErrDatabaseRequest
	}

	err = r.dbm.Execute(queryShiftDownVote, idx, qid)
	if err != nil {
		log.Warn("{DeleteQuizVote} in query: " + queryShiftDownVote)
		log.Error(err)
		return domain.ErrDatabaseRequest
	}

	return nil
}
