package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type succeed struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

var (
	responseSuccess = succeed{Success: true}
	responseFailure = succeed{Success: false}
)

/* GetUser returns user info */
func GetUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logUrl(r)

	// get form value
	user_id, _ := strconv.Atoi(r.FormValue("user_id"))
	login_name := r.FormValue("login_name")

	user := User{}
	var err error

	// get user
	if user_id != 0 {
		user, err = getUser(user_id)
	} else if login_name != "" {
		user, err = getUserWithLoginName(login_name)
	} else {
		respondError(w, "ユーザが見つかりませんでした")
		return
	}

	if err != nil {
		respondError(w, "ユーザの取得に失敗しました")
		return
	}

	// return json
	respondJson(w, user)
}

/* CreateUser creates user */
func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logUrl(r)

	// get form value
	login_name := r.FormValue("login_name")

	if login_name == "" {
		respondError(w, "ログイン名が必要です")
		return
	}

	if m, _ := regexp.MatchString(`^[0-9a-zA-Z-]*$`, login_name); !m {
		respondError(w, "ログイン名には英数字、ハイフンのみ利用できます")
		return
	}

	// create user
	user, err := createUser(login_name)
	if err != nil {
		respondError(w, "ユーザの作成に失敗しました")
		return
	}

	respondJson(w, user)
}

/* KadaiIndex returns kadais of the specified user that have not been done */
func KadaiIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logUrl(r)

	// get form value
	user_id, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil || user_id == 0 {
		respondError(w, "課題取得中にエラーが発生しました")
		return
	}

	// get kadais
	kadais, err := kadaiIndex(user_id)
	if err != nil {
		respondError(w, "課題取得中にエラーが発生しました")
		return
	}

	respondJson(w, kadais)
}

/* UpdateKadai inserts kadai information into db */
func CreateKadai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logUrl(r)

	// get form value
	user_id, err := strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		respondError(w, "課題作成中にエラーが発生しました")
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	if title == "" || content == "" {
		respondError(w, "タイトル、内容を埋めてください")
		return
	}

	draft := r.FormValue("draft")

	kadai, err := createKadai(user_id, title, content, draft)
	if err != nil {
		respondError(w, "課題作成中にエラーが発生しました")
		return
	}

	respondJson(w, kadai)
}

/* UpdateKadai updates kadai informations in db */
func UpdateKadai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logUrl(r)

	kadai_id, err := strconv.Atoi(r.FormValue("kadai_id"))
	if err != nil {
		respondError(w, "課題の更新中にエラーが発生しました")
		return
	}

	updateData := map[string]string{}

	title := r.FormValue("title")
	if title != "" {
		updateData["title"] = title
	}

	content := r.FormValue("content")
	if content != "" {
		updateData["content"] = content
	}

	draft := r.FormValue("draft")
	if draft != "" {
		updateData["draft"] = draft
	}

	if len(updateData) == 0 {
		respondError(w, "最低一つのフィールドを更新する必要があります")
		return
	}

	kadai, err := updateKadai(kadai_id, updateData)
	if err != nil {
		respondError(w, "課題の更新中にエラーが発生しました")
		return
	}

	respondJson(w, kadai)
}

/* KadaiDone updates `done` field of specified kadai */
func KadaiDone(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logUrl(r)

	kadai_id, err := strconv.Atoi(r.FormValue("kadai_id"))
	if err != nil {
		respondError(w, "処理中にエラーが発生しました")
		return
	}

	err = kadaiDone(kadai_id)
	if err != nil {
		respondError(w, "データの更新に失敗しました")
		return
	}

	respondJson(w, responseSuccess)
}

type ErrorJson struct {
	Error string `json:"error"`
}

func respondError(w http.ResponseWriter, err interface{}) {
	data := ErrorJson{err.(string)}

	respondJson(w, data)
}

func respondJson(w http.ResponseWriter, data interface{}) {
	log.Printf("response: %v", data)

	// return json
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(data)
}

func logUrl(r *http.Request) {
	r.ParseForm()
	log.Printf("access: %v", r.URL)
	log.Printf("form: %v", r.Form)
}
