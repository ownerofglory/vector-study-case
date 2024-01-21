package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/ownerofglory/vector-study-case/internal/core/port"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHandleFindPair(t *testing.T) {
	ctrl := gomock.NewController(t)
	numServiceMock := port.NewMockNumService(ctrl)

	numServiceMock.EXPECT().
		FindPairForTarget(gomock.Any(), gomock.Any(), gomock.Any()).
		MinTimes(0).
		Return(true)

	testCases := []struct {
		name       string
		a, b       string
		target     string
		statusCode int
	}{
		{
			name:       "When a and b are valid comma separated lists of ints and target is an int status OK is returned",
			a:          "1,5",
			b:          "2",
			target:     "3",
			statusCode: http.StatusOK,
		},
		{
			name:       "When a is invalid 'Bad Request' status is returned",
			a:          "",
			b:          "2",
			target:     "3",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "When b is invalid 'Bad Request' status is returned",
			a:          "1,5",
			b:          "invalid",
			target:     "3",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "When target is float 'Bad Request' status is returned",
			a:          "1,5",
			b:          "2",
			target:     "3.5",
			statusCode: http.StatusBadRequest,
		},
	}

	ginHandler := NewHandler(numServiceMock)

	for _, tc := range testCases {
		gin.SetMode(gin.TestMode)
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		queryParams := make(url.Values)
		queryParams.Add("a", tc.a)
		queryParams.Add("b", tc.b)
		queryParams.Add("target", tc.target)

		ctx.Request = &http.Request{
			Form: queryParams,
			URL:  &url.URL{RawQuery: queryParams.Encode()},
		}

		t.Run(tc.name, func(t *testing.T) {
			ginHandler.HandleFindPair(ctx)
		})
	}
}
