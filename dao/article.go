package dao

import "github.com/jinzhu/gorm"

type Article struct {
	Model
	Title   string `gorm:"type:varchar(100);DEFAULT ''"`
	Desc    string
	Content string `gorm:"type:text"`
	Tags    []Tag `gorm:"many2many:tag_articles;association_autoupdate:false;"`
	DataBaseModel
}

// 获取文章
func GetArticles(pageSize, pageNum int, v interface{}) (arts []Article, err error) {
	if pageSize > 0 && pageNum > 0 {
		err = DB.Order("created_on desc").Where("deleted_on = 0").Where(v).Find(&arts).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = DB.Order("created_on desc").Where("deleted_on = 0").Where(v).Find(&arts).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound  {
		return
	}
	return
}

// 通过tagd获取所有文章
func GetArticlesByTagId(tagId int) ([]Article, error) {
	var (
		tag      Tag
		articles []Article
	)
	err := DB.Where("id = ?", tagId).First(&tag).Error
	if err != nil {
		return nil, err
	}
	err = DB.Model(&tag).Association("Articles").Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound  {
		return nil, err
	}
	return articles, nil
}

// 更新文章
func UpdateArticle(artId int, art Article) error {
	err := DB.Model(&Article{}).Where("id=? and deleted_on = 0", artId).Updates(&art).Error
	if err != nil  {
		return err
	}
	return nil
}

// 添加一片文章 -.-这里踩了gorm多对多的很多坑....
func AddArticle(art Article) (error, int) {
	err := DB.Create(&art).Error
	if err != nil  {
		return err, 0
	}
	return nil, art.ID
}

// 删除文章
func DelArticle(id int) error {
	// 创建一个事务
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			// 回滚操作
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	// 先删除文章
	err := tx.Where("id=? and deleted_on = 0", id).Delete(&Article{}).Error
	if err != nil  {
		tx.Rollback()
		return err
	}

	// 清空关联关系表中当前文章的关系
	err = tx.Model(&Article{Model: Model{ID: id}}).Association("Tags").Clear().Error
	if err != nil  {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// 获取单篇文章
func GetArticle(artId int) (arts []Article, err error) {
	err = DB.Order("created_on desc").Where("id=? and deleted_on = 0", artId).First(&arts).Error
	if err != nil && err != gorm.ErrRecordNotFound  {
		return nil, err
	}
	return arts, nil
}

func ExistArticleByName() {

}
