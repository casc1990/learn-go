package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_"github.com/Unknwon/com"
	_"os"
	_"path"
	"fmt"
	"strconv"
)

const (
	_DB_NAME = "beeblog"
	_DB_TYPE = "mysql"
	_Sql_NAME = "data/beeblog.db"
	_Sdb_Type = "sqlite3"
)

//分类表
type Category struct {
	Id 				int64 //int,int64的id字段会主动被识别为主键
	Title 			string  //string类型，不指定长度，默认255
	CrateTime 		time.Time   `orm:"index"`  //使用tag，告诉创建时建索引
	Views 			int64			`orm:"index"`  //同上
	TopicTime 		time.Time `orm:"index"` //最后文字发表时间
	TopicCount 		int64           //分类中有多少篇文章
	TopicLastUserId int64   //最后一个浏览的用户id

}

//文章表
type Topic struct {
	Id 			int64    //文章id
	Uid 		int64    //用户id
	Title 		string   //标题
	Content 	string  `orm:"size(50000)"` //指定字符串长度（文章内容）
	Attachment 	string   //附件
	Created 	time.Time `orm:"index"`
	Updated 	time.Time `orm:"index"`
	Views 		int64     `orm:"index"` //浏览次数
	Autoher 	string //作者
	ReplyTime  time.Time  `orm:"index"`  //回复时间
	ReplyCount  int64   //回复次数
	ReplyLastUserId int64  //最后回复用户ID
}

//注册数据库
func RegisterDB() {
	/*
	//如何是sqlit3数据库，就要判断数据库文件目前是否存在了
	if !com.IsExist(_DB_NAME) {  //目录不存在，创建目前
		os.MkdirAll(path.Dir(_Sql_NAME),os.ModePerm)
		os.Create(_Sql_NAME)
	}
	*/

	//注册模型，接收指针参数
	orm.RegisterModel(new(Category),new(Topic))
	//注册驱动
	orm.RegisterDriver(_DB_TYPE,orm.DRMySQL)
	//注册连接
	orm.RegisterDataBase("default",_DB_TYPE,"root:123456@tcp(127.0.0.1:3306)/beeblog?charset=utf8",30,30)
}

//添加分类
func AddCategory(name string) error {
	//初始化or实例
	o := orm.NewOrm()
	//实例化一个分类结构体
	cate := &Category{Title: name,CrateTime: time.Now(),TopicTime:time.Now()}

	fmt.Println("匹配数据。。。。。。。。。。。。。")
	//查询category表，匹配title等于name变量的
	qs := o.QueryTable("category")
	err := qs.Filter("title",name).One(cate)
	//如果匹配到,说明分类已经存在
	if err == nil {
		fmt.Println("此分类已经存在")
		return err
	}
	//不存在，插入
	fmt.Println("数据入库......")
	_,err = o.Insert(cate)
	if err != nil {
		fmt.Println("插入失败")
		return  err
	}
	return nil
}

func DelCategory(id string) error {
	cid,err := strconv.ParseInt(id,10,64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_,err = o.Delete(cate)
	return err
}

func GetAllCategory() ([]*Category,error) {
	o := orm.NewOrm()
	cates := make([]*Category,0)
	qs := o.QueryTable("category")
	//qs.All查询所有
	_,err := qs.All(&cates)
	return cates,err
}
