package gin_item

import (
	"Go_backend_tut/common"
	"Go_backend_tut/modules/item/bussiness"
	"Go_backend_tut/modules/item/model"
	"Go_backend_tut/modules/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewSQLStorage(db)
		biz := bussiness.NewCreateItemBussiness(store)

		if err := biz.CreateNewItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
