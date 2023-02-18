package v1

import (
	"github.com/labstack/echo/v4"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/usecases"
)

type Api struct {
	userUsecase usecases.UserUsecase
}

func NewApi(usecase usecases.UserUsecase) *Api {
	return &Api{
		userUsecase: usecase,
	}
}

func (e *Api) PostLogin(ctx echo.Context) error {
	// TODO implement me
	panic("implement me")
}

func (e *Api) GetUserGetId(ctx echo.Context, id UserId) error {
	// TODO implement me
	panic("implement me")
}

func (e *Api) PostUserRegister(ctx echo.Context) error {
	// TODO implement me
	panic("implement me")
}

func (e *Api) PutFriendDeleteUserId(ctx echo.Context, userId UserId) error {
	// TODO implement me
	panic("implement me")
}

func (e *Api) PutFriendSetUserId(ctx echo.Context, userId UserId) error {
	// TODO implement me
	panic("implement me")
}

func (e *Api) PostPostCreate(ctx echo.Context) error {
	// TODO implement me
	panic("implement me")
}

func (e *Api) PutPostDeleteId(ctx echo.Context, id PostId) error {
	// TODO implement me
	panic("implement me")
}

func (e *Api) GetPostFeed(ctx echo.Context, params GetPostFeedParams) error {
	// TODO implement me
	panic("implement me")
}

func (e *Api) GetPostGetId(ctx echo.Context, id PostId) error {
	// TODO implement me
	panic("implement me")
}

func (e *Api) PutPostUpdate(ctx echo.Context) error {
	// TODO implement me
	panic("implement me")
}

func (e *Api) GetUserSearch(ctx echo.Context, params GetUserSearchParams) error {
	// TODO implement me
	panic("implement me")
}
