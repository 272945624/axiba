package database

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init() {
	orm.RegisterModel(new(UserAdmin))
	orm.RegisterModel(new(ProductCategory))
	orm.RegisterModel(new(ProductAttribute))
	orm.RegisterModel(new(ProductStock))
	orm.RegisterModel(new(ProductOperation))
	orm.RegisterModel(new(ProductStockOperation))
	orm.RegisterModel(new(ProductStockBuffer))
	orm.RegisterModel(new(ProductStockBufferOperation))
	orm.RegisterModel(new(ProductPurchase))
}

type DBOP struct{}

func (ua *DBOP) Create(o orm.Ormer, instance interface{}) (int64, error) {
	return o.Insert(instance)
}

func (ua *DBOP) Read(o orm.Ormer, instance interface{}, arg ...string) error {
	return o.Read(instance, arg...)
}

func (ua *DBOP) Update(o orm.Ormer, instance interface{}, arg ...string) (int64, error) {
	return o.Update(instance, arg...)
}

func (ua *DBOP) Delete(o orm.Ormer, instance interface{}, arg ...string) (int64, error) {
	return o.Delete(instance, arg...)
}

type UserAdmin struct {
	DBOP
	Id             int64     `orm:"pk;auto"`
	UserName       string    `orm:"size(45)"`
	WxId           string    `orm:"size(45)"`
	UpdateTime     time.Time `orm:"auto_now;type(datetime)"`
	Permission     int       `orm:"size(10)"`
	PermissionDesc string    `orm:"size(45)"`
}

func (ua *UserAdmin) TableName() string {
	return "user_admin"
}

type ProductCategory struct {
	DBOP
	Id          int64  `orm:"pk;auto"`
	Name        string `orm:"size(100)"`
	Description string `orm:"size(100)"`
	Deleted     int
	Unit        string `orm:"size(45)"`
}

func (pc *ProductCategory) TableName() string {
	return "product_category"
}

type ProductAttribute struct {
	DBOP
	Id             int64 `orm:"pk;auto"`
	ProductId      int64
	AttributeName  string `orm:"size(45)"`
	AttributeValue string `orm:"size(45)"`
	ValueType      string `orm:"size(45)"`
	ValueLength    int    `orm:"size(11)"`
}

func (pa *ProductAttribute) TableName() string {
	return "product_attribute"
}

type ProductStock struct {
	DBOP
	Id         int64 `orm:"pk;auto"`
	ProductId  int64
	StockCount int64
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

func (ps *ProductStock) TableName() string {
	return "product_stock"
}

type ProductOperation struct {
	DBOP
	Id              int64 `orm:"pk;auto"`
	ProductId       int64
	Operation_type  int       `orm:"size(2)"`
	OperationTime   time.Time `orm:"auto_now;type(datatime)"`
	OperationUserId int64
	Description     string
	Typedesc        string `orm:"size(45)"`
	UserAdminId     int64
}

func (po *ProductOperation) TableName() string {
	return "product_operation"
}

type ProductStockOperation struct {
	DBOP
	Id            int64 `orm:"pk;auto"`
	StockId       int64
	OperationType string
	TypeDesc      string    `orm:"size(45)"`
	UpdateTime    time.Time `orm:"auto_now;type(datatime)"`
	UserAdminId   int64
	UserWxId      string `orm:"size(45)"`
	ExtOpId       int64
}

func (pso *ProductStockOperation) TableName() string {
	return "product_stock_operation"
}

type ProductStockBuffer struct {
	DBOP
	Id               int64 `orm:"pk;auto"`
	ProductId        int64
	StockBufferCount int64
	UpdateTime       time.Time `orm:"auto_now;type(datatime)"'`
}

func (psb *ProductStockBuffer) TableName() string {
	return "product_stock_buffer"
}

type ProductStockBufferOperation struct {
	DBOP
	Id            int64 `orm:"pk;auto"`
	StockBufferId int64
	OperationType int `orm:"size(2)"`
	Description   string
	TypeDesc      string    `orm:"size(45)"`
	UpdateTime    time.Time `orm:"auto_now;type(datatime)"`
	UserAdminId   int64
	ExtOpId       int64
}

func (psbo *ProductStockBufferOperation) TableName() string {
	return "product_stock_buffer_operation"
}

type ProductPurchase struct {
	DBOP
	Id          int64     `orm:"pk;auto"`
	CreateTime  time.Time `orm:"auto_now_add;type(datatime)"`
	ProductId   int64
	UnitPrice   int64
	UnitCount   int64
	Description string
}

func (pp *ProductPurchase) TableName() string {
	return "product_purchase"
}
