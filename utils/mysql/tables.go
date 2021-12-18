package mysql

// User
/*
create table user
(
    id        varchar(200)              not null comment '用户id（主键）'
        primary key,
    username  varchar(50)               null comment '用户名称',
    real_name varchar(45)               null comment '用户真实名称',
    role_id   tinyint(20)     default 1 not null comment '用户角色，1表示普通用户',
    password  varchar(80)               not null comment '用户密码',
    phone     varchar(11)               null comment '用户电话',
    balance   bigint unsigned default 0 not null comment '用户余额',
    status    tinyint(20)     default 1 not null comment '用户状态，1表示正常，0表示暂停',
    created   char(50)                  null comment '创建时间',
    updated   char(50)                  null comment '更新时间'
)
    charset = utf8;
*/
type User struct {
	Id       string  `json:"id"        gorm:"primaryKey;column:id"`
	Username string  `json:"username"  gorm:"column:username"`
	RealName string  `json:"real_name" gorm:"column:real_name"`
	RoleId   int     `json:"role_id"   gorm:"column:role_id"`
	Password string  `json:"password"  gorm:"column:password"`
	Phone    string  `json:"phone"     gorm:"column:phone"`
	Balance  float64 `json:"balance"   gorm:"column:balance"`
	Status   int     `json:"status"    gorm:"column:status"`
	Created  string  `json:"created"   gorm:"column:created"`
	Updated  string  `json:"updated"   gorm:"column:updated"`
}

func (User) TableName() string {
	return "user"
}

/*
Product
create table product
(
    id                 bigint unsigned auto_increment comment '商品编号'
        primary key,
    category_id        bigint         null comment '类目编号',
    title              varchar(50)    null comment '商品标题',
    description        varchar(80)    null comment '商品描述',
    price              decimal(20, 2) null comment '商品价格',
    amount             int(10)        null comment '商品数量',
    sales              int(10)        null comment '商品销量',
    main_image         varchar(80)    null comment '商品主图',
    delivery           varchar(30)    null comment '商品发货',
    assurance          varchar(30)    null comment '商品保障',
    name               varchar(30)    null comment '商品名称',
    weight             double(20, 0)  null comment '商品重量',
    brand              varchar(10)    null comment '商品品牌',
    origin             varchar(80)    null comment '商品产地',
    shelf_life         int(20)        null comment '商品保质期',
    net_weight         double(20, 0)  null comment '商品净含量',
    use_way            varchar(20)    null comment '使用方式',
    packing_way        varchar(20)    null comment '包装方式',
    storage_conditions varchar(20)    null comment '存储条件',
    detail_image       varchar(80)    null comment '详情图片',
    status             int(10)        null comment '商品状态',
    created            varchar(50)    null comment '创建时间',
    updated            varchar(50)    null comment '更新时间'
)
    charset = utf8;
*/
type Product struct {
	Id                int64   `json:"id"                 gorm:"column:id"`
	CategoryId        int64   `json:"category_id"        gorm:"column:category_id"`
	Title             string  `json:"title"              gorm:"column:title"`
	Description       string  `json:"description"        gorm:"column:description"`
	Price             float64 `json:"price"              gorm:"column:price"`
	Amount            int     `json:"amount"             gorm:"column:amount"`
	Sales             int     `json:"sales"              gorm:"column:sales"`
	MainImage         string  `json:"main_image"         gorm:"column:main_image"`
	Delivery          string  `json:"delivery"           gorm:"column:delivery"`
	Assurance         string  `json:"assurance"          gorm:"column:assurance"`
	Name              string  `json:"name"               gorm:"column:name"`
	Weight            float64 `json:"weight"             gorm:"column:weight"`
	Brand             string  `json:"brand"              gorm:"column:brand"`
	Origin            string  `json:"origin"             gorm:"column:origin"`
	ShelfLife         int     `json:"shelf_life"         gorm:"column:shelf_life"`
	NetWeight         float64 `json:"net_weight"         gorm:"column:net_weight"`
	UseWay            string  `json:"use_way"            gorm:"column:use_way"`
	PackingWay        string  `json:"packing_way"        gorm:"column:packing_way"`
	StorageConditions string  `json:"storage_conditions" gorm:"column:storage_conditions"`
	DetailImage       string  `json:"detail_image"       gorm:"column:detail_image"`
	Status            int     `json:"status"             gorm:"column:status"`
	Created           string  `json:"created"            gorm:"column:created"`
	Updated           string  `json:"updated"            gorm:"column:updated"`
}

func (Product) TableName() string {
	return "product"
}

/*
Category
create table category
(
    id        bigint unsigned auto_increment comment '类目id'
        primary key,
    name      char(50) null comment '类目名称',
    parent_id bigint   null comment '父级类目id',
    level     int(5)   null comment '类目层级',
    sort      int(5)   null comment '类目排序',
    created   char(20) null comment '创建时间',
    updated   char(20) null comment '更新时间'
)
    charset = utf8;
*/
type Category struct {
	Id       int64  `json:"id"        gorm:"column:id"`
	Name     string `json:"name"      gorm:"column:name"`
	ParentId int64  `json:"parent_id" gorm:"column:parent_id"`
	Level    int    `json:"level"     gorm:"column:level"`
	Sort     int    `json:"sort"      gorm:"column:sort"`
	Created  string `json:"created"   gorm:"column:created"`
	Updated  string `json:"updated"   gorm:"column:updated"`
}

func (Category) TableName() string {
	return "category"
}

/*
Order
create table `order`
(
    id           bigint unsigned auto_increment comment '订单id'
        primary key,
    product_item varchar(20)    null comment '商品项',
    total_price  decimal(20, 2) null comment '合计',
    status       char(20)       null comment '订单状态',
    address_id   bigint         null comment '地址id',
    user_id      varchar(200)   null comment '用户id',
    nick_name    char(50)       null comment '用户昵称',
    created      char(50)       null comment '创建时间',
    updated      char(50)       null comment '更新时间'
)
    charset = utf8;
*/
type Order struct {
	Id          int64   `json:"id"           gorm:"column:id"`
	ProductItem string  `json:"product_item" gorm:"column:product_item"`
	TotalPrice  float64 `json:"total_price"  gorm:"column:total_price"`
	Status      string  `json:"status"       gorm:"column:status"`
	AddressId   int64   `json:"address_id"   gorm:"column:address_id"`
	UserId      string  `json:"user_id"      gorm:"column:user_id"`
	NickName    string  `json:"nick_name"    gorm:"column:nick_name"`
	Created     string  `json:"created"      gorm:"column:created"`
	Updated     string  `json:"updated"      gorm:"column:updated"`
}

func (Order) TableName() string {
	return "order"
}
