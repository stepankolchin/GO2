package graph

import (
	"context"
	"fmt"

	"example.com/prac11-graphql/graph/model"
)

// QueryResolver implementation
func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	return Tasks, nil
}

func (r *queryResolver) Task(ctx context.Context, id string) (*model.Task, error) {
	for _, t := range Tasks {
		if t.ID == id {
			return t, nil
		}
	}
	return nil, nil
}

// MutationResolver implementation
func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTaskInput) (*model.Task, error) {
	id := fmt.Sprintf("t_%03d", len(Tasks)+1)
	task := &model.Task{
		ID:          id,
		Title:       input.Title,
		Description: input.Description,
		Done:        false,
	}
	Tasks = append(Tasks, task)
	return task, nil
}

func (r *mutationResolver) UpdateTask(ctx context.Context, id string, input model.UpdateTaskInput) (*model.Task, error) {
	for _, t := range Tasks {
		if t.ID == id {
			if input.Title != nil {
				t.Title = *input.Title
			}
			if input.Description != nil {
				t.Description = input.Description
			}
			if input.Done != nil {
				t.Done = *input.Done
			}
			return t, nil
		}
	}
	return nil, fmt.Errorf("task not found")
}

func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (bool, error) {
	for i, t := range Tasks {
		if t.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}
