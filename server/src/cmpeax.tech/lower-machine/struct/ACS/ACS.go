package ACS

type AcsPassData struct {
	EmployeeName string
	EmployeeID   string
	Department   string
	Faceicon     []byte
	Bands        string
	Shoes        string
}

//0x01 人员通过检测
type Acs0x01 struct {
	Code     string
	IsPass   string
	PassData AcsPassData
	Token    string
}

//0x02 下位机登陆/心跳包
type Acs0x02 struct {
	Code   string
	Result string
	Token  string
}

type Acs0x03 struct {
	Code  string
	Token string
}

type Acs0x04 struct {
	Code              string
	Token             string
	DeviceID          string
	MachineNum        string
	CompanyName       string
	AdminPwd          string
	LocalIP           string
	SingleBandsHigher string
	SingleBandsLower  string
	DoubleBandsHigher string
	DoubleBandsLower  string
	ShoesHigher       string
	ShoesLower        string
	OutBrakeType      string
	PointLight        string
}

func NewACS0x01(token string) *Acs0x01 {
	return &Acs0x01{
		Token: token,
	}
}
func NewACS0x02(token string) *Acs0x02 {
	return &Acs0x02{
		Token: token,
	}
}
func NewACS0x03(token string) *Acs0x03 {
	return &Acs0x03{
		Token: token,
	}
}
func NewACS0x04(token string) *Acs0x04 {
	return &Acs0x04{
		Token: token,
	}
}
