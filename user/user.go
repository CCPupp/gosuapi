package user

//Complete [Likely will need a rework to account for compact users.]
type User struct {
	AvatarURL                        string                `json:"avatar_url"`
	CountryCode                      string                `json:"country_code"`
	DefaultGroup                     string                `json:"default_group"`
	ID                               int                   `json:"id"`
	IsActive                         bool                  `json:"is_active"`
	IsBot                            bool                  `json:"is_bot"`
	IsDeleted                        bool                  `json:"is_deleted"`
	IsOnline                         bool                  `json:"is_online"`
	IsSupporter                      bool                  `json:"is_supporter"`
	LastVisit                        string                `json:"last_visit"`
	PmFriendsOnly                    bool                  `json:"pm_fiends_only"`
	ProfileColor                     string                `json:"profile_colour"`
	Username                         string                `json:"username"`
	CoverURL                         string                `json:"cover_url"`
	Discord                          string                `json:"discord"`
	HasSupported                     bool                  `json:"has_supported"`
	Interests                        string                `json:"interests"`
	JoinDate                         string                `json:"join_date"`
	Kudosus                          Kudosu                `json:"kudosu"`
	Location                         string                `json:"location"`
	MaxBlocks                        int                   `json:"max_blocks"`
	MaxFriends                       int                   `json:"max_friends"`
	Occupation                       string                `json:"occupation"`
	Playmode                         string                `json:"playmode"`
	Playstyle                        []string              `json:"playstyle"`
	PostCount                        int                   `json:"post_count"`
	ProfileOrder                     []string              `json:"profile_order"`
	Title                            string                `json:"title"`
	Twitter                          string                `json:"twitter"`
	Website                          string                `json:"website"`
	Country                          Country_info          `json:"country"`
	Cover                            Cover_info            `json:"cover"`
	AccountHistory                   []string              `json:"account_history"`
	ActiveTournamentBanner           ProfileBanner         `json:"active_tournament_banner"`
	Badges                           []Badge               `json:"badges"`
	FavoriteBeatmapsetCount          int                   `json:"favorite_beatmapset_count"`
	FollowerCount                    int                   `json:"follower_count"`
	GraveyardBeatmapsetCount         int                   `json:"graveyard_beatmapset_count"`
	Groups                           []Group               `json:"groups"`
	LovedBeatmapsetCount             int                   `json:"loved_beatmapset_count"`
	MonthlyPlaycounts                []Playcount           `json:"monthly_playcounts"`
	Page                             Page_info             `json:"page"`
	PreviousUsernames                []string              `json:"previous_usernames"`
	RankedAndApprovedBeatmapsetCount int                   `json:"ranked_and_approved_beatmapset_count"`
	ReplaysWatchedCounts             []ReplaysWatchedCount `json:"replays_watched_counts"`
	ScoresFirstCount                 int                   `json:"scores_first_count"`
	Statistics                       Statistic             `json:"statistics"`
	SupportLevel                     int                   `json:"support_level"`
	UnrankedBeatmapsetCount          int                   `json:"unranked_beatmapset_count"`
	UserAchievements                 []Achievement         `json:"user_achievements"`
	RankHistory                      RankHistory_info      `json:"rank_history"`
}

type ProfileBanner struct {
	ID           int    `json:"id"`
	TournamentID int    `json:"tournament_id"`
	Image        string `json:"image"`
}

//Complete
type Kudosu struct {
	Total     int `json:"total"`
	Available int `json:"available"`
}

//Complete
type Country_info struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

//Complete
type Cover_info struct {
	CustomURL string `json:"custom_url"`
	URL       string `json:"url"`
	ID        int    `json:"id"`
}

//Complete
type Badge struct {
	Awarded_At  string `json:"awarded_at" `
	Description string `json:"description" `
	Image_URL   string `json:"image_url"  `
	URL         string `json:"url"`
}

//Complete
type Group struct {
	ID          int    `json:"id"`
	Identifier  string `json:"identifier"`
	Name        string `json:"name"`
	ShortName   string `json:"short_name"`
	Description string `json:"description"`
	Colour      string `json:"colour"`
}

//Complete
type Playcount struct {
	StartDate string `json:"start_date"`
	Count     int    `json:"count"`
}

//Complete
type Page_info struct {
	HTML string `json:"html"`
	Raw  string `json:"raw"`
}

//Complete
type ReplaysWatchedCount struct {
	StartDate string `json:"start_date"`
	Count     int    `json:"count"`
}

//Complete
type Statistic struct {
	Level                  Level_info `json:"level"`
	Pp                     float64    `json:"pp"`
	GlobalRank             int        `json:"Global_rank"`
	RankedScore            int        `json:"ranked_score"`
	HitAccuracy            float64    `json:"hit_accuracy"`
	PlayCount              int        `json:"play_count"`
	PlayTime               int        `json:"play_time"`
	TotalScore             int        `json:"total_score"`
	TotalHits              int        `json:"total_hits"`
	MaximumCombo           int        `json:"maximum_combo"`
	ReplaysWatchedByOthers int        `json:"replays_watched_by_others"`
	IsRanked               bool       `json:"is_ranked"`
	GradeCounts            GradeCount `json:"grade_counts"`
	Rank                   Rank_info  `json:"rank_info"`
}

//Complete
type Level_info struct {
	Current  int `json:"current"`
	Progress int `json:"progress"`
}

//Complete
type GradeCount struct {
	SS  int `json:"ss"`
	SSH int `json:"ssh"`
	S   int `json:"s"`
	SH  int `json:"sh"`
	A   int `json:"a"`
}

//Complete
type Rank_info struct {
	Global  int `json:"global"`
	Country int `json:"country"`
}

//Complete
type Achievement struct {
	AchievedAt    string `json:"achieved_at"`
	AchievementID int    `json:"achievement_id"`
}

//Complete
type RankHistory_info struct {
	Mode string `json:"mode"`
	Data []int  `json:"data"`
}
