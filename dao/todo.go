package dao

import (
	"github.com/TinyKitten/sugoibot/models"
)

// AddTask タスクを追加する
func AddTask(memberCode, taskName string) (id int64, err error) {
	db, err := getDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into todo(member_code, task_name, completed) values(?,?,?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(memberCode, taskName, 0)
	if err != nil {
		return 0, err
	}
	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// UpdateTaskStatus Todoを追加する
func UpdateTaskStatus(id int64, flag bool) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("update todo set completed=? where id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(flag, id)
	if err != nil {
		return err
	}
	return nil
}

// GetAllTask タスクを全件取得する
func GetAllTask() (*[]models.Todo, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select id, member_code, task_name, completed from todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	todos := []models.Todo{}
	for rows.Next() {
		todo := models.Todo{}
		err = rows.Scan(&todo.ID, &todo.MemberCode, &todo.TaskName, &todo.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return &todos, nil
}

// GetTaskByMemberCode メンバーコードに該当するTODOを全件取得する
func GetTaskByMemberCode(memberCode string) (*[]models.Todo, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select id, member_code, task_name, completed from todo where member_code='" + memberCode + "'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	todos := []models.Todo{}
	for rows.Next() {
		todo := models.Todo{}
		err = rows.Scan(&todo.ID, &todo.MemberCode, &todo.TaskName, &todo.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return &todos, nil
}

// GetTaskByID メンバーコードに該当するTODOを全件取得する
func GetTaskByID(id int64) (*models.Todo, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare("select id, member_code, task_name, completed from todo where id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	todo := &models.Todo{}

	err = stmt.QueryRow(id).Scan(&todo.ID, &todo.MemberCode, &todo.TaskName, &todo.Completed)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// DeteleTask タスクを削除する
func DeteleTask(id int64) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("delete from todo where id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
