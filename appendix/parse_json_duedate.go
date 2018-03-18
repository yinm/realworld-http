package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type DueDate struct {
	time.Time
}

func (d *DueDate) UnmarshalJSON(raw []byte) error {
	epoch, err := strconv.Atoi(string(raw))
	if err != nil {
		return err
	}
	d.Time = time.Unix(int64(epoch), 0)
	return nil
}

type ToDo struct {
	Task string  `json:"task"`
	Time DueDate `json:"due"`
}

var jsonString = []byte(`[
	{"task": "幼稚園登園", "due": 1486600200},
	{"task": "エリクソン研究会にいく", "due": 1486634400}
]`)

func main() {
	var todos []ToDo
	err := json.Unmarshal(jsonString, &todos)
	if err != nil {
		panic(err)
	}
	for _, todo := range todos {
		fmt.Printf("%s: %v\n", todo.Task, todo.Time)
	}
}
