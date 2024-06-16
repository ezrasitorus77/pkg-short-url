package logdb

import (
	"encoding/json"
	"net/http"

	"github.com/ezrasitorus77/pkg-short-url/models"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
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

func (lr LogReqBuilder) record(r *fasthttp.Request, httpCode int, reqBodyString, respBodyString string) {
	var errStr string

	defer fasthttp.ReleaseRequest(r)

	if lr.Err != nil {
		errStr = lr.Err.Error()
	}

	lr.db.Create(&models.LogDBRequest{
		UserID:     0,
		URL:        string(r.URI().Path()),
		Query:      string(r.URI().QueryString()),
		Method:     string(r.Header.Method()),
		Headers:    r.Header.String(),
		ReqBody:    reqBodyString,
		HTTPStatus: int8(httpCode),
		RespCode:   lr.Resp.RC,
		RespBody:   respBodyString,
		Error:      errStr,
	})
}

func (lr LogReqBuilder) setResponse(c *fiber.Ctx) error {
	var (
		resp            models.Response = lr.Resp
		responseBody    []byte
		requestBody     []byte
		newRequest      *fasthttp.Request = fasthttp.AcquireRequest()
		originalRequest *fasthttp.Request = c.Request()
	)

	c.JSON(resp)

	responseBody, _ = json.Marshal(resp)
	requestBody, _ = json.Marshal(lr.ReqBody)

	originalRequest.Header.CopyTo(&newRequest.Header)
	originalRequest.URI().CopyTo(newRequest.URI())

	go lr.record(newRequest, c.Context().Response.StatusCode(), string(requestBody), string(responseBody))

	return nil
}

func (lr LogReqBuilder) OK(c *fiber.Ctx) error {
	c.Status(http.StatusOK)

	return lr.setResponse(c)
}

func (lr LogReqBuilder) Created(c *fiber.Ctx) error {
	c.Status(http.StatusCreated)

	return lr.setResponse(c)
}

func (lr LogReqBuilder) Unauthorized(c *fiber.Ctx) error {
	c.Status(http.StatusUnauthorized)

	return lr.setResponse(c)
}

func (lr LogReqBuilder) ISE(c *fiber.Ctx) error {
	c.Status(http.StatusInternalServerError)

	return lr.setResponse(c)
}

func (lr LogReqBuilder) BadRequest(c *fiber.Ctx) error {
	c.Status(http.StatusBadRequest)

	return lr.setResponse(c)
}

func (lr LogReqBuilder) Write(c *fiber.Ctx, httpCode int) error {
	c.Status(httpCode)

	return lr.setResponse(c)
}
