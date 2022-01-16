package main

import (
	"fmt"
	"github.com/nicourrrn/fin-manager/config"
	"github.com/nicourrrn/fin-manager/pkg/db"
	"github.com/nicourrrn/fin-manager/pkg/db/loaders"
	"github.com/nicourrrn/fin-manager/pkg/models"
	"log"
)

func main() {
	conn, err := db.NewConnection()
	if err != nil {
		log.Fatalln(err)
	}
	repo := loaders.NewUserRepo(config.UserCacheLifeTime)
	user, err := repo.LoadUser(conn, 1)
	models.NewUser("Login", 100)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(user.Login)
}
