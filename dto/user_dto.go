package dto

import "ginessential/model"

// dto，Data Transfer Object，数据传输对象
// 只返回给前端用户名称和手机号
type UserDto struct {
	Name string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name: user.Name,
		Telephone: user.Telephone,
	}
}