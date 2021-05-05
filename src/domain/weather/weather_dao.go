package weather

import (
	"context"
	"encoding/json"
	"gitlab.com/chertokdmitry/surfweatherbot/src/server"
	pb "gitlab.com/chertokdmitry/surfproto/src/pb"
	"google.golang.org/grpc/grpclog"
)

func Get(spotId int64) ([]Hour, string) {
	conn := server.GetConn()
	defer conn.Close()

	client := pb.NewApiClient(conn)
	request := &pb.WeatherRequest{
		SpotId: spotId,
	}

	response, err := client.WeatherBySpotId(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	w := Weather{}
	w.SpotId = response.SpotId
	w.Title = response.Title

	data := []byte(response.Hourly)
	hours := make([]Hour, 0)
	json.Unmarshal(data, &hours)

	return hours, w.Title
}
