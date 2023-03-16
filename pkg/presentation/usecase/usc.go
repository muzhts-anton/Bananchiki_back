package presusc

import (
	"context"

	"banana/pkg/domain"
	"banana/pkg/presentation/delivery/grpc"
)

type presUsecase struct {
	presClient grpc.ParsingClient
}

func (pu *presUsecase) uploadPres(s *domain.Presentation) error {
	pu.presClient.Split(context.Background(), &grpc.Pres{
		Url: s.Url,
		Id: s.CreatorId,
	})
	return nil
}
