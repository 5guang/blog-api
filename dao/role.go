package dao

type Role struct {
	Model
	Name string `gorm:"type:varchar(10);"`
	// association_autoupdate关闭自动更新
	User    []User `gorm:"many2many:user_role;association_autoupdate:false;"`
	Permissions []Permission `gorm:"many2many:permission_role;association_autoupdate:false;"`
}