package models

type User struct {
	UserId           int    `gorm:"user_id" json:"user_id"`
	Active           int    `gorm:"active" json:"active"`
	Stop             int    `gorm:"stop" json:"stop"`
	Level            int    `gorm:"level" json:"level"`
	Username         string `gorm:"username" json:"username"`
	AuthAccount      string `gorm:"auth_account" json:"auth_account"` // 第三方賬號
	Pwd              string `gorm:"pwd" json:"pwd"`
	Email            string `gorm:"email" json:"email"`
	Realname         string `gorm:"realname" json:"realname"`
	Nickname         string `gorm:"nickname" json:"nickname"`
	Sex              int    `gorm:"sex" json:"sex"`
	Mobile           string `gorm:"mobile" json:"mobile"`
	IsAuth           int    `gorm:"is_auth" json:"is_auth"`
	Tel              string `gorm:"tel" json:"tel"`
	Teltype          int    `gorm:"teltype" json:"teltype"`
	Regtime          int    `gorm:"regtime" json:"regtime"`
	Regip            string `gorm:"regip" json:"regip"`
	Logincount       int    `gorm:"logincount" json:"logincount"`
	Loginip          string `gorm:"loginip" json:"loginip"`
	Newmsg           int    `gorm:"newmsg" json:"newmsg"`
	UserSessionTime  int    `gorm:"user_session_time" json:"user_session_time"` // 最近一次登录时间
	UserSessionPage  int    `gorm:"user_session_page" json:"user_session_page"`
	UserLastvisit    int    `gorm:"user_lastvisit" json:"user_lastvisit"`
	Allowagency      int    `gorm:"allowagency" json:"allowagency"`
	Allowmodify      int    `gorm:"allowmodify" json:"allowmodify"` // 是否允許修改
	UserTotaltime    int    `gorm:"user_totaltime" json:"user_totaltime"`
	AgencyId         int    `gorm:"agencyId" json:"agencyId"`
	Foundpwdtime     int    `gorm:"foundpwdtime" json:"foundpwdtime"`
	Editmobiletime   int    `gorm:"editmobiletime" json:"editmobiletime"`
	Invoicetype      string `gorm:"invoicetype" json:"invoicetype"` // 发票类型: 1 捐赠 2 普通发票 3 统一发票 4 對中發票 6手機載具
	Role             string `gorm:"role" json:"role"`               // company仲介公司
	RoleApprove      string `gorm:"roleApprove" json:"roleApprove"`
	EmailSetting     string `gorm:"email_setting" json:"email_setting"` // email接收設置
	IsSendEdm        int    `gorm:"is_send_edm" json:"is_send_edm"`     // 是否接收edm
	RegClient        string `gorm:"reg_client" json:"reg_client"`       // 注册终端
	CountryCode      string `gorm:"country_code" json:"country_code"`   // 國家編碼
	City             string `gorm:"city" json:"city"`                   // 城市名
	RegType          string `gorm:"reg_type" json:"reg_type"`           // 注册方式 'mobile','facebook','google'
	Signature        string `gorm:"signature" json:"signature"`
	ImageId          int    `gorm:"image_id" json:"image_id"`                       // 用户的头像id
	RegionId         int    `gorm:"region_id" json:"region_id"`                     // 用户所在地区
	LineNotifyToken  string `gorm:"line_notify_token" json:"line_notify_token"`     // line notify 推送的token
	ImToken          string `gorm:"im_token" json:"im_token"`                       // 即时通讯的token
	ImTokenUpdatedAt int    `gorm:"im_token_updated_at" json:"im_token_updated_at"` // token注册时间
}
