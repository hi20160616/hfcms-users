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

type Role struct {
	RoleId, ParentId, State, Deleted int
	RoleCode, RoleName, Description  string
	UpdateTime                       time.Time
}

type Roles struct {
	Collection []*Role
}

type RoleQuery struct {
	db       *sql.DB
	limit    *int
	offset   *int
	query    string
	clauses  [][4]string // [ ["name", "=", "jack", "and"], ["title", "like", "anything", ""] ]
	order    string
	args     []interface{}
	keywords []string
}

func (dc *DatabaseClient) InsertRole(ctx context.Context, role *Role) error {
	q := `INSERT INTO roles(
		parent_id, code, name, description
		) VALUES (?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		parent_id=?, code=?, name=?, description=?
		`
	aq := &RoleQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query,
		role.ParentId, role.RoleCode, role.RoleName, role.Description,
		role.ParentId, role.RoleCode, role.RoleName, role.Description,
	)
	return errors.WithMessage(err, "mariadb: roles: Insert error")
}

func (dc *DatabaseClient) UpdateRole(ctx context.Context, role *Role) error {
	q := `UPDATE roles SET
		parent_id=?, code=?, name=?, description=?, state=?
		WHERE id=?`
	uq := &RoleQuery{db: dc.db, query: q}
	_, err := uq.db.Exec(uq.query,
		role.ParentId, role.RoleCode, role.RoleName, role.Description,
		role.State, role.RoleId)
	return err
}

// DeleteRole2 is true delete from database instead of DeleteRole just update the row
func (dc *DatabaseClient) DeleteRole2(ctx context.Context, id int) error {
	q := `DELETE FROM roles WHERE id=?`
	aq := &RoleQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRole is soft delete, not delete from database,
// but update deleted field to 1
// DeleteRole is cooperate with All(ctx), that just return
// all rows except deleted is 1
func (dc *DatabaseClient) DeleteRole(ctx context.Context, id int) error {
	q := `UPDATE roles SET deleted=? WHERE id=?`
	aq := &RoleQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, 1, id)
	return err
}

func (dc *DatabaseClient) UndeleteRole(ctx context.Context, id int) error {
	q := `UPDATE roles SET deleted=? WHERE id=?`
	aq := &RoleQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, 0, id)
	return err
}

func (dc *DatabaseClient) QueryRole() *RoleQuery {
	return &RoleQuery{db: dc.db, query: `SELECT * FROM roles`}
}

// All2 will display all rows even if deleted field value is 1
func (aq *RoleQuery) All2(ctx context.Context) (*Roles, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	// rows, err := aq.db.Query("SELECT * FROM roles WHERE title like ?", "%%test%%")
	rows, err := aq.db.Query(aq.query, aq.args...)
	if err != nil {
		return nil, err
	}
	return mkRole(rows)
}

// All will display all lines that the deleted field value is 0
func (aq *RoleQuery) All(ctx context.Context) (*Roles, error) {
	return aq.Where([4]string{"deleted", "=", "0"}).All2(ctx)
}

func (aq *RoleQuery) First(ctx context.Context) (*Role, error) {
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
func (aq *RoleQuery) Where(cs ...[4]string) *RoleQuery {
	aq.clauses = append(aq.clauses, cs...)
	return aq
}

func (aq *RoleQuery) Order(condition string) *RoleQuery {
	aq.order = condition
	return aq
}

func (aq *RoleQuery) Limit(limit int) *RoleQuery {
	aq.limit = &limit
	return aq
}

func (aq *RoleQuery) Offset(offset int) *RoleQuery {
	aq.offset = &offset
	return aq
}

func (aq *RoleQuery) prepareQuery(ctx context.Context) error {
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

func mkRole(rows *sql.Rows) (*Roles, error) {
	var code, name, description sql.NullString
	var update_time sql.NullTime
	var id, state, deleted, parent_id sql.NullInt32
	var roles = &Roles{}
	for rows.Next() {
		if err := rows.Scan(&id, &parent_id, &code, &name, &description,
			&state, &deleted, &update_time); err != nil {
			return nil, errors.WithMessage(err, "mkRole rows.Scan error")
		}
		roles.Collection = append(roles.Collection, &Role{
			RoleId:     int(id.Int32),
			ParentId:   int(parent_id.Int32),
			RoleCode:   code.String,
			RoleName:   name.String,
			State:      int(state.Int32),
			Deleted:    int(deleted.Int32),
			UpdateTime: update_time.Time,
		})
	}
	// TODO: to confirm code below can make sence.
	if err := rows.Err(); err != nil {
		return nil, errors.WithMessage(err, "mkRole error")
	}
	return roles, nil
}
