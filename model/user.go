package model

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

type (
	User struct {
		ID                  int64     `xorm:"'id' not null BIGINT(22) pk autoincr"`
		UserName            string    `xorm:"'user_name' unique(unique_user_name) not null VARCHAR(255)"`
		Nickname            string    `xorm:"'nickname' not null VARCHAR(255)"`
		Avatar              string    `xorm:"'avatar' not null VARCHAR(255)"`
		Password            string    `xorm:"'password' not null VARCHAR(255)"`
		RoleID              int       `xorm:"'role_id' not null INT(8)"`
		DepartmentID        int64     `xorm:"'department_id' not null INT(22) index(index_department)"`
		DepartmentName      string    `xorm:"'department_name' not null VARCHAR(255)"`
		Position            string    `xorm:"'position' not null VARCHAR(255)"`
		PhoneNumber         string    `xorm:"'phone_number' unique(unique_phone_number) not null VARCHAR(255)"`
		InternalPhoneNumber string    `xorm:"not null VARCHAR(255)"`
		Disable             int8      `xorm:"not null tinyint"`
		CreatedAt           time.Time `xorm:"'created_at' not null created DATETIME"`
		UpdatedAt           time.Time `xorm:"'updated_at' not null updated DATETIME"`
		LastLoginAt         time.Time `xorm:"'last_login_at' DATETIME"`
		DeletedAt           time.Time `xorm:"'deleted_at' not null deleted DATETIME default('0001-01-01 00:00:00') unique(unique_phone_number) unique(unique_user_name)"`
	}

	userStatic struct{}
)

var UserStatic = new(userStatic)

func (User) TableName() string {
	return "t_users"
}

func (*userStatic) DisableOrEnableUser(ctx context.Context, id int64, isDisable int8) error {

	session := GetSession(ctx)
	defer session.Close()

	_, err := session.ID(id).MustCols("disable").Update(&User{
		Disable: isDisable,
	})

	return err
}

func (*userStatic) UpdateUserInfo(ctx context.Context, id int64, nickname string, phoneNumber string) error {

	session := GetSession(ctx)
	defer session.Close()

	_, err := session.Where("id = ? AND disable = 0", id).Cols("nickname", "phone_number").Update(&User{
		Nickname:    nickname,
		PhoneNumber: phoneNumber,
	})

	return err
}

func (*userStatic) UpdateUserInternalPhoneNumber(ctx context.Context, userID int64, internalPhoneNumber string) error {

	session := GetSession(ctx)
	defer session.Close()

	_, err := session.ID(userID).Cols("internal_phone_number").Update(&User{
		InternalPhoneNumber: internalPhoneNumber,
	})

	return err
}

func (*userStatic) UpdateLoginTime(ctx context.Context, userID int64) error {

	session := GetSession(ctx)
	defer session.Close()

	_, err := session.ID(userID).Cols("last_login_at").Update(&User{
		LastLoginAt: time.Now(),
	})

	return err
}

func (*userStatic) Insert(ctx context.Context, user *User) error {

	session := GetSession(ctx)
	defer session.Close()

	var prevUser User
	found, err := session.Where("user_name = ?", user.UserName).Get(&prevUser)
	if err != nil {
		return err
	}

	if found {
		user.ID = prevUser.ID
		user.CreatedAt = prevUser.CreatedAt
		user.UpdatedAt = prevUser.UpdatedAt
		return nil
	}

	_, err = session.AllCols().Insert(user)

	return errors.WithStack(err)
}

func (*userStatic) Delete(ctx context.Context, userName string) error {

	session := GetSession(ctx)
	defer session.Close()

	_, err := session.Where("user_name = ?", userName).Delete(&User{})
	return errors.WithStack(err)
}

func (*userStatic) GetByName(ctx context.Context, name string) (*User, error) {

	session := GetSession(ctx)
	defer session.Close()

	user := User{}

	found, err := session.Where("user_name = ? AND disable = 0", name).Get(&user)
	if err != nil {
		return nil, err
	}

	if found {
		return &user, nil
	}

	return nil, nil
}

func (*userStatic) GetByDepartment(ctx context.Context, departmentID int64) ([]User, error) {

	session := GetSession(ctx)
	defer session.Close()

	var users []User
	err := session.Where("department_id = ? AND disable = 0", departmentID).Find(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (*userStatic) GetByPhone(ctx context.Context, phone string) (*User, error) {

	session := GetSession(ctx)
	defer session.Close()

	user := User{}

	found, err := session.Where("phone_number = ? AND disable = 0", phone).Get(&user)
	if err != nil {
		return nil, err
	}

	if found {
		return &user, nil
	}

	return nil, nil
}

func (*userStatic) GetByID(ctx context.Context, id int64) (*User, error) {

	session := GetSession(ctx)
	defer session.Close()

	user := User{}

	found, err := session.Where("id = ? AND disable = 0", id).Get(&user)
	if err != nil {
		return nil, err
	}

	if found {
		return &user, nil
	}

	return nil, nil
}

func (*userStatic) UpdateUserPhoneNumberByID(ctx context.Context, id int64, phoneNumber string) error {

	session := GetSession(ctx)
	defer session.Close()

	_, err := session.Where("id = ? AND disable = 0", id).Cols("phone_number").Update(&User{
		PhoneNumber: phoneNumber,
	})

	return errors.WithStack(err)
}

func (*userStatic) UpdateUserRoleByID(ctx context.Context, userID int64, roleID int64) error {

	session := GetSession(ctx)
	defer session.Close()

	_, err := session.Where("id = ? AND disable = 0", userID).Update(&User{
		RoleID: int(roleID),
	})

	return errors.WithStack(err)
}

func (*userStatic) UpdateUserDepartmentByID(ctx context.Context, userID int64, departmentID int64, departmentName string) error {

	session := GetSession(ctx)
	defer session.Close()

	_, err := session.Where("id = ? AND disable = 0", userID).Cols("department_id", "department_name").Update(&User{
		DepartmentID:   departmentID,
		DepartmentName: departmentName,
	})

	return errors.WithStack(err)
}

func (*userStatic) UpdateUserPositionByID(ctx context.Context, userID int64, position string) error {

	session := GetSession(ctx)
	defer session.Close()

	_, err := session.Where("id = ? AND disable = 0", userID).Update(&User{
		Position: position,
	})

	return errors.WithStack(err)
}

func (*userStatic) GetAllUsers(ctx context.Context, page, pageSize int) ([]User, error) {

	session := GetSession(ctx)
	defer session.Close()

	var users []User

	err := session.Asc("id").Limit(pageSize, (page-1)*pageSize).Find(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (*userStatic) GetByRoleID(ctx context.Context, roleID int64) (*User, error) {

	session := GetSession(ctx)
	defer session.Close()

	user := User{}

	found, err := session.Where("role_id = ? AND disable = 0", roleID).Get(&user)
	if err != nil {
		return nil, err
	}

	if found {
		return &user, nil
	}

	return nil, nil
}

func (*userStatic) GetByDepartmentD(ctx context.Context, departmentID int64) (*User, error) {

	session := GetSession(ctx)
	defer session.Close()

	user := User{}

	found, err := session.Where("department_id = ? AND disable = 0", departmentID).Get(&user)
	if err != nil {
		return nil, err
	}

	if found {
		return &user, nil
	}

	return nil, nil
}

func (*userStatic) UpdateUserPassword(ctx context.Context, userID int64, password string) error {

	session := GetSession(ctx)
	defer session.Close()

	_, err := session.ID(userID).Cols("password").Update(&User{
		Password: password,
	})

	return err
}
