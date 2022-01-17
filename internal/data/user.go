package data

import (
	"context"
	"log"
	"regexp"
	"time"

	_ "github.com/hi20160616/hfcms-users/api/users/v1"
	_ "github.com/hi20160616/hfcms-users/configs"
	"github.com/hi20160616/hfcms-users/internal/biz"
	"github.com/hi20160616/hfcms-users/internal/data/db/mariadb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ biz.UserRepo = new(userRepo)

type userRepo struct {
	data *Data
	log  *log.Logger
}

func NewUserRepo(data *Data, logger *log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.Default(),
	}
}

// parent=categories/*/users
// TODO parent=tags/*/users
// parent=users/*/users
func (ur *userRepo) ListUsers(ctx context.Context, parent string) (*biz.Users, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	// bizas := &biz.Users{}
	// bizas.Collection = append(bizas.Collection, &biz.User{
	//         UserId:     1,
	//         Username:   "test",
	//         Password:   "PWD",
	//         Realname:   "Real",
	//         Nickname:   "N1ke",
	//         AvatarUrl:  "bigway.jpg",
	//         Phone:      "13912345678",
	//         UserIP:     "127.0.0.1",
	//         State:      0,
	//         Deleted:    0,
	//         CreateTime: timestamppb.Now(),
	//         UpdateTime: timestamppb.Now(),
	// })
	bizas := &biz.Users{}
	us := &mariadb.Users{}
	var err error
	re := regexp.MustCompile(`^(departments|roles|usergroups)/(.+)/users$`)
	x := re.FindStringSubmatch(parent)
	y, err := regexp.MatchString(parent, `^users$`)
	if err != nil {
		return nil, err
	}
	if len(x) != 3 && y {
		us, err = ur.data.DBClient.DatabaseClient.QueryUser().All(ctx)
	} else {
		clause := [4]string{}
		switch x[1] {
		case "departments":
			clause = [4]string{"department_id", "=", x[2], "and"}
		case "roles":
			clause = [4]string{"role_id", "=", x[2], "and"}
		case "usergroups":
			clause = [4]string{"usergroup_id", "=", x[2], "and"}
		}
		us, err = ur.data.DBClient.DatabaseClient.QueryUser().
			Where(clause).All(ctx)
	}
	if err != nil {
		return nil, err
	}
	for _, u := range us.Collection {
		bizas.Collection = append(bizas.Collection, &biz.User{
			UserId:     u.Id,
			Username:   u.Username,
			Password:   u.Password,
			Realname:   u.Realname,
			Nickname:   u.Nickname,
			AvatarUrl:  u.AvatarUrl,
			Phone:      u.Phone,
			UserIP:     u.UserIP,
			State:      u.State,
			Deleted:    u.Deleted,
			CreateTime: timestamppb.New(u.CreateTime),
			UpdateTime: timestamppb.New(u.UpdateTime),
		})
	}
	return bizas, nil
}

func (ar *userRepo) GetUser(ctx context.Context, name string) (*biz.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	return nil, nil
	// name=users/211228101711.111111000001
	// re := regexp.MustCompile(`^users/([\d.]+)$`)
	// x := re.FindStringSubmatch(name)
	// if len(x) != 2 {
	//         return nil, errors.New("name cannot match regex express")
	// }
	// id := x[1]
	// clause := [4]string{"id", "=", id}
	// a, err := ar.data.DBClient.DatabaseClient.QueryUser().
	//         Where(clause).First(ctx)
	// if err != nil {
	//         return nil, err
	// }
	// c := ar.getCate(ctx, a.CategoryId)
	// attrs, err := ar.getAttrs(ctx, id)
	// if err != nil {
	//         return nil, err
	// }
	// tags, err := ar.getTags(ctx, id)
	// if err != nil {
	//         return nil, err
	// }
	// return &biz.User{
	//         UserId:     a.Id,
	//         Title:      a.Title,
	//         Content:    a.Content,
	//         CategoryId: a.CategoryId,
	//         UserId:     a.UserId,
	//         Category:   c,
	//         Attributes: attrs,
	//         Tags:       tags,
	//         UpdateTime: timestamppb.New(a.UpdateTime),
	// }, nil
}

func (ar *userRepo) SearchUsers(ctx context.Context, name string) (*biz.Users, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	return nil, nil
	// re := regexp.MustCompile(`^users/(.+)/search$`)
	// x := re.FindStringSubmatch(name)
	// if len(x) != 2 {
	//         return nil, errors.New("name cannot match regex express")
	// }
	// kws := strings.Split(
	//         strings.TrimSpace(strings.ReplaceAll(x[1], "　", " ")), ",")
	// cs := [][4]string{}
	// for _, kw := range kws {
	//         cs = append(cs,
	//                 // cs will be filtered by Where(clauses...)
	//                 // the last `or` `and` in clause will cut off.
	//                 // so, every clause need `or` or `and` for last element.
	//                 [4]string{"title", "like", kw, "or"},
	//                 [4]string{"content", "like", kw, "or"},
	//         )
	// }
	// as, err := ar.data.DBClient.DatabaseClient.QueryUser().
	//         Where(cs...).All(ctx)
	// if err != nil {
	//         return nil, err
	// }
	// bizas := &biz.Users{Collection: []*biz.User{}}
	// for _, a := range as.Collection {
	//         c := ar.getCate(ctx, a.CategoryId)
	//         bizas.Collection = append(bizas.Collection, &biz.User{
	//                 UserId:     a.Id,
	//                 Title:      a.Title,
	//                 Content:    a.Content,
	//                 CategoryId: a.CategoryId,
	//                 Category:   c,
	//                 UserId:     a.UserId,
	//                 UpdateTime: timestamppb.New(a.UpdateTime),
	//         })
	// }
	// return bizas, nil
}

func (ar *userRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	return nil, nil
	// user.UserId = time.Now().Format("060102150405.000000") +
	//         fmt.Sprintf("%05d", user.UserId)
	// if err := ar.data.DBClient.DatabaseClient.
	//         InsertUser(ctx, &mariadb.User{
	//                 Id:         user.UserId,
	//                 Title:      user.Title,
	//                 Content:    user.Content,
	//                 CategoryId: user.CategoryId,
	//                 UserId:     user.UserId,
	//         }); err != nil {
	//         return nil, err
	// }
	// return user, nil
}

func (ar *userRepo) UpdateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	return nil, nil
	// if err := ar.data.DBClient.DatabaseClient.
	//         UpdateUser(ctx, &mariadb.User{
	//                 Id:         user.UserId,
	//                 Title:      user.Title,
	//                 Content:    user.Content,
	//                 CategoryId: user.CategoryId,
	//                 UserId:     user.UserId,
	//         }); err != nil {
	//         return nil, err
	// }
	// return user, nil
}

func (ar *userRepo) DeleteUser(ctx context.Context, name string) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	return nil, nil
	// re := regexp.MustCompile(`^users/([\d.]+)/delete$`)
	// x := re.FindStringSubmatch(name)
	// if len(x) != 2 {
	//         return &emptypb.Empty{}, errors.New("name cannot match regex express")
	// }
	// return &emptypb.Empty{}, ar.data.DBClient.DatabaseClient.DeleteUser(ctx, x[1])
}
