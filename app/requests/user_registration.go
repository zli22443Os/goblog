package requests

import (
	"goblog/app/models/user"

	"github.com/thedevsaddam/govalidator"
)

func ValidateUserRegistrationForm(data user.User) map[string][]string {
	rules := govalidator.MapData{
		"name":             []string{"required", "alpha_num", "between:3,20"},
		"email":            []string{"required", "min:4", "max:30", "email"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
	}

	messages := govalidator.MapData{
		"name":             []string{"required:用户名 为必填项", "alpha_num:格式错误，只允许数字和英文", "between:用户名长度需要在 3~20之间"},
		"email":            []string{"required:Email 为必填项", "min:Email 至少4位", "max:Email 最长30位", "email:格式错误， 请提供有效的邮箱"},
		"password":         []string{"required:密码 为必填项", "min:密码 至少6位"},
		"password_confirm": []string{"required:确认密码 为必填项"},
	}

	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}

	errs := govalidator.New(opts).ValidateStruct()

	if data.Password != data.PasswordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次密码输入不匹配！")
	}

	return errs
}
