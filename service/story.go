package service

import (
	"example/request"
	"github.com/jmoiron/sqlx"
)

type StoryToDB struct {
	Title   string
	Content string
}

type Story struct {
	StoryId   int `db:"story_id"`
	Title     string
	Content   string
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func (story Story) ToDB() StoryToDB {
	return StoryToDB{
		Title:   story.Title,
		Content: story.Content,
	}
}

type StoryCommonRequests struct {
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

func (c StoryCommonRequests) GetAll() ([]Story, error) {

	var result []Story

	items, err := c.CommonRequests.GetAll()

	if err != nil {
		return result, err
	}

	for items.Next() {
		result = append(result, CastToStory(items))
	}

	return result, nil
}

func (c StoryCommonRequests) GetOne(storyId any) (Story, error) {

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

var StoryCommonReq = StoryCommonRequests{
	request.CommonRequests{
		Table:          "story",
		PrimaryKey:     "story_id",
		DatasourceName: "root:1234@tcp(127.0.0.1:3306)/story",
	},
}
