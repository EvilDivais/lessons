package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

type User struct {
	Name    string
	Address string
	Age     int
}

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  "1273930063:AAFiwfWvxMxyBGWFhdFSPOIp4CBx-M633B4",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/setuser", func(m *tb.Message) {
		b.Send(m.Sender, "Привет, Меня зовут Путин я Ваш Бот и врач Лор, Как зовут тебя мой маленьки любитель налогов?!")
	})
	user := &User{}
	b.Handle(tb.OnText, func(m *tb.Message) {
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
			user.Age, _ = strconv.Atoi(m.Text)
			b.Send(m.Sender, fmt.Sprintf("Надеюсь ты не собираешься дожить до пенсии гаденышь, %d. это много, раб не должен только жить ", user.Age))
			return
		}

	})
	b.Start()

}
