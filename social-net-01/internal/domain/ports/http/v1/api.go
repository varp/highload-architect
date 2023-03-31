package v1

import (
	"github.com/labstack/echo/v4"
	domainUser "go.vardan.dev/highload-architect/social-net-01/internal/domain/models/user"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/usecases/user"
)

type Api struct {
	userUsecase user.Usecase
}

func NewApi(usecase user.Usecase) *Api {
	return &Api{
		userUsecase: usecase,
	}
}

func (a *Api) PostLogin(ctx echo.Context) error {
	var apiLoginBody PostLoginJSONRequestBody

	if err := ctx.Bind(&apiLoginBody); err != nil {
		return err
	}

	if err := a.userUsecase.Login(*apiLoginBody.Id, *apiLoginBody.Password); err != nil {
		return err
	}

	return nil
}

func (a *Api) GetUserGetId(ctx echo.Context, id UserId) error {
	domainUserModel, err := a.userUsecase.Get(id)
	if err != nil {
		return err
	}

	return nil
}

func (a *Api) PostUserRegister(ctx echo.Context) error {
	var apiUser PostUserRegisterJSONRequestBody
	err := ctx.Bind(&apiUser)
	if err != nil {
		return err
	}

	domainUserModel := domainUser.New(*apiUser.Age, *apiUser.Biography, *apiUser.City, *apiUser.FirstName,
		*apiUser.SecondName, *apiUser.Password)

	err = a.userUsecase.Register(domainUserModel)
	if err != nil {
		return err
	}

	return nil
}

func (a *Api) PutFriendDeleteUserId(ctx echo.Context, userId UserId) error {
	// TODO implement me
	panic("implement me")
}

func (a *Api) PutFriendSetUserId(ctx echo.Context, userId UserId) error {
	// TODO implement me
	panic("implement me")
}

func (a *Api) PostPostCreate(ctx echo.Context) error {
	// TODO implement me
	panic("implement me")
}

func (a *Api) PutPostDeleteId(ctx echo.Context, id PostId) error {
	// TODO implement me
	panic("implement me")
}

func (a *Api) GetPostFeed(ctx echo.Context, params GetPostFeedParams) error {
	// TODO implement me
	panic("implement me")
}

func (a *Api) GetPostGetId(ctx echo.Context, id PostId) error {
	// TODO implement me
	panic("implement me")
}

func (a *Api) PutPostUpdate(ctx echo.Context) error {
	// TODO implement me
	panic("implement me")
}

func (a *Api) GetUserSearch(ctx echo.Context, params GetUserSearchParams) error {
	// TODO implement me
	panic("implement me")
}
