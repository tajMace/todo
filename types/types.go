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

func (tm *TaskManager) NewTaskOnManager(text, due string) {
	tm.Tasks = append(tm.Tasks, NewTask(tm.NextID, text, due))
	tm.NextID++
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

func (t *Task) MarkDone() {
	t.Done = true
}
