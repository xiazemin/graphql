package graph

import "context"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r *queryResolver) People(ctx context.Context, id string) (*People, error) {
	return &People{}, nil	// 替换panic（避免运行过程中退出，利于我们观察执行过程）
}


type People struct {

}