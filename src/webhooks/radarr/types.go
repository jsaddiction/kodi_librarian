package radarr

type TestEvent struct {
	ApplicationUrl string `json:"applicationUrl"`
	EventType      string `json:"eventType"`
	InstanceName   string `json:"instanceName"`
	Movie          struct {
		FolderPath  string   `json:"folderPath"`
		ID          int      `json:"id"`
		ReleaseDate string   `json:"releaseDate"`
		Tags        []string `json:"tags"`
		Title       string   `json:"title"`
		TmdbID      int      `json:"tmdbId"`
		Year        int      `json:"year"`
	} `json:"movie"`
	Release struct {
		CustomFormatScore int    `json:"customFormatScore"`
		Indexer           string `json:"indexer"`
		Quality           string `json:"quality"`
		QualityVersion    int    `json:"qualityVersion"`
		ReleaseGroup      string `json:"releaseGroup"`
		ReleaseTitle      string `json:"releaseTitle"`
		Size              int    `json:"size"`
	} `json:"release"`
	RemoteMovie struct {
		ImdbID string `json:"imdbId"`
		Title  string `json:"title"`
		TmdbID int    `json:"tmdbId"`
		Year   int    `json:"year"`
	} `json:"remoteMovie"`
}
