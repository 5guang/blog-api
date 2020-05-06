package dao

type DataBaseModel struct {
	CreatedBy string `json:"created_by" gorm:"type:varchar(100);DEFAULT ''"`
	UpdatedBy string `json:"updated_by" gorm:"type:varchar(100);DEFAULT ''"`
	DeletedBy string `json:"deleted_by" gorm:"type:varchar(100);DEFAULT ''"`
	State int `json:"state" gorm:"type:tinyint(3);unsigned DEFAULT '1'"`
}
type Model struct {
	ID int `json:"id" gorm:"primary_key"`
	CreatedOn int `json:"created_on"`
	UpdatedOn int `json:"modified_on"`
	DeletedOn int `json:"deleted_on"`
}
