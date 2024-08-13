package foo

import (
	"context"
	"fmt"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
	"github.com/fatih/structs"
	"github.com/pkg/errors"
	"go-project-template/internal/app/biz/foo"
	"go-project-template/internal/pkg/database"
)

type bizFooRepo struct {
	db *database.Main
}

func New(db *database.Main) (foo.Repo, error) {
	repo := &bizFooRepo{db: db}
	err := repo.migrate()
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (repo *bizFooRepo) tableName() string {
	return "foo"
}

func (repo *bizFooRepo) migrate() error {
	return migrate(repo.db.Driver, repo.db.Conn)
}

func (repo *bizFooRepo) GetFooList(ctx context.Context) ([]foo.Foo, error) {
	cond, vals, err := builder.BuildSelect(repo.tableName(), nil, []string{"*"})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	rows, err := repo.db.Conn.QueryContext(ctx, cond, vals...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	var foos []foo.Foo
	if err := scanner.Scan(rows, &foos); err != nil {
		return nil, errors.WithStack(err)
	}

	return foos, nil
}

func (repo *bizFooRepo) GetFoo(ctx context.Context, id int64) (*foo.Foo, error) {
	cond, vals, err := builder.BuildSelect(repo.tableName(), map[string]interface{}{
		"id": id,
	}, []string{"*"})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	rows, err := repo.db.Conn.QueryContext(ctx, cond, vals...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	var foo foo.Foo
	if err := scanner.Scan(rows, &foo); err != nil {
		return nil, errors.WithStack(err)
	}

	return &foo, nil
}

func (repo *bizFooRepo) CreateFoo(ctx context.Context, foo *foo.Foo) error {
	dataMap := structs.Map(foo)
	delete(dataMap, "id")
	cond, vals, err := builder.BuildInsert(repo.tableName(), []map[string]interface{}{dataMap})
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Println(cond, vals)

	result, err := repo.db.Conn.ExecContext(ctx, cond, vals...)
	if err != nil {
		return errors.WithStack(err)
	}
	foo.ID, err = result.LastInsertId()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
