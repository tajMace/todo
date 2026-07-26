// taj-mace
// global project types

package types

/* ===== TaskManager + Helpers ====== */
type TaskManager struct {
	NextID int    `json:"next_id"`
	Tasks  []Task `json:"tasks"`
}

func NewTaskManager(tasks []Task) TaskManager {
	return TaskManager{NextID: len(tasks), Tasks: tasks}
}

/* ===== Task + Helpers ===== */
type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Due  string `json:"due"`
	Done bool   `json:"done"`
}

// constructor - always sets done to 'false'
func NewTask(id int, text, due string) Task {
	return Task{
		ID:   id,
		Text: text,
		Due:  due,
		Done: false,
	}
}
