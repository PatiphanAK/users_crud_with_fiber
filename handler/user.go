package handler

import (
	"tatar/book/.gen/mydb/model"
	api "tatar/book/oapi"
	"tatar/book/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUsers(c *fiber.Ctx, params api.GetUsersParams) error {
	limit := 10
	offset := 0

	if params.Limit != nil {
		limit = *params.Limit
	}

	if params.Offset != nil {
		offset = *params.Offset
	}

	users, err := h.userService.GetUsers(limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch users")
	}

	apiUsers := make([]api.User, len(users))
	for i, user := range users {
		id := int(user.ID)
		apiUsers[i] = api.User{
			Id:    &id,
			Name:  &user.Name,
			Email: &user.Email,
			Uuid: &user.UUID,
		}
	}

	return c.JSON(
		fiber.Map{
			"total":len(users),
			"data":apiUsers,
		},
	)
}

func (h *UserHandler) GetUserById(c *fiber.Ctx, uuid string) error {
	user, err := h.userService.GetUserByUUID(uuid)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	id := int(user.ID)
	apiUser := api.User{
		Id:    &id,
		Name:  &user.Name,
		Email: &user.Email,
	}

	return c.JSON(apiUser)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx, params api.CreateUserParams) error {
	c.BodyParser(&params)
	if params.Name == nil || params.Email == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Name and email are required")
	}

	user := &model.User{
		Name:  *params.Name,
		Email: *params.Email,
	}
	createdUser, err := h.userService.CreateUser(user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create user")
	}

	apiUser := api.User{
		Name:  &createdUser.Name,
		Email: &createdUser.Email,
	}

	return c.Status(fiber.StatusCreated).JSON(apiUser)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx, uuid string, params api.UpdateUserParams) error {
	c.BodyParser(&params)
	existingUser, err := h.userService.GetUserByUUID(uuid)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	if params.Name == nil || params.Email == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Name and email are required")
	}

	existingUser.Name = *params.Name
	existingUser.Email = *params.Email

	updatedUser, err := h.userService.UpdateUser(existingUser)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update user")
	}

	apiUser := api.User{
		Name:  &updatedUser.Name,
		Email: &updatedUser.Email,
	}

	return c.JSON(apiUser)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx, uuid string) error {
	err := h.userService.DeleteUser(uuid)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete user")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
