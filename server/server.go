package server

import (
	context "context"
	"io"

	"github.com/k0kubun/pp"
)

type Server struct {
	MetrcisClient ActantTestClient
}

func NewServer(mc ActantTestClient) *Server {
	return &Server{
		MetrcisClient: mc,
	}
}

func (s *Server) SendTelemetryWithHttp(context context.Context, in *SomeData) (*BaseReply, error) {
	return &BaseReply{}, nil
}
func (s *Server) SendTelemetry(context context.Context, in *SomeData) (*BaseReply, error) {
	pp.Println("Server SendTelemetry inbound PointsCloud", in.PointsCloud)
	pp.Println("Server SendTelemetry inbound MotionFrames", in.MotionFrames)
	return &BaseReply{}, nil
}

func (s *Server) RouteTelemetry(stream ActantTest_RouteTelemetryServer) error {

	for {
		sessionData, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		pp.Println(sessionData.DataFrames)
		pp.Println(sessionData.MotionFrames)

		var points []*PointsCloud
		var mFrames []*MotionFrame
		for _, p := range sessionData.GetDataFrames() {
			points = append(points, p.GetArSessionFrame().GetPointsCloud())
		}

		mFrames = append(mFrames, sessionData.GetMotionFrames()...)
		msg := &SomeData{
			PointsCloud:  points,
			MotionFrames: mFrames,
		}
		br, err := s.MetrcisClient.SendTelemetry(stream.Context(), msg)
		if err != nil {
			return err
		}
		pp.Println(br)
	}
}
