package dao

import (
	"model"
	"time"
	"log"
	"strings"
	"errors"
	"util"
	"os"
	"path/filepath"
)

func AddMove(m model.Move) (error) {
	insert := "INSERT INTO Move(Content,ImgPath,User,Time) VALUES(?,?,?,?)"
	_, err := Modify(insert, m.Content, m.ImgPath, m.User, m.Time)
	if err != nil {
		log.Println(err)
	}
	return err

}

func AddMoveComment(m model.Move, c model.Comment) (error) {
	id, err := AddComment(c)
	if err != nil {
		return err
	}
	mcr := model.MCR{Mid: m.Id, Cid: id, Time: time.Now()}
	err = AddMCR(mcr)
	return err
}

func GetMoves(start int, end int, args ...string) ([]model.Move) {
	query := "SELECT * FROM Move "
	if len(args) != 0 {
		for _, arg := range args {
			query += arg
		}
	}
	query += " LIMIT ?,?"
	results := Get(query, &model.Move{}, start, end)
	if len(results) == 0 {
		return nil
	}
	moves := make([]model.Move, 0)
	for _, res := range results {
		v, ok := res.(model.Move)
		if ok {
			moves = append(moves, v)
		}
	}
	return moves
}

func GetMove(id int) model.Move {
	var move model.Move
	query := "SELECT * FROM Move WHERE Id=?"
	results := Get(query, &model.Move{}, id)
	if len(results) == 0 {
		return move
	}
	move = results[0].(model.Move)
	return move
}

func GetMoveComments(m model.Move) []model.Comment {
	comments := make([]model.Comment, 0)
	mcrs := GetMCRS(m)
	for _, m := range mcrs {
		comment := GetComment(m.Cid, "AND CommentType=1")
		if comment.Id != 0 {
			comments = append(comments, comment)
		}
	}
	return comments
}

func UpdateMove(m model.Move) (int, error) {
	update := "UPDATE Move SET Content=?,User=? WHERE Id=?"
	return Modify(update, m.Content, m.User, m.Id)
}

func UpdateMoveComment(m model.Move, c model.Comment) (int, error) {
	return UpdateComment(c)
}

func DelMove(m model.Move) (int, error) {
	m = GetMove(m.Id)

	moveImgPath := strings.SplitAfterN(m.ImgPath, "/", 3)[2]
	if moveImgPath == "" {
		return 0, errors.New("portrait path not found")
	}
	staticPath, err := util.GetStaticPath()
	if err != nil {
		return 0, errors.New("static path not found")
	}

	mip := filepath.Join(staticPath, moveImgPath)
	if _, err := util.PathExists(mip); err == nil {
		os.Remove(mip)
	} else {
		return 0, errors.New("file delete failed")
	}

	mcrs := GetMCRS(m)
	for _, m := range mcrs {
		c := model.Comment{Id: m.Cid}
		_, err := DelComment(c)
		if err != nil {
			return 0, err
		}
		DelMCR(m)
	}
	del := "DELETE FROM Move WHERE Id=?"
	return Modify(del, m.Id)
}

func DelMoveComment(m model.Move, c model.Comment) (int, error) {
	mcr := model.MCR{Mid: m.Id, Cid: c.Id}
	_, err := DelMCR(mcr)
	if err != nil {
		return 0, err
	}
	return DelComment(c)
}

func GetMoveCount() (int) {
	query := "SELECT COUNT(*) FROM Move"
	results := Get(query, &model.Count{})
	if len(results) == 0 {
		return 0
	}
	count := results[0].(model.Count)
	return count.Count
}
