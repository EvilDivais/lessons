package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

const admin = 129969284

type User struct {
	Name    string
	Address string
	Age     int
}

func (u *User) Info() string {
	return fmt.Sprintf("Name: %s, Address: %s, Age: %d.", u.Name, u.Address, u.Age)
}

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  "1273930063:AAG_n8nurSdu66Gpua06pbJygdEMp0dEZak",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}
	var (
		// Universal markup builders.
		menu = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
		// Reply buttons.
		btnHelp     = menu.Text("ℹ Help \n Помощь")
		btnSettings = menu.Text("⚙ Settings \n Настройки  ")
	)

	menu.Reply(
		menu.Row(btnHelp),
		menu.Row(btnSettings),
	)
	// On reply button pressed (message)
	b.Handle(&btnHelp, func(m *tb.Message) {
		b.Send(m.Sender, "для общение с Луноликим Господином используй следующие просьбы \n/hello - для начала диалога, \n Если ты сам Император то посмотреть список твоих \"друзей\" /users  ")
	})
	b.Handle(&btnSettings, func(m *tb.Message) {
		b.Send(m.Sender, "для изменений в настройках свяжись с Императором")
	})
	users := map[int]*User{}
	b.Handle("/users", func(m *tb.Message) {

		if !isAdmin(m.Sender.ID) {
			b.Send(m.Sender, "Ты чё сука! Ты не Путин, ТЫ ПЕТУХ")
			return
		}
		for _, user := range users {

			b.Send(m.Sender, user.Info())
		}
	})

	b.Handle("/hello", func(m *tb.Message) {
		users[m.Sender.ID] = &User{}
		b.Send(m.Sender, "Привет, Меня зовут Путин я Ваш Бот и врач Лор, Как зовут тебя мой маленьки любитель налогов?!")
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		user := users[m.Sender.ID]
		if user == nil {
			//user = &User{}
			//users[m.Sender.ID] = user
			b.Send(m.Sender, "Привет мой маленький любитель экстремизма, для общение с Луноликим Господином используй следующие просьбы \n/hello - для начала диалога, \n  Если ты сам Император то посмотреть список твоих \"друзей\" /users  ")
			return
		}
		if user.Name == "" {
			user.Name = m.Text
			b.Send(m.Sender, fmt.Sprintf("Привет моя доеная корова, %s. Где живет мой Холоп", user.Name))
			return
		}

		if user.Address == "" {
			user.Address = m.Text
			b.Send(m.Sender, fmt.Sprintf("Надеюсь в, %s. мало оппозиции, а то я придумаю новые налоги. Кстати сколько тебе лет? ", user.Address))
			return
		}

		if 0 == user.Age {
			user.Age, err = strconv.Atoi(m.Text)
			if err != nil {
				b.Send(m.Sender, "Ты дебил? Еще раз напишешь строку я вызову ментов и тебе подкинут наркоту!!! Пиши правду УБЛЮДОК")
				return
			}
			b.Send(m.Sender, fmt.Sprintf("Надеюсь ты не собираешься дожить до пенсии гаденышь, %d. это много, раб не должен только жить ", user.Age))
			b.Send(m.Sender, fmt.Sprintf("Теперь я о тебе все знаю маленький засранец, Ты %s.Напиши мне кто из твоих друзей за Навального?", user.Info()), menu)
			return
		}

	})

	b.Start()

}

func isAdmin(checkId int) bool {
	if checkId == admin {
		return true
	}

	return false
}
