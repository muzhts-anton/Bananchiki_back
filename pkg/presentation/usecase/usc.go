package presusc

import (
	"banana/pkg/domain"
	"banana/pkg/presentation/delivery/grpc"
	"banana/pkg/utils/hash"
	"banana/pkg/utils/log"

	"context"
)

type PresUsecase struct {
	presGrpcClient grpc.ParsingClient
	presRepo       domain.PresRepository
}

func InitPresUscase(gc grpc.ParsingClient, pr domain.PresRepository) *PresUsecase {
	return &PresUsecase{
		presGrpcClient: gc,
		presRepo:       pr,
	}
}

func (pu *PresUsecase) CreatePres(name string, cid uint64) (uint64, error) {
	presId, err := pu.presRepo.CreatePres(cid)
	if err != nil {
		return 0, err
	}

	gslides, err := pu.presGrpcClient.Split(context.Background(), &grpc.Pres{
		Name: name,
		Id:   presId,
	})
	if err != nil {
		log.Error(err)
	}
	if gslides == nil {
		return 0, domain.ErrGrpc
	}

	slides := make([]domain.ConvertedSlide, 0, gslides.Num)
	for _, s := range gslides.Slide {
		slides = append(slides, domain.ConvertedSlide{
			Name:   s.Name,
			Idx:    s.Idx,
			Width:  s.ImageWidth,
			Height: s.ImageHeight,
		})
	}

	err = pu.presRepo.CreateCovertedSlides(presId, slides)
	if err != nil {
		return 0, err
	}

	return presId, nil
}

func (pu *PresUsecase) GetPres(cid, pid uint64) (p domain.PresApiResponse, err error) {
	pres, err := pu.presRepo.GetPres(pid)
	if err != nil {
		return domain.PresApiResponse{}, err
	}
	if pres.CreatorId != cid {
		return domain.PresApiResponse{}, domain.ErrPermissionDenied
	}

	p.QuizNum = pres.QuizNum
	p.SlideNum = pres.SlideNum
	p.Url = pres.Url
	p.Code = pres.Code
	p.Name = pres.Name
	p.Hash = hash.EncodeToHash(p.Code)

	pres.Slides, err = pu.presRepo.GetConvertedSlides(domain.SildeTypeConvertedSlide, pres.Id)
	if err != nil {
		return domain.PresApiResponse{}, err
	}

	pres.Quizzes, err = pu.presRepo.GetQuizzes(domain.SlideTypeQuiz, pres.Id)
	if err != nil {
		return domain.PresApiResponse{}, err
	}

	p.Slides = make([]domain.SlideApiResponse, 0)
	var totalidx uint32 = uint32(len(pres.Slides) + len(pres.Quizzes))
	var tmpidx, qi, si uint32
	for tmpidx < totalidx {
		if qi < uint32(len(pres.Quizzes)) && tmpidx == pres.Quizzes[qi].Idx {
			kind := ""
			if pres.Quizzes[qi].AnswerTime > 0 || pres.Quizzes[qi].Cost > 0 {
				kind = domain.SlideTypeTimedQuiz
			} else {
				kind = domain.SlideTypeQuiz
			}
			p.Slides = append(p.Slides, domain.SlideApiResponse{
				Idx:         pres.Quizzes[qi].Idx,
				Kind:        kind,
				QuizId:      pres.Quizzes[qi].Id,
				Type:        pres.Quizzes[qi].Type,
				Question:    pres.Quizzes[qi].Question,
				Vote:        pres.Quizzes[qi].Votes,
				Background:  pres.Quizzes[qi].Background,
				FontColor:   pres.Quizzes[qi].FontColor,
				FontSize:    pres.Quizzes[qi].FontSize,
				GraphColor:  pres.Quizzes[qi].GraphColor,
				Runout:      pres.Quizzes[qi].Runout,
				AnswerTime:  pres.Quizzes[qi].AnswerTime,
				ResultAfter: pres.Quizzes[qi].ResultAfter,
				Cost:        pres.Quizzes[qi].Cost,
				ExtraPts:    pres.Quizzes[qi].ExtraPts,
			})
			qi++
		} else if si < uint32(len(pres.Slides)) {
			p.Slides = append(p.Slides, domain.SlideApiResponse{
				Idx:    pres.Slides[si].Idx,
				Kind:   domain.SildeTypeConvertedSlide,
				Name:   pres.Slides[si].Name,
				Width:  pres.Slides[si].Width,
				Height: pres.Slides[si].Height,
			})
			si++
		}
		tmpidx++
	}

	if p.SlideNum != 0 {
		p.Width = p.Slides[0].Width
		p.Height = p.Slides[0].Height
	}

	return
}

func (pu *PresUsecase) ChangePresName(uid, pid uint64, name string) error {
	ownerId, err := pu.presRepo.GetPresOwner(pid)
	if err != nil {
		return err
	}

	if ownerId != uid {
		return domain.ErrPermissionDenied
	}

	if len(name) > 128 {
		return domain.ErrInvalidPresName
	}

	return pu.presRepo.ChangePresName(pid, name)
}

func (pu *PresUsecase) DeletePres(uid, pid uint64) error {
	ownerId, err := pu.presRepo.GetPresOwner(pid)
	if err != nil {
		return err
	}

	if ownerId != uid {
		return domain.ErrPermissionDenied
	}

	return pu.presRepo.DeletePres(pid)
}
