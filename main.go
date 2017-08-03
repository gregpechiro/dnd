package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cagnosolutions/adb"
	"github.com/cagnosolutions/web"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

var tmpl *web.TmplCache
var mux *web.Mux
var db *adb.DB = adb.NewDB()
var driver bolt.DriverPool

const NEO4JURL = "bolt://neo4j:admin@localhost:7687"

func init() {
	var err error
	driver, err = bolt.NewDriverPool(NEO4JURL, 10)
	if err != nil {
		panic(err)
	}

	// web.DEFAULT_ERR_ROUTE = web.Route{"GET", "/error/:code", func(w http.ResponseWriter, r *http.Request) {
	// 	code, err := strconv.Atoi(r.FormValue(":code"))
	// 	if err != nil {
	// 		code = 500
	// 	}
	// 	var page = ""
	// 	switch web.GetRole(r) {
	// 	case "DEVELOPER", "ADMIN":
	// 		page = HTTP_ERROR_ADMIN
	// 	case "EMPLOYEE":
	// 		page = HTTP_ERROR_EMPLOYEE
	// 	default:
	// 		page = HTTP_ERROR_DEFAULT
	// 	}
	// 	w.Header().Set("Content-Type", "text/html; utf-8")
	// 	fmt.Fprintf(w, page, code, http.StatusText(int(code)), code)
	// 	return
	// }}

	tmpl = web.NewTmplCache()
	mux = web.NewMux()

	defaultUsers()

}

func main() {

	fmt.Println("DID YOU REGISTER ANY NEW ROUTES?")

	mux.AddRoutes(login, loginPost, logout)
	mux.AddSecureRoutes(ADMIN, adminHome, adminCharacterEdit, adminCharacterSave, adminLocationEdit, adminLocationSave)

	/*mux.AddSecureRoutes(ADMIN, adminHome, adminCampaignEdit, adminCampaignSave)
	mux.AddSecureRoutes(ADMIN, adminHome, adminQuestEdit, adminQuestSave)

	mux.AddSecureRoutes(USER, home)*/

	log.Fatal(http.ListenAndServe(":8080", mux))
}

var login = web.Route{"GET", "/login", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "login.tmpl", nil)
}}

var loginPost = web.Route{"POST", "/login", func(w http.ResponseWriter, r *http.Request) {

	user, err := GetOnlyOneUserByCustom(map[string]interface{}{
		"Email":    r.FormValue("email"),
		"Password": r.FormValue("password"),
		"Active":   true,
	})

	if err != nil {
		if err != NoUserFound && err != MultipleUserFound {
			log.Printf("\nmain.go >> GetOnlyOneUserByCustom() >> %v\n", err)
		}
		web.SetErrorRedirect(w, r, "/login", "Incorrect email or password")
		return
	}

	sess := web.Login(w, r, user.Role)
	sess["id"] = user.Id
	sess["email"] = user.Email
	web.PutMultiSess(w, r, sess)
	redirect := "/"
	if user.Role == "ADMIN" {
		redirect = "/admin"
	}
	web.SetSuccessRedirect(w, r, redirect, "Welcome "+user.FirstName)
	return
}}

var logout = web.Route{"GET", "/logout", func(w http.ResponseWriter, r *http.Request) {
	web.Logout(w)
	web.SetSuccessRedirect(w, r, "/login", "Successfully logged out")
	return
}}

var home = web.Route{"GET", "/", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "home.tmpl", nil)
}}
