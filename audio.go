package hiddentunes

type Audio struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	ArtistName        string `json:"artist_name"`
	AlbumName         string `json:"album_name"`
	Audio             string `json:"audio"`
	AudioDownload     string `json:"audiodownload"`
	StatsRateListened int    `json:"rate_listened_total"`
}

type APIResponse struct {
	Headers struct {
		Status       string `json:"status"`
		Code         int    `json:"code"`
		ErrorMessage string `json:"error_message"`
		Warnings     string `json:"warnings"`
		ResultsCount int    `json:"results_count"`
	} `json:"headers"`
	Results []struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		ArtistName    string `json:"artist_name"`
		AlbumName     string `json:"album_name"`
		Audio         string `json:"audio"`
		AudioDownload string `json:"audiodownload"`
		Stats         struct {
			RateListenedTotal int `json:"rate_listened_total"`
		} `json:"stats"`
		AudioDownloadAllowed bool `json:"audiodownload_allowed"`
	} `json:"results"`
}
