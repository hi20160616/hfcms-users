package mariadb

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type Department struct {
	DepartmentId, ParentId, State, Deleted      int
	DepartmentCode, DepartmentName, Description string
	UpdateTime                                  time.Time
}

type Departments struct {
	Collection []*Department
}

type DepartmentQuery struct {
	db       *sql.DB
	limit    *int
	offset   *int
	query    string
	clauses  [][4]string // [ ["name", "=", "jack", "and"], ["title", "like", "anything", ""] ]
	order    string
	args     []interface{}
	keywords []string
}

func (dc *DatabaseClient) InsertDepartment(ctx context.Context, department *Department) error {
	q := `INSERT INTO departments(
		parent_id, code, name, description
		) VALUES (?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		parent_id=?, code=?, name=?, description=?
		`
	aq := &DepartmentQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query,
		department.ParentId, department.DepartmentCode, department.DepartmentName, department.Description,
		department.ParentId, department.DepartmentCode, department.DepartmentName, department.Description,
	)
	return errors.WithMessage(err, "mariadb: departments: Insert error")
}

func (dc *DatabaseClient) UpdateDepartment(ctx context.Context, department *Department) error {
	q := `UPDATE departments SET
		parent_id=?, code=?, name=?, description=?, state=?
		WHERE id=?`
	uq := &DepartmentQuery{db: dc.db, query: q}
	_, err := uq.db.Exec(uq.query,
		department.ParentId, department.DepartmentCode, department.DepartmentName, department.Description,
		department.State, department.DepartmentId)
	return err
}

// DeleteDepartment2 is true delete from database instead of DeleteDepartment just update the row
func (dc *DatabaseClient) DeleteDepartment2(ctx context.Context, id int) error {
	q := `DELETE FROM departments WHERE id=?`
	aq := &DepartmentQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteDepartment is soft delete, not delete from database,
// but update deleted field to 1
// DeleteDepartment is cooperate with All(ctx), that just return
// all rows except deleted is 1
func (dc *DatabaseClient) DeleteDepartment(ctx context.Context, id int) error {
	q := `UPDATE departments SET deleted=? WHERE id=?`
	aq := &DepartmentQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, 1, id)
	return err
}

func (dc *DatabaseClient) UndeleteDepartment(ctx context.Context, id int) error {
	q := `UPDATE departments SET deleted=? WHERE id=?`
	aq := &DepartmentQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, 0, id)
	return err
}

func (dc *DatabaseClient) QueryDepartment() *DepartmentQuery {
	return &DepartmentQuery{db: dc.db, query: `SELECT * FROM departments`}
}

// All2 will display all rows even if deleted field value is 1
func (aq *DepartmentQuery) All2(ctx context.Context) (*Departments, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	// rows, err := aq.db.Query("SELECT * FROM departments WHERE title like ?", "%%test%%")
	rows, err := aq.db.Query(aq.query, aq.args...)
	if err != nil {
		return nil, err
	}
	return mkDepartment(rows)
}

// All will display all lines that the deleted field value is 0
func (aq *DepartmentQuery) All(ctx context.Context) (*Departments, error) {
	return aq.Where([4]string{"deleted", "=", "0"}).All2(ctx)
}

func (aq *DepartmentQuery) First(ctx context.Context) (*Department, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes.Collection) == 0 {
		return nil, ErrNotFound
	}
	return nodes.Collection[0], nil
}

// cs: {["name", "=", "jack", "and"], ["title", "like", "anything", ""]}
// the last `or` or `and` in clause will cut off after prepareQuery().
// so, every clause need `or` or `and` for last element.
func (aq *DepartmentQuery) Where(cs ...[4]string) *DepartmentQuery {
	aq.clauses = append(aq.clauses, cs...)
	return aq
}

func (aq *DepartmentQuery) Order(condition string) *DepartmentQuery {
	aq.order = condition
	return aq
}

func (aq *DepartmentQuery) Limit(limit int) *DepartmentQuery {
	aq.limit = &limit
	return aq
}

func (aq *DepartmentQuery) Offset(offset int) *DepartmentQuery {
	aq.offset = &offset
	return aq
}

func (aq *DepartmentQuery) prepareQuery(ctx context.Context) error {
	if aq.clauses != nil {
		aq.query += " WHERE "
		for i, c := range aq.clauses {
			// TODO: 2nd clause cannot be tied together automaticly
			// the last `or` or `and` in clause will cut off there.
			// so, every clause need `or` or `and` for last element.
			if i == len(aq.clauses)-1 {
				aq.query += fmt.Sprintf(" %s %s ?", c[0], c[1])
			} else {
				aq.query += fmt.Sprintf(" %s %s ? %s", c[0], c[1], c[3])
			}
			if strings.ToLower(c[1]) == "like" {
				c[2] = fmt.Sprintf("%%%s%%", c[2])
			} else {
				c[2] = fmt.Sprintf("%s", c[2])
			}
			aq.args = append(aq.args, c[2])
		}
	}
	if aq.order != "" {
		aq.query += " ORDER BY ?"
		aq.args = append(aq.args, aq.order)
	}
	if aq.limit != nil {
		aq.query += " LIMIT ?"
		a := strconv.Itoa(*aq.limit)
		aq.args = append(aq.args, a)
	}
	if aq.offset != nil {
		aq.query += ", ?"
		a := strconv.Itoa(*aq.offset)
		aq.args = append(aq.args, a)
	}
	return nil
}

func mkDepartment(rows *sql.Rows) (*Departments, error) {
	var code, name, description sql.NullString
	var update_time sql.NullTime
	var id, state, deleted, parent_id sql.NullInt32
	var departments = &Departments{}
	for rows.Next() {
		if err := rows.Scan(&id, &parent_id, &code, &name, &description,
			&state, &deleted, &update_time); err != nil {
			return nil, errors.WithMessage(err, "mkDepartment rows.Scan error")
		}
		departments.Collection = append(departments.Collection, &Department{
			DepartmentId:   int(id.Int32),
			ParentId:       int(parent_id.Int32),
			DepartmentCode: code.String,
			DepartmentName: name.String,
			State:          int(state.Int32),
			Deleted:        int(deleted.Int32),
			UpdateTime:     update_time.Time,
		})
	}
	// TODO: to confirm code below can make sence.
	if err := rows.Err(); err != nil {
		return nil, errors.WithMessage(err, "mkDepartment error")
	}
	return departments, nil
}
