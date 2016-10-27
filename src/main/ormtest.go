package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"main/orm"
	"main/pojo"
)

var mysqlUrl = "test:test@tcp(127.0.0.1:3306)/test?charset=utf8"

/*
插入测试用例
*/
func addUsers(userInfoDao *orm.UserInfoDao) error {
	lin := pojo.UserInfo{100, "linlin", 23, "15868884065", "linsony0@163.com"}
	zhi := pojo.UserInfo{101, "zhizhi", 26, "15867884066", "linsony0@163.com"}
	dede := pojo.UserInfo{103, "dede", 27, "15867984067", "linsony0@163.com"}

	_, lerr := userInfoDao.AddUserInfo(&lin)
	_, zerr := userInfoDao.AddUserInfo(&zhi)
	_, derr := userInfoDao.AddUserInfo(&dede)

	if lerr != nil || zerr != nil || derr != nil {
		fmt.Println(lerr)
		return lerr
	}
	return nil
}

/*
查询测试用例  单个和列表
*/
func selectTest(userInfoDao *orm.UserInfoDao) error {
	usr, ierr := userInfoDao.GetUserInfoById(100)
	fmt.Println(usr, ierr)

	users, userr := userInfoDao.GetUserInfoListByEmail("linsony0@163.com")
	fmt.Println(users, userr)

	if ierr != nil {
		return ierr
	} else {
		return userr
	}
}

/*
更改测试用例
*/
func updateTest(userInfoDao *orm.UserInfoDao) error {
	lin := pojo.UserInfo{100, "linlin", 23, "15868884065", "linsony0@163.com"}
	lin.Name = "hahah"
	lin.Age = 25
	lin.Email = "lin@163.com"
	c, err := userInfoDao.UpdateUserInfoByIdAndMobile(&lin)
	fmt.Println("updated ", c, err)
	return err
}

/*
删除测试用例
*/
func deleteById(userInfoDao *orm.UserInfoDao) error {
	c, err := userInfoDao.DeleteById(100)
	fmt.Println("deleted ", c, err)
	return err
}

func main() {
	db, err := sql.Open("mysql", mysqlUrl)
	if err != nil {
		fmt.Println("connect local test mysql error", err)
		return
	}

	defer db.Close()

	userInfoDao := orm.UserInfoDao{*db}

	// updateTest(&userInfoDao)
	deleteById(&userInfoDao)

}
