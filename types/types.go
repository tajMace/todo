// taj-mace
// global project types

package types

type TaskManager struct {
	NextID int    `json:"next_id"`
	Tasks  []Task `json:"tasks"`
}

type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Due  string `json:"due"`
	Done bool   `json:"done"`
}
