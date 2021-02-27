package mysql

// sqlCourseTable is the table name
const (
	sqlCourseTable = "courses"
)

// DTO, reference in our database columns, we can implement here validations
type sqlCourse struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Duration string `db:"duration"`
}