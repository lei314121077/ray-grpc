package saymsg

import (
	"log"
	"golang.org/x/net/context"
	pb "pb/grpconepb"

)

//@name SendMsg
//@param pb
//@paran content
//@param name
//@return string
func SendMsg(c pb.GrpconepbClient, ctx context.Context, name string) string{
	r, err := c.SayHi(ctx, &pb.Say{Name: name, Msg:"hello world!"})
	if err != nil {
		log.Fatalf("未响应!: %v", err)
	}
	log.Printf("Greeting: %s", r.Body)
	return r.Body
}

