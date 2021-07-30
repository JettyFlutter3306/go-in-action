package utils

import "fmt"

/*
注意:  首字母是小写是私有权限不可让别的包访问!
       必须得是开头大写才可以在别的包访问!
*/
func GetConnection() {

	fmt.Println("执行数据库连接操作!")
}
