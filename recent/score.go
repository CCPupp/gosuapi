package recent

import (
	"os/user"

	"github.com/CCPupp/gosuapi/beatmap"
	"github.com/CCPupp/gosuapi/beatmapset"
)

type Score struct {
	Id         int                          `json:"id"`
	BestId     int                          `json:"best_id"`
	UserId     int                          `json:"user_id"`
	Accuracy   float64                      `json:"accuracy"`
	Mods       []string                     `json:"mods"`
	Score      int                          `json:"score"`
	MaxCombo   int                          `json:"max_combo"`
	Perfect    bool                         `json:"perfect"`
	Passed     bool                         `json:"passed"`
	Pp         float64                      `json:"pp"`
	Rank       string                       `json:"rank"`
	CreatedAt  string                       `json:"created_at"`
	Mode       string                       `json:"mode"`
	ModeInt    int                          `json:"mode_int"`
	Replay     bool                         `json:"replay"`
	Beatmap    beatmap.Beatmap              `json:"beatmap"`
	BeatmapSet beatmapset.BeatmapsetCompact `json:"beatmapset"`
	Weight     int                          `json:"weight"`
	User       user.User                    `json:"user"`
}
