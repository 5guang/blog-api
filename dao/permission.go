package dao

type Permission struct {
	Model
	Name string `gorm:"type:varchar(10);"`
	Roles []Role `gorm:"many2many:permission_role;association_autoupdate:false;"`
}
