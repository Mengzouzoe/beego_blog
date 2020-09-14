package models

import (
	"os"
	"path"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/unknwon/com"
)

const (
	_DB_NAME      = "data/myapp.db"
	_MYSQL_DRIVER = "mysql"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterDriver(_MYSQL_DRIVER, orm.DRMySQL)

	orm.RegisterModel(new(Category), new(Topic))

	orm.RegisterDataBase("default", "mysql", "root:password@tcp(127.0.0.1:3306)/beego_blog1?charset=utf8", 30)
}
