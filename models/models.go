package models

import (
	"os"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"

	_ "github.com/mattn/go-sqlite3"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id           int64
	Title        string
	Create       time.Time `orm:"index"`
	TopicCount   int64
	CreateFormat string
}

type Topic struct {
	Id                 int64
	Title              string
	Category           string
	Content            string `orm:"size(5000)"`
	Attachment         string
	Created            time.Time `orm:"index"`
	Updated            time.Time `orm:"index"`
	Views              int64     `orm:"index"`
	Author             string
	CreateFormat string
	UpdateFormat string
}

func RegisterDB() {
	if _, err := os.Stat(_DB_NAME); err != nil {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{
		Title:        name,
		Create:       time.Now(),
		TopicCount:   0,
		CreateFormat: time.Now().Format("2006年1月2日15点04分05秒"),
	}
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

func DelCategories(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func AddTopic(title, cate, content string) (err error) {
	o := orm.NewOrm()
	topic := &Topic{
		Title:              title,
		Category:           cate,
		Content:            content,
		Created:            time.Now(),
		Updated:            time.Now(),
		Author:             beego.AppConfig.String("author"),
		CreateFormat: time.Now().Format("2006年1月2日15点04分05秒"),
	}
	qs := o.QueryTable("topic")
	err = qs.Filter("title", title).One(topic)
	if err == nil {
		return
	}
	_, err = o.Insert(topic)

	category := &Category{
		Title: cate,
	}

	err = o.QueryTable("category").Filter("title", cate).One(category)
	if err != nil {
		return
	}
	category.TopicCount++
	_, err = o.Update(category)

	return
}

func GetAllTopics(isDesc bool, cate string) (topics []*Topic, err error) {
	o := orm.NewOrm()
	topics = make([]*Topic, 0)
	qs := o.QueryTable("topic")
	if isDesc {
		if len(cate) > 0 {
			qs = qs.Filter("category", cate)
		}
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

func GetTopic(tid string) (topic *Topic, err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return
	}
	o := orm.NewOrm()
	topic = new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return
	}
	topic.Views++
	_, err = o.Update(topic)
	return
}

func ModifyTopic(tid, title, category, content string) (err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return
	}
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if err = o.Read(topic); err == nil {
		topic.Title = title
		topic.Category = category
		topic.Content = content
		topic.Updated = time.Now()
		topic.UpdateFormat = time.Now().Format("2006年1月2日15点04分05秒")
		o.Update(topic)
		return
	}
	return
}

func DeleteTopic(tid string) (err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return
	}
	o := orm.NewOrm()

	topic := &Topic{Id: tidNum}
	_, err = o.Delete(topic)
	return
}
