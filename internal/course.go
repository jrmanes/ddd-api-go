package mooc

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidCourseID = errors.New("invalid Course ID")

// CourseID represents the course unique identifier
type CourseID struct {
	value string
}

// CourseName represents the course unique identifier
type CourseName struct {
	value string
}

// CourseDuration represents the course unique identifier
type CourseDuration struct {
	value string
}

// NewCourseID instantiate the VO for CourseID
func NewCourseID(value string) (CourseID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return CourseID{}, fmt.Errorf("%w: %s", ErrInvalidCourseID, value)
	}

	return CourseID{
		value: v.String(),
	},nil
}

// CourseRepository defines the expected behavior from a course storage
type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CourseRepository

// Course is the data structure that represents a course.
type Course struct {
	id       CourseID
	name     CourseName
	duration CourseDuration
}

// NewCourse creates a new course
func NewCourse(id, name, duration string) Course {
	return Course{
		id:       id,
		name:     name,
		duration: duration,
	}
}

// Name returns the course name
func (c Course) Name() string {
	return c.name
}

// ID returns the course id
func (c Course) ID() string {
	return c.id
}

// Duration returns the course duration
func (c Course) Duration() string {
	return c.duration
}