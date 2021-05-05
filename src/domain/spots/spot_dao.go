package spots

import (
	"context"
	"encoding/json"
	pb "gitlab.com/chertokdmitry/surfproto/src/pb"
	"gitlab.com/chertokdmitry/surfweatherbot/src/message"
	"gitlab.com/chertokdmitry/surfweatherbot/src/server"
	"google.golang.org/grpc/grpclog"
	"log"

	"strconv"
)

var SpotIds []int

func Get(regionId int64) []Spot {
	conn := server.GetConn()
	defer conn.Close()
	client := pb.NewApiClient(conn)

	request := &pb.SpotsRequest{
		Id: regionId,
	}

	response, err := client.GetSpotsByRegionId(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	var spots []Spot

	err = json.Unmarshal([]byte(response.Spots), &spots)
	if err != nil {
		log.Println(err)
	}

	return spots
}

func GetAll() []Spot {
	conn := server.GetConn()
	defer conn.Close()
	client := pb.NewApiClient(conn)

	request := &pb.AllSpotsRequest{
		Request: "AllSpotsRequest",
	}

	response, err := client.GetAllSpots(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	var spots []Spot

	err = json.Unmarshal([]byte(response.Spots), &spots)
	if err != nil {
		log.Println(err)
	}

	return spots
}

func GetSubcriptionList() string {
	spots := GetAll()
	var res string

	for _, spot := range spots {
		res = res + "\n" + strconv.Itoa(spot.Id) + " " + spot.Title
		SpotIds = append(SpotIds, spot.Id)
	}

	return message.SubMessage + "\n" + res
}
