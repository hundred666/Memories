package dao

import (
	"model"
)

func AddComment(c model.Comment) (int, error) {
	insert := "INSERT INTO Comment(Content,CommentTime,CommentType,User,IP,UA) VALUES(?,?,?,?,?,?)"
	return Modify(insert, c.Content, c.CommentTime, c.CommentType, c.User, c.IP, c.UA)
}

func GetComments(start int, end int, args ...string) ([]model.Comment) {
	query := "SELECT * FROM Comment "
	if len(args) != 0 {
		for _, arg := range args {
			query += arg
		}
	}
	query += " LIMIT ?,?"
	results := Get(query, &model.Comment{}, start, end)
	if len(results) == 0 {
		return nil
	}
	comments := make([]model.Comment, 0)
	for _, res := range results {
		v, ok := res.(model.Comment)
		if ok {
			comments = append(comments, v)
		}
	}
	return comments
}
func GetComment(id int, args ...string) (model.Comment) {
	var comment model.Comment
	query := "SELECT * FROM Comment WHERE Id=? "
	if len(args) != 0 {
		for _, arg := range args {
			query += arg
		}
	}
	results := Get(query, &model.Comment{}, id)
	if len(results) == 0 {
		return comment
	}
	comment = results[0].(model.Comment)
	return comment
}

func UpdateComment(c model.Comment) (int, error) {
	update := "UPDATE Comment SET Content=? ,User=? WHERE Id=?"
	return Modify(update, c.Content, c.User, c.Id)
}

func DelComment(c model.Comment) (int, error) {
	mcr := model.MCR{Cid: c.Id}
	DelMCR(mcr)
	del := "DELETE FROM Comment WHERE Id=?"
	return Modify(del, c.Id)
}

func GetCommentCount() (int) {
	query := "SELECT COUNT(*) FROM Comment"
	results := Get(query, &model.Count{})
	if len(results) == 0 {
		return 0
	}
	count := results[0].(model.Count)
	return count.Count
}
