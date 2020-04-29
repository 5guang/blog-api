package util

import (
	"blog/dao"
	"blog/models/response"
)

func DaoTag2ResTag(tags []dao.Tag) []response.ResTag {
	resTags := []response.ResTag{}
	for _, tag := range tags {
		resTags = append(resTags, response.ResTag{
			Creator:     tag.CreatedBy,
			Updater:     tag.UpdatedBy,
			ID:          tag.ID,
			CreatedTime: Timestamp2str(int64(tag.CreatedOn)),
			UpdatedTime: Timestamp2str(int64(tag.UpdatedOn)),
			Name:        tag.Name,
		})
	}
	return resTags
}

func DaoArticle2ResArticle(arts []dao.Article) []response.ResArticle {
	// 如果没有结果返回一个空数组
	resArts := []response.ResArticle{}
	for _, article := range arts {
		resArts = append(resArts, response.ResArticle{
			Title:       article.Title,
			Desc:        article.Desc,
			Content:     article.Content,
			ArticleId:   article.ID,
			Creator:     article.CreatedBy,
			State:       article.State,
			CreatedTime: Timestamp2str(int64(article.CreatedOn)),
		})
	}
	return resArts
}
