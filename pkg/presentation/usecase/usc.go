package presusc

import (
	"banana/pkg/domain"
	"banana/pkg/presentation/delivery/grpc"
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

func (pu *PresUsecase) CreatePres(url string) (uint64, error) {
	var cid uint64 = 1
	presId, err := pu.presRepo.CreatePres(cid)
	if err != nil {
		return 0, err
	}

	gslides, err := pu.presGrpcClient.Split(context.Background(), &grpc.Pres{
		Url: url,
		Id:  presId,
	})
	if err != nil {
		log.Error(err)
	}

	err = pu.presRepo.UpdatePresUrl(presId, gslides.Url)
	if err != nil {
		return 0, err
	}

	slides := make([]domain.ConvertedSlide, 0)
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

	p.QuizNum = pres.QuizNum
	p.SlideNum = pres.SlideNum
	p.Url = pres.Url

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
		if tmpidx == pres.Quizzes[qi].Idx {
			p.Slides = append(p.Slides, domain.SlideApiResponse{
				Idx:        pres.Quizzes[qi].Idx,
				Kind:       domain.SlideTypeQuiz,
				QuizId:     pres.Quizzes[qi].Id,
				Type:       pres.Quizzes[qi].Type,
				Question:   pres.Quizzes[qi].Question,
				Vote:       pres.Quizzes[qi].Votes,
				Background: pres.Quizzes[qi].Background,
				FontColor:  pres.Quizzes[qi].FontColor,
				FontSize:   pres.Quizzes[qi].FontSize,
				GraphColor: pres.Quizzes[qi].GraphColor,
			})
			qi++
		} else {
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

	return
}
