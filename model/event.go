package model

import (
	"github.com/jinzhu/gorm"
	"github.com/schweigert/inscreveai/view/html"
)

type Event struct {
	gorm.Model
	Description string `json:"description"`
	Name        string `json:"name"`
	UserInfoId  uint
}

func (event *Event) Card() html.Dom {
	return html.DivTag(
		`class="col-sm-6"`,
		html.DivTag(
			`class="card"`,
			html.DivTag(
				`class="card-body"`,
				html.PTag(
					`card-text`,
					html.Text(event.Description),
				),
			),
		),
	)
}

func AllEventsCards() []html.Dom {
	list := []html.Dom{}
	events := []Event{}
	db := Db()
	db.Close()

	db.Find(&events)

	for _, el := range events {
		list = append(list, el.Card())
	}
	return list
}
