package dao

import (
	"model"
	"log"
)

func AddAnnounce(a model.Announce) (error) {
	insert := "INSERT INTO Announce(Content,Display,Prior,User,Time) VALUES(?,?,?,?,?)"
	_, err := Modify(insert, a.Content, a.Display, a.Prior, a.User, a.Time)
	if err != nil {
		log.Println(err)
	}
	return err
}

func DelAnnounce(a model.Announce) (int, error) {
	del := "DELETE FROM Announce WHERE Id=?"
	return Modify(del, a.Id)
}

func GetLatestAnnounce() (model.Announce) {
	var announce model.Announce
	query := "SELECT * FROM Announce WHERE Id=  (SELECT MAX(Id) FROM Announce) "
	results := Get(query, &model.Announce{})
	if len(results) == 0 {
		return announce
	}
	announce = results[0].(model.Announce)
	return announce
}

func GetAnnounce(id int, args ...string) (model.Announce) {
	var announce model.Announce
	query := "SELECT * FROM Announce WHERE Id=? "
	if len(args) != 0 {
		for _, arg := range args {
			query += arg
		}
	}
	results := Get(query, &model.Announce{}, id)
	if len(results) == 0 {
		return announce
	}
	announce = results[0].(model.Announce)
	return announce
}

func GetAnnounces(start int, end int, args ...string) ([]model.Announce) {
	query := "SELECT * FROM Announce "
	if len(args) != 0 {
		for _, arg := range args {
			query += arg
		}
	}
	query += " LIMIT ?,?"
	results := Get(query, &model.Announce{}, start, end)
	if len(results) == 0 {
		return nil
	}
	announces := make([]model.Announce, 0)
	for _, res := range results {
		v, ok := res.(model.Announce)
		if ok {
			announces = append(announces, v)
		}
	}
	return announces
}

func UpdateAnnounce(a model.Announce) (int, error) {
	update := "UPDATE Announce SET Content=?,Display=? ,User=? WHERE Id=?"
	return Modify(update, a.Content, a.Display, a.User, a.Id)
}
