package graph

import "example.com/prac11-graphql/graph/model"

// In-memory хранилище
var Tasks = []*model.Task{
	{ID: "t_001", Title: "Первая задача", Description: strPtr("Учебный пример"), Done: false},
	{ID: "t_002", Title: "Вторая задача", Description: strPtr("GraphQL API"), Done: true},
}

func strPtr(s string) *string {
	return &s
}
