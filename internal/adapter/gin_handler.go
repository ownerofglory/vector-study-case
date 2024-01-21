package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/ownerofglory/vector-study-case/internal/core/port"
	"github.com/ownerofglory/vector-study-case/internal/util"
	"net/http"
	"strconv"
)

// swagger:parameters findPairRequest
type NumRequest struct {
	// in: query
	// required: true
	// example: 1,2,3,4
	A string `form:"a"`
	// in: query
	// required: true
	// example: 5,6,7
	B string `form:"b"`
	// in: query
	// required: true
	Target int `form:"target"`
}

type httpHandler struct {
	numService port.NumService
}

func NewHandler(numService port.NumService) *httpHandler {
	return &httpHandler{numService: numService}
}

func (h *httpHandler) HandleFindPair(c *gin.Context) {
	var req NumRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.String(http.StatusBadRequest, "request invalid: %v", err)
		return
	}

	a, err := util.StrToIntSlice(req.A)
	if err != nil {
		c.String(http.StatusBadRequest, "unable to process the request data: %v", err)
		return
	}
	b, err := util.StrToIntSlice(req.B)
	if err != nil {
		c.String(http.StatusBadRequest, "unable to process the request data: %v", err)
		return
	}
	target := req.Target

	result := h.numService.FindPairForTarget(a, b, target)

	c.String(http.StatusOK, strconv.FormatBool(result))
}
