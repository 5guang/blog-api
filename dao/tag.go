package dao

import "github.com/jinzhu/gorm"

type Tag struct {
	DataBaseModel
	Model
	Articles []Article `gorm:"many2many:tag_articles;"`
	Name     string    `json:"name"`
}

func GetTags(pageSize, pageNum int, v interface{}) (tags []Tag, err error) {
	if pageSize > 0 && pageNum > 0 {
		err = DB.Order("created_on desc").Where(v).Find(&tags).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = DB.Order("created_on desc").Where(v).Find(&tags).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, nil
}

func GetTagsByArticleId(artId int) ([]Tag, error) {
	var (
		article Article
		tags    []Tag
	)
	err := DB.Where("id = ?", artId).First(&article).Error
	// 当没有找到对应的记录时补鞥呢往下进行 所以不能使用&& err != gorm.ErrRecordNotFound
	if err != nil   {
		return nil, err
	}
	err = DB.Model(&article).Association("Tags").Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, nil
}

func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := DB.Select("id").Where("name=? and deleted_on = 0", name).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}
func ExistTagByID(id int) (bool, error) {
	var tag Tag

	err := DB.Select("id").Where("id=? and deleted_on = 0", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound  {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

func AddTag(name string, state int, createdBy string) error {
	err := DB.Create(&Tag{
		Name: name,
		DataBaseModel: DataBaseModel{
			CreatedBy: createdBy,
			State:     state,
		},
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateTag(id int, v Tag) (err error) {
	err = DB.Model(&Tag{}).Where("id=? and deleted_on = 0", id).Updates(v).Error
	if err != nil  {
		return err
	}
	return nil
}

func DeleteTag(id int) (err error) {
	err = DB.Where("id=? and deleted_on = 0", id).Delete(&Tag{}).Error
	if err != nil  {
		return err
	}
	return nil

}

func GetTagTotal(maps interface{}) (int, error) {
	var count int
	if err := DB.Model(&Tag{}).Where(maps).Count(&count).Error; err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}
	return count, nil
}

func ExistByName() {

}
