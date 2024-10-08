package v1

import (
	"gin-api/server/controller/actuator_ctrl"
	"gin-api/server/router/restful"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	app := restful.New(c)
	req := actuator_ctrl.HealthReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		app.StatusBadRequest(err.Error())
		return
	}
	resp, err := req.Exec()
	if err != nil {
		app.StatusInternalServerError(err.Error())
		return
	}

	app.Ok(resp)
}

func Info(c *gin.Context) {
	app := restful.New(c)
	req := actuator_ctrl.InfoReq{}
	if err := c.ShouldBindQuery(&req); err != nil {
		app.StatusBadRequest(err.Error())
		return
	}
	resp, err := req.Exec()
	if err != nil {
		app.StatusInternalServerError(err.Error())
		return
	}

	app.Ok(resp)
}
