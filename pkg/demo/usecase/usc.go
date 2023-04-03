package demousc

import (
	"banana/pkg/domain"
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

	vm, err := du.demoRepo.GetPresViewMode(presId)
	if err != nil {
		return domain.CurrentDemoSlide{}, err
	}

	if vm == false {
		return domain.CurrentDemoSlide{ViewMode: vm}, nil
	}

	out.Slide, err = du.demoRepo.GetCurrentDemoSlide(presId)
	if err != nil {
		return domain.CurrentDemoSlide{}, err
	}
	
	if out.Slide.Kind == domain.SildeTypeConvertedSlide {
		out.Height = out.Slide.Height
		out.Width = out.Slide.Width
	}

	return
}
