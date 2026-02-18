package main

import (
	"basics/data_structure/tracker_lite/tracker"
	"fmt"
)

func main() {
	tr := tracker.NewTracker()

	pid := tr.CreateProject("Landing Site", "Promo page", nil)
	t1, _ := tr.AddTask(pid, "HTML", "Make markup", 0)
	_, _ = tr.AddTask(pid, "CSS", "Style page", 0)

	_ = tr.UpdateTaskStatus(pid, t1, tracker.Done)
	fmt.Println(tr.ProjectProgress(pid)) // 50

	// Возвращает список задач, которые находятся в процессе выполнения
	for _, t := range tr.ListTasks(pid, tracker.Todo) {
		fmt.Println(t.Title, t.Status)
	}
}
