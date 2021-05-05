package cameras

import (
	"context"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	pb "gitlab.com/chertokdmitry/surfproto/src/pb"
	"gitlab.com/chertokdmitry/surfweatherbot/src/message"
	"gitlab.com/chertokdmitry/surfweatherbot/src/server"
	"gitlab.com/chertokdmitry/surfweatherbot/src/tg"
	"gitlab.com/chertokdmitry/surfweatherbot/src/utils/logger"
	"google.golang.org/grpc/grpclog"
	"log"
)

func GetBySpot(id int64) []int64 {
	conn := server.GetConn()
	defer conn.Close()
	client := pb.NewApiClient(conn)

	request := &pb.CamerasRequest{
		SpotId: id,
	}

	response, err := client.GetCamerasBySpotId(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	var cameras []int64

	err = json.Unmarshal([]byte(response.Cameras), &cameras)
	if err != nil {
		log.Println(err)
	}

	return cameras
}

func SendImage(id string) {
	url := message.ImageUrl + id + ".jpg"
	msgFile := tgbotapi.NewPhotoUpload(tg.GetChatId(), url)

	if _, err := tg.GetBot().Send(msgFile); err != nil {
		logger.Error(message.ErrSendInline, err)
	}
}
