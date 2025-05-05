package domain

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
