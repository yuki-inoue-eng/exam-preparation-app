package routes

import (
	"exam-preparation-app/app/controller"
	"exam-preparation-app/app/trace"
	"net/http"
	"os"
)

// RegisterTheHandler 各種ハンドラをhttp.Handleに登録します。
func RegisterTheHandler() {
	http.HandleFunc("/auth/", controller.LoginHander)
	http.Handle("/login", &controller.TemplateHandler{
		Filename:   "login.html",
		Controller: &controller.LoginController{},
		Tracer:     trace.New(os.Stdout),
	})
	http.HandleFunc("/signup/post", controller.SigunpHandler)
	http.Handle("/signup", &controller.TemplateHandler{
		Filename:   "sigunup.html",
		Controller: &controller.SignupController{},
		Tracer:     trace.New(os.Stdout),
	})
}
