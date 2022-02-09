package usersapi

import (
	"context"
	"faceit_users/api/httputils"
	"faceit_users/internal/usersapi/models/request"
	"faceit_users/internal/usersapi/models/response"
	mngpkg "faceit_users/pkg/mongo"
	"faceit_users/pkg/mongo/entities"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	usersRepository         mngpkg.UserRepository
	eventsRepository        mngpkg.EventsRepository
	eventsHistoryRepository mngpkg.EventsRepository
}

func NewUserController(user mngpkg.UserRepository, events, eventsHistory mngpkg.EventsRepository) UserController {
	return UserController{
		usersRepository:         user,
		eventsRepository:        events,
		eventsHistoryRepository: eventsHistory,
	}
}

// CreateUser godoc
// @Summary creates a new user by request body.
// @Description create new user.
// @Tags users
// @Accept json
// @Produce json
// @Param user body request.UserCreateRequest true "User Payload"
// @Success 201
// @Failure 409 {string} string "User already exist"
// @Failure 500 {string} string
// @Router /users [post]
func (u *UserController) CreateUser(c echo.Context) error {
	req := new(request.UserCreateRequest)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	existUser, err := u.usersRepository.FindOneByNickname(ctx, req.NickName)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if existUser != nil {
		return c.String(http.StatusConflict, "User already exist")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	err = u.usersRepository.UseSession(ctx, func(sc mongo.SessionContext) error {
		newUser := entities.NewUser(req.FirstName, req.LastName, req.NickName, string(hashedPassword), req.Email, req.Country)
		err := u.usersRepository.InsertOne(sc, newUser)
		if err != nil {
			return err
		}
		event := entities.NewEvent(req.NickName, "Created", "User", "", newUser.ToCreatedEvent())
		err = u.eventsRepository.InsertOne(sc, event)
		if err != nil {
			return err
		}
		err = u.eventsHistoryRepository.InsertOne(sc, event)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}

// DeleteUser godoc
// @Summary deletes a user by param.
// @Description delete user.
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204
// @Success 400 {string} string "User Id required"
// @Success 404 {string} string "User not found"
// @Failure 500
// @Router /users/{id} [delete]
func (u *UserController) DeleteUser(c echo.Context) error {
	id := httputils.ParamStringOrDefaultValue(c, "id", "")
	if id == "" {
		return c.String(http.StatusBadRequest, "User Id required")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	user, err := u.usersRepository.FindOne(ctx, id)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if user == nil {
		return c.String(http.StatusNotFound, "User not found")
	}

	err = u.usersRepository.UseSession(ctx, func(sc mongo.SessionContext) error {
		err := u.usersRepository.DeleteOne(sc, id)
		if err != nil {
			return err
		}
		event := entities.NewEvent(user.NickName, "Deleted", "User", "", user.ToDeletedEvent())
		err = u.eventsRepository.InsertOne(sc, event)
		if err != nil {
			return err
		}
		err = u.eventsHistoryRepository.InsertOne(sc, event)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// GetUsers godoc
// @Summary get users by country.
// @Description fetch users with filter.
// @Tags users
// @Accept json
// @Produce json
// @Param country query string true "User Country ex: UK"
// @Param pageIndex query int false "Default is 1"
// @Param pageSize query int false "Default is 20"
// @Success 200 {object} map[string]interface{}
// @Failure 500
// @Router /users [get]
func (u *UserController) GetUsers(c echo.Context) error {

	param := httputils.ParamStringOrDefaultValue(c, "country", "")
	pageIndex := httputils.QueryIntOrDefaultValue(c, "pageIndex", 0) - 1
	pageSize := httputils.QueryIntOrDefaultValue(c, "pageSize", 20)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	users, err := u.usersRepository.Get(ctx, param, pageIndex*pageSize, pageSize)

	resp := []response.UsersResponse{}
	for _, value := range users {
		resp = append(resp, response.NewUserResponse(value.Id, value.FirstName, value.LastName, value.NickName, value.Email, value.Country))
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// UpdateUser godoc
// @Summary update user by request body.
// @Description updates a user.
// @Tags users
// @Accept json
// @Produce json
// @Param user body request.UserUpdateRequest true "User Payload"
// @Param id path string true "User ID"
// @Success 200
// @Success 400 {string} string "User Id required"
// @Success 404 {string} string "User not found"
// @Failure 500
// @Router /users/{id} [put]
func (u *UserController) UpdateUser(c echo.Context) error {
	id := httputils.ParamStringOrDefaultValue(c, "id", "")
	if id == "" {
		return c.String(http.StatusBadRequest, "User Id required")
	}
	req := new(request.UserUpdateRequest)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	existUser, err := u.usersRepository.FindOne(ctx, id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if existUser == nil {
		return c.String(http.StatusNotFound, "User not found")
	}
	existUser.UpdateFields(req.FirstName, req.LastName, req.Country)
	err = u.usersRepository.UseSession(ctx, func(sc mongo.SessionContext) error {
		err := u.usersRepository.UpdateOne(sc, id, *existUser)
		if err != nil {
			return err
		}
		event := entities.NewEvent(existUser.NickName, "Updated", "User", "", existUser.ToUpdatedEvent())
		err = u.eventsRepository.InsertOne(sc, event)
		if err != nil {
			return err
		}
		err = u.eventsHistoryRepository.InsertOne(sc, event)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
