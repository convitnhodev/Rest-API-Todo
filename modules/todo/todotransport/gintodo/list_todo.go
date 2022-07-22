package gintodo

import (
	"API_ToDo_Golang/component"
	"API_ToDo_Golang/modules/todo/todobiz"
	"API_ToDo_Golang/modules/todo/todostorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func List(atx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging component.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Fullfill()

		store := todostorage.NewSqlStore(atx.GetMainDbConnection())
		biz := todobiz.NewListTodoStore(store)

		data, err := biz.ListItems(c.Request.Context(), nil, &paging)

		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, &data)

	}
}
