package repository

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/softplan/tenkai-api/pkg/dbms/model"
	model2 "github.com/softplan/tenkai-api/pkg/dbms/model"
)

//UserDAOInterface UserDAOInterface
type UserDAOInterface interface {
	CreateUser(user model2.User) error
	DeleteUser(id int) error
	AssociateEnvironmentUser(userID int, environmentID int) error
	ListAllUsers(email string) ([]model2.LightUser, error)
	CreateOrUpdateUser(user model2.User) error
	FindByEmail(email string) (model2.User, error)
	FindByID(id string) (model2.User, error)
	FindByUsersIDFilteredByIntersectionEnv(userID, userRequesterID int) (model2.User, error)
}

//UserDAOImpl UserDAOImpl
type UserDAOImpl struct {
	Db *gorm.DB
}

//CreateUser - Creates a new user
func (dao UserDAOImpl) CreateUser(user model2.User) error {
	return dao.Db.Create(&user).Error
}

//DeleteUser - Delete user
func (dao UserDAOImpl) DeleteUser(id int) error {

	var user model2.User

	if err := dao.Db.First(&user, id).Error; err != nil {
		return err
	}

	//Remove all associations
	if err := dao.Db.Model(&user).Association("Environments").Clear().Error; err != nil {
		return err
	}

	if err := dao.Db.Unscoped().Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

//AssociateEnvironmentUser - Associate an environment with a user
func (dao UserDAOImpl) AssociateEnvironmentUser(userID int, environmentID int) error {
	var user model2.User
	var environment model2.Environment

	if err := dao.Db.First(&user, userID).Error; err != nil {
		return err
	}

	if err := dao.Db.First(&environment, environmentID).Error; err != nil {
		return err
	}
	if err := dao.Db.Model(&user).Association("Environments").Append(&environment).Error; err != nil {
		return err
	}

	return nil
}

//ListAllUsers - List all users
func (dao UserDAOImpl) ListAllUsers(email string) ([]model2.LightUser, error) {
	users := make([]model2.LightUser, 0)
	var rows *sql.Rows
	var err error
	clause := "deleted_at IS NULL AND email like '" + email + "%'"

	if rows, err = dao.Db.Table("users").Select([]string{"id", "email"}).Where(clause).Rows(); err != nil {
		return nil, err
	}
	for rows.Next() {
		user := model2.LightUser{}
		rows.Scan(&user.ID, &user.Email)
		users = append(users, user)
	}
	return users, nil
}

func (dao UserDAOImpl) isEditUser(user model2.User) (*model2.User, error) {
	var loadUser model2.User
	if err := dao.Db.Preload("Environments").Where(model2.User{Email: user.Email}).First(&loadUser).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return nil, err
		}
		return nil, nil
	}
	return &loadUser, nil
}

func (dao UserDAOImpl) editUser(user model2.User, loadUser *model2.User) error {
	//will remove roles associated with removed environments
	environmentsRemoved := []model.Environment{}
	for _, env := range loadUser.Environments {
		find := false
		for _, item := range user.Environments {
			if item.ID == env.ID {
				find = true
				break
			}
		}
		if !find {
			environmentsRemoved = append(environmentsRemoved, env)
		}
	}
	if len(environmentsRemoved) > 0 {
		uer := model2.UserEnvironmentRole{}
		for _, env := range environmentsRemoved {
			if err := dao.Db.Where("environment_id = ? and user_id = ?", env.ID, user.ID).Delete(uer).Error; err != nil {
				return err
			}
		}
	}

	//Remove all associations
	if err := dao.Db.Model(&loadUser).Association("Environments").Clear().Error; err != nil {
		return err
	}
	//Associate Envs
	for _, element := range user.Environments {
		var environment model2.Environment
		if err := dao.Db.First(&environment, element.ID).Error; err != nil {
			return err
		}
		if err := dao.Db.Model(&loadUser).Association("Environments").Append(&environment).Error; err != nil {
			return err
		}
	}

	return nil
}

func (dao UserDAOImpl) createUser(user model2.User) error {

	envsToAssociate := user.Environments
	user.Environments = nil

	if err := dao.Db.Create(&user).Error; err != nil {
		return err
	}

	//Associate Envs
	for _, element := range envsToAssociate {
		var environment model2.Environment
		if err := dao.Db.First(&environment, element.ID).Error; err != nil {
			return err
		}
		if err := dao.Db.Model(&user).Association("Environments").Append(&environment).Error; err != nil {
			return err
		}
	}
	return nil
}

//CreateOrUpdateUser - Create or update a user
func (dao UserDAOImpl) CreateOrUpdateUser(user model2.User) error {

	loadUser, err := dao.isEditUser(user)
	if err != nil {
		return err
	}

	if loadUser != nil {
		return dao.editUser(user, loadUser)
	}

	return dao.createUser(user)

}

//FindByEmail FindByEmail
func (dao UserDAOImpl) FindByEmail(email string) (model2.User, error) {
	var user model2.User
	if err := dao.Db.Where(&model2.User{Email: email}).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

//FindByID FindByID
func (dao UserDAOImpl) FindByID(id string) (model2.User, error) {
	var user model2.User
	if err := dao.Db.Preload("Environments").Where("id = ?", id).Take(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

//FindByUsersIDFilteredByIntersectionEnv func
func (dao UserDAOImpl) FindByUsersIDFilteredByIntersectionEnv(userID, userRequesterID int) (model2.User, error) {
	sql := fmt.Sprintf(
		`select
			*
		from
			user_environment ue
		join environments e on
			e.id = ue.environment_id
		join users u on
			u.id = ue.user_id
		where
			ue.user_id = %d
			and e.id in (
			select
				ue2.environment_id
			from
				user_environment ue2
			where
				ue2.user_id = %d);
		`, userID, userRequesterID,
	)
	user := model2.User{}
	rows, err := dao.Db.Raw(sql).Rows()
	if err != nil {
		return model2.User{}, err
	}

	defer rows.Close()
	env := model2.Environment{}
	envs := []model2.Environment{}
	for rows.Next() {
		dao.Db.ScanRows(rows, &user)
		dao.Db.ScanRows(rows, &env)
		envs = append(envs, env)
	}
	user.ID = uint(userID)
	user.Environments = envs
	return user, err
}
