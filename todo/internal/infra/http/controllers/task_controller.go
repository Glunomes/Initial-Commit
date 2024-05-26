package controllers

import (
	"log"
	"net/http"
	"strconv"
	"fmt"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
	"github.com/go-chi/chi/v5"
)

type TaskConroller struct {
	taskService app.TaskService
}

func NewTaskController(ts app.TaskService) TaskConroller {
	return TaskConroller{
		taskService: ts,
	}
}

func (c TaskConroller) Save() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
				user := r.Context().Value(UserKey).(domain.User)

				task, err := requests.Bind(r, requests.TaskRequest{}, domain.Task{})
				if err != nil {
						log.Printf("TaskConroller: %s", err)
						BadRequest(w, err)
						return
				}

				task.UserId = user.Id
				task.Status = domain.New
				task, err = c.taskService.Save(task)
				if err != nil {
						log.Printf("TaskConroller: %s", err)
						InternalServerError(w, err)
						return
				}

				var tDto resources.TaskDto
				tDto = tDto.DomainToDto(task)
				Created(w, tDto)
		}
}


func (c TaskConroller) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		taskId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			log.Printf("TaskConroller: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		task, err := requests.Bind(r, requests.TaskRequest{}, domain.Task{})
		if err != nil {
			log.Printf("TaskConroller: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		task.Id = taskId
		updatedTask, err := c.taskService.Update(task)
		if err != nil {
			log.Printf("TaskConroller: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		dto := resources.TaskDto{}.DomainToDto(updatedTask)
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(fmt.Sprintf("%+v\n", dto)))
		if err != nil {
			log.Printf("TaskConroller: %s", err)
			return
		}
	}
}

func (c TaskConroller) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		taskId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			log.Printf("TaskConroller: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = c.taskService.Delete(taskId)
		if err != nil {
			log.Printf("TaskConroller: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func (c TaskConroller) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		taskId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			log.Printf("TaskConroller: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		task, err := c.taskService.GetByID(taskId)
		if err != nil {
			log.Printf("TaskConroller: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		dto := resources.TaskDto{}.DomainToDto(task)
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(fmt.Sprintf("%+v\n", dto)))
		if err != nil {
			log.Printf("TaskConroller: %s", err)
			return
		}
	}
}
