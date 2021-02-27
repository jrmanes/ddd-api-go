package mooc

import "context"

// CourseRepository defines the expected behavior from a course storage
type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CourseRepository

// Course is the data structure that represents a course.
type Course struct {
	id       string
	name     string
	duration string
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