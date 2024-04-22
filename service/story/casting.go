package story

import (
	"example/request"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jmoiron/sqlx"
)

type CommonRequests struct {
	request.CommonRequests
}

func CastToStory(item *sqlx.Rows) Story {
	var p Story
	scanErr := item.StructScan(&p)
	if scanErr != nil {
		return Story{}
	}
	return p
}

func (c CommonRequests) Select(options ...exp.Expression) (any, error) {

	var result []Story

	items, err := c.CommonRequests.Select(options...)

	if err != nil {
		return result, err
	}

	for items.Next() {
		result = append(result, CastToStory(items))
	}

	return result, nil
}

func (c CommonRequests) GetOne(storyId any) (any, error) {

	var result Story

	items, err := c.CommonRequests.GetOne(storyId)

	if err != nil {
		return result, err
	}
	for items.Next() {
		result = CastToStory(items)
	}

	return result, nil
}
