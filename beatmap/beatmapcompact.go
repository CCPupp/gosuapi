package beatmap

type BeatmapCompact struct {
	BeatmapsetID     int     `json:"beatmapset_id"`
	DifficultyRating float64 `json:"difficulty_rating"`
	ID               int     `json:"id"`
	// Mode             gamemode.GameMode `json:"mode"`
	Status      string `json:"status"`
	TotalLength int    `json:"total_length"`
	UserID      int    `json:"user_id"`
	Version     string `json:"version"`
	// Optional Attributes
	// WIP
}
