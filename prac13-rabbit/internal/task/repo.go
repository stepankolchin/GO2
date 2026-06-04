package task

import "sync"

type Repo struct {
	mu    sync.RWMutex
	tasks map[string]Task
}

func NewRepo() *Repo {
	return &Repo{
		tasks: make(map[string]Task),
	}
}

func (r *Repo) Create(t Task) (Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tasks[t.ID] = t
	return t, nil
}
