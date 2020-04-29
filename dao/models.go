package dao

type DataBaseModel struct {
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	DeletedBy string `json:"deleted_by"`
	State int `json:"state"`
}
type Model struct {
	ID int `json:"id" gorm:"primary_key"`
	CreatedOn int `json:"created_on"`
	UpdatedOn int `json:"modified_on"`
	DeletedOn int `json:"deleted_on"`
}
