package sayapi

import (
	"context"
	pb "pb/grpconepb"
	r "ray/common"
	)

type Server struct{}

func (s *Server) SayHi(ctx context.Context, in *pb.Say) (*pb.Response, error) {
	result := r.MapToJsStr(r.RD{in.Name: in.Msg})
	return &pb.Response{Body:result}, nil

}





