package userinfo

import (
	"github.com/nwyy/wechat/mp/card/code"
	"github.com/nwyy/wechat/mp/core"
)

type CustomField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type UInfo struct {
	CustomFieldList []CustomField `json:"custom_field_list"`
	CommonFieldList []CustomField `json:"common_field_list"`
}

type UserInfo struct {
	Openid          string        `json:"openid"`
	Nickname        string        `json:"nickname"`
	Sex             string        `json:"sex"`
	Bonus             int         `json:"bonus"`
	Balance             int         `json:"balance"`
	UserIF   UInfo  `json:"user_info"`
}

// 拉取会员信息（积分查询）接口
func Get(clt *core.Client, id *code.CardItemIdentifier) (info *UserInfo, err error) {
	var result struct {
		core.Error
		UserInfo
	}

	incompleteURL := "https://api.weixin.qq.com/card/membercard/userinfo/get?access_token="
	if err = clt.PostJSON(incompleteURL, id, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.UserInfo
	return
}
