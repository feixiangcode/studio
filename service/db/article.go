// Copyright 2018. All rights reserved.
// This file is part of blog project
// Created by duguying on 2018/1/25.

package db

import (
	"duguying/blog/g"
	"duguying/blog/modules/models"
)

func PageArticle(page uint, pageSize uint) (total uint, list []*models.Article, err error) {
	total = 0
	errs := g.Db.Table("articles").Where("status=?", 1).Count(&total).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return 0, nil, errs[0]
	}

	list = []*models.Article{}
	errs = g.Db.Table("articles").Where("status=?", 1).Order("time desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return 0, nil, errs[0]
	}

	return total, list, nil
}

func ArticleToContent(articles []*models.Article) (articleContent []*models.WrapperArticleContent) {
	articleContent = []*models.WrapperArticleContent{}
	for _, article := range articles {
		articleContent = append(articleContent, article.ToArticleContent())
	}
	return articleContent
}

func ArticleToTitle(articles []*models.Article) (articleTitle []*models.WrapperArticleTitle) {
	articleTitle = []*models.WrapperArticleTitle{}
	for _, article := range articles {
		articleTitle = append(articleTitle, article.ToArticleTitle())
	}
	return articleTitle
}

func HotArticleTitle(num uint) (articleTitle []*models.WrapperArticleTitle, err error) {
	list := []*models.Article{}
	errs := g.Db.Table("articles").Order("count desc").Limit(num).Find(&list).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	}
	articleTitle = []*models.WrapperArticleTitle{}
	for _, article := range list {
		articleTitle = append(articleTitle, article.ToArticleTitle())
	}
	return articleTitle, nil
}

type ArchInfo struct {
	Date   string `json:"date"`
	Number uint   `json:"number"`
	Year   uint   `json:"year"`
	Month  uint   `json:"month"`
}

func MonthArch() (archInfos []*ArchInfo, err error) {
	archInfos = []*ArchInfo{}
	errs := g.Db.Table("articles").Select("DATE_FORMAT(time,'%Y年%m月') as date,count(*) as number ,year(time) as year, month(time) as month").Where("status=?", 1).Group("date").Order("year desc, month desc").Find(&archInfos).GetErrors()
	if len(errs) > 0 && errs[0] != nil {
		return nil, errs[0]
	}
	return archInfos, nil
}