package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DB stores the database session imformation. Needs to be initialized once
type DB struct {
	session    *mgo.Session
	collection *mgo.Collection
}

type Profile struct {
	//ID               string `json:"id" bson:"_id"`
	Full_Name        string `json:"full_name" bson:"full_name"`
	User_Name        string `json:"user" bson:"user"`
	Registation      string `json:"registration" bson:"registration"`
	Exam_Roll        string `json:"roll" bson:"roll"`
	Email            string `json:"email" bson:"email"`
	Password         string `json:"password" bson:"password"`
	Confirm_Password string `json:"confirm_password" bson:"confirm_password"`
}

type Result struct {
	Registation string `json:"registation" bson:"registation"`
	Res         string `json:"result" bson:"result"`
}

// GetProfile fetches a Profile with a given registration
func (db *DB) GetProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	var profile Profile
	err := db.collection.Find(bson.M{"registration": bson.M{"$eq": vars["registration"]}}).One(&profile)

	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(profile)
		w.Write(response)
	}
}

func (db *DB) GetResult(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	var result Result
	err := db.collection.Find(bson.M{"registration": bson.M{"$eq": vars["registration"]}}).One(&result)

	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(result)
		w.Write(response)
	}
}

// PostProfile adds a new Profile to our MongoDB collection
func (db *DB) PostProfile(w http.ResponseWriter, r *http.Request) {
	var profile Profile
	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &profile)

	err := db.collection.Insert(profile)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(profile)
		w.Write(response)
	}
}

type SignIn struct {
	Email string `json:"email" bson:"email"`
	//Password string `json:"password" bson:"password"`
}

// PostSignIN adds a new movie to our MongoDB collection
func (db *DB) PostSignIn(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var signin SignIn
	var profile Profile
	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &profile)

	er := db.collection.Find(bson.M{"email": bson.M{"$eq": vars["email"]}}).One(&signin)
	//err := db.collection.Find(bson.M{"password": bson.M{"$eq": vars["password"]}}).One(&signin)

	if er != nil {
		w.Write([]byte(er.Error()))
	} else {

		//w.Header().Set("Content-Type", "application/json")

		//response, _ := json.Marshal(profile)
		//w.WriteHeader(http.StatusOK)
		//w.Write(response)
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(profile)
		w.Write(response)
	}
}

func main() {

	session, err := mgo.Dial("127.0.0.1")
	c := session.DB("JU_CSE").C("profiles")

	db := &DB{session: session, collection: c}
	if err != nil {
		panic(err)
	}
	defer session.Close()

	r11 := session.DB("JU_CSE").C("Result11")
	db11 := &DB{session: session, collection: r11}
	if err != nil {
		panic(err)
	}

	r12 := session.DB("JU_CSE").C("Result12")
	db12 := &DB{session: session, collection: r12}
	if err != nil {
		panic(err)
	}

	r21 := session.DB("JU_CSE").C("Result21")
	db21 := &DB{session: session, collection: r21}
	if err != nil {
		panic(err)
	}

	r22 := session.DB("JU_CSE").C("Result22")
	db22 := &DB{session: session, collection: r22}
	if err != nil {
		panic(err)
	}

	r31 := session.DB("JU_CSE").C("Result31")
	db31 := &DB{session: session, collection: r31}
	if err != nil {
		panic(err)
	}

	r32 := session.DB("JU_CSE").C("Result32")
	db32 := &DB{session: session, collection: r32}
	if err != nil {
		panic(err)
	}

	r41 := session.DB("JU_CSE").C("Result41")
	db41 := &DB{session: session, collection: r41}
	if err != nil {
		panic(err)
	}

	r42 := session.DB("JU_CSE").C("Result42")
	db42 := &DB{session: session, collection: r42}
	if err != nil {
		panic(err)
	}

	defer session.Close()
	// Create a new router
	r := mux.NewRouter()
	// Attach an elegant path with handler
	r.HandleFunc("/v1/profile/{registration:[a-zA-Z0-9]*}",
		db.GetProfile).Methods("GET")

	r.HandleFunc("/v1/result11/{registration:[a-zA-Z0-9]*}",
		db11.GetResult).Methods("GET")

	r.HandleFunc("/v1/result12/{registration:[a-zA-Z0-9]*}",
		db12.GetResult).Methods("GET")

	r.HandleFunc("/v1/result21/{registration:[a-zA-Z0-9]*}",
		db21.GetResult).Methods("GET")

	r.HandleFunc("/v1/result22/{registration:[a-zA-Z0-9]*}",
		db22.GetResult).Methods("GET")

	r.HandleFunc("/v1/result31/{registration:[a-zA-Z0-9]*}",
		db31.GetResult).Methods("GET")

	r.HandleFunc("/v1/result32/{registration:[a-zA-Z0-9]*}",
		db32.GetResult).Methods("GET")

	r.HandleFunc("/v1/result41/{registration:[a-zA-Z0-9]*}",
		db41.GetResult).Methods("GET")

	r.HandleFunc("/v1/result42/{registration:[a-zA-Z0-9]*}",
		db42.GetResult).Methods("GET")

	r.HandleFunc("/v1/profile", db.PostProfile).Methods("POST")
	r.HandleFunc("/v1/signIn", db.PostSignIn).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    "10.10.10.101:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
