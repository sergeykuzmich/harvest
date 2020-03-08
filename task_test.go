package harvest_api_client

import (
	"net/http"
	"testing"
)

func TestGetTask(t *testing.T) {
	a := testAPI()
	taskResponse := mockResponse("tasks", "task-example.json")
	a.BaseURL = taskResponse.URL
	task, err := a.GetTask(2086199, Defaults())
	if err != nil {
		t.Fatal(err)
	}

	if task == nil {
		t.Fatal("testTask() returned nil instead of task")
	}
	if task.Name != "Admin" {
		t.Errorf("Incorrect Task Name '%s'", task.Name)
	}
	if task.ID != 2086199 {
		t.Errorf("Incorrect Task ID '%v'", task.ID)
	}
}

func TestGetTasks(t *testing.T) {
	a := testAPI()
	taskResponse := mockResponse("tasks", "tasks-example.json")
	a.BaseURL = taskResponse.URL
	tasks, err := a.GetTasks(Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(tasks) != 2 {
		t.Errorf("Incorrect number of tasks '%v'", len(tasks))
	}
	if tasks[0].Name != "Admin" {
		t.Errorf("Incorrect Task Name '%s'", tasks[0].Name)
	}
	if tasks[1].ID != 2086200 {
		t.Errorf("Incorrect Task ID '%v'", tasks[1].ID)
	}
}

func TestCreateTask(t *testing.T) {
	a := testAPI()
	taskResponse := mockResponse("tasks", "2086200.json")
	a.BaseURL = taskResponse.URL

	valid_task := Task{
		Name: "Test",
	}

	_, err := a.CreateTask(&valid_task, Defaults())
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateInvalidTask(t *testing.T) {
	a := testAPI()
	taskResponse := mockErrorResponse(http.StatusUnprocessableEntity)
	a.BaseURL = taskResponse.URL

	invalid_task := Task{}

	_, err := a.CreateTask(&invalid_task, Defaults())
	if err == nil {
		t.Fatal("An error expected")
	}
}

func TestUpdateTask(t *testing.T) {
	a := testAPI()
	taskResponse := mockResponse("tasks", "2086200-updated.json")
	a.BaseURL = taskResponse.URL

	input_task := Task{
		ID:   2086200,
		Name: "New Name",
	}

	task, err := a.UpdateTask(&input_task, Defaults())
	if err != nil {
		t.Fatal(err)
	}

	if task.Name != input_task.Name {
		t.Fatal("Task weren't updated")
	}
}

func TestDeleteTask(t *testing.T) {
	a := testAPI()
	taskResponse := mockResponse("tasks", "2086200-DELETE.json")
	a.BaseURL = taskResponse.URL
	err := a.DeleteTask(2086199, Defaults())
	if err != nil {
		t.Fatal(err)
	}
}
