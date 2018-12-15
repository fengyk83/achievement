package models

import (
	"github.com/astaxie/beego/orm"
)

type Exam struct {
	Id int
	Name string
	Time string
	Remark string
	Type int
}

func NewExam() *Exam  {
	return &Exam{}
}

func (this *Exam)GetExam(gradeid , classzid int) []Exam {
	var Exam  []Exam
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.Select("exam.name,exam.time,exam.remark,,exam.type").
		From("exam").
		Where("gradeid = ?").
		And("classzid = ?").
		OrderBy("time").
		Desc()
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql,gradeid,classzid).QueryRows(&Exam)
	return Exam
}

func (this *Exam)GetAllExam() []Exam  {
	var Exam  []Exam
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.Select("exam.*").
		From("exam").
		OrderBy("time").
		Desc()
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&Exam)
	return Exam
}

