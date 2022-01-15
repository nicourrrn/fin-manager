package main

import (
	"fmt"
	"github.com/nicourrrn/fin-manager/config"
	"github.com/nicourrrn/fin-manager/pkg/db"
	"github.com/nicourrrn/fin-manager/pkg/db/loaders"
	"log"
)

func main() {
	conn, err := db.NewConnection()
	if err != nil {
		log.Fatalln(err)
	}
	repo := loaders.NewUserRepo(config.UserCacheLifeTime)
	user, err := repo.LoadUser(conn, 1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(user.Login)
}
