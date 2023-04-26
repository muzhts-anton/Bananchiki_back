package demorep

import (
	"banana/pkg/domain"
	"banana/pkg/utils/cast"
	"banana/pkg/utils/database"
	"banana/pkg/utils/hash"
	"banana/pkg/utils/log"
)

type dbDemoRepository struct {
	dbm *database.DBManager
}

func InitDemoRep(manager *database.DBManager) domain.DemoRepository {
	return &dbDemoRepository{
		dbm: manager,
	}
}

func (ur *dbDemoRepository) GetPresIdByHash(h string) (uint64, error) {
	resp, err := ur.dbm.Query(queryGetAllPres)
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

func (ur *dbDemoRepository) GetPresViewMode(pid uint64) (bool, error) {
	resp, err := ur.dbm.Query(queryGetPresVm, pid)
	if err != nil {
		log.Warn("{GetPresViewMode} in query: " + queryGetPresVm)
		log.Error(err)
		return false, err
	}
	if len(resp) == 0 {
		return false, domain.ErrCodeNotFound
	}

	return cast.ToBool(resp[0][0]), nil
}

func (ur *dbDemoRepository) GetCurrentDemoSlide(pid uint64) (out domain.SlideApiResponse, err error) {
	resp, err := ur.dbm.Query(queryGetCurrentDemoSlideType, pid)
	if err != nil {
		log.Warn("{GetCurrentDemoSlide} in query: " + queryGetCurrentDemoSlideType)
		log.Error(err)
		return domain.SlideApiResponse{}, err
	}
	if len(resp) == 0 {
		return domain.SlideApiResponse{}, domain.ErrDatabaseRequest
	}

	itemId := cast.ToUint64(resp[0][1])
	out.Idx = uint32(cast.ToUint16(resp[0][2]))
	if cast.ToString(resp[0][0]) == domain.SildeTypeConvertedSlide {
		resp, err = ur.dbm.Query(queryGetConvertedSlide, itemId)
		if err != nil {
			log.Warn("{GetCurrentDemoSlide} in query: " + queryGetConvertedSlide)
			log.Error(err)
			return domain.SlideApiResponse{}, err
		}

		out.Kind = domain.SildeTypeConvertedSlide
		out.Name = cast.ToString(resp[0][0])
		out.Width = uint32(cast.ToUint16(resp[0][1]))
		out.Height = uint32(cast.ToUint16(resp[0][2]))
	} else {
		resp, err = ur.dbm.Query(queryGetQuiz, itemId)
		if err != nil {
			log.Warn("{GetCurrentDemoSlide} in query: " + queryGetQuiz)
			log.Error(err)
			return domain.SlideApiResponse{}, err
		}

		out.QuizId = cast.ToUint64(resp[0][0])
		out.Type = cast.ToString(resp[0][1])
		out.Question = cast.ToString(resp[0][2])
		out.Background = cast.ToString(resp[0][3])
		out.FontColor = cast.ToString(resp[0][4])
		out.FontSize = cast.ToString(resp[0][5])
		out.GraphColor = cast.ToString(resp[0][6])
		out.Runout = cast.ToBool(resp[0][7])
		out.AnswerTime = cast.ToUint64(resp[0][8])
		out.Cost = cast.ToUint64(resp[0][9])

		if out.Cost > 0 || out.AnswerTime > 0 {
			out.Kind = domain.SlideTypeTimedQuiz
		} else {
			out.Kind = domain.SlideTypeQuiz
		}

		out.Vote = make([]domain.Vote, 0)
		tresp, terr := ur.dbm.Query(queryGetVotes, itemId)
		if terr != nil {
			log.Warn("{GetCurrentDemoSlide} in query: " + queryGetVotes)
			log.Error(err)
		}
		for _, vote := range tresp {
			out.Vote = append(out.Vote, domain.Vote{
				Idx:    uint32(cast.ToUint16(vote[0])),
				Option: cast.ToString(vote[1]),
				Votes:  cast.ToUint64(vote[2]),
				Color:  cast.ToString(vote[3]),
			})
		}
	}

	return
}

func (ur *dbDemoRepository) GetPresCreatorId(pid uint64) (uint64, error) {
	resp, err := ur.dbm.Query(queryGetCreatorId, pid)
	if err != nil {
		log.Warn("{GetPresCreatorId} in query: " + queryGetCreatorId)
		log.Error(err)
		return 0, err
	}
	if len(resp) == 0 {
		return 0, domain.ErrDatabaseRequest
	}

	return cast.ToUint64(resp[0][0]), nil
}

func (ur *dbDemoRepository) DemoGo(pid uint64, idx uint32) error {
	err := ur.dbm.Execute(queryDemoGo, idx, pid)
	if err != nil {
		log.Warn("{DemoGo} in query: " + queryDemoGo)
		log.Error(err)
	}

	return err
}

func (ur *dbDemoRepository) GetViewMode(pid uint64) (bool, error) {
	resp, err := ur.dbm.Query(queryGetViewMode, pid)
	if err != nil {
		log.Warn("{GetViewMode} in query: " + queryGetViewMode)
		log.Error(err)
		return false, err
	}

	if len(resp) == 0 {
		return false, domain.ErrDatabaseRequest
	}
	return cast.ToBool(resp[0][0]), nil
}

func (ur *dbDemoRepository) DemoStop(pid uint64) error {
	err := ur.dbm.Execute(queryDemoStop, pid)
	if err != nil {
		log.Warn("{DemoStop} in query: " + queryDemoStop)
		log.Error(err)
	}

	return err
}

func (ur *dbDemoRepository) GetPresEmotions(pid uint64) (domain.PresEmotions, error) {
	resp, err := ur.dbm.Query(queryGetPresEmotions, pid)
	if err != nil {
		log.Warn("{GetPresEmotions} in query: " + queryGetPresEmotions)
		log.Error(err)
		return domain.PresEmotions{}, err
	}
	if len(resp) == 0 {
		return domain.PresEmotions{}, domain.ErrDatabaseRequest
	}

	return domain.PresEmotions{
		Like:     cast.ToUint64(resp[0][0]),
		Love:     cast.ToUint64(resp[0][1]),
		Laughter: cast.ToUint64(resp[0][2]),
		Surprise: cast.ToUint64(resp[0][3]),
		Sad:      cast.ToUint64(resp[0][4]),
	}, nil
}

func (ur *dbDemoRepository) GetPresQuestions(pid uint64) ([]domain.Question, error) {
	resp, err := ur.dbm.Query(queryGetPresQuestions, pid)
	if err != nil {
		log.Warn("{GetPresQuestions} in query: " + queryGetPresQuestions)
		log.Error(err)
		return nil, err
	}

	out := make([]domain.Question, 0)
	for _, q := range resp {
		out = append(out, domain.Question{
			Idx:    cast.ToUint64(q[0]),
			Option: cast.ToString(q[1]),
			Likes:  cast.ToUint64(q[2]),
		})
	}

	return out, nil
}

func (ur *dbDemoRepository) ZeroingReactions(pid uint64) error {
	err := ur.dbm.Execute(queryZeroingReations, pid)
	if err != nil {
		log.Warn("{ZeroingReactions} in query: " + queryZeroingReations)
		log.Error(err)
		return err
	}
	log.Debug("zeroing reactions")
	return nil
}

func (ur *dbDemoRepository) SetAllVotes(pid uint64, value int) error {
	resp, err := ur.dbm.Query(queryGetAllQuizzes, pid)
	if err != nil {
		log.Warn("{setAllVotes} in query: " + queryGetAllQuizzes)
		log.Error(err)
		return err
	}

	for _, q := range resp {
		err = ur.dbm.Execute(querySetAllVotes, value, cast.ToUint64(q[0]))
		if err != nil {
			log.Warn("{setAllVotes} in query: " + querySetAllVotes)
			log.Error(err)
			return err
		}
	}

	return nil
}

func (ur *dbDemoRepository) DeletePresQuestions(pid uint64) error {
	err := ur.dbm.Execute(queryDeletePresQuestions, pid)
	if err != nil {
		log.Warn("{DeletePresQuestions} in query: " + queryDeletePresQuestions)
		log.Error(err)
		return err
	}

	return nil
}
