package services

import (
	"context"
	"encoding/json"
	"fmt"
	"sau-na/conf"
	"sau-na/global"
	"sau-na/model"
)

func Register(ctx context.Context, code string) (*model.User, error) {
	user := new(model.User)

	// codeからtokenを取得しにいく
	token, err := global.GoogleAuthConf.Exchange(ctx, code)
	if err != nil {
		return user, err
	}

	// tokenの有効性を確認
	client := global.GoogleAuthConf.Client(context.Background(), token)
	res, err := client.Get(conf.GoogleOAuthScopeUserInfo)
	if err != nil {
		return user, err
	}

	userInfo := make(map[string]interface{})
	err = json.NewDecoder(res.Body).Decode(&userInfo)

	fmt.Print(userInfo)

	return user, err
}
