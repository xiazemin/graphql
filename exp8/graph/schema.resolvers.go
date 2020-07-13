package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/xiazemin/graphql/exp8/graph/generated"
	"github.com/xiazemin/graphql/exp8/graph/model"
)

func (r *queryResolver) People(ctx context.Context, id string) (*model.People, error) {
	//return &model.People{}, nil	// 替换panic（避免运行过程中退出，利于我们观察执行过程）
	//panic(fmt.Errorf("not implemented"))

	// GetPeopleByID即使我们要实现的获取people的函数
	err, people := model.GetPeopleByID(id)
	fmt.Println(err,id)
	return people, err
}

func (r *queryResolver) Peoples(ctx context.Context, first *int, after *string) (*model.PeopleConnection, error) {
	//panic(fmt.Errorf("not implemented"))
	//from := -1
	if after != nil {
		b, err := base64.StdEncoding.DecodeString(*after)
		if err != nil {
			return &model.PeopleConnection{}, err
		}
		_,err = strconv.Atoi(strings.TrimPrefix(string(b), "cursor"))
		if err != nil {
			return &model.PeopleConnection{}, err
		}
		//from = i
	}
	count := 0
	startID := ""
	hasPreviousPage := true
	hasNextPage := true
	// 获取edges
	edges := []*model.PeopleEdge{}
	/*
	db, err := bolt.Open("./data/data.db", 0600, nil)
	defer db.Close()
	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(peopleBucket)).Cursor()

		判断是否还有前向页
		k, v := c.First()
		if from == -1 || strconv.Itoa(from) == string(k) {
			startID = string(k)
			hasPreviousPage = false
		}

		if from == -1 {
			for k, _ := c.First(); k != nil; k, _ = c.Next() {
				_, people := GetPeopleByID(string(k), db)
				edges = append(edges, PeopleEdge{
					Node:   people,
					Cursor: encodeCursor(string(k)),
				})
				count++
				if count == *first {
					break
				}
			}
		} else {
			for k, _ := c.First(); k != nil; k, _ = c.Next() {
				if strconv.Itoa(from) == string(k) {
					k, _ = c.Next()
					startID = string(k)
				}
				if startID != "" {
					_, people := GetPeopleByID(string(k), db)
					edges = append(edges, PeopleEdge{
						Node:   people,
						Cursor: encodeCursor(string(k)),
					})
					count++
					if count == *first {
						break
					}
				}
			}
		}

		k, v = c.Next()
		if k == nil && v == nil {
			hasNextPage = false
		}
		return nil
	})
	*/
	edges = append(edges, &model.PeopleEdge{
		Node:   &model.People{
			ID:"7896",
		},
		Cursor: encodeCursor(string("123456")),
	})
	count++

	if count == 0 {
		return &model.PeopleConnection{}, nil
	}
	// 获取pageInfo
	pageInfo := model.PageInfo{
		HasPreviousPage: hasPreviousPage,
		HasNextPage:     hasNextPage,
		StartCursor:     encodeCursor(startID),
		EndCursor:       encodeCursor(edges[count-1].Node.ID),
	}

	return &model.PeopleConnection{
		PageInfo:   &pageInfo,
		Edges:      edges,
		TotalCount: count,
	}, nil
}

// 编码游标（游标指向当前节点）
func encodeCursor(k string) string {
	i, _ := strconv.Atoi(k)
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("cursor%d", i)))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
