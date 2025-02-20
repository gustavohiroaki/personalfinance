package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gustavohiroaki/personalfinance/routers/api"
)

func InitRouter(engine *gin.Engine) *gin.Engine {
	r := engine

	r.GET("/ping", api.PingHandler)
	r.POST("/transactions/batch", api.CreateBatchTransaction)
	r.POST("/transactions", api.CreateTransaction)
	r.GET("/transactions", api.GetTransactions)
	r.DELETE("/transactions/:id", api.DeleteTransaction)
	r.PATCH("/transactions/:id", api.UpdateTransaction)

	r.POST("/corporate-events", api.CreateCorporateEvent)
	r.POST("/corporate-events/batch", api.CreateBatchCorporateEvent)
	r.GET("/corporate-events", api.GetCorporateEvents)
	r.DELETE("/corporate-events/:id", api.DeleteCorporateEvent)
	r.PATCH("/corporate-events/:id", api.UpdateCorporateEvent)

	r.GET("/metrics/position", api.GetPosition)
	r.GET("/metrics/position/:id", api.GetPositionByAsset)

	return r
}
