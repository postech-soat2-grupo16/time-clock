package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time-clock/adapters/user"
	"time-clock/interfaces"
	"time-clock/util"
)

type UserController struct {
	useCase interfaces.UserUserCase
}

func NewUserController(useCase interfaces.UserUserCase, r *chi.Mux) *UserController {
	controller := UserController{useCase: useCase}
	r.Route("/users", func(r chi.Router) {
		r.Post("/", controller.Create())
		r.Get("/{registration}", controller.GetByRegistration())
		r.Post("/{registration}/clock-in", controller.ClockIn())
		r.Post("/{registration}/report", controller.Report())
		r.Get("/health", controller.Health())
	})
	return &controller
}

// @Summary	Create a new user
//
// @Tags		Users
//
// @ID			create-user
// @Produce	json
// @Param		data	body		user.User	true	"User data"
// @Success	200		{object}	user.User
// @Failure	400
// @Router		/users [post]
func (c *UserController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userAdapter := user.User{}
		err := json.NewDecoder(r.Body).Decode(&userAdapter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userAlreadyExist, err := c.useCase.GetByRegistration(userAdapter.Registration)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if userAlreadyExist != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newUser, err := c.useCase.Create(userAdapter.Name, userAdapter.Email, userAdapter.Registration, userAdapter.Password)
		if err != nil {
			if util.IsDomainError(err) {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(err)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
	}
}

// @Summary	Get user by registration
//
// @Tags		Users
//
// @ID			get-user-by-registration
// @Produce	json
// @Param		id	path		string	true	"User Registration"
// @Success	200	{object}	user.User
// @Failure	404
// @Router		/users/{registration} [get]
func (c *UserController) GetByRegistration() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		registration := chi.URLParam(r, "registration")
		userFound, err := c.useCase.GetByRegistration(registration)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if userFound == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(userFound)
	}
}

func (c *UserController) ClockIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (c *UserController) Report() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// @Summary	Health check
//
// @Tags		Users
//
// @ID			health-check
// @Produce	json
// @Success	200	{object}	string
// @Router		/user/health [get]
func (c *UserController) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("OK")
	}
}
