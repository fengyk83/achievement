package models

import "github.com/astaxie/beego/orm"
type User struct {
	Id int
	Account string
	Password string
	Name string
	Type int
	Student     *Student   `orm:"rel(one)"` //一对一
}

//登陆判断
func (u *User)LoginJudge(account string,password string)  []User {
	var users []User
	qb,_:=orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("user.account").
		From("user").
		Where("account = ?").
		And("password = ?")
	//返回sql语句
	sql := qb.String()
	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Raw(sql,account,password).QueryRows(&users)
	return users
}

func NewUser() *User  {
	return &User{}
}

func (u *User)GetMenager() []User  {
	var users []User
	qb,_:=orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("user.*").
		From("user").
		Where("type = 3")
	//返回sql语句
	sql := qb.String()
	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&users)
	return users
}

