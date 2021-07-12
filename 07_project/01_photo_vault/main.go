package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	host     = "localhost"
	port     = 5432
	psqluser = "postgres"
	password = "password"
	dbname   = "myDatabase"
)

var psql *sql.DB
var tpl *template.Template

type user struct {
	Name     string
	Email    string
	Password string
	Dob      string
	Gender   string
}
type userinfo struct {
	Namex       string
	Emailx      string
	Passwordx   string
	Dobx        string
	Genderx     string
	Photosx     []string
	Sizex       float64
	Profilepicx string
}

func init() {
	var err error
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, psqluser, password, dbname)
	psql, err = sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	if err = psql.Ping(); err != nil {
		panic(err)
	}
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/userhomepage", userhomepage)
	http.Handle("/users/", http.StripPrefix("/users", http.FileServer(http.Dir("./users"))))
	http.Handle("/icons/", http.StripPrefix("/icons", http.FileServer(http.Dir("./styles/icons"))))
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/logout", logout)
	http.Handle("/styles/", http.StripPrefix("/styles", http.FileServer(http.Dir("./styles"))))
	http.Handle("/favicon/", http.StripPrefix("/favicon", http.FileServer(http.Dir("./favicon"))))
	http.ListenAndServe(":8080", nil)
}
func signup(w http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" {
		fmt.Println("from post method")
		err := req.ParseForm()
		req.ParseMultipartForm(32)
		if err != nil {
			log.Fatalln(err)
		}
		name := req.FormValue("name")
		email := req.FormValue("email")
		pass := req.FormValue("password")
		dob := req.FormValue("date")
		gender := req.FormValue("gender")
		f, k, err := req.FormFile("profilepic")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Println("formdata", name, email, pass, dob, gender)
		defer f.Close()
		//testing for already taken
		if rowExists("SELECT id FROM userid WHERE email=$1", email) {
			fmt.Fprint(w, "emailid already taken")
			return
		}
		//encrypting password
		password, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println("password generated", password)
		strpass := string(password)
		// fmt.Println("generated password converted to string format", strpass)
		uid := uuid.NewV4()
		c := &http.Cookie{
			Name:     "session",
			Value:    uid.String(),
			HttpOnly: true,
		}
		//creating files in storage
		os.Mkdir("./users/"+email, 0777)
		os.Mkdir("./users/"+email+"/photos", 0777)
		os.Mkdir("./users/"+email+"/profilepic", 0777)
		pic, err := os.Create(filepath.Join("./users/"+email+"/profilepic/", k.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer pic.Close()
		io.Copy(pic, f)
		http.SetCookie(w, c)
		//saving user details and cookie in database
		psql.Exec("insert into userid(name,email,password,dob,gender,cookie) values ($1,$2,$3,$4,$5,$6)", name, email, strpass, dob, gender, c.Value)
		http.Redirect(w, req, "/login", http.StatusSeeOther)

		return
	}
	tpl.ExecuteTemplate(w, "signup.html", nil)

}
//
func rowExists(query string, args ...interface{}) bool {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	// fmt.Println("checking query data :-",query)
	err := psql.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		fmt.Printf("error checking if row exists '%s' %v", args, err)

	}
	return exists
}
func login(w http.ResponseWriter, req *http.Request) {
	// fmt.Println("from login func")
	c, err := req.Cookie("session")

	if err == nil {
		if rowExists("select cookie from userid where cookie=$1", c.Value) {
			usr := user{}
			row := psql.QueryRow("select name,email,password,dob,gender,cookie from userid where cookie=$1", c.Value)
			row.Scan(&usr.Name, &usr.Email, &usr.Password, &usr.Dob, &usr.Gender)
			fmt.Println(usr)
			// fmt.Println("auto login")
			http.Redirect(w, req, "/userhomepage", http.StatusSeeOther)
			return
		}
	}
	// if rowExists("select cookie from users where cookie=$1", c.Value) {
	if req.Method == "POST" {
		email := req.FormValue("email")
		password := req.FormValue("password")
		// fmt.Println("entered password", password)
		row := psql.QueryRow("select password from userid where email=$1", email)
		var encpassword string
		row.Scan(&encpassword)
		hash := []byte(encpassword)
		// fmt.Println("password from psql data", hash)
		err = bcrypt.CompareHashAndPassword(hash, []byte(password))
		if err == nil {
			uid := uuid.NewV4()
			c := &http.Cookie{
				Name:     "session",
				Value:    uid.String(),
				HttpOnly: true,
			}
			http.SetCookie(w, c)
			psql.Exec("update userid set cookie = $1 where email = $2", c.Value, email)
			http.Redirect(w, req, "/userhomepage", http.StatusSeeOther)
			// fmt.Println("user found hurraaa...!!!!!")
			return
		} else {
			fmt.Println(err, "some error occur in login")
			tpl.ExecuteTemplate(w, "login.html", "email and password does not match or user does not exists")
			return
		}
	}
	// }
	tpl.ExecuteTemplate(w, "login.html", nil)

}
func upload(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		c, err := req.Cookie("session")

		if err == nil {
			if rowExists("select cookie from userid where cookie=$1", c.Value) {
				usr := user{}
				row := psql.QueryRow("select name,email,password,dob,gender from userid where cookie=$1", c.Value)
				row.Scan(&usr.Name, &usr.Email, &usr.Password, &usr.Dob, &usr.Gender)
				// fmt.Println(usr)
				f, k, err := req.FormFile("uploadpic")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				defer f.Close()
				n, err := os.Create(filepath.Join("./users/"+usr.Email+"/photos", k.Filename))
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				defer n.Close()
				io.Copy(n, f)
				http.Redirect(w, req, "/userhomepage", http.StatusSeeOther)
				return

			}
		}

	}
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
func userhomepage(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err == nil {
		if rowExists("select cookie from userid where cookie=$1", c.Value) {
			usr := userinfo{}
			row := psql.QueryRow("select upper(name),email,password,dob,gender from userid where cookie=$1", c.Value)
			row.Scan(&usr.Namex, &usr.Emailx, &usr.Passwordx, &usr.Dobx, &usr.Genderx)
			//getting profilepic address
			i := 0
			err := filepath.Walk("./users/"+usr.Emailx+"/profilepic/", func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if i == 1 {
					usr.Profilepicx = path
				}
				i++
				return nil
			})
			if err != nil {
				log.Println(err)
			}

			var arr string
			var s int64
			err = filepath.Walk("./users/"+usr.Emailx+"/photos/",
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					arr = arr + "|" + path
					fmt.Println(path, info.Size())
					s += info.Size()
					return nil
				})
			if err != nil {
				log.Println(err)
			}

			usr.Sizex = math.Round(float64(s) / (1048576))
			// fmt.Println(s,r)
			// fmt.Println(arr)
			splarr := strings.Split(arr, "|")
			usr.Photosx = splarr[2:]
			// fmt.Println(usr.Photosx)
			tpl.ExecuteTemplate(w, "userhomepage.html", usr)
			return

		}
	}
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
func logout(w http.ResponseWriter, req *http.Request) {
	_, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c := &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
