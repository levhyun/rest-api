package service

import (
	"context"
	"rest-api/domain/user/domain/ent"
	"rest-api/domain/user/domain/repository"
	"rest-api/domain/user/presentation/dto"
	"strconv"
)

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

type UserService struct {
	userRepository *repository.UserRepository
}

func (ur *UserService) Create(ctx context.Context, userReqeust *dto.UserRequest) error {
	_, err := ur.userRepository.Save(ctx, dto.ToEntity(0, userReqeust))
	return err
}

func (ur *UserService) ReadAll(ctx context.Context) ([]*ent.User, error) {
	user, err := ur.userRepository.FindAll(ctx)
	return user, err
}

func (ur *UserService) Read(ctx context.Context, id string) (*ent.User, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	user, err := ur.userRepository.FindById(ctx, intId)
	return user, err
}

func (ur *UserService) Update(ctx context.Context, id string, userReqeust *dto.UserRequest) error {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	_, err2 := ur.userRepository.Update(ctx, dto.ToEntity(intId, userReqeust))
	return err2
}

func (ur *UserService) Delete(ctx context.Context, id string) error {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return ur.userRepository.RemoveById(ctx, intId)
}
