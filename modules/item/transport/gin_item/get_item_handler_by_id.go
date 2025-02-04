package gin_item

import (
	"Go_backend_tut/common"
	"Go_backend_tut/modules/item/bussiness"
	"Go_backend_tut/modules/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetItemById(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		strconv.Atoi(c.Param("id"))
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		store := storage.NewSQLStorage(db)
		biz := bussiness.NewGetItemBussinessById(store)

		data, err := biz.GetItemBuzById(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
