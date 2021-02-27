package courses

import (
	"testing"

	"github.com/jrmanes/ddd-api-go/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/mock"
)

func TestHandler_Create(t *testing.T) {
	courseRepository := new(storagemocks.CourseRepository)
	courseRepository.On("Save", mock.Anything, mock.AnythingOfType("mooc.Course"))

	t.Run("given an invalid request it returns 400", func(t *testing.T) {

	})

	t.Run("given a valid request it returns 201", func(t *testing.T) {

	})

}