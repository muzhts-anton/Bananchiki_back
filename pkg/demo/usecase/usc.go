package demousc

import (
	"banana/pkg/domain"
	"banana/pkg/utils/log"
)

type demoUsecase struct {
	demoRepo domain.DemoRepository
}

func InitDemoUsc(dr domain.DemoRepository) domain.DemoUsecase {
	return &demoUsecase{
		demoRepo: dr,
	}
}

func (du demoUsecase) GetCurrentDemoSlide(hash string) (out domain.CurrentDemoSlide, err error) {
	presId, err := du.demoRepo.GetPresIdByHash(hash)
	if err != nil {
		return domain.CurrentDemoSlide{}, err
	}

	out.ViewMode, err = du.demoRepo.GetPresViewMode(presId)
	if err != nil {
		return domain.CurrentDemoSlide{}, err
	}

	if out.ViewMode == false {
		return
	}

	out.Slide, err = du.demoRepo.GetCurrentDemoSlide(presId)
	if err != nil {
		return domain.CurrentDemoSlide{}, err
	}

	if out.Slide.Kind == domain.SildeTypeConvertedSlide {
		out.Height = out.Slide.Height
		out.Width = out.Slide.Width
		out.Url = domain.PresentationSlidesPath
	}

	out.Emotions, err = du.demoRepo.GetPresEmotions(presId)
	if err != nil {
		return domain.CurrentDemoSlide{}, err
	}

	out.Questions, err = du.demoRepo.GetPresQuestions(presId)
	if err != nil {
		return domain.CurrentDemoSlide{}, err
	}

	return
}

func (du demoUsecase) ShowDemoGo(presId, userId uint64, idx uint32) error {
	cid, err := du.demoRepo.GetPresCreatorId(presId)
	if err != nil {
		return err
	}

	if cid != userId {
		return domain.ErrPermissionDenied
	}

	viewMode, err := du.demoRepo.GetViewMode(presId)
	if err != nil{
		return err
	}

	log.Debug("viewMode")

	if !viewMode{
		log.Debug("ZeroingVotes")
		du.demoRepo.SetAllVotes(presId, 0)
	}

	return du.demoRepo.DemoGo(presId, idx)
}

func (du demoUsecase) ShowDemoSop(presId, userId uint64) error {
	cid, err := du.demoRepo.GetPresCreatorId(presId)
	if err != nil {
		return err
	}

	if cid != userId {
		return domain.ErrPermissionDenied
	}

	err = du.demoRepo.ZeroingReactions(presId)
	if err != nil {
		return err
	}

	err = du.demoRepo.SetAllVotes(presId, 1)
	if err != nil {
		return err
	}

	err = du.demoRepo.DeletePresQuestions(presId)
	if err != nil {
		return err
	}

	return du.demoRepo.DemoStop(presId)
}
