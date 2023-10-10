package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/evgentrue/bot_template/internal/storage"
	"github.com/evgentrue/bot_template/internal/usecase"
	"github.com/evgentrue/bot_template/internal/usecase/provider"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Item struct {
	Name string
	Sum  int
	Cur  string
}

func main() {
	// https://t.me/ex_2_go_bot
	bot, err := tgbotapi.NewBotAPI("6516198339:AAH7kfNtyYH3-TiODIFJ8LgFBMGs352mfvU")
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	s := storage.NewMemoryStorage()
	fixer := provider.NewFixerProvider()
	//currate := provider.NewCurrateProvider()
	calc := usecase.New(fixer)

	//Новая midi клавиатура 1000 USD
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			str := strings.Split(update.Message.Text, " ")

			name := strings.Join(str[:len(str)-2], " ")

			sum := str[len(str)-2]
			cur := str[len(str)-1]
			sumint, err := strconv.Atoi(sum)
			if err != nil {
				fmt.Println(err)
			}
			r, err := calc.Calculate(sumint, cur)
			if err != nil {
				continue
			}

			s.AddItem(storage.Item{Name: name, Sum: r, Currency: "RUB"})
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}

}

// 0. Репу форкнуть, склонировать
//1. Распарсить строку // string, regexp, strings.Contains
//2. Создать InMemoryStorage (map + lock) // Продукты 1000 RUB map + lock (* interface)
//2.1 создать структуру с мапой (лежит в internal/storage)
//2.2 сделать методы у структруры чтобы класть данные
//2.3 сделать тест к методу
//3. Перевод из валюты в рубли по курсу (1 запрос к апи) // 1 net query // http package , module go
//3.1 добавить пакет usecase сделать в нем файлик calculate.
//3.2 сделать структуру usecase с метододом вычисления валюты(метод принимает сумму и валюту)
//3.3 использовать пакет http для перевода в рубли(сделать get или post)
//3.4 разобрать ответ от сервера
//4. Конкурентные походы по сети за курсом 4 запроса параллельно, и выбираешь минимальный курс (interface + solid) // goroutine, channel, interface
//5. Команду вывода в боте, sum // map
//6. InMemory убираем, добавляем DB // clean arch, repository, models, sql
//7. Команду вывода с фильтром по месяцу
//8. Actions, linters // ci/cd

//go1.21.1
