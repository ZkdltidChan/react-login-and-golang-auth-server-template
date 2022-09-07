package utils

import (
	"math"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PageData struct {
	TotalCount int64       `json:"total"`
	Data       interface{} `json:"data_list"`
	Page       int         `json:"page_index"`
	PageCount  int         `json:"page_counts"`
	Size       int         `json:"size"`
	Sort       string      `json:"sort"`
	Searchs    []Search    `json:"searchs"`
}

type Search struct {
	Column string `json:"column"`
	Action string `json:"action"`
	Query  string `json:"query"`
}

func PageOperation(c *gin.Context, db *gorm.DB, data interface{}) PageData {
	var (
		size       = 10
		page       = 1
		sort       = "id desc"
		searchs    []Search
		totalCount int64
	)

	query := c.Request.URL.Query()

	for k, v := range query {
		queryValue := v[len(v)-1]

		switch k {
		case "size":
			size, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		}

		if strings.Contains(k, ".") {
			searchKeys := strings.Split(k, ".")
			search := Search{Column: searchKeys[0], Action: searchKeys[1], Query: queryValue}

			searchs = append(searchs, search)
		}
	}

	db.Count(&totalCount)

	pageCount := int(math.Ceil(float64(totalCount) / float64(size)))
	if page < 1 {
		page = 1
	}
	if page >= pageCount {
		page = pageCount
	}

	db.Offset((page - 1) * size).Limit(size).Find(data)

	return PageData{
		Data:       data,
		Page:       page,
		Size:       size,
		Sort:       sort,
		TotalCount: totalCount,
		PageCount:  pageCount,
		Searchs:    searchs,
	}
}
