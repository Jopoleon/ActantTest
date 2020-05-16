package main

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/Jopoleon/ActantTest/server"
)

//// Create a random number of random points
//r := rand.New(rand.NewSource(time.Now().UnixNano()))
//pointCount := int(r.Int31n(100)) + 2 // Traverse at least two points
//var points []*pb.Point
//for i := 0; i < pointCount; i++ {
//	points = append(points, randomPoint(r))
//}
//log.Printf("Traversing %d points.", len(points))
//stream, err := client.RecordRoute(ctx)
//if err != nil {
//	log.Fatalf("%v.RecordRoute(_) = _, %v", client, err)
//}
//for _, point := range points {
//	if err := stream.Send(point); err != nil {
//		log.Fatalf("%v.Send(%v) = %v", stream, point, err)
//	}
//}
//reply, err := stream.CloseAndRecv()
//if err != nil {
//	log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
//}
//log.Printf("Route summary: %v", reply)
func main() {
	mc := NewMetricsClient("localhost", "8899")

	br, err := mc.SendTelemetry(context.Background(), data)
	if err != nil {
		logrus.Error("SendTelemetry error: ", err)
		return
	}
	stream, err := mc.RouteTelemetry(context.Background())
	if err != nil {
		logrus.Error("RouteTelemetry error: ", err)
		return
	}

	if err := stream.Send(&server.SessionData{
		ResultArSession:         nil,
		VideoStartTimestamp:     123213,
		SessionStartTimestamp:   2131,
		IntrinsicMatrix:         nil,
		MotionFps:               4214,
		UserId:                  0421144212411,
		SessionId:               "",
		DeviceModel:             "",
		DeviceOs:                "",
		IsDetailFeaturesSession: false,
		OrbFrameWidth:           0,
		OrbFrameHeight:          0,
		ArFrameWidth:            0,
		ArFrameHeight:           0,
		ExtraData:               nil,
		DataFrames:              nil,
		Planes:                  nil,
		DetectedAssets:          nil,
		MotionFrames: []*server.MotionFrame{{
			AccelerationX: 123,
			AccelerationY: 555,
			AccelerationZ: 555,
			GyroX:         56,
			GyroY:         66,
			GyroZ:         67,
			MagnetometerX: 77,
			MagnetometerY: 78,
			MagnetometerZ: 88,
			Timestamp:     89,
			ClientTs:      0,
			ServerTs:      0,
			FrameIndex:    0,
		}},
		LoadWorld: nil,
		HwFrames:  nil,
	}); err != nil {
		logrus.Error("RouteTelemetry error: ", err)
		return
	}

	logrus.Info(br)
}

var data = &server.SomeData{
	PointsCloud: []*server.PointsCloud{{
		Centre: &server.VectorFloat3{
			X: 1,
			Y: 66,
			Z: 56,
		},
		Size: &server.VectorFloat3{
			X: 1,
			Y: 66,
			Z: 56,
		},
		PointsCount: 123424,
		Points: []*server.Point{{
			Id: 21324124,
			Pos: &server.VectorFloat3{
				X: 34231,
				Y: 6326,
				Z: 4234234,
			},
			Color: nil,
		}},
	},
	},
	MotionFrames: []*server.MotionFrame{{
		AccelerationX: 123,
		AccelerationY: 555,
		AccelerationZ: 555,
		GyroX:         56,
		GyroY:         66,
		GyroZ:         67,
		MagnetometerX: 77,
		MagnetometerY: 78,
		MagnetometerZ: 88,
		Timestamp:     89,
		ClientTs:      0,
		ServerTs:      0,
		FrameIndex:    0,
	}}}
