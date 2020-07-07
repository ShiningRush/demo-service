package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
	"github.com/shiningrush/demo-service/internal/application/handler"
	"github.com/shiningrush/demo-service/internal/utils/dao"
	"log"
	"net/http"
)

func main() {
	err := dao.InitMySQL("root:root@tcp(127.0.0.1:3306)/demo_service?charset=utf8&parseTime=True&loc=Local")
	checkErr(err)
	defer func() {
		if err := dao.CloseMySQL(); err != nil {
			fmt.Println("close mysql conn failed: ", err)
		}
	}()

	h := handler.NewAccountHandler()
	restful.DefaultContainer.Add(h.WebService())
	err = http.ListenAndServe(":8080", restful.DefaultContainer)
	if err != nil {
		log.Fatal(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
