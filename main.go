package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"kannon.gyozatech.dev/generated/sqlc"
)

func main() {
	ps_db, err := sql.Open("postgres", "host=localhost user=postgres dbname=test_cannon port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}

	dbi := sqlc.New(ps_db)
	msg, err := dbi.CreateMessage(context.Background(), sqlc.CreateMessageParams{
		TemplateID:  "pippo",
		Domain:      "ciao@ciao.com",
		Subject:     "ciao",
		SenderEmail: "ciao@ciao.com",
		SenderAlias: "",
		MessageID:   "asdadsda_asddasds",
	})
	if err != nil {
		fmt.Printf("%v", err)
	}

	data, _ := dbi.CreatePool(context.Background(), sqlc.CreatePoolParams{
		ScheduledTime: time.Now(),
		MessageID:     msg.ID,
		Emails:        []string{"email@1.com", "asdasd@email.com"},
	})
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("geenrated %v", data)

}
