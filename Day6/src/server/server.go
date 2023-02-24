package main

import (
	"container/list"
	"day6/database"
	"day6/logo/logorand"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
)

var (
	db      *pg.DB = nil
	is_auth        = false
	page           = 1
	in_page        = 3
	th      list.List
	mu      sync.Mutex
)

func main() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func() {
		db = database.CreateConnection(":5432", "postgres", "1234", "postgres")
		err := database.CreateTable(db)
		defer db.Close()
		if err != nil {
			fmt.Println(err)
		}
		r := mux.NewRouter()
		r.HandleFunc("/", firstHandleFunc)
		r.HandleFunc("/article/{id}", articleHandleFunc)
		r.HandleFunc("/logo", logoHandlerFunc)
		r.HandleFunc("/admin", adminHandleFunc)
		r.HandleFunc("/next", nextHandleFunc)
		r.HandleFunc("/prev", prevHandleFunc)
		go logorand.InitClient("amazing_logo.png", 300, 300)
		err = http.ListenAndServe(":8888", Throttle(100, 1, r))
		if err != nil {
			fmt.Println(err)
		}
	}()
	sig := <-sigc
	fmt.Println("Received signal:", sig)
	os.Exit(0)
}

func Throttle(rate int, per int64, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ThrottleRequest(rate, per) {
			http.Error(w, "429 Too Many Requests", http.StatusTooManyRequests)
			fmt.Println("429 Too Many Requests")
			return
		}

		h.ServeHTTP(w, r)
	})
}

func ThrottleRequest(rate int, per int64) bool {
	mu.Lock()
	defer mu.Unlock()
	th.PushBack(time.Now().Unix())
	for th.Front().Value.(int64) < time.Now().Unix()-per {
		th.Remove(th.Front())
	}
	if th.Len() > rate {
		return true
	}
	return false
}

func logoHandlerFunc(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "amazing_logo.png")
}

func firstHandleFunc(w http.ResponseWriter, r *http.Request) {
	var articles []database.Article
	for k := 0; k < in_page; k++ {
		er, add := database.GetArticle(db, (k+1)+3*(page-1))
		if er == nil {
			articles = append(articles, add)
		} else {
			break
		}
	}
	for k, art := range articles {
		if art.Title == "" {
			articles[k].Title = "empty:("
		}
	}
	t, err := template.ParseFiles("docs/index.html")
	if err != nil {
		fmt.Println("template.ParseFiles(\"index.html\")", err)
		return
	}
	data := struct {
		LogoPath    string
		Articles    []database.Article
		PreviousURL string
		NextURL     string
	}{
		LogoPath:    "/logo",
		Articles:    articles,
		PreviousURL: "/prev",
		NextURL:     "/next",
	}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println("t.Execute(w, data)", err)
		return
	}
}

func prevHandleFunc(w http.ResponseWriter, r *http.Request) {
	if page != 1 {
		page--
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func nextHandleFunc(w http.ResponseWriter, r *http.Request) {
	er, _ := database.GetArticle(db, page*3+1)
	if er == nil {
		page++
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func articleHandleFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("strconv.Atoi(idStr)", err)
		return
	}
	err, article := database.GetArticle(db, id)
	if err != nil {
		fmt.Println("database.GetArticle(db, id)", err)
		return
	}
	t, err := template.ParseFiles("docs/article.html")
	if err != nil {
		fmt.Println("template.ParseFiles(\"article.html\")", err)
		return
	}
	data := struct {
		Article database.Article
		BackURL string
	}{
		Article: article,
		BackURL: "/",
	}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println("t.Execute(w, data)", err)
		return
	}
}

func adminHandleFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Println("MethodGet")
		t, _ := template.ParseFiles("docs/login.html")
		t.Execute(w, nil)
		return
	}
	if r.Method == http.MethodPost {
		if is_auth {
			title := r.FormValue("title")
			summary := r.FormValue("summary")
			body := r.FormValue("Body")
			var er error
			fmt.Println(er, "   openai")
			newArticle := database.Article{
				Title:   title,
				Summary: summary,
				Body:    body,
			}
			err := database.InsertArticle(db, &newArticle)
			if err != nil {
				fmt.Println("database.InsertArticle(db, newArticle)", err)
				return
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
			is_auth = false
		}
		fmt.Println("MethodPost")
		username := r.FormValue("username")
		password := r.FormValue("password")
		if isAuthorized(username, password) {
			t, _ := template.ParseFiles("docs/admin.html")
			data := struct {
				RandomString string
			}{
				RandomString: logorand.RandomString(false),
			}
			t.Execute(w, data)
			is_auth = true
			return
		} else {
			r.Method = http.MethodGet
			adminHandleFunc(w, r)
		}
		return
	}
}

func isAuthorized(username, password string) bool {
	fmt.Println(username, password)
	return username == "admin" && password == "secret"
}
