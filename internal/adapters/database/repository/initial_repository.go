package repository

// This is an example repository, you can delete it if you wish.

import "context"

type InitialRepository struct {
}

func NewInitialRepository() *InitialRepository {
	return &InitialRepository{}
}

func (i *InitialRepository) Initial(ctx context.Context) {

}
