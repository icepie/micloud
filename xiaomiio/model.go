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
	Password  string `json:"password"`
	P2PID     string `json:"p2p_id"`
	Rssi      int    `json:"rssi"`
	FamilyID  int    `json:"family_id"`
	ResetFlag int    `json:"reset_flag"`
	Extra     any    `json:"extra,omitempty"`
}

type RespRet struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}

type GetDevicesReq struct {
	GetVirtualModel  bool `json:"getVirtualModel"`
	GetHuamiDevices  int  `json:"getHuamiDevices"`
	GetSplitDevice   bool `json:"get_split_device"`
	SupportSmartHome bool `json:"support_smart_home"`
}

type PropParam struct {
	Did   string `json:"did"`
	Siid  int    `json:"siid"`
	Piid  int    `json:"piid"`
	Value any    `json:"value,omitempty"`
}

type ActionParam struct {
	Did  string `json:"did"`
	Siid int    `json:"siid"`
	Aiid int    `json:"aiid"`
	In   []any  `json:"in"`
	Out  []any  `json:"out,omitempty"`
}

type PropParams []PropParam

type PropRet struct {
	PropParam
	Code    int `json:"code"`
	ExeTime int `json:"exe_time"`
}

type PropRets []PropRet

type PropParamsReq struct {
	Params PropParams `json:"params"`
	// Method string     `json:"method,omitempty"`
}

type ActionRet struct {
	ActionParam
	Miid        int `json:"miid"`
	Code        int `json:"code"`
	ExeTime     int `json:"exe_time"`
	WithLatency int `json:"withLatency"`
}

type ActionParamReq struct {
	Param ActionParam `json:"params"`
}

type BatchDeviceDatasReq struct {
	Did   string   `json:"did"`
	Props []string `json:"props"`
}

type BatchDeviceDatasRet = map[string]map[string]any

type GetDeviceDataReq struct {
	Did       string `json:"did"`
	Uid       string `json:"uid"`
	Key       string `json:"key"`
	Type      string `json:"type"`
	TimeStart int    `json:"time_start"`
	TimeEnd   int    `json:"time_end"`
	Group     string `json:"group"`
	Limit     int    `json:"limit"`
}

type GetDeviceDataRet []struct {
	UID   string `json:"uid"`
	Did   string `json:"did"`
	Type  string `json:"type"`
	Value string `json:"value"`
	Time  int    `json:"time"`
	Key   string `json:"key"`
}

type SetDeviceDataReq struct {
	Uid   string      `json:"uid"`
	Did   string      `json:"did"`
	Time  int         `json:"time"`
	Type  string      `json:"type"`
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type GetHomeListReq struct {
	Limit      int  `json:"limit"`
	Fg         bool `json:"fg"`
	FetchShare bool `json:"fetch_share"`
}

type GetHomeListRet struct {
	Homelist []struct {
		ID                  string        `json:"id"`
		Name                string        `json:"name"`
		Bssid               string        `json:"bssid"`
		Dids                []interface{} `json:"dids"`
		Icon                string        `json:"icon"`
		Shareflag           int           `json:"shareflag"`
		PermitLevel         int           `json:"permit_level"`
		Status              int           `json:"status"`
		Background          string        `json:"background"`
		SmartRoomBackground string        `json:"smart_room_background"`
		Longitude           float64       `json:"longitude"`
		Latitude            float64       `json:"latitude"`
		CityID              int           `json:"city_id"`
		Address             string        `json:"address"`
		CreateTime          int           `json:"create_time"`
		Roomlist            []struct {
			ID         string   `json:"id"`
			Name       string   `json:"name"`
			Bssid      string   `json:"bssid"`
			Parentid   string   `json:"parentid"`
			Dids       []string `json:"dids"`
			Icon       string   `json:"icon"`
			Background string   `json:"background"`
			Shareflag  int      `json:"shareflag"`
			CreateTime int      `json:"create_time"`
		} `json:"roomlist"`
		UID int `json:"uid"`
	} `json:"homelist"`
}
