package controller

import (
	"html/template"
	"log"
	"net/http"

	"../model"
)

type User struct {
	temp *template.Template
}

func (u User) RegisterRoutes() {
	http.HandleFunc("/", u.ServeBase)
	http.HandleFunc("/user/showsessions", u.ShowSessions)
	http.HandleFunc("/userJoin", u.userJoin)

}

func (u User) ServeBase(w http.ResponseWriter, r *http.Request) {

	t := u.temp.Lookup("index.html")
	if t != nil {

		session, err := Store.Get(r, "host")
		if err != nil {
			log.Println(err)
		}
		host := session.Values["host"]
		if host == nil {
			err = t.Execute(w, struct {
				Email string
			}{""})
			if err != nil {
				log.Println(err)
			}
		} else {
			err = t.Execute(w, struct {
				Email string
			}{host.(string)})
			if err != nil {
				log.Println(err)
			}
		}

		if err != nil {
			log.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (u User) ShowSessions(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t := u.temp.Lookup("view_sessions.html")
		if t != nil {

			data, err := model.ReadSessions()
			if err != nil {
				log.Println(err)
			}

			err = t.Execute(w, data)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func (u User) userJoin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		f := r.Form

		data := model.HostType{
			Port1:   f.Get("port1"),
			Port2:   f.Get("port2"),
			Image1:  f.Get("image1"),
			Image2:  f.Get("image2"),
			Email:   f.Get("email"),
			Channel: f.Get("channel"),
		}
		log.Println(data)

		t := u.temp.Lookup("agora.html")
		if t != nil {
			err = t.Execute(w, data)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
