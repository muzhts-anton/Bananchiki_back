package presrep

import (
	"banana/pkg/database"
	"banana/pkg/domain"
	"banana/pkg/utils/cast"
	"banana/pkg/utils/log"
)

type dbPresRepository struct {
	dbm *database.DBManager
}

func InitPresRep(manager *database.DBManager) domain.PresRepository {
	return &dbPresRepository{
		dbm: manager,
	}
}

func (r *dbPresRepository) CreatePres(q *domain.Presentation) error {
	return nil
}

func (r *dbPresRepository) GetPres(pid uint64) (domain.Presentation, error) {
	resp, err := r.dbm.Query(queryGetPres, pid)
	if err != nil {
		log.Warn("{GetPres} in query: " + queryGetPres)
		log.Error(err)
		return domain.Presentation{}, domain.ErrDatabaseRequest
	}
	if len(resp) == 0 {
		log.Warn("{GetPres}")
		log.Error(domain.ErrDatabaseRange)
		return domain.Presentation{}, domain.ErrDatabaseRange
	}

	return domain.Presentation{
		Id:        cast.ToUint64(resp[0][0]),
		CreatorId: cast.ToUint64(resp[0][1]),
		Url:       cast.ToString(resp[0][2]),
		SlideNum:  cast.ToUint32(resp[0][3]),
		QuizNum:   cast.ToUint32(resp[0][4]),
		Slides:    nil,
		Quizzes:   nil,
	}, nil
}

func (r *dbPresRepository) GetConvertedSlides(t string, pid uint64) (slides []domain.ConvertedSlide, err error) {
	resp, err := r.dbm.Query(queryGetConvertedSlides, t, pid)
	if err != nil {
		log.Warn("{GetConvertedSlides} in query: " + queryGetConvertedSlides)
		log.Error(err)
		return nil, domain.ErrDatabaseRequest
	}

	slides = make([]domain.ConvertedSlide, 0)
	for _, slide := range resp {
		slides = append(slides, domain.ConvertedSlide{
			Idx:    cast.ToUint32(slide[0]),
			Name:   cast.ToString(slide[1]),
			Width:  cast.ToUint32(slide[2]),
			Height: cast.ToUint32(slide[3]),
		})
	}

	return
}

func (r *dbPresRepository) GetQuizzes(t string, pid uint64) (quizzes []domain.Quiz, err error) {
	resp, err := r.dbm.Query(queryGetQuizzes, t, pid)
	if err != nil {
		log.Warn("{GetQuizzes} in query: " + queryGetQuizzes)
		log.Error(err)
		return nil, domain.ErrDatabaseRequest
	}

	quizzes = make([]domain.Quiz, 0)
	for i, slide := range resp {
		quizzes = append(quizzes, domain.Quiz{
			Id:         cast.ToUint64(slide[0]),
			Idx:        cast.ToUint32(slide[1]),
			Type:       cast.ToString(slide[2]),
			Question:   cast.ToString(slide[3]),
			Background: cast.ToString(slide[4]),
			FontColor:  cast.ToString(slide[5]),
			FontSize:   cast.ToString(slide[6]),
			GraphColor: cast.ToString(slide[7]),
		})

		quizzes[i].Votes = make([]domain.Vote, 0)
		tresp, terr := r.dbm.Query(queryGetVotes, quizzes[i].Id)
		if terr != nil {
			log.Warn("{GetQuizzes} in query: " + queryGetVotes)
			log.Error(err)
		}
		for _, vote := range tresp {
			quizzes[i].Votes = append(quizzes[i].Votes, domain.Vote{
				Idx:    cast.ToUint32(vote[0]),
				Option: cast.ToString(vote[1]),
				Votes:  cast.ToUint64(vote[2]),
				Color:  cast.ToString(vote[3]),
			})
		}
	}

	return
}
