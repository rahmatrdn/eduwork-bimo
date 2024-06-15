package models

import (
	"database/sql"
	"eduwork-bimo/Final-Project-Task-Manager/entities"
	_ "github.com/go-sql-driver/mysql"
)

type TaskModel struct {
	db *sql.DB
}

func NewTaskModel(db *sql.DB) *TaskModel {
	return &TaskModel{db: db}
}

func (m TaskModel) GetAll(userID int) ([]entities.Task, error) {
	rows, err := m.db.Query("SELECT id, title, completed, user_id FROM tasks WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []entities.Task{}
	for rows.Next() {
		var task entities.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Completed, &task.UserID)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (m TaskModel) Create(task *entities.Task) error {
	result, err := m.db.Exec("INSERT INTO tasks (title, completed, user_id) VALUES (?, ?, ?)", task.Title, task.Completed, task.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	task.ID = int(id)
	return nil
}

func (m TaskModel) Update(task *entities.Task) error {
	_, err := m.db.Exec("UPDATE tasks SET title = ?, completed = ? WHERE id = ? AND user_id = ?", task.Title, task.Completed, task.ID, task.UserID)
	return err
}

func (m TaskModel) Delete(taskID int, userID int) error {
	_, err := m.db.Exec("DELETE FROM tasks WHERE id = ? AND user_id = ?", taskID, userID)
	if err != nil {
		return err
	}
	return nil
}
