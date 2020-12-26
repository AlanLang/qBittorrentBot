package qbt

//Torrent holds a basic torrent object from qbittorrent
type Torrent struct {
	AddedOn       int    `json:"added_on"`
	Category      string `json:"category"`
	Completed     int64  `json:"completed"`
	CompletionOn  int64  `json:"completion_on"`
	Dlspeed       int    `json:"dlspeed"`
	Eta           int    `json:"eta"`
	ForceStart    bool   `json:"force_start"`
	Hash          string `json:"hash"`
	Name          string `json:"name"`
	NumComplete   int    `json:"num_complete"`
	NumIncomplete int    `json:"num_incomplete"`
	NumLeechs     int    `json:"num_leechs"`
	NumSeeds      int    `json:"num_seeds"`
	Priority      int    `json:"priority"`
	Progress      int    `json:"progress"`
	Ratio         int    `json:"ratio"`
	SavePath      string `json:"save_path"`
	SeqDl         bool   `json:"seq_dl"`
	Size          int64  `json:"size"`
	State         string `json:"state"`
	SuperSeeding  bool   `json:"super_seeding"`
	Upspeed       int    `json:"upspeed"`
	Uploaded      int64  `json:"uploaded"`
}

//Tracker holds a tracker object from qbittorrent
type Tracker struct {
	Msg      string `json:"msg"`
	NumPeers int    `json:"num_peers"`
	Status   string `json:"status"`
	URL      string `json:"url"`
}

//WebSeed holds a webseed object from qbittorrent
type WebSeed struct {
	URL string `json:"url"`
}

//TorrentFile holds a torrent file object from qbittorrent
type TorrentFile struct {
	IsSeed   bool   `json:"is_seed"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
	Progress int    `json:"progress"`
	Size     int    `json:"size"`
}

//Sync holds the sync response struct which contains
//the server state and a map of infohashes to Torrents
type Sync struct {
	Categories  []string `json:"categories"`
	FullUpdate  bool     `json:"full_update"`
	Rid         int      `json:"rid"`
	ServerState struct {
		ConnectionStatus  string `json:"connection_status"`
		DhtNodes          int    `json:"dht_nodes"`
		DlInfoData        int    `json:"dl_info_data"`
		DlInfoSpeed       int    `json:"dl_info_speed"`
		DlRateLimit       int    `json:"dl_rate_limit"`
		Queueing          bool   `json:"queueing"`
		RefreshInterval   int    `json:"refresh_interval"`
		UpInfoData        int    `json:"up_info_data"`
		UpInfoSpeed       int    `json:"up_info_speed"`
		UpRateLimit       int    `json:"up_rate_limit"`
		UseAltSpeedLimits bool   `json:"use_alt_speed_limits"`
	} `json:"server_state"`
	Torrents map[string]Torrent `json:"torrents"`
}

type AddForm struct {
	
}
