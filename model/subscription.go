package model

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/schweigert/inscreveai/view/html"
)

type Subscription struct {
	gorm.Model
	EventID    uint
	UserInfoID uint
	Waiting    bool
	Confirmed  bool
	Event      Event
	UserInfo   UserInfo
}

func (sub *Subscription) AdminRow() html.Dom {
	user := &UserInfo{Name: "Sem nome"}
	db := Db()
	defer db.Close()
	db.Model(sub).Related(&user)

	return html.DivTag(
		`class="col-12 mb-1 text-center"`,
		html.DivTag(
			`class="col-12"`,
			html.Text(`<img class="rounded-circle m-0" style="max-width: 36px" src="`+user.Picture+`">`),
			html.SafeText(user.Name),
			html.If(
				!sub.Waiting,
				html.IfElse(
					sub.Confirmed,
					html.ITag(`class="fas fa-check ml-2"`),
					html.ITag(`class="fas fa-times ml-2"`),
				),
			),
		),
		html.If(
			sub.Waiting,
			html.DivTag(
				`class="row"`,
				html.DivTag(
					`class="col-6"`,
					html.DivTag(
						`class="float-left"`,
						html.FormTag(
							`action="/subscription/`+strconv.Itoa(int(sub.ID))+`/accept" method="POST" class="form-inline"`,
							html.ButtonTag(
								`class="btn btn-outline-warning" type="submit"`,
								html.ITag(`class="fas fa-check"`),
								html.Text("aceitar"),
							),
						),
					),
				),
				html.DivTag(
					`class="col-6"`,
					html.DivTag(
						`class="float-right"`,
						html.FormTag(
							`action="/subscription/`+strconv.Itoa(int(sub.ID))+`/reject" method="POST" class="form-inline"`,
							html.ButtonTag(
								`class="btn btn-outline-warning" type="submit"`,
								html.ITag(`class="fas fa-times"`),
								html.Text("rejeitar"),
							),
						),
					),
				),
			),
		),
	)
}

func (sub *Subscription) PenddingCard() html.Dom {
	event := &Event{Name: "Sem nome"}
	db := Db()
	defer db.Close()
	db.Model(sub).Related(&event)
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
				html.Text("Você está esperando a aprovação desde: "),
				html.Text(sub.CreatedAt),
			),
		),
	)
}

func (sub *Subscription) ConfirmedCard() html.Dom {
	event := &Event{Name: "Sem nome"}
	db := Db()
	defer db.Close()
	db.Model(sub).Related(&event)
	return html.DivTag(
		`class="col-sm-6 mt-3"`,
		html.DivTag(
			`class="card text-white bg-success mb-3"`,
			html.DivTag(
				`class="card-header text-center"`,
				html.StrongTag(
					``,
					html.SafeText(event.Name),
				),
			),
			html.DivTag(
				`class="card-body"`,
				html.Text("Confirmado em: "),
				html.Text(sub.UpdatedAt),
			),
		),
	)
}

func (sub *Subscription) RefusedCard() html.Dom {
	event := &Event{Name: "Sem nome"}
	db := Db()
	defer db.Close()
	db.Model(sub).Related(&event)
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
				html.Text("Cancelado em: "),
				html.Text(sub.UpdatedAt),
			),
		),
	)
}

func PenddingSubscriptionCards(subs []Subscription) []html.Dom {
	doms := []html.Dom{}

	for _, sub := range subs {
		if sub.Waiting {
			doms = append(doms, sub.PenddingCard())
		} else if sub.Confirmed {
			doms = append(doms, sub.ConfirmedCard())
		} else {
			doms = append(doms, sub.RefusedCard())
		}
	}

	return doms
}
