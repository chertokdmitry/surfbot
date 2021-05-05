package subscription

import (
	"context"
	"encoding/json"
	pb "gitlab.com/chertokdmitry/surfproto/src/pb"
	"gitlab.com/chertokdmitry/surfweatherbot/src/message"
	"gitlab.com/chertokdmitry/surfweatherbot/src/server"
	"gitlab.com/chertokdmitry/surfweatherbot/src/tg"
	"google.golang.org/grpc/grpclog"
	"log"
)

// insert new sub to postgres
func Insert(hour string) {
	sub := &Sub{0, tg.GetChatId(), int64(SpotId), "", hour}
	jsonReq, err := json.Marshal(sub)

	conn := server.GetConn()
	defer conn.Close()
	client := pb.NewApiClient(conn)

	request := &pb.CreateSubscriptionRequest{
		Sub: string(jsonReq),
	}

	_, err = client.CreateSubscription(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
}

// delete sub
func Delete(subId int64) {
	conn := server.GetConn()
	defer conn.Close()
	client := pb.NewApiClient(conn)

	request := &pb.DeleteSubRequest{
		SubId: subId,
	}

	_, err := client.DeleteSub(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
}

// get list of current subs
func List() []Sub {
	conn := server.GetConn()
	defer conn.Close()
	client := pb.NewApiClient(conn)

	request := &pb.SubsRequest{
		ChatId: tg.GetChatId(),
	}

	response, err := client.GetSubsByChatId(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	var subs []Sub

	err = json.Unmarshal([]byte(response.Subs), &subs)
	if err != nil {
		log.Println(err)
	}

	return subs
}

func GetSpotId() {
	IsSpotId = true
	IsHour = false

	message.Send(message.GetSpot)
}

func GetHour() {
	IsSpotId = false
	IsHour = true
}

func VarsFlash() {
	IsSpotId = false
	IsHour = false
}
