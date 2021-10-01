package article

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/types"
)

func Get(idstr string) (Article, error) {
	var article Article

	id := types.StringToInt(idstr)

	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}

func GetAll() ([]Article, error) {
	var articles []Article

	if err := model.DB.Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

// Create 创建文章通article.ID来判断是否创建成功
func (article *Article) Create() (err error) {
	result := model.DB.Create(&article)

	if err = result.Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}
