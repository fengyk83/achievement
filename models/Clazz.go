package models

import "github.com/astaxie/beego/orm"

type Clazz struct {
	id int
	name string
}

func NewClazz() *Clazz  {
	return &Clazz{}
}

func (this *Clazz)GetClazz() []Clazz  {
	var clazz  []Clazz
	qb,_:=orm.NewQueryBuilder("mysql")
	qb.Select("clazz.*").
		From("clazz")
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&clazz)
	return clazz
}