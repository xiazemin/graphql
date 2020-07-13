package schema
import (
	"errors"
	"github.com/graphql-go/graphql"
)
var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"addGoods":&graphql.Field{
				Type:goodsType,
				Args:graphql.FieldConfigArgument{
					"input":&graphql.ArgumentConfig{
						Type:goodsInputType,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					input,isOk := p.Args["input"].(map[string]string)
					if !isOk{
						err := errors.New("Field 'addGoods' is missing required arguments: input. ")
						return nil,err
					}
					result := Goods{
						Name:input["name"].(string),
						Price:input["price"].(float64),
						Url:input["url"].(string),
					}
					// 处理数据
					return result,err
				},
			},
		},
	},
)

var goodsInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "goodsInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"price": &graphql.InputObjectFieldConfig{
				Type: graphql.Float,
			},
			"url": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// 无需处理参数
			"goodsList": &graphql.Field{
				Type:goodsListType,
				// 处理结构体的回调函数，直接返回处理完成的结构体即可
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return result, err
				},
			},
			// 参数是id
			"goods": &graphql.Field{
				Type: goodsType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// 获取参数
					idQuery, isOK := p.Args["id"].(string)
					if isOK {
						return result, nil
					}
					err := errors.New("Field 'goods' is missing required arguments: id. ")
					return nil, err
				},
			},
		},
	},
)

type Goods struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price float64`json:"price"`
	Url   string `json:"url"`
}

var goodsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Goods",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
			"url": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
var goodsListType = graphql.NewList(goodsType)