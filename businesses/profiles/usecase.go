package profiles

import (
	"backend-golang/app/middlewares"
	"context"
)

type profileUsecase struct {
	profileRepository Repository
	jwtAuth           *middlewares.JWTConfig
}

func NewProfileUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &profileUsecase{
		profileRepository: repository,
		jwtAuth:           jwtAuth,
	}
}

func (usecase *profileUsecase) GetByID(ctx context.Context, id string) (Domain, error) {
	return usecase.profileRepository.GetByID(ctx, id)
}

func (usecase *profileUsecase) UpdateProfileUser(ctx context.Context, profileDomain *Domain, id string) (Domain, error) {
	return usecase.profileRepository.UpdateProfileUser(ctx, profileDomain, id)
}

func (usecase *profileUsecase) UploadProfileImage(ctx context.Context, profileDomain *Domain, avatarPath string, id string) (Domain, string, error) {
	return usecase.profileRepository.UploadProfileImage(ctx, profileDomain, avatarPath, id)
}
