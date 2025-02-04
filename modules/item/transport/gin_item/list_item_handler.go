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

func GetItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		paging.Process()

		var filter model.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStorage(db)
		biz := bussiness.NewListItemBussiness(store)

		result, err := biz.ListItemBuz(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
