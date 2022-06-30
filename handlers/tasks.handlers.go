package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/faztweb/golang-restapi-crud/data"
	"github.com/faztweb/golang-restapi-crud/models"
	"github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
	}

	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(data.TasksData) + 1
	data.TasksData = append(data.TasksData, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.TasksData)
}

func GetOneTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}

	for _, task := range data.TasksData {
		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updatedTask models.Task

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please Enter Valid Data")
	}
	json.Unmarshal(reqBody, &updatedTask)

	for i, t := range data.TasksData {
		if t.ID == taskID {
			data.TasksData = append(data.TasksData[:i], data.TasksData[i+1:]...)

			updatedTask.ID = t.ID
			data.TasksData = append(data.TasksData, updatedTask)

			// w.Header().Set("Content-Type", "application/json")
			// json.NewEncoder(w).Encode(updatedTask)
			fmt.Fprintf(w, "The task with ID %v has been updated successfully", taskID)
		}
	}

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid User ID")
		return
	}

	for i, t := range data.TasksData {
		if t.ID == taskID {
			data.TasksData = append(data.TasksData[:i], data.TasksData[i+1:]...)
			fmt.Fprintf(w, "The task with ID %v has been remove successfully", taskID)
		}
	}
}
