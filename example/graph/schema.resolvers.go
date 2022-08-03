package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/graph/generated"
	"example/graph/model"
	"example/utils"
	"fmt"
	"math/rand"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.InputUser) (*model.User, error) {
	utils.Info("[CreateUser] start, req: %s!", utils.ToJsonString(input))
	user, err := r.Query().User(ctx, input.Name)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, fmt.Errorf(`already create user, name: %s`, input.Name)
	}
	user = &model.User{
		ID:     fmt.Sprintf("%d", rand.Int()),
		Name:   input.Name,
		Height: input.Height,
		Hobby:  input.Hobby,
		Age:    input.Age,
	}
	r.users = append(r.users, user)
	return user, nil
}

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.InputTodo) (*model.Todo, error) {
	utils.Info("[CreateTodo] start, req: %s!", utils.ToJsonString(input))
	user, err := r.Query().User(ctx, input.UserName)
	if err != nil {
		return nil, fmt.Errorf(`not found user, name: %s`, input.UserName)
	}
	todo := &model.Todo{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Text: input.Text,
		User: user,
	}
	if input.Done != nil {
		todo.Done = *input.Done
	}
	if input.Extra != nil {
		todo.Extra = &model.ExtraInfo{
			RemindTime: &input.Extra.RemindTime,
		}
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	utils.Info("[Todo] start, id: %s!", id)
	for _, elem := range r.todos {
		if elem.ID == id {
			return elem, nil
		}
	}
	return nil, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, name string) (*model.User, error) {
	utils.Info("[User] start, name: %s!", name)
	for _, elem := range r.users {
		if elem.Name == name {
			return elem, nil
		}
	}
	return nil, nil
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, name string) (*model.User, error) {
	return r.User(ctx, name)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
