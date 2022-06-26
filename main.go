package main

import (
	"fmt"
	"go.uber.org/dig"
	"github.com/google/uuid"
)

// DIの実験
// Uber/digがsingletonとしてDIしてくれるかを見ている
// 複数のインスタンスから別々のRepositoryのインスタンスが生成されてしまうと、
// InMemoryではデータを正しく操作できない

type (
	Foo struct {
		repo	IRepo
	}

	Bar struct {
		repo	IRepo
	}
)

type IFoo interface {
	use()	string
}

type IBar interface {
	use()	string
}

type IRepo interface {
	call()	string
}

func NewFoo(repo IRepo) IFoo {
	return Foo {
		repo:	repo,
	}
}

func NewBar(repo IRepo) IBar {
	return Bar {
		repo:	repo,
	}
}

func (foo Foo) use() string {
	return foo.repo.call()
}

func (bar Bar) use() string {
	return bar.repo.call()
}

type Repo struct{
	id		uuid.UUID
}

func NewRepo() IRepo {
	return Repo {
		id:		uuid.New(),
	}
}

func (repo Repo) call() string {
	return repo.id.String()
}

func main() {
	c := dig.New()
	c.Provide(NewFoo)
	c.Provide(NewBar)
	c.Provide(NewRepo)
	err := c.Invoke(func(foo IFoo) {
		fmt.Println(foo.use())
	})
	if err != nil {
		fmt.Println(err)
	}
	err2 := c.Invoke(func(bar IBar) {
		fmt.Println(bar.use())
	})
	if err2 != nil {
		fmt.Println(err2)
	}
}
