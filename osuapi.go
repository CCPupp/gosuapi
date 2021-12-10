package gosuapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/CCPupp/gosuapi/beatmap"
	"github.com/CCPupp/gosuapi/recent"
	"github.com/CCPupp/gosuapi/user"
)

var Client ClientData

var ClientToken Token
var UserToken Token

type ClientData struct {
	ID     int
	Secret string
}

//Complete
type UserRequest struct {
	Grant_type    string `json:"grant_type"`
	Client_id     int    `json:"client_id"`
	Client_secret string `json:"client_secret"`
	Redirect_uri  string `json:"redirect_uri"`
	Code          string `json:"code"`
}

//Complete
type ClientRequest struct {
	Grant_type    string `json:"grant_type"`
	Client_id     int    `json:"client_id"`
	Client_secret string `json:"client_secret"`
	Scope         string `json:"scope"`
}

//Complete
type Token struct {
	Token_type    string `json:"token_type"`
	Expires_in    int    `json:"expires_in"`
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
}

// Takes an ID or Username and returns a User struct
func GetUserById(id, mode string) user.User {
	url := "https://osu.ppy.sh/api/v2/users/" + id + "/" + mode
	var body = handleRequest(url)
	var user user.User
	jsonErr := json.Unmarshal(body, &user)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return user
}

// Takes an ID or Username and returns a list of Recent Events
func GetUserRecentById(id string) []recent.Event {
	url := "https://osu.ppy.sh/api/v2/users/" + id + "/recent_activity"
	var body = handleRequest(url)
	var events []recent.Event
	jsonErr := json.Unmarshal(body, &events)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return events
}

// Takes an ID and returns a Beatmap struct
func GetBeatmapById(id int) beatmap.Beatmap {
	idString := strconv.Itoa(id)
	url := "https://osu.ppy.sh/api/v2/beatmaps/" + idString
	var body = handleRequest(url)
	var beatmap beatmap.Beatmap
	jsonErr := json.Unmarshal(body, &beatmap)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return beatmap
}

func SetClientToken() {
	url := "https://osu.ppy.sh/oauth/token"
	var jsonStr, _ = json.Marshal(ClientRequest{
		Client_id:     Client.ID,
		Client_secret: Client.Secret,
		Grant_type:    "client_credentials",
		Scope:         "public"})
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonStr))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var token Token
	jsonErr := json.Unmarshal(body, &token)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	ClientToken = token
}

func handleRequest(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	osuClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}

	req.Header.Add("Authorization", "Bearer "+ClientToken.Access_token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, getErr := osuClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body
}
