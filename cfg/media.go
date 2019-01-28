package cfg

type MediaCfgStruct struct {
	Name        string `toml:"name"`
	Type        string `toml:"type"`
	Host        string `toml:"host"`
	Port        int    `toml:"port"`
	Username    string `toml:"username"`
	Password    string `toml:"password"`
	NeedAuth    bool   `toml:"auth"`
	Uri         string `toml:"uri"`
	ContentType string `toml:"contentType"`
	CorpId      string `toml:"corpid"`
	CorpSecret  string `toml:"corpsecret"`
	AgentId     int    `toml:"agentid"`
}
