package appUser

import "github.com/google/uuid"

type ToDB struct {
	Name  string
	Bio   string
	Image string
}

type User struct {
	Id        uuid.UUID
	Name      string
	Bio       string
	Image     string
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func (user User) ToDB() ToDB {
	return ToDB{
		Name:  user.Name,
		Bio:   user.Bio,
		Image: user.Image,
	}
}
