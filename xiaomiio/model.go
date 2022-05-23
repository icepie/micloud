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
	Extra     interface{} `json:"extra,omitempty"`
}

type RespRet struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type GetDevicesReq struct {
	GetVirtualModel  bool `json:"getVirtualModel"`
	GetHuamiDevices  int  `json:"getHuamiDevices"`
	GetSplitDevice   bool `json:"get_split_device"`
	SupportSmartHome bool `json:"support_smart_home"`
}

type PropParam struct {
	Did   string      `json:"did"`
	Siid  int         `json:"siid"`
	Piid  int         `json:"piid"`
	Value interface{} `json:"value,omitempty"`
}

type ActionParam struct {
	Did  string        `json:"did"`
	Siid int           `json:"siid"`
	Aiid int           `json:"aiid"`
	In   []interface{} `json:"in"`
	Out  []interface{} `json:"out,omitempty"`
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
	Method string     `json:"method,omitempty"`
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
