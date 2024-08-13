package foo

import "context"

type Repo interface {
	GetFooList(ctx context.Context) ([]Foo, error)
	GetFoo(ctx context.Context, id int64) (*Foo, error)
	CreateFoo(ctx context.Context, foo *Foo) error
}

type Biz struct {
	repo Repo
}

func NewBiz(repo Repo) *Biz {
	return &Biz{
		repo: repo,
	}
}

func (biz *Biz) CreateFoo(ctx context.Context, foo *Foo) error {
	return biz.repo.CreateFoo(ctx, foo)
}

func (biz *Biz) GetFooList(ctx context.Context) ([]Foo, error) {
	return biz.repo.GetFooList(ctx)
}

func (biz *Biz) GetFoo(ctx context.Context, id int64) (*Foo, error) {
	return biz.repo.GetFoo(ctx, id)
}
