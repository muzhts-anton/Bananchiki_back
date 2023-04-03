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

		out.Kind = domain.SlideTypeQuiz
		out.QuizId = cast.ToUint64(resp[0][0])
		out.Type = cast.ToString(resp[0][1])
		out.Question = cast.ToString(resp[0][2])
		out.Background = cast.ToString(resp[0][3])
		out.FontColor = cast.ToString(resp[0][4])
		out.FontSize = cast.ToString(resp[0][5])
		out.GraphColor = cast.ToString(resp[0][6])
	}

	return
}
