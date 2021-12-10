package recent

type Event struct {
	CreatedAt string `json:"created_at"`
	Id        int    `json:"id"`
	Type      string `json:"type"`
}

type RankEvent struct {
	CreatedAt string `json:"created_at"`
	Id        int    `json:"id"`
	Type      string `json:"type"`
	// TODO: FIX WHEN PEPPY FIXES
	ScoreRank string       `json:"scoreRank"`
	Rank      int          `json:"rank"`
	Mode      string       `json:"mode"`
	Beatmap   BeatmapEvent `json:"Beatmap"`
	User      UserEvent    `json:"User"`
}

type UserEvent struct {
	CreatedAt string `json:"created_at"`
	Id        int    `json:"id"`
	Type      string `json:"type"`
	Username  string `json:"username"`
	URL       string `json:"url"`
}

type BeatmapEvent struct {
	CreatedAt string `json:"created_at"`
	Id        int    `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	URL       string `json:"url"`
}
