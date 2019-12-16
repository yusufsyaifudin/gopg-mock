package main

import (
	"fmt"
	"testing"

	goPgMock "github.com/yusufsyaifudin/gopg-mock"
)

// a successful case
func TestShouldInsertUserProduct(t *testing.T) {
	db, mock, err := goPgMock.NewGoPGDBTest()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	want := &UserProduct{
		UserID:    2,
		ProductID: 3,
	}

	mock.ExpectQuery("INSERT INTO user_products (user_id, product_id) VALUES (?0, ?1)").
		WithArgs(2, 3).
		Returns(goPgMock.NewResult(1, 1, want), nil)

		// now we execute our method
	actual, err := insertUserProduct(db, 2, 3)

	// we make sure that all expectations were met
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if want.UserID != actual.UserID {
		t.Errorf("expected user id = %d, got %d", want.UserID, actual.UserID)
		return
	}

	if want.ProductID != actual.ProductID {
		t.Errorf("expected product id = %d, got %d", want.ProductID, actual.ProductID)
		return
	}
}

// a failing test case
func TestShouldNotInsertUserProduct(t *testing.T) {
	db, mock, err := goPgMock.NewGoPGDBTest()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	errMsg := fmt.Errorf("error insert to db")
	mock.ExpectQuery("INSERT INTO user_products (user_id, product_id) VALUES (?0, ?1)").
		WithArgs(2, 3).
		Returns(goPgMock.NewResult(1, 1, nil), errMsg)

	// now we execute our method
	_, err = insertUserProduct(db, 2, 3)

	// we make sure that the operation must return error as we define it before
	if err == nil {
		t.Errorf(`the error message: "%s", must be returned"`, errMsg.Error())
		return
	}

}
