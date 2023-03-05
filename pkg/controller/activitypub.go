package controller

import (
	"strconv"

	"github.com/approvers/qip/pkg/domain"

	"github.com/approvers/qip/pkg/application/user"

	"github.com/approvers/qip/pkg/utils/id"

	"github.com/approvers/qip/pkg/activitypub"
	"github.com/approvers/qip/pkg/activitypub/types"
	"github.com/approvers/qip/pkg/application"
	"github.com/approvers/qip/pkg/repository"
)

// ActivityPubController ActivityPubの通信に応答する部分
type ActivityPubController struct {
	repo    repository.UserRepository
	usecase user.UserUseCase
}

func NewActivityPubController(r repository.UserRepository) *ActivityPubController {
	return &ActivityPubController{
		repo:    r,
		usecase: *application.NewUserUseCase(r),
	}
}

// GetUser 自インスタンスのユーザーを取得
func (c ActivityPubController) GetUser(uid string) *types.PersonResponseJSONLD {
	// snowflakeかUsernameか判別
	_, err := strconv.Atoi(uid)
	var user *domain.User
	if err != nil {
		// UserNameのとき
		user, err = c.usecase.FindLocalByUserName(uid)
		if err != nil {
			return nil
		}
	} else {
		// SnowflakeIDのとき
		user, err = c.usecase.FindByID(id.SnowFlakeID(uid))
		if err != nil {
			return nil
		}
	}

	if user == nil {
		return nil
	}

	n := ""
	if user.IconImageURL == nil {
		user.IconImageURL = &n
	}
	if user.HeaderImageURL == nil {
		user.HeaderImageURL = &n
	}

	arg := types.PersonResponseArgs{
		ID:             string(user.id),
		UserName:       user.name,
		UserScreenName: user.ScreenName,
		Summary:        user.Summary,
		Icon: struct {
			Url       string
			Sensitive bool
			Name      interface{}
		}{
			Url: *user.IconImageURL,
		},
		Image: struct {
			Url       string
			Sensitive bool
			Name      interface{}
		}{
			Url: *user.HeaderImageURL,
		},
		Tag:                       nil,
		ManuallyApprovesFollowers: false,
		PublicKey:                 user.publicKey,
	}

	res := activitypub.Person(arg)
	return &res
}
