package app

import "github.com/Ellan98/ding-water-service/user/app/command/query"

/*
在golang中首字母大写 意味着包外也可以进行访问
如下在对外可访问的 结构体中 若含有小写开头字段，则包外不可访问该结构体内字段
*/

type Application struct {
	Queries Queries
}

type Queries struct {
	PostChatCompletion query.PostChatCompletionHandler
}
