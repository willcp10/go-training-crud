package controller

import (
	"encoding/json"
	"fmt"
	"go-training-crud/internal/adapters/controller/dto"
	"go-training-crud/internal/application/cmd"
	"go-training-crud/internal/application/qry"
	"go-training-crud/internal/domain"
	"io"
	"net/http"
	"strconv"
)

const (
	headerKeyContentType = "Content-Type"
	headerValueApplicationJSON = "application/json"
)

type UserController interface {
	ping(w http.ResponseWriter, r *http.Request)
	usersHandler(w http.ResponseWriter, r *http.Request)
	userHandler(w http.ResponseWriter, r *http.Request)
}

var _ UserController = new(UserControllerImpl)

type UserControllerImpl struct {
	userCommandService cmd.UserCommandService
	userQueryService   qry.UserQueryService
}

func MapRoutes(
	userCommandService cmd.UserCommandService,
	userQueryService qry.UserQueryService,
) error {
	userController := UserControllerImpl{
		userCommandService: userCommandService,
		userQueryService: userQueryService,
	}

	http.HandleFunc("/ping", userController.ping)
	http.HandleFunc("/users", userController.usersHandler)
	http.HandleFunc("/users/{id}", userController.userHandler)

	return nil
}

func (uc *UserControllerImpl) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func (uc *UserControllerImpl) usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		uc.getAllUsers(w, r)
	case http.MethodPost:
		uc.createUser(w, r)
	case http.MethodPut:
		uc.updateUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (uc *UserControllerImpl) getAllUsers(w http.ResponseWriter, _ *http.Request) {
	userDTOSlice := []dto.UserDTO{}
	for _, u := range uc.userQueryService.FindAllUsers() {
		userDTOSlice = append(userDTOSlice, dto.FromDomain(u))
	}

	response, err := json.Marshal(userDTOSlice)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "error writting response: %s", err.Error()))
		return
	}

	w.Header().Set(headerKeyContentType, headerValueApplicationJSON)
	w.Write(response)
}

func (uc *UserControllerImpl) createUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(fmt.Appendf(nil, "error reading body: %s", err.Error()))
		return
	}

	var userDTO dto.UserDTO
	err = json.Unmarshal(body, &userDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(fmt.Appendf(nil, "error decoding json: %s", err.Error()))
		return
	}
	
	cmd := cmd.NewCreateUserCommand(userDTO.ToDomain())
	err = uc.userCommandService.CreateUser(cmd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "error executing create command: %s", err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (uc *UserControllerImpl) updateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(fmt.Appendf(nil, "error reading body: %s", err.Error()))
		return
	}

	var userDTO dto.UserDTO
	err = json.Unmarshal(body, &userDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(fmt.Appendf(nil, "error decoding json: %s", err.Error()))
		return
	}

	cmd := cmd.NewUpdateUserCommand(userDTO.ToDomain())
	err = uc.userCommandService.UpdateUser(cmd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "error executing update command: %s", err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uc *UserControllerImpl) userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		uc.getUserByID(w, r)
	case http.MethodDelete:
		uc.deleteUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (uc *UserControllerImpl) getUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(fmt.Appendf(nil, "error getting id from request: %s", err.Error()))
		return
	}
	domainID := domain.NewID(id)
	user, err := uc.userQueryService.FindUserByID(domainID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(fmt.Appendf(nil, "error finding by id: %s", err.Error()))
		return
	}

	userDTO := dto.FromDomain(user)
	response, err := json.Marshal(userDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "error writting response: %s", err.Error()))
		return
	}

	w.Header().Set(headerKeyContentType, headerValueApplicationJSON)
	w.Write(response)
}

func (uc *UserControllerImpl) deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(fmt.Appendf(nil, "error getting id from request: %s", err.Error()))
		return
	}
	domainID := domain.NewID(id)

	cmd := cmd.NewDeleteUserCommand(domainID)
	err = uc.userCommandService.DeleteUser(cmd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "error executing delete command: %s", err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
