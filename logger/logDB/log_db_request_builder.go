package logdb

import (
	"encoding/json"
	"net/http"

	"github.com/ezrasitorus77/pkg-short-url/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type (
	LogReqBuilder struct {
		LogRequestData
		db *gorm.DB
	}

	LogRequestData struct {
		Resp    models.Response
		ReqBody interface{}
		Err     error
	}
)

func NewLogRequest(db *gorm.DB) LogReqBuilder {
	return LogReqBuilder{
		db: db,
	}
}

func (lr LogReqBuilder) record(c *fiber.Ctx, reqBodyString, respBodyString string) {
	var (
		headers string = c.Request().Header.String()
		errStr  string
	)

	if lr.Err != nil {
		errStr = lr.Err.Error()
	}

	lr.db.Create(&models.LogDBRequest{
		UserID:     0,
		URL:        string(c.Request().URI().Path()),
		Query:      string(c.Request().URI().QueryString()),
		Method:     string(c.Method()),
		Headers:    headers,
		ReqBody:    reqBodyString,
		HTTPStatus: int8(c.Context().Response.StatusCode()),
		RespCode:   lr.Resp.RC,
		RespBody:   respBodyString,
		Error:      errStr,
	})
}

func (lr LogReqBuilder) setResponse(c *fiber.Ctx) {
	var (
		resp         models.Response = lr.Resp
		responseBody []byte
		requestBody  []byte
	)

	c.JSON(resp)

	responseBody, _ = json.Marshal(resp)
	requestBody, _ = json.Marshal(lr.ReqBody)

	go lr.record(c, string(requestBody), string(responseBody))
}

func (lr LogReqBuilder) OK(c *fiber.Ctx) {
	c.Status(http.StatusOK)

	lr.setResponse(c)
}

func (lr LogReqBuilder) Created(c *fiber.Ctx) {
	c.Status(http.StatusCreated)

	lr.setResponse(c)
}

func (lr LogReqBuilder) Unauthorized(c *fiber.Ctx) {
	c.Status(http.StatusUnauthorized)

	lr.setResponse(c)
}

func (lr LogReqBuilder) ISE(c *fiber.Ctx) {
	c.Status(http.StatusInternalServerError)

	lr.setResponse(c)
}

func (lr LogReqBuilder) BadRequest(c *fiber.Ctx) {
	c.Status(http.StatusBadRequest)

	lr.setResponse(c)
}

func (lr LogReqBuilder) Write(c *fiber.Ctx, httpCode int) {
	c.Status(httpCode)

	lr.setResponse(c)
}
