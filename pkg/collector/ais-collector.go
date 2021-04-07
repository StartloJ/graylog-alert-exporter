package collector

type Alerter struct {
	Module      string
	AmoName     string
	AlarmName   string
	description string
	node        string
	mcZone      string
	systemName  string
	emsName     string
	emsIp       string
	siteCode    string
	region      string
	severity    string
	nodeIp      string
	networkType string
}
