package gintodo

import (
	 "API_ToDo_Golang/component"
	 "API_ToDo_Golang/modules/todo/todobiz"
	 "API_ToDo_Golang/modules/todo/todomodel"
	 "API_ToDo_Golang/modules/todo/todostorage"
	 "github.com/gin-gonic/gin"
	 "net/http"
)

func CreateTodo(atx component.AppContext) gin.HandlerFunc {
	 return func(c *gin.Context) {
		  var data todomodel.ToDoItemCreate
		  if err := c.ShouldBind(&data); err != nil {
			   c.JSON(400, gin.H{
				    "error": err.Error(),
			   })
			   return
		  }
		  
		  store := todostorage.NewSqlStore(atx.GetMainDbConnection())
		  biz := todobiz.NewCreateTodoStore(store)
		  
		  if err := biz.CreateTodo(c.Request.Context(), &data); err != nil {
			   c.JSON(400, gin.H{
				    "error": err.Error(),
			   })
			   return
		  }
		  
		  c.JSON(http.StatusOK, data)
		  
	 }
}
