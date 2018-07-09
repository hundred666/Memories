package main

import (
	"net/http"
	"util"
	"log"
	"controller/index"
	"controller/comment"
	"controller/portrait"
	"controller/move"
	"controller/admin"
	"controller/user"
)

func main() {

	indexHandler := index.IndexHandler{}
	commentHandler := comment.CommentHandler{}
	portraitsHandler := portrait.PortraitHandler{}
	moveHandler := move.MoveHandler{}
	userHandler := user.UserHandler{}
	adminHandler := admin.AdminHandler{}

	http.HandleFunc("/", indexHandler.Index)
	http.HandleFunc("/comment/addComment", commentHandler.AddComment)
	http.HandleFunc("/comment/getComments", commentHandler.GetComments)
	http.HandleFunc("/comment/getComment", commentHandler.GetComment)
	http.HandleFunc("/comment/getAllComments", commentHandler.GetAllComments)

	http.HandleFunc("/portrait/getPortraits", portraitsHandler.GetPortraits)
	http.HandleFunc("/portrait/viewPortraits", portraitsHandler.View)
	http.HandleFunc("/portrait/addPortrait", portraitsHandler.Add)

	http.HandleFunc("/move/addMove", moveHandler.AddMove)
	http.HandleFunc("/move/addMoveComment", moveHandler.AddComment)
	http.HandleFunc("/move/getMoves", moveHandler.GetMoves)
	http.HandleFunc("/move/getMove", moveHandler.GetMove)
	http.HandleFunc("/move/getMoveComments", moveHandler.GetComments)
	http.HandleFunc("/move/viewMoves", moveHandler.ViewMoves)
	http.HandleFunc("/move/viewMoveDetail", moveHandler.ViewMoveDetail)

	//user
	http.HandleFunc("/user/login", userHandler.Login)
	http.HandleFunc("/user/register", userHandler.Register)
	http.HandleFunc("/user/getUsers", userHandler.GetUsers)
	http.HandleFunc("/user/getUser", userHandler.GetUser)
	http.HandleFunc("/user/logout", userHandler.Logout)

	//admin
	http.HandleFunc("/admin", adminHandler.Index)
	http.HandleFunc("/admin/updateComment", adminHandler.UpdateComment)
	http.HandleFunc("/admin/delComment", adminHandler.DelComment)
	http.HandleFunc("/admin/updateMove", adminHandler.UpdateMove)
	http.HandleFunc("/admin/delMove", adminHandler.DelMove)
	http.HandleFunc("/admin/updateMoveComment", adminHandler.UpdateMoveComment)
	http.HandleFunc("/admin/delMoveComment", adminHandler.DelMoveComment)
	http.HandleFunc("/admin/updateUser", adminHandler.UpdateUser)
	http.HandleFunc("/admin/delUser", adminHandler.DelUser)

	staticPath, err := util.GetStaticPath()
	if err != nil {
		log.Fatal(err)
	}
	fsh := http.FileServer(http.Dir(staticPath))
	http.Handle("/static/", http.StripPrefix("/static/", fsh))
	log.Println("server start")

	http.ListenAndServe(":8000", nil)

}
