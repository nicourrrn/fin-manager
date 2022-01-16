package db_test

import (
	"fmt"
	"github.com/nicourrrn/fin-manager/pkg/db"
	"github.com/nicourrrn/fin-manager/pkg/db/models"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestNewConnection(t *testing.T) {
	conn, err := db.NewConnection()
	defer conn.Close()
	if err != nil {
		log.Fatalln(err)
	}
	uRepo := models.NewUserRepo(time.Minute)
	user, err := uRepo.LoadUser(conn, 1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("User login: " + user.Login)
	tRepo := models.NewTransactionRepo(time.Second * 30)
	categories, err := models.LoadCategories(conn)
	transaction, err := tRepo.LoadTransaction(conn, categories, 1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Transaction total " + strconv.Itoa(transaction.Total))
	fmt.Println("Transaction category " + *transaction.Category)
}
