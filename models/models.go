package models

import(
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

)

const (
	_DB_NAME ="root:123456@/project?charset=utf8"
	_SQLITE3_DRIVER ="mysql"
)
type FengNei struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"` //创建索引
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id int64
	Uid int64
	Title string
	Content string `orm:"size(5000)"` //创建大小
	Attachment string
	Created time.Time `orm:"index"`
	Updated time.Time `orm:"index"`
	Views int64       `orm:"index"`
	Author string
	ReplyTime time.Time`orm:"index"`
	ReplyCount int64
	RepleyLastUserId int64
}

func RegisterDB () {
	orm.RegisterModel(new(FengNei),new(Topic))
	orm.RegisterDataBase("default",_SQLITE3_DRIVER,_DB_NAME)
}

func AddTopic(title,content string) error  {
	o := orm.NewOrm()

	topic := &Topic{
		Title:title,
		Content:content,
		Created: time.Now(),
		Updated:time.Now(),
		ReplyTime:time.Now(),
	}
	_,err :=o.Insert(topic)
	return err
}

func Addfenlei(name string) error {
	t := time.Now()
	o := orm.NewOrm()
	cate := &FengNei{Title:name,Created:t,TopicTime:t}
	qs := o.QueryTable("feng_nei")
	err := qs.Filter("title",name).One(cate)
	if err == nil{
		return err
	}
	_,err = o.Insert(cate)
	if err != nil{
		return err
	}
	return nil
}


func GetAllfenlei() ([]*FengNei,error) {
	o := orm.NewOrm()
	cates := make([]*FengNei,0)
	qs := o.QueryTable("feng_nei")
	_,err := qs.All(&cates)
	return cates,err
}

func GetAlltopic() ([] *Topic,error)  {
	o := orm.NewOrm()
	topics := make([]*Topic,0)
	qs := o.QueryTable("topic")
	_,err := qs.All(&topics)
	return topics,err

}