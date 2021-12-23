package monitor

const (
	KGuestRoleId = 1
	KAdminRoleId = 2
)

// User
/*
create table user
(
    id        varchar(200)              not null comment '用户id（主键）'
        primary key,
    username  varchar(50)               null comment '用户名称',
    role_id   tinyint(20)     default 1 not null comment '用户角色，1表示普通用户 2表示管理员',
    password  varchar(80)               not null comment '用户密码',
    phone     varchar(11)               null comment '用户电话',
    status    tinyint(20)     default 1 not null comment '用户状态，1表示正常，0表示暂停',
    created   char(50)                  null comment '创建时间',
)
    charset = utf8;
*/
type User struct {
	Id       string `json:"id"       gorm:"column:id;primary_key"`
	Username string `json:"username" gorm:"column:username"`
	RoleId   int    `json:"role_id"  gorm:"column:role_id"`
	Password string `json:"password" gorm:"column:password"`
	Phone    string `json:"phone"    gorm:"column:phone"`
	Status   int    `json:"status"   gorm:"column:status"`
	Created  string `json:"created"  gorm:"column:created"`
}

func (User) TableName() string {
	return "user"
}
