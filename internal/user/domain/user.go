package domain

import (
	"time"
)

// 贫血模型 状态
type User struct {
	Model           string
	Prompt          string
	Reply           string
	SearchEnabled   bool
	ThinkingEnabled bool
}

/*
	富血模型 状态 + 行为
		状态：
			type User struct {
				ID       int
				Name     string
				password string
			}

			// 把注册逻辑封装进 User 本身
		行为：
			func (u *User) Register() error {
				if u.Name == "" {
					return errors.New("用户名不能为空")
				}
				// 其他校验逻辑
				return nil
			}
*/

type Profile struct {
	//账户
	Account string
	//性别 1:女 2:男
	Gender string
	//邮箱
	Email string
	//密码
	PassWord string
	//账户状态
	State string
	//昵称
	Nickame string
	//出生日期
	Brithday time.Time
	//注册日期
	RegistDate time.Time
	//是否在线
	Online int
}
