package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"fmt"
)


func  GetAllTopics(isDesc bool) ([]*Topic,error) {
	o := orm.NewOrm()
	topics := make([]*Topic,0)
	qs := o.QueryTable("topic")
	//按创建时间排序(qs.OrderBy排序，-倒序)
	var err error
	if isDesc {
		_, err = qs.OrderBy("-created").All(&topics)
	}else {
		_,err = qs.All(&topics)
	}
	return topics,err

}

func AddTopic(title, content string) error {
	o := orm.NewOrm()
	topic := &Topic{
		Title: title,
		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
		ReplyTime: time.Now(),
	}

	fmt.Println("数据入库。。。。")
	_,err := o.Insert(topic)
	return err
}

