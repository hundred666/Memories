package dao

import "model"

func AddUser(user model.User) (error) {
	insert := "INSERT INTO User(Name,Password,RegisterTime,LoginTime) VALUES(?,?,?,?)"
	_, err := Modify(insert, user.Name, user.Password, user.RegisterTime, user.LoginTime)
	return err
}

func UpdateUserLogin(user model.User) (error) {
	update := "UPDATE User set LoginTime=? WHERE Id=?"
	_, err := Modify(update, user.LoginTime, user.Id)
	return err
}

func UpdateUser(user model.User) (error) {
	update := "UPDATE User set Name=?, Password=?,Permission=? WHERE Id=?"
	_, err := Modify(update, user.Name, user.Password, user.Permission, user.Id)
	return err
}

func DelUser(user model.User) (int, error) {
	del := "DELETE FROM User WHERE Id=?"
	return Modify(del, user.Id)
}

func GetUserByName(Name string) (model.User) {
	query := "SELECT * FROM User WHERE Name=?"
	user := model.User{}
	results := Get(query, &model.User{}, Name)
	if len(results) == 0 {
		return user
	}
	user = results[0].(model.User)
	return user
}

func GetUserById(id int) (model.User) {
	query := "SELECT * FROM User WHERE Id=?"
	user := model.User{}
	results := Get(query, &model.User{}, id)
	if len(results) == 0 {
		return user
	}
	user = results[0].(model.User)
	return user
}

func UserLogin(user model.User) (bool) {
	query := "SELECT * FROM User WHERE Name=? AND Password=?"
	results := Get(query, &model.User{}, user.Name, user.Password)
	if len(results) == 0 {
		return false
	}
	return true
}

func AdminPermission(user model.User) (bool) {
	query := "SELECT * FROM User WHERE Name=? AND Password=? AND Permission>4"
	results := Get(query, &model.User{}, user.Name, user.Password)
	if len(results) == 0 {
		return false
	}
	return true
}

func UploadPermission(user model.User) (bool) {
	query := "SELECT * FROM User WHERE Name=? AND Password=? AND Permission>0"
	results := Get(query, &model.User{}, user.Name, user.Password)
	if len(results) == 0 {
		return false
	}
	return true
}

func UserExist(Name string) (bool) {
	query := "SELECT * FROM User WHERE Name=?"
	results := Get(query, &model.User{}, Name)
	if len(results) == 0 {
		return false
	}
	return true
}

func GetUsers(start int, end int, args ...string) []model.User {
	query := "SELECT * FROM User "
	if len(args) != 0 {
		for _, arg := range args {
			query += arg
		}
	}
	query += " LIMIT ?,?"
	results := Get(query, &model.User{}, start, end)
	if len(results) == 0 {
		return nil
	}
	users := make([]model.User, 0)
	for _, res := range results {
		v, ok := res.(model.User)
		if ok {
			users = append(users, v)
		}
	}
	return users
}

func GetUserCount() (int) {
	query := "SELECT COUNT(*) FROM User"
	results := Get(query, &model.Count{})
	if len(results) == 0 {
		return 0
	}
	count := results[0].(model.Count)
	return count.Count
}
