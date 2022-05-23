package xiaomiio

type Device struct {
	Did         string `json:"did"`
	Token       string `json:"token"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	Name        string `json:"name"`
	Pid         string `json:"pid"`
	Localip     string `json:"localip"`
	Mac         string `json:"mac"`
	Ssid        string `json:"ssid"`
	Bssid       string `json:"bssid"`
	ParentID    string `json:"parent_id"`
	ParentModel string `json:"parent_model"`
	ShowMode    int    `json:"show_mode"`
	Model       string `json:"model"`
	AdminFlag   int    `json:"adminFlag"`
	ShareFlag   int    `json:"shareFlag"`
	PermitLevel int    `json:"permitLevel"`
	IsOnline    bool   `json:"isOnline"`
	Desc        string `json:"desc"`
	Prop        struct {
		Power string `json:"power"`
	} `json:"prop,omitempty"`
	UID    int64 `json:"uid"`
	PdID   int   `json:"pd_id"`
	Method []struct {
		AllowValues string `json:"allow_values"`
		Name        string `json:"name"`
	} `json:"method,omitempty"`
	Password  string      `json:"password"`
	P2PID     string      `json:"p2p_id"`
	Rssi      int         `json:"rssi"`
	FamilyID  int         `json:"family_id"`
	ResetFlag int         `json:"reset_flag"`
	Extra     interface{} `json:"extra"`
}

type RespRet struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}
