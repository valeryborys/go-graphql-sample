package schema

import (
	"vborys/repo"

	"github.com/graphql-go/graphql"
)

var User = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"Id": &graphql.Field{
				Type: graphql.String,
			},
			"FirstName": &graphql.Field{
				Type: graphql.String,
			},
			"LastName": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var UserInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "UserInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"FirstName": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"LastName": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

func DefineSchema() graphql.SchemaConfig {
	return graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"users": &graphql.Field{
					Name:    "users",
					Type:    graphql.NewList(User),
					Resolve: repo.FindAllUsers,
				},
				"user": &graphql.Field{
					Name: "user",
					Type: User,
					Args: graphql.FieldConfigArgument{
						"Id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: repo.GetById,
				},
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"addUser": &graphql.Field{
					Name:    "addUser",
					Type:    User,
					Resolve: repo.CreateUser,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{
							Type: UserInput,
						},
					},
				},
				"updateUser": &graphql.Field{
					Name:    "updateUser",
					Type:    User,
					Resolve: repo.UpdateUser,
					Args: graphql.FieldConfigArgument{
						"Id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"FirstName": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"LastName": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
				},
				"deleteUser": &graphql.Field{
					Name:    "deleteUser",
					Type:    graphql.String,
					Resolve: repo.DeleteById,
					Args: graphql.FieldConfigArgument{
						"Id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
				},
			},
		})}
}
