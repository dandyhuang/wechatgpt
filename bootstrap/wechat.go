package bootstrap

import (
	"os"

	"wechatbot/handler/wechat"

	"github.com/eatmoreapple/openwechat"
	log "github.com/sirupsen/logrus"
)

func StartWebChat() {
	log.Info("Start WebChat Bot")
	bot := openwechat.DefaultBot(openwechat.Desktop)
	bot.MessageHandler = wechat.Handler
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	reloadStorage := openwechat.NewJsonFileHotReloadStorage("token.json")
	//selector
	//NORMAL("0", "正常"),
	//NEW_MSG("2", "有新消息"),
	//MOD_CONTACT("4", "有人修改了自己的昵称或你修改了别人的备注"),
	//ADD_OR_DEL_CONTACT("6", "存在删除或者新增的好友信息"),
	//ENTER_OR_LEAVE_CHAT("7", "进入或离开聊天界面");
	err := bot.HotLogin(reloadStorage)
	if err != nil {
		log.Info("hot login err:", err)
		err := os.Remove("token.json")
		if err != nil {
			log.Error("Remover token json:", err)
			return
		}

		//reloadStorage = openwechat.NewJsonFileHotReloadStorage("token.json")
		err = bot.Login()
		if err != nil {
			log.Error("login err:", err)
			return
		}
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		log.Fatal(err)
		return
	}

	friends, err := self.Friends()
	for i, friend := range friends {
		log.Println(i, friend)
	}

	groups, err := self.Groups()
	for i, group := range groups {
		log.Println(i, group)
	}

	err = bot.Block()
	if err != nil {
		log.Fatal(err)
		return
	}
}
