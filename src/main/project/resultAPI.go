package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"

	//"go.mongodb.org/mongo-driver/bson"

	mgo "gopkg.in/mgo.v2"

	//"gopkg.in/mgo.v2"
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
	Registation      string `json:"registation" bson:"registation"`
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
	err := db.collection.Find(bson.M{"registation": bson.M{"$eq": vars["registation"]}}).One(&profile)

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
	err := db.collection.Find(bson.M{"registation": bson.M{"$eq": vars["registation"]}}).One(&result)

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
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

var SECRET_KEY = []byte("gosecretkey")

//var client DB.Collection

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}

// PostSignIN adds a new movie to our MongoDB collection
func (db *DB) PostSignIn(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	//var user SignIn
	var dbUser Profile
	vars := mux.Vars(request)
	json.NewDecoder(request.Body).Decode(&dbUser)
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := db.collection.Find(bson.M{"email": bson.M{"$eq": vars["email"]}}).One(&dbUser)
	//collection := db.("JU_CSE").Collection("profiles")
	er := db.collection.Find(bson.M{"password": bson.M{"$eq": vars["password"]}}).One(&dbUser)
	//err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser)

	if (err == nil) && (er == nil) {
		//response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"response":"ok Password!"}`))
		return
	}
	if (err != nil) && (er != nil) {
		//response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"response":"Wrong Password!"}`))
		return
	}
	/*userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)

	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

	if passErr != nil {
		log.Println(passErr)
		response.Write([]byte(`{"response":"Wrong Password!"}`))
		return
	}
	jwtToken, err := GenerateJWT()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	response.Write([]byte(`{"token":"` + jwtToken + `"}`))*/

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
	r.HandleFunc("/v1/profile/{registation:[a-zA-Z0-9]*}",
		db.GetProfile).Methods("GET")

	r.HandleFunc("/v1/result11/{registation:[a-zA-Z0-9]*}",
		db11.GetResult).Methods("GET")

	r.HandleFunc("/v1/result12/{registation:[a-zA-Z0-9]*}",
		db12.GetResult).Methods("GET")

	r.HandleFunc("/v1/result21/{registation:[a-zA-Z0-9]*}",
		db21.GetResult).Methods("GET")

	r.HandleFunc("/v1/result22/{registation:[a-zA-Z0-9]*}",
		db22.GetResult).Methods("GET")

	r.HandleFunc("/v1/result31/{registation:[a-zA-Z0-9]*}",
		db31.GetResult).Methods("GET")

	r.HandleFunc("/v1/result32/{registation:[a-zA-Z0-9]*}",
		db32.GetResult).Methods("GET")

	r.HandleFunc("/v1/result41/{registation:[a-zA-Z0-9]*}",
		db41.GetResult).Methods("GET")

	r.HandleFunc("/v1/result42/{registation:[a-zA-Z0-9]*}",
		db42.GetResult).Methods("GET")

	r.HandleFunc("/v1/profile", db.PostProfile).Methods("POST")
	r.HandleFunc("/v1/signIn", db.PostSignIn).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
