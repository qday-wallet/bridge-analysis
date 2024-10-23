package server

import (
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qday-wallet/bridge-analysis/common/driver"
	"github.com/qday-wallet/bridge-analysis/config"
	"github.com/qday-wallet/bridge-analysis/db"
	"github.com/sirupsen/logrus"
	"github.com/sunjiangjun/xlog"
	"github.com/tidwall/gjson"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

type Handler struct {
	db  *db.DB
	log *logrus.Entry
}

func NewHandler(cfg *config.DB, log *xlog.XLog) *Handler {

	conn, err := driver.Open(cfg.User, cfg.Password, cfg.Addr, cfg.DbName, cfg.Port, log)
	if err != nil {
		panic(err)
	}

	pg := db.NewDB(conn, log)

	return &Handler{
		db:  pg,
		log: log.WithField("module", "handler"),
	}
}

func (h *Handler) Monitor(ctx *gin.Context) {
	b, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}
	startTime := gjson.ParseBytes(b).Get("startTime").String()
	start, err := time.ParseInLocation(TimeFormat, startTime, time.UTC)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}
	endTime := gjson.ParseBytes(b).Get("endTime").String()
	end, err := time.ParseInLocation(TimeFormat, endTime, time.UTC)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}
	array := gjson.ParseBytes(b).Get("status").Array()
	list := make([]int64, 0, 2)
	for _, v := range array {
		list = append(list, v.Int())
	}

	txs, err := h.db.QueryTxs(start, end, list)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}

	h.Success(ctx, string(b), txs, ctx.Request.RequestURI)
}

func (h *Handler) QueryTxs(ctx *gin.Context) {
	b, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}
	address := gjson.ParseBytes(b).Get("address").String()
	pageSize := gjson.ParseBytes(b).Get("pageSize").Int()
	pageNumber := gjson.ParseBytes(b).Get("pageNumber").Int()

	txs, total, err := h.db.QueryTxByFrom(address, pageSize, pageNumber)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}

	mp := make(map[string]any, 2)
	mp["data"] = txs
	mp["total"] = total

	h.Success(ctx, string(b), mp, ctx.Request.RequestURI)
}

func (h *Handler) Income(ctx *gin.Context) {
	b, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}
	startTime := gjson.ParseBytes(b).Get("startTime").String()
	start, err := time.ParseInLocation(TimeFormat, startTime, time.UTC)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}
	endTime := gjson.ParseBytes(b).Get("endTime").String()
	end, err := time.ParseInLocation(TimeFormat, endTime, time.UTC)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}

	points, err := h.db.AssetIncome(start, end)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}

	h.Success(ctx, string(b), points, ctx.Request.RequestURI)
}

func (h *Handler) Pay(ctx *gin.Context) {
	b, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}
	startTime := gjson.ParseBytes(b).Get("startTime").String()
	start, err := time.ParseInLocation(TimeFormat, startTime, time.UTC)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}
	endTime := gjson.ParseBytes(b).Get("endTime").String()
	end, err := time.ParseInLocation(TimeFormat, endTime, time.UTC)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}

	points, err := h.db.AssetPay(start, end)
	if err != nil {
		h.Error(ctx, "", ctx.Request.RequestURI, err.Error())
		return
	}

	h.Success(ctx, string(b), points, ctx.Request.RequestURI)
}

const (
	SUCCESS = 0
	FAIL    = 1
)

func (h *Handler) Success(c *gin.Context, req string, resp interface{}, path string) {
	req = strings.Replace(req, "\t", "", -1)
	req = strings.Replace(req, "\n", "", -1)
	if v, ok := resp.(string); ok {
		resp = strings.Replace(v, "\n", "", -1)
	}
	h.log.Printf("path=%v,req=%v,resp=%v\n", path, req, resp)
	mp := make(map[string]interface{})
	mp["code"] = SUCCESS
	mp["message"] = "ok"
	mp["data"] = resp
	c.JSON(200, mp)
}

func (h *Handler) Error(c *gin.Context, req string, path string, err string) {
	req = strings.Replace(req, "\t", "", -1)
	req = strings.Replace(req, "\n", "", -1)
	h.log.Errorf("path=%v,req=%v,err=%v\n", path, req, err)
	mp := make(map[string]interface{})
	mp["code"] = FAIL
	mp["message"] = err
	mp["data"] = ""
	c.JSON(200, mp)
}
