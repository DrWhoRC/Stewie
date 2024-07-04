package main

import (
	"fim/core"
	chatmodel "fim/fim_chat/models"
	groupmodel "fim/fim_group/models"
	usermodel "fim/fim_user/models"
	"flag"
)

type Options struct {
	DB bool
}

func main() {
	var opt Options

	//这行代码使用flag包的BoolVar函数来定义一个命令行标志。
	//这个标志的名称是"db"，默认值是false，并且会将解析的值存储到opt.DB中。
	//标志的描述也是"db"。

	//flag.Parse()：解析命令行参数，并将解析的值赋给之前定义的标志变量
	//（在这个例子中是opt.DB）。

	//if opt.DB ：检查opt.DB的值。如果为true（表示用户在命令行中使用了-db标志），
	//则执行大括号内的代码。
	flag.BoolVar(&opt.DB, "db", false, "db")
	flag.Parse()

	if opt.DB {
		db := core.InitMysql()
		db.AutoMigrate(
			&usermodel.FriendModel{},
			&usermodel.UserModel{},
			&usermodel.UserConfigModel{},
			&usermodel.FriendVerifyModel{},

			&groupmodel.GroupModel{},
			&groupmodel.GroupMembersModel{},
			&groupmodel.GroupMsgModel{},
			&groupmodel.GroupVerifyModel{},

			&chatmodel.ChatModel{},
		)
	}
}
