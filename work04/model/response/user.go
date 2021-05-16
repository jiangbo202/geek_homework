/**
 * @Author: jiangbo
 * @Description:
 * @File:  user.go
 * @Version: 1.0.0
 * @Date: 2021/01/02 5:30 下午
 */

package response

import (
	"jiang.geek/work04/model/db_model"
)

type UserResponse struct {
	Id        uint   `json:id`
	Name      string `json:name`
	Telephone string `json:telephone`
}

func ToUserResponse(user db_model.User) UserResponse {
	return UserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
