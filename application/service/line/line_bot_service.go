package line

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	_ "github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	_ "github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"linebot-go/global"
	"log"
	"net/http"
)

// 建立一個 map 來儲存每個用戶的 ChatSession
var userSessions = make(map[string]string)

type BotService struct {
	bot  *messaging_api.MessagingApiAPI
	blob *messaging_api.MessagingApiBlobAPI
}

func NewBotService() *BotService {
	channelToken := global.ServerConfig.LineConfig.ChannelToken
	bot, err := messaging_api.NewMessagingApiAPI(channelToken)
	if err != nil {
		log.Fatal(err)
	}
	blob, err := messaging_api.NewMessagingApiBlobAPI(channelToken)
	if err != nil {
		log.Fatal(err)
	}
	return &BotService{
		bot:  bot,
		blob: blob,
	}
}

func (b *BotService) CallbackHandler(c *gin.Context) {
	channelSecret := global.ServerConfig.LineConfig.ChannelSecret
	cb, err := webhook.ParseRequest(channelSecret, c.Request)
	log.Printf("CallbackHandler. cb: %+v", cb)
	if err != nil {
		if errors.Is(err, linebot.ErrInvalidSignature) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	for k := range cb.Events {
		event := cb.Events[k]
		switch e := event.(type) {
		//訊息事件
		case webhook.MessageEvent:

			switch message := e.Message.(type) {
			// Handle only on text message
			case webhook.TextMessageContent:
				req := message.Text
				// 取得用戶 ID
				//var uID string
				//switch source := e.Source.(type) {
				//case *webhook.UserSource:
				//	uID = source.UserId
				//case *webhook.GroupSource:
				//	uID = source.UserId
				//case *webhook.RoomSource:
				//	uID = source.UserId
				//}

				if err := b.replyText(e.ReplyToken, req); err != nil {
					log.Print(err)
				}
				continue

			// Handle only on Sticker message
			case webhook.StickerMessageContent:
				var kw string
				for _, k := range message.Keywords {
					kw = kw + "," + k
				}

				outStickerResult := fmt.Sprintf("收到貼圖訊息: %s, pkg: %s kw: %s  text: %s", message.StickerId, message.PackageId, kw, message.Text)
				if err := b.replyText(e.ReplyToken, outStickerResult); err != nil {
					log.Print(err)
				}

			// Handle only image message
			case webhook.ImageMessageContent:
				//log.Println("Got img msg ID:", message.Id)
				//
				////Get image binary from LINE server based on message ID.
				//content, err := blob.GetMessageContent(message.Id)
				//if err != nil {
				//	log.Println("Got GetMessageContent err:", err)
				//}
				//defer content.Body.Close()
				//data, err := io.ReadAll(content.Body)
				//if err != nil {
				//	log.Fatal(err)
				//}
				//ret, err := GeminiImage(data)
				//if err != nil {
				//	ret = "無法辨識圖片內容，請重新輸入:" + err.Error()
				//}
				//if err := replyText(e.ReplyToken, ret); err != nil {
				//	log.Print(err)
				//}
			// Handle only video message
			case webhook.VideoMessageContent:
				log.Println("Got video msg ID:", message.Id)

			default:
				log.Printf("Unknown message: %v", message)
			}
		//追蹤事件
		case webhook.FollowEvent:
			log.Printf("message: Got followed event")
		case webhook.PostbackEvent:
			data := e.Postback.Data
			log.Printf("Unknown message: Got postback: " + data)
		case webhook.BeaconEvent:
			log.Printf("Got beacon: " + e.Beacon.Hwid)
		}
	}
}

func (b *BotService) replyText(replyToken, text string) error {
	if _, err := b.bot.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: replyToken,
			Messages: []messaging_api.MessageInterface{
				&messaging_api.TextMessage{
					Text: text,
				},
			},
		},
	); err != nil {
		return err
	}
	return nil
}
