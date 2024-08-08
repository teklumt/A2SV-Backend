package usecase

import (
	"clean_architecture/domain"
	"errors"
)

type TaskUsecase struct {
	TaskRepo domain.TaskRepository
}

func NewTaskUsecase(taskRepo domain.TaskRepository) *TaskUsecase {
	return &TaskUsecase{TaskRepo: taskRepo}
}

func (uc *TaskUsecase) CreateTask(task domain.Task) error {
	if task.Title == "" || task.Description == "" {
		return errors.New("missing required fields")
	}
	_, err := uc.TaskRepo.CreateTask(task)
	return err
}

func (uc *TaskUsecase) GetTasks() ([]domain.Task, error) {
	return uc.TaskRepo.GetTasks()
}


func (uc *TaskUsecase) GetTaskByID(id string, creter string, Role_ string) (domain.Task, error) {
	return uc.TaskRepo.GetTaskByID(id, creter, Role_)
}


func (uc *TaskUsecase) GetMyTasks(username string) ([]domain.Task, error) {
	return uc.TaskRepo.GetMyTasks(username)
}

func (uc *TaskUsecase) DeleteTask(id string) error {
	_, err := uc.TaskRepo.DeleteTask(id)
	return err
}


func (uc *TaskUsecase) UpdateTask(id string, task domain.Task) error {
	if task.Title == "" || task.Description == "" {
		return errors.New("missing required fields")
	}
	_, err := uc.TaskRepo.UpdateTask(id, task)
	return err
}