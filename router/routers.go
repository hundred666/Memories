package router

import (
	"net/http"
	"handler/index"
	"handler/comment"
	"handler/portrait"
	"handler/move"
	"handler/user"
	"handler/admin"
	"util"
	"log"
)

var Mux *http.ServeMux

func InitRT() {
	log.Println("init router")

	mux := http.NewServeMux()

	indexHandler := index.IndexHandler{}
	commentHandler := comment.CommentHandler{}
	portraitsHandler := portrait.PortraitHandler{}
	moveHandler := move.MoveHandler{}
	userHandler := user.UserHandler{}
	adminHandler := admin.AdminHandler{}

	mux.HandleFunc("/", indexHandler.Index)
	mux.HandleFunc("/comment/addComment", commentHandler.AddComment)
	mux.HandleFunc("/comment/getComments", commentHandler.GetComments)
	mux.HandleFunc("/comment/getComment", commentHandler.GetComment)
	mux.HandleFunc("/comment/getAllComments", commentHandler.GetAllComments)

	mux.HandleFunc("/portrait/getPortraits", portraitsHandler.GetPortraits)
	mux.HandleFunc("/portrait/viewPortraits", portraitsHandler.View)
	mux.HandleFunc("/portrait/addPortrait", portraitsHandler.Add)

	mux.HandleFunc("/move/addMove", moveHandler.AddMove)
	mux.HandleFunc("/move/addMoveComment", moveHandler.AddComment)
	mux.HandleFunc("/move/getMoves", moveHandler.GetMoves)
	mux.HandleFunc("/move/getMove", moveHandler.GetMove)
	mux.HandleFunc("/move/getMoveComments", moveHandler.GetComments)
	mux.HandleFunc("/move/viewMoves", moveHandler.ViewMoves)
	mux.HandleFunc("/move/viewMoveDetail", moveHandler.ViewMoveDetail)

	//user
	mux.HandleFunc("/user/login", userHandler.Login)
	mux.HandleFunc("/user/register", userHandler.Register)
	mux.HandleFunc("/user/getUsers", userHandler.GetUsers)
	mux.HandleFunc("/user/getUser", userHandler.GetUser)
	mux.HandleFunc("/user/logout", userHandler.Logout)

	//admin
	mux.HandleFunc("/admin", adminHandler.Index)
	mux.HandleFunc("/admin/updateComment", adminHandler.UpdateComment)
	mux.HandleFunc("/admin/delComment", adminHandler.DelComment)
	mux.HandleFunc("/admin/updateMove", adminHandler.UpdateMove)
	mux.HandleFunc("/admin/delMove", adminHandler.DelMove)
	mux.HandleFunc("/admin/updateMoveComment", adminHandler.UpdateMoveComment)
	mux.HandleFunc("/admin/delMoveComment", adminHandler.DelMoveComment)
	mux.HandleFunc("/admin/updateUser", adminHandler.UpdateUser)
	mux.HandleFunc("/admin/delUser", adminHandler.DelUser)

	staticPath, err := util.GetStaticPath()
	if err != nil {
		log.Fatal(err)
	}
	fsh := http.FileServer(http.Dir(staticPath))
	mux.Handle("/static/", http.StripPrefix("/static/", fsh))

	Mux = mux
}
