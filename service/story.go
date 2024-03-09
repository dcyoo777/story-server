package service

import (
	"example/request"
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

const StoryPrimaryKey string = "story_id"

var StoryCommon = request.CommonRequests{
	Table:          "story",
	PrimaryKey:     "story_id",
	DatasourceName: "root:1234@tcp(127.0.0.1:3306)/story",
}
