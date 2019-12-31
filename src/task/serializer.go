package task

type TaskResponse struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	UUID        string `json:"uuid"`
}

func (task *Task) Response() TaskResponse {
	return TaskResponse{
		Description: task.Description,
		Done:        task.Done,
		Name:        task.Name,
		Slug:        task.Slug,
		UUID:        task.UUID,
	}
}

func Map(tasks []Task) []TaskResponse {
	mappedTasks := make([]TaskResponse, 0, len(tasks))

	for _, task := range tasks {
		mappedTasks = append(mappedTasks, task.Response())
	}

	return mappedTasks
}
