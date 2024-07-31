package db

import "task_manager/model"

var Database  = []model.Task{
	{
		ID:          "1",
		Title:       "Task 1",
		Description: "This is Task 1",
		Status:      "Pending",
	},
	{
		ID:          "2",
		Title:       "Task 2",
		Description: "This is Task 2",
		Status:      "Completed",
	},
	{
		ID:          "3",
		Title:       "Task 3",
		Description: "This is Task 3",
		Status:      "Pending",
	},
	
	{
		ID:          "4",
		Title:       "Task 4",
		Description: "This is Task 4",
		Status:      "Completed",
	},
}