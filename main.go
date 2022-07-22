package main

import (
	 "API_ToDo_Golang/component"
	 "API_ToDo_Golang/modules/todo/todotransport/gintodo"
	 "github.com/gin-gonic/gin"
	 "gorm.io/driver/mysql"
	 "gorm.io/gorm"
	 "log"
)

func main() {
	 dsn := "foot_delivery:Thaothaothao2230@tcp(localhost:3306)/foot_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	 db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	 if err := runService(db); err != nil {
		  log.Fatalln(err)
	 }
}
func runService(db *gorm.DB) error {
	 r := gin.Default()
	 appCtx := component.NewAppContext(db)
	 
	 todo := r.Group("/todo")
	 {
		  todo.POST("", gintodo.CreateTodo(appCtx))
	 }
	 
	 return r.Run()
}
