package model

type Player struct {
	Id       uint   `json:"id"`
	Name     string `json:"name" validate:"required, min=2, max=100"`
	Lastname string `json:"lastname" validate:"required, min=2, max=100"`
	Phone    string `json:"phone" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

func (b *Player) TableName() string {
	return "player"
}
