package tracker

import (
	"errors"
	"time"
)

type Status string

const (
	Todo       Status = "TODO"
	InProgress Status = "IN_PROGRESS"
	Done       Status = "DONE"
)

type Member struct {
	ID   int
	Name string
}

type Task struct {
	ID        int
	Title     string
	Desc      string
	Assignee  int
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Project struct {
	ID          int
	Name        string
	Description string
	Start       time.Time
	Tasks       []Task
	Team        []Member
}

var (
	ErrNotFound   = errors.New("not found")
	ErrDuplicated = errors.New("duplicated")
)

type Tracker struct {
	nextProjID int
	nextTaskID int
	projects   map[int]*Project
}

func NewTracker() *Tracker {
	return &Tracker{
		nextProjID: 1,
		nextTaskID: 1,
		projects:   make(map[int]*Project),
	}
}

// CreateProject Создание нового проекта в системе, возвращает id проекта
func (tr *Tracker) CreateProject(name, desc string, team []Member) int {
	id := tr.nextProjID // берем следующий доступный ID проекта
	tr.nextProjID++     // увеличиваем счетчик для следующего проекта

	// создаем новый проект с указателем для хранения в map по ссылке
	project := &Project{
		ID:          tr.nextProjID,
		Name:        name,
		Description: desc,
		Start:       time.Now(),
		Tasks:       []Task{},
		Team:        team,
	}

	tr.projects[id] = project
	return id
}

// AddTask Добавление новой таски в существующий проект, возвращает id таски
func (tr *Tracker) AddTask(projectID int, title, desc string, assignee int) (int, error) {
	project, ok := tr.projects[projectID]
	if !ok {
		return 0, ErrNotFound
	}

	id := tr.nextTaskID
	tr.nextTaskID++

	task := Task{
		ID:        id,
		Title:     title,
		Desc:      desc,
		Assignee:  assignee,
		Status:    Todo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	project.Tasks = append(project.Tasks, task)
	return id, nil
}

// UpdateTaskStatus Обновление статуса задачи
func (tr *Tracker) UpdateTaskStatus(projectID, taskID int, st Status) error {
	project, ok := tr.projects[projectID]
	if !ok {
		return ErrNotFound
	}

	for i := range project.Tasks {
		if project.Tasks[i].ID == taskID {

			project.Tasks[i].Status = st
			project.Tasks[i].UpdatedAt = time.Now()

			return nil
		}
	}

	return ErrNotFound
}

// ListTasks Список всех задач проекта определенного статуса
func (tr *Tracker) ListTasks(projectID int, st Status) []Task {
	project, ok := tr.projects[projectID]
	if !ok {
		return nil
	}

	var result []Task
	if st == "" {
		return result
	}

	for _, task := range project.Tasks {
		if task.Status == st {
			result = append(result, task)
		}
	}
	return result
}

// ProjectProgress Процент выполненных задач
func (tr *Tracker) ProjectProgress(projectID int) float64 {
	project, ok := tr.projects[projectID]
	if !ok {
		return 0
	}
	var completed, countTask int

	for _, task := range project.Tasks {
		if task.Status == Done {
			completed++
		}
		countTask++
	}
	result := (float64(completed) * 100) / float64(countTask)
	return result
}
