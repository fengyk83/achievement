package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Score struct {
	Examid int
	Studentid int
	Courseid int
	Score int
	Name string
}

func NewScore() *Score  {
	return &Score{}
}

func (this *Score)GetId(name string) int  {
	var id int
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.Select("course.id").
		From("course").
		Where("name = ?")
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql,name).QueryRow(&id)
	return id
}

//批量添加成绩
func (this *Score)AddBatchSCore(score []Score) error  {
	var error error
	qb,_:=orm.NewQueryBuilder("mysql")
	qb.InsertInto("escore","examid","studentid","courseid","score").
		Values("?","?","?","?" )
	sql := qb.String()
	o := orm.NewOrm()
	for _, scores := range score {
		_,error =o.Raw(sql,scores.Examid,scores.Studentid,scores.Courseid,scores.Score).Exec()
	}
	return error
}


//获取学生的成绩
func (this *Score)GetAll(examids,studentids int) []Score  {
	var id []Score
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.Select("course.name","escore.score").
		From("escore").
		InnerJoin("course").
		Where("escore.courseid = course.id").
		And("escore.studentid = ?").
		And("escore.examid = ?")
	sql := qb.String()
	o := orm.NewOrm()
	_,error:=o.Raw(sql,studentids,examids).QueryRows(&id)
	fmt.Println(error)
	fmt.Println(id)
	return id
}