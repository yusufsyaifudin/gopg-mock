package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type UserProduct struct {
	UserID    int64
	ProductID int64
}

func insertUserProduct(db orm.DB, userID, productID int64) (result *UserProduct, err error) {
	result = &UserProduct{}
	_, err = db.QueryContext(context.Background(),
		result,
		"INSERT INTO user_products (user_id, product_id) VALUES (?0, ?1)",
		userID, productID)

	return
}

func main() {
	var (
		pgOptions = &pg.Options{}
		err       error
	)

	// @NOTE: the real connection is not required for tests
	pgOptions, err = pg.ParseURL("postgres://user:pass@localhot:5432?sslmode=disable")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error db master options: %s\n", err.Error())
		return
	}

	db := pg.Connect(pgOptions)
	defer func() {
		_ = db.Close()
	}()

	userID := int64(1)
	productID := int64(10)
	userProduct, err := insertUserProduct(db, userID, productID)
	if err != nil {
		panic(err)
	}

	fmt.Println(userProduct)
}
