package initializer

import (
	"context"
	"fmt"
	"rest-api/domain/user/domain/ent"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase() (*ent.Client, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		"root",
		"rnwkgus",
		"localhost",
		"3306",
		"go_rest_api",
	)

	client, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}
