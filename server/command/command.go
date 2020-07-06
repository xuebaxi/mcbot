package command

type command func([]string) string

//Command 是普通用户命令列表
var Command map[string]command = map[string]command{}

//Admin 是管理员命令列表
var Admin map[string]command = map[string]command{}
