package v1

import (
	"github.com/gin-gonic/gin"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/models"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/repos"
	"go.vardan.dev/highload-architect/social-net-01/internal/domain/usecases"
	"go.vardan.dev/highload-architect/social-net-01/pkg/tools"
)

type Api struct {
	userUsecase usecases.UserUsecase
}

func NewApi(usecase usecases.UserUsecase) *Api {
	return &Api{
		userUsecase: usecase,
	}
}

func (a *Api) PostLogin(c *gin.Context) {
	var apiLoginBody PostLoginJSONRequestBody

	if err := c.Bind(&apiLoginBody); err != nil {
		c.Status(400)
		return
	}

	if err := a.userUsecase.Login(*apiLoginBody.Id, *apiLoginBody.Password); err != nil {
		c.String(500, err.Error())
		return
	}

	var apiKey *models.ApiKey
	apiKeyRepo := repos.NewApiKeyRepository(tools.NewUUIDGenerator())
	apiKey, err := apiKeyRepo.GetByUserId(*apiLoginBody.Id)

	if err != nil {
		c.String(500, err.Error())
	}

	createApiKey := func() *models.ApiKey {
		apiKey, err = apiKeyRepo.Create(*apiLoginBody.Id)
		if err != nil {
			c.String(500, err.Error())
		}

		return apiKey
	}

	if apiKey != nil {
		if apiKey.Expired() {
			errDelete := apiKeyRepo.Delete(apiKey.Key)
			if errDelete != nil {
				c.String(500, errDelete.Error())
			}
			apiKey = createApiKey()
		}
	} else {
		apiKey = createApiKey()
	}

	c.JSON(200, gin.H{"token": apiKey.Key})
}

func (a *Api) GetUserGetId(c *gin.Context, id UserId) {
	//domainUserModel, err := a.userUsecase.Get(id)
	//if err != nil {
	//	return err
	//}

	panic("implement me")
}

func (a *Api) PostUserRegister(c *gin.Context) {
	var apiUser PostUserRegisterJSONRequestBody
	err := c.Bind(&apiUser)
	if err != nil {
		c.Status(400)
		return
	}

	domainUserModel := models.NewUser(*apiUser.Age, *apiUser.Biography, *apiUser.City, *apiUser.FirstName,
		*apiUser.SecondName, *apiUser.Password)

	err = a.userUsecase.Register(domainUserModel)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.JSON(201, gin.H{"user_id": domainUserModel.Id})
}

func (a *Api) PutFriendDeleteUserId(c *gin.Context, userId UserId) {
	// TODO implement me
	panic("implement me")
}

func (a *Api) PutFriendSetUserId(c *gin.Context, userId UserId) {
	// TODO implement me
	panic("implement me")
}

func (a *Api) PostPostCreate(c *gin.Context) {
	// TODO implement me
	panic("implement me")
}

func (a *Api) PutPostDeleteId(c *gin.Context, id PostId) {
	// TODO implement me
	panic("implement me")
}

func (a *Api) GetPostFeed(c *gin.Context, params GetPostFeedParams) {
	// TODO implement me
	panic("implement me")
}

func (a *Api) GetPostGetId(c *gin.Context, id PostId) {
	// TODO implement me
	panic("implement me")
}

func (a *Api) PutPostUpdate(c *gin.Context) {
	// TODO implement me
	panic("implement me")
}

func (a *Api) GetUserSearch(c *gin.Context, params GetUserSearchParams) {
	// TODO implement me
	panic("implement me")
}
