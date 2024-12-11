package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type App struct {
	Data *Data
}

func NewApp(data *Data) *App {
	if data == nil {
		data = &Data{
			Todos: []todo{},
		}
	}
	return &App{Data: data}
}

func (a *App) Add(task string) error {
	t := todo{
		Task:   task,
		Status: false,
	}
	a.Data.Todos = append(a.Data.Todos, t)
	return nil
}

func (a *App) Complete(task string) error {
	for i, t := range a.Data.Todos {
		if t.Task == task {
			a.Data.Todos[i].Status = true
			return nil
		}
	}
	return errors.New("task doesn't exsit")
}

func (a *App) Read() {
	status := map[bool]string{
		false: "not done",
		true:  "completed",
	}
	for i, t := range a.Data.Todos {
		fmt.Printf("%d: (task: %s) (status: %s)\n", i, t.Task, status[t.Status])
	}
}

func (a *App) Delete(task string) error {
	for i, t := range a.Data.Todos {
		if t.Task == task {
			a.Data.Todos = append(a.Data.Todos[:i], a.Data.Todos[i+1:]...)
			return nil
		}
	}
	return errors.New("task doesn't exit")
}

func (a *App) NewDump() func() error {
	const filename = "todos.json"

	return func() error {
		f, err := os.Create(filename)
		if err != nil {
			return errors.New("Failed to open file")
		}
		defer f.Close()

		if err := json.NewEncoder(f).Encode(a.Data); err != nil {
			return errors.New("Failed to save data")
		}
		return nil
	}
}

func (a *App) NewLoad() func() error {
	const filename = "todos.json"

	return func() error {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()

		if err := json.NewDecoder(f).Decode(&a.Data); err != nil {
			return err
		}
		return nil
	}
}
