package beatmap

type Beatmap struct {
	BeatmapCompact
	Accuracy      float64 `json:"accuracy"`
	AR            float64 `json:"ar"`
	BeatmapsetID  int     `json:"beatmapset_id"`
	BPM           float64 `json:"bpm"`
	Convert       bool    `json:"convert"`
	CountCircles  int     `json:"count_circles"`
	CountSliders  int     `json:"count_sliders"`
	CountSpinners int     `json:"count_spinners"`
	CS            float64 `json:"cs"`
	// DeletedAt timestamp `json:"deleted_at"`
	Drain       float64 `json:"drain"`
	HitLength   int     `json:"hit_length"`
	IsScoreable bool    `json:"is_scoreable"`
	// LastUpdated timestamp `json:"last_updated"`
	ModeInt   int    `json:"mode_int"`
	Passcount int    `json:"passcount"`
	Playcount int    `json:"playcount"`
	Ranked    int    `json:"ranked"`
	URL       string `json:"url"`
}
