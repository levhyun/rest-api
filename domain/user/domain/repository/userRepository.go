package repository

import (
	"context"
	"rest-api/domain/user/domain/ent"
)

func NewUserRepository(db *ent.Client) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type UserRepository struct {
	db *ent.Client
}

func (ur *UserRepository) Save(ctx context.Context, user *ent.User) (*ent.User, error) {
	return ur.db.User.Create().SetName(user.Name).Save(ctx)
}

func (ur *UserRepository) FindAll(ctx context.Context) ([]*ent.User, error) {
	return ur.db.User.Query().All(ctx)
}

func (ur *UserRepository) FindById(ctx context.Context, id int) (*ent.User, error) {
	return ur.db.User.Get(ctx, id)
}

func (ur *UserRepository) Update(ctx context.Context, user *ent.User) (*ent.User, error) {
	return ur.db.User.UpdateOneID(user.ID).SetName(user.Name).Save(ctx)
}

func (ur *UserRepository) RemoveById(ctx context.Context, id int) error {
	return ur.db.User.DeleteOneID(id).Exec(ctx)
}
