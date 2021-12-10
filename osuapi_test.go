package gosuapi_test

import (
	"fmt"
	"testing"

	"github.com/CCPupp/gosuapi"
)

func setup() {
	gosuapi.Client.ID = 1234         // SET YOUR ID
	gosuapi.Client.Secret = "Secret" // SET YOUR SECRET
}
func TestCreateClient(t *testing.T) {
	setup()
	gosuapi.SetClientToken()
	if gosuapi.ClientToken.Access_token == "" {
		t.Error("Failed to set client token.")
	}
}

func TestGetUserById(t *testing.T) {
	setup()
	gosuapi.SetClientToken()

	user := gosuapi.GetUserById("2", "osu")
	if user.Username != "peppy" {
		t.Error("Either peppy namechanged or something went wrong.")
	}
}

func TestGetBeatmapById(t *testing.T) {
	setup()
	gosuapi.SetClientToken()

	beatmap := gosuapi.GetBeatmapById(75)
	fmt.Println(beatmap)
	if beatmap.URL != "https://osu.ppy.sh/beatmaps/75" {
		t.Error("Either peppy deleted DISCO PRINCE or something went wrong.")
	}
}
