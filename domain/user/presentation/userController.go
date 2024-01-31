package presentation

import (
	"rest-api/domain/user/presentation/dto"
	"rest-api/global/response"

	"rest-api/domain/user/service"

	"github.com/gofiber/fiber/v2"
)

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

type UserController struct {
	userService *service.UserService
}

func (uc *UserController) Create(c *fiber.Ctx) error {
	body := new(dto.UserRequest)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&response.GeneralResponse{
			Status:  fiber.StatusBadRequest,
			Message: "다시 잘해서 보내셈",
		})
	}

	if err := uc.userService.Create(c.Context(), body); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&response.GeneralResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "서버 처리하다 에러남, " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&response.GeneralResponse{
		Status:  fiber.StatusOK,
		Message: "유저 생성됨",
	})
}

func (uc *UserController) ReadAll(c *fiber.Ctx) error {
	users, err := uc.userService.ReadAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&response.GeneralResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "서버 처리하다 에러남, " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&response.GeneralResponse{
		Status:  fiber.StatusOK,
		Message: "유저 리스트",
		Data:    users,
	})
}

func (uc *UserController) Read(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := uc.userService.Read(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&response.GeneralResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "서버 처리하다 에러남, " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&response.GeneralResponse{
		Status:  fiber.StatusOK,
		Message: "유저",
		Data:    user,
	})
}

func (uc *UserController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	body := new(dto.UserRequest)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&response.GeneralResponse{
			Status:  fiber.StatusBadRequest,
			Message: "다시 잘해서 보내셈",
		})
	}

	if err := uc.userService.Update(c.Context(), id, body); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&response.GeneralResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "서버 처리하다 에러남, " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&response.GeneralResponse{
		Status:  fiber.StatusOK,
		Message: "유저 정보 수정됨",
	})
}

func (uc *UserController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := uc.userService.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&response.GeneralResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "서버 처리하다 에러남, " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&response.GeneralResponse{
		Status:  fiber.StatusOK,
		Message: "유저 삭제됨",
	})
}
