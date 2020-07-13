package mutation
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

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)


