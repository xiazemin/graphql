package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/xiazemin/graphql/exp11/graph/generated"
	"github.com/xiazemin/graphql/exp11/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	//panic(fmt.Errorf("not implemented"))
	return &model.Todo{input.UserID, input.Text, true, nil}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	//panic(fmt.Errorf("not implemented"))
	return []*model.Todo{&model.Todo{"123", "test todo", true, nil}}, nil
}

func (r *queryResolver) Todo(ctx context.Context) (*model.Todo, error) {
	//panic(fmt.Errorf("not implemented"))
	return &model.Todo{"1234", "to do func", true, nil}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
