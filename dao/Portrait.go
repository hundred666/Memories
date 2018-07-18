package dao

import (
	"model"
	"strings"
	"errors"
	"os"
	"path/filepath"
	"util"
)

func AddPortrait(p model.Portrait) (error) {
	insert := "INSERT INTO Portrait(Name,Path,Time,User) VALUES(?,?,?,?)"
	_, err := Modify(insert, p.Name, p.Path, p.Time, p.User)
	return err
}

func DelPortrait(p model.Portrait) (int, error) {
	portraitPath := strings.SplitAfterN(p.Path, "/", 3)[2]
	if portraitPath == "" {
		return 0, errors.New("portrait path not found")
	}
	staticPath, err := util.GetStaticPath()
	if err != nil {
		return 0, errors.New("static path not found")
	}

	err = os.Remove(filepath.Join(staticPath, portraitPath))
	if err != nil {
		return 0, errors.New("file delete failed")
	}

	del := "DELETE FROM Portrait WHERE Id=?"
	return Modify(del, p.Id)
}

func UpdatePortrait(p model.Portrait) (int, error) {
	update := "UPDATE Portrait SET Name=?,User=? WHERE Id=?"
	return Modify(update, p.Name, p.User, p.Id)
}

func GetPortraits(start int, end int, args ...string) ([]model.Portrait) {
	query := "SELECT * FROM Portrait "
	if len(args) != 0 {
		for _, arg := range args {
			query += arg
		}
	}
	query += " LIMIT ?,?"
	results := Get(query, &model.Portrait{}, start, end)
	if len(results) == 0 {
		return nil
	}
	portraits := make([]model.Portrait, 0)
	for _, res := range results {
		v, ok := res.(model.Portrait)
		if ok {
			portraits = append(portraits, v)
		}
	}
	return portraits
}

func GetPortraitCount() (int) {
	query := "SELECT COUNT(*) FROM Portrait"
	results := Get(query, &model.Count{})
	if len(results) == 0 {
		return 0
	}
	count := results[0].(model.Count)
	return count.Count
}
