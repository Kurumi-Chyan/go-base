package main

//func params(w http.ResponseWriter, r *http.Request) {
//	values := r.URL.Query()
//	fmt.Fprintf(w, "nmmsl is %v\n", values)
//}
//
//func testForm(w http.ResponseWriter, r *http.Request) {
//	err := r.ParseForm()
//	if err != nil {
//		fmt.Fprintf(w, "nmmsl is %v\n", r.Form)
//	}
//	fmt.Fprintf(w, "nmmsl no err is %v\n", r.Form)
//}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

func SignUp(ctx *Context) {
	req := &signUpReq{}
	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
		return
	}
}

func main() {
	Server := NewHttpServer("test-server")
	//Server.Route("/params", params)
	Server.Route("get", "/singup", SignUp)
	Server.Start(":8080")
}
