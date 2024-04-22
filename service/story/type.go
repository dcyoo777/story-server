package story

import "github.com/google/uuid"

type ToDB struct {
	Title   string
	Content string
	Place   string
	Start   string
	End     string
}

type Story struct {
	Id        uuid.UUID
	Title     string
	Content   string
	Place     string
	Start     string
	End       string
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func (story Story) ToDB() ToDB {
	return ToDB{
		Title:   story.Title,
		Content: story.Content,
		Place:   story.Place,
		Start:   story.Start,
		End:     story.End,
	}
}
