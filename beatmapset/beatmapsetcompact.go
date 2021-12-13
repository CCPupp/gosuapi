package beatmapset

type BeatmapsetCompact struct {
	Artist        string `json:"artist"`
	ArtistUnicode string `json:"artist_unicode"`
	Covers        Covers `json:"covers"`
	Creator       string `json:"creator"`
	FavoriteCount int    `json:"favourite_count"`
	ID            int    `json:"id"`
	NSFW          bool   `json:"nsfw"`
	PlayCount     int    `json:"play_count"`
	PreviewURL    string `json:"preview_url"`
	Source        string `json:"source"`
	Status        string `json:"status"`
	Title         string `json:"title"`
	TitleUnicode  string `json:"title_unicode"`
	UserID        int    `json:"user_id"`
	Video         bool   `json:"video"`
}

type Covers struct {
	Cover       string `json:"cover"`
	Cover2x     string `json:"cover@2x"`
	Card        string `json:"card"`
	Card2x      string `json:"card@2x"`
	List        string `json:"list"`
	List2x      string `json:"list@2x"`
	SlimCover   string `json:"slimcover"`
	SlimCover2x string `json:"slimcover@2x"`
}

var StatusValue = map[string]string{
	"-2": "graveyard",
	"-1": "wip",
	"0":  "pending",
	"1":  "ranked",
	"2":  "approved",
	"3":  "qualified",
	"4":  "loved",
}

func GetRankedStatus(mapset BeatmapsetCompact) string {
	return StatusValue[mapset.Status]
}
