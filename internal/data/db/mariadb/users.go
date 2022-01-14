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

type User struct {
	Id, Name, Title, Content string
	UserId, CategoryId       int
	UpdateTime               time.Time
}

type Users struct {
	Collection []*User
}

type UserQuery struct {
	db       *sql.DB
	limit    *int
	offset   *int
	query    string
	clauses  [][4]string // [ ["name", "=", "jack", "and"], ["title", "like", "anything", ""] ]
	order    string
	args     []interface{}
	keywords []string
}

func (dc *DatabaseClient) InsertUser(ctx context.Context, user *User) error {
	q := `INSERT INTO users(
		id, title, content, category_id, user_id
		) VALUES (?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		title=?, content=?, category_id=?, user_id=?`
	aq := &UserQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query,
		user.Id, user.Title, user.Content, user.CategoryId, user.UserId,
		user.Title, user.Content, user.CategoryId, user.UserId)
	if err != nil {
		return errors.WithMessage(err, "mariadb: Insert error")
	}
	return nil
}

func (dc *DatabaseClient) UpdateUser(ctx context.Context, user *User) error {
	q := `UPDATE users SET title=?, content=?, category_id=?, user_id=?
		WHERE id=?`
	aq := &UserQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, user.Title, user.Content,
		user.CategoryId, user.UserId, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) DeleteUser(ctx context.Context, id string) error {
	q := `DELETE FROM users WHERE id=?`
	aq := &UserQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) QueryUser() *UserQuery {
	return &UserQuery{db: dc.db, query: "SELECT * FROM users"}
}

func (aq *UserQuery) All(ctx context.Context) (*Users, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	rows, err := aq.db.Query(aq.query, aq.args...)
	// rows, err := aq.db.Query("SELECT * FROM users WHERE title like ?", "%%test%%")
	if err != nil {
		return nil, err
	}
	return mkUser(rows)
}

func (aq *UserQuery) First(ctx context.Context) (*User, error) {
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
func (aq *UserQuery) Where(cs ...[4]string) *UserQuery {
	aq.clauses = append(aq.clauses, cs...)
	return aq
}

func (aq *UserQuery) Order(condition string) *UserQuery {
	aq.order = condition
	return aq
}

func (aq *UserQuery) Limit(limit int) *UserQuery {
	aq.limit = &limit
	return aq
}

func (aq *UserQuery) Offset(offset int) *UserQuery {
	aq.offset = &offset
	return aq
}

func (aq *UserQuery) prepareQuery(ctx context.Context) error {
	if aq.clauses != nil {
		aq.query += " WHERE "
		for i, c := range aq.clauses {
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

func mkUser(rows *sql.Rows) (*Users, error) {
	var id, title, content sql.NullString
	var update_time sql.NullTime
	var user_id, category_id int
	var users = &Users{}
	for rows.Next() {
		if err := rows.Scan(&id, &title, &content, &category_id, &user_id, &update_time); err != nil {
			return nil, errors.WithMessage(err, "mkUser rows.Scan error")
		}
		users.Collection = append(users.Collection, &User{
			Id:         id.String,
			Title:      title.String,
			Content:    content.String,
			CategoryId: category_id,
			UserId:     user_id,
			UpdateTime: update_time.Time,
		})
	}
	// TODO: to confirm code below can make sence.
	if err := rows.Err(); err != nil {
		return nil, errors.WithMessage(err, "mkUser error")
	}
	return users, nil
}
