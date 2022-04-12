package graph

import "go-graphql-todo/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos     []*model.Todo
	users     []*model.User
	usersChan chan *model.User

	lastTodoId int
	lastUserId int
}

func NewResolver() *Resolver {
	users := make([]*model.User, 0)
	users = append(users, &model.User{ID: "1", Name: "phillip"})
	users = append(users, &model.User{ID: "2", Name: "hans"})
	return &Resolver{
		todos:      make([]*model.Todo, 0),
		users:      users,
		lastUserId: 3,
		usersChan:  make(chan *model.User),
	}
}
