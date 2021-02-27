package courses

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jrmanes/ddd-api-go/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Create(t *testing.T) {
	courseRepository := new(storagemocks.CourseRepository)
	courseRepository.On("Save", mock.Anything, mock.AnythingOfType("mooc.Course")).Return(nil)

	// setup gin in test mode
	gin.SetMode(gin.TestMode)
	r := gin.New()
	// create the method post for /courses handler
	r.POST("/courses", CreateHandler(courseRepository))

	// create the test cases
	t.Run("given an invalid request it returns 400", func(t *testing.T) {
		createCourseReq := createRequest{
			ID:   "c01d7cf6-ec3f-47f0-9556-a5d6e9009a43",
			Name: "Test",
		}

		b, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		// create the http request
		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(b))
		require.NoError(t, err)

		// use the standard library in order to emulate a new request using httptest
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		// convection, lets close the body for management
		defer res.Body.Close()

		// we validate that the status code is as expected 400
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		createCourseReq := createRequest{
			ID:       "c01d7cf6-ec3f-47f0-9556-a5d6e9009a43",
			Name:     "Test",
			Duration: "1 month",
		}

		b, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		// create the http request
		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(b))
		require.NoError(t, err)

		// use the standard library in order to emulate a new request using httptest
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		// convection, lets close the body for management
		defer res.Body.Close()

		// we validate that the status code is as expected 201
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})

}
