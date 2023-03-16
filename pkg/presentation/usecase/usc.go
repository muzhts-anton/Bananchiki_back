package presusc

import (
	"context"
	"fmt"

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
	tmp, err := pu.PresClient.Split(context.Background(), &grpc.Pres{
		Url: s.Url,
		Id:  s.CreatorId,
	})
	fmt.Println(tmp, err)
	return nil
}
