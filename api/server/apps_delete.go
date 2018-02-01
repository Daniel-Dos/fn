package server

import (
	"net/http"

	"github.com/fnproject/fn/api"
	"github.com/fnproject/fn/api/common"
	"github.com/fnproject/fn/api/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleAppDelete(c *gin.Context) {
	ctx := c.Request.Context()
	log := common.Logger(ctx)

	app := &models.App{Name: c.MustGet(api.AppName).(string)}

	err = s.datastore.RemoveApp(ctx, app.Name)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "App deleted"})
}
