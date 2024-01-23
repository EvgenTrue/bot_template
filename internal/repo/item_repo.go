package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/evgentrue/bot_template/internal/storage"
	"github.com/jackc/pgx"
)

type ItemRepo struct {
	db *pgx.Conn
}

func NewItemRepo(d *pgx.Conn) *ItemRepo {          ///  ?
	return &ItemRepo{
		db: d,
	}
}

func (i *ItemRepo) AddItem(ctx context.Context, item storage.Item) {                   /// добавляем в базу и заполняем 3 поля 

	query := `
	insert into item (name, sum, currency) values ($1, $2, $3)
`
	i.db.ExecEx(ctx, query, nil, item.Name, item.Sum, item.Currency)        ///  ??

}
func (i *ItemRepo) GetItem(ctx context.Context, id int) storage.Item {              ///ищем в базе по id  нужную покупку 
	query := `													 
	select id, name, currency, sum from item where id=$1
	`															// создаем переменную для удобства
	row := i.db.QueryRowEx(ctx, query, nil, id)                             /// создаем переменную для поиска переменной по айди
	var item storage.Item
	err := row.Scan(&item.Id, &item.Name, &item.Currency, &item.Sum)  //// проверка на ошибку 
	if err != nil {
		fmt.Println(err)
	}

	return item
}

func (i *ItemRepo) GetLast(ctx context.Context) storage.Item {									// найти последнее значение
	ctx, cancel:= context.WithTimeout(ctx, time.Second)
	defer cancel()
	query := `
	select id, name, currency, sum from item order by id desc limit 1       
	`																		/// создаем переменную для возврата последнего значения с лимитом времени
	row := i.db.QueryRowEx(ctx, query, nil)												//  записываем в переменную для удобства
	var item storage.Item													// создаем переменную для записи в структуру
	err := row.Scan(&item.Id, &item.Name, &item.Currency, &item.Sum)        // проверка на ошибку
	if err != nil {
		fmt.Println(err)
	}

	return item																// возвращаем значение
}