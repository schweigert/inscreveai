package model

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/schweigert/inscreveai/view/html"
)

type Event struct {
	gorm.Model
	Description   string `json:"description"`
	Name          string `json:"name"`
	UserInfoID    uint
	Subscriptions []Subscription
}

func (event *Event) AdminCard() html.Dom {
	return html.DivTag(
		`class="col-sm-6 mt-3"`,
		html.DivTag(
			`class="card text-white bg-danger mb-3"`,
			html.DivTag(
				`class="card-header text-center"`,
				html.StrongTag(
					``,
					html.SafeText(event.Name),
				),
			),
			html.DivTag(
				`class="card-body"`,
				html.DivTag(
					`class="row"`,
					SubscriptionAdminList(event.Subscriptions)...,
				),
				html.FormTag(
					`action="/event/`+strconv.Itoa(int(event.ID))+`/delete" method="POST"`,
					html.DivTag(
						`class="text-center mt-3 mb-3"`,
						html.ButtonTag(
							`class="btn btn-outline-warning" type="submit"`,
							html.ITag(`class="fas fa-trash"`),
							html.Text("deletar evento"),
						),
					),
				),
			),
		),
	)
}

func (event *Event) Card() html.Dom {
	return html.DivTag(
		`class="col-sm-6 mt-3"`,
		html.DivTag(
			`class="card text-white bg-warning mb-3"`,
			html.DivTag(
				`class="card-header text-center"`,
				html.StrongTag(
					``,
					html.SafeText(event.Name),
				),
			),
			html.DivTag(
				`class="card-body"`,
				html.PTag(
					`class="card-text"`,
					html.SafeText(event.Description),
				),
				html.FormTag(
					`action="/event/`+strconv.Itoa(int(event.ID))+`/subscribe" method="POST"`,
					html.DivTag(
						`class="text-center mt-3 mb-3"`,
						html.ButtonTag(
							`class="btn btn-outline-danger" type="submit"`,
							html.ITag(`class="fas fa-check"`),
							html.Text("inscrever-se"),
						),
					),
				),
			),
		),
	)
}

func SubscriptionAdminList(subs []Subscription) []html.Dom {
	ret := []html.Dom{}

	for _, sub := range subs {
		ret = append(ret, sub.AdminRow())
	}
	return ret
}

func AllAdminCards(user *UserInfo) []html.Dom {
	list := []html.Dom{}
	events := []Event{}
	db := Db()
	defer db.Close()

	db.Where("user_info_id = ?", user.ID).Preload("Subscriptions").Find(&events)

	for _, el := range events {
		list = append(list, el.AdminCard())
	}
	return list
}

func AllEventsCards(query string) []html.Dom {
	list := []html.Dom{}
	events := []Event{}
	db := Db()
	defer db.Close()

	if query != "" {
		db.Where("name = ? OR description = ?", query, query).Find(&events)
	} else {
		db.Find(&events)
	}

	for _, el := range events {
		list = append(list, el.Card())
	}
	return list
}
