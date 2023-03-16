package presusc

import (
	"context"

	"banana/pkg/domain"
	"banana/pkg/presentation/delivery/grpc"
)

type PresUsecase struct {
	PresClient grpc.ParsingClient
}

func InitPresUscase(uc grpc.ParsingClient) *PresUsecase {
	return &PresUsecase{
		PresClient: uc,
	}
}

func (pu *PresUsecase) UploadPres(s *domain.Presentation) error {
	pu.PresClient.Split(context.Background(), &grpc.Pres{
		Url: s.Url,
		Id:  s.CreatorId,
	})
	return nil
}
