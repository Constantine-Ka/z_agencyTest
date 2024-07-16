package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gopkg.in/reform.v1"
	"strings"
	"zeroagencytest/pkg/utils/adapters"
	"zeroagencytest/pkg/utils/logging"
)

func New(db *reform.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) Getlist(ctx context.Context, page, limit int) ([]News, error) {
	logger := logging.GetLogger()

	end := page * limit
	start := end - limit
	query := "SELECT `Id`, `Title`, `Content`, GROUP_CONCAT(c.CategoryId) as category_str FROM News n LEFT JOIN NewsCategories c ON (n.Id=c.NewsId) GROUP BY Id  LIMIT ?, ?;"
	var out []News
	result, err := r.DB.QueryContext(ctx, query, start, limit)
	if err != nil {
		logger.Error("query:", query, err)
		return nil, err
	}
	for result.Next() {
		var news News
		var categorystring sql.NullString
		err = result.Scan(&news.Id, &news.Title, &news.Content, &categorystring)
		if err != nil {
			logger.Error("query:", query, err)
			return nil, err
		}
		if categorystring.String != "" {
			news.Categories = adapters.StrToIntSLice(categorystring.String)
		} else {
			news.Categories = []int{}
		}
		out = append(out, news)
	}

	return out, nil
}
func (r *Repository) UpdateElem(ctx context.Context, id int, data News) error {
	logger := logging.GetLogger()
	if data.Categories != nil {
		query := "DELETE FROM `NewsCategories` WHERE `NewsId`=?"
		_, err := r.DB.ExecContext(ctx, query, id)
		if err != nil {
			logger.Error("query:", query, err)
			return err
		}
		var insertArr []string
		for _, cat := range data.Categories {
			insertArr = append(insertArr, fmt.Sprintf("(%d,%d)", id, cat))
		}
		query = fmt.Sprintf("INSERT INTO `NewsCategories`(`NewsId`, `CategoryId`) VALUES %s", strings.Join(insertArr, ","))
		_, err = r.DB.ExecContext(ctx, query)
		if err != nil {
			logger.Error("query:", query, err)
			return err
		}
	}

	//Вопрос, как понять что в JSON-не 1)не приходит поле впринципе; 2)приходит пустое поле?
	query := "UPDATE `News` SET "
	var setArr []string
	if data.Id != 0 {
		setArr = append(setArr, fmt.Sprintf("`Id`=%d", data.Id))
	}
	if data.Title != "" {
		setArr = append(setArr, fmt.Sprintf("`Title`='%s'", data.Title))
	}
	if data.Content != "" {
		setArr = append(setArr, fmt.Sprintf("`Content`='%s'", data.Content))
	}
	query = query + strings.Join(setArr, ",") + " WHERE `id`=?"
	_, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		logger.Error("query:", query, err)
		return err
	}
	return nil
}
