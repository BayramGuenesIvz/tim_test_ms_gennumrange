package data

type ApplConfStruct struct {
	Dummy      string
	DB_UserPwd string
	SERVICEDB  string
	PORTDB     string
}

var ApplConf ApplConfStruct

var (
	ThisCommunnicationProtokoll string
	ServiceLocationID           string
	ServiceDB                   string
	PortDB                      string
	ThisServerAddr              string
	ThisPort                    string
	//PortPresentation            string
	ThisConfLocation string
	IsDBInfoFromEnv  bool
)
