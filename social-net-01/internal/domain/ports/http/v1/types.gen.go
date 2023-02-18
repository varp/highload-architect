// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package v1

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Post Пост пользователя
type Post struct {
	// AuthorUserId Идентификатор пользователя
	AuthorUserId *UserId `json:"author_user_id,omitempty"`

	// Id Идентификатор поста
	Id *PostId `json:"id,omitempty"`

	// Text Текст поста
	Text *PostText `json:"text,omitempty"`
}

// PostId Идентификатор поста
type PostId = string

// PostText Текст поста
type PostText = string

// User defines model for User.
type User struct {
	// Age Возраст
	Age *int `json:"age,omitempty"`

	// Biography Интересы
	Biography *string `json:"biography,omitempty"`

	// City Город
	City *string `json:"city,omitempty"`

	// FirstName Имя
	FirstName *string `json:"first_name,omitempty"`

	// Id Идентификатор пользователя
	Id *UserId `json:"id,omitempty"`

	// SecondName Фамилия
	SecondName *string `json:"second_name,omitempty"`
}

// UserId Идентификатор пользователя
type UserId = string

// N5xx defines model for 5xx.
type N5xx struct {
	// Code Код ошибки. Предназначен для классификации проблем и более быстрого решения проблем.
	Code *int `json:"code,omitempty"`

	// Message Описание ошибки
	Message string `json:"message"`

	// RequestId Идентификатор запроса. Предназначен для более быстрого поиска проблем.
	RequestId *string `json:"request_id,omitempty"`
}

// PostLoginJSONBody defines parameters for PostLogin.
type PostLoginJSONBody struct {
	Id       *string `json:"id,omitempty"`
	Password *string `json:"password,omitempty"`
}

// PostPostCreateJSONBody defines parameters for PostPostCreate.
type PostPostCreateJSONBody struct {
	// Text Текст поста
	Text PostText `json:"text"`
}

// GetPostFeedParams defines parameters for GetPostFeed.
type GetPostFeedParams struct {
	Offset *float32 `form:"offset,omitempty" json:"offset,omitempty"`
	Limit  *float32 `form:"limit,omitempty" json:"limit,omitempty"`
}

// PutPostUpdateJSONBody defines parameters for PutPostUpdate.
type PutPostUpdateJSONBody struct {
	// Id Идентификатор поста
	Id PostId `json:"id"`

	// Text Текст поста
	Text PostText `json:"text"`
}

// PostUserRegisterJSONBody defines parameters for PostUserRegister.
type PostUserRegisterJSONBody struct {
	Age        *int    `json:"age,omitempty"`
	Biography  *string `json:"biography,omitempty"`
	City       *string `json:"city,omitempty"`
	FirstName  *string `json:"first_name,omitempty"`
	Password   *string `json:"password,omitempty"`
	SecondName *string `json:"second_name,omitempty"`
}

// GetUserSearchParams defines parameters for GetUserSearch.
type GetUserSearchParams struct {
	// FirstName Условие поиска по имени
	FirstName string `form:"first_name" json:"first_name"`

	// LastName Условие поиска по фамилии
	LastName string `form:"last_name" json:"last_name"`
}

// PostLoginJSONRequestBody defines body for PostLogin for application/json ContentType.
type PostLoginJSONRequestBody PostLoginJSONBody

// PostPostCreateJSONRequestBody defines body for PostPostCreate for application/json ContentType.
type PostPostCreateJSONRequestBody PostPostCreateJSONBody

// PutPostUpdateJSONRequestBody defines body for PutPostUpdate for application/json ContentType.
type PutPostUpdateJSONRequestBody PutPostUpdateJSONBody

// PostUserRegisterJSONRequestBody defines body for PostUserRegister for application/json ContentType.
type PostUserRegisterJSONRequestBody PostUserRegisterJSONBody