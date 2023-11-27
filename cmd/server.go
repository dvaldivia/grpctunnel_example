/*
Copyright © 2023 Daniel Valdivia <c@dvm.sh>
*/
package cmd

import (
	"context"
	pb "github.com/dvaldivia/grpctunnel_example/stubs"
	"github.com/jhump/grpctunnel"
	"github.com/jhump/grpctunnel/tunnelpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the network server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("server started")

		handler := grpctunnel.NewTunnelServiceHandler(
			grpctunnel.TunnelServiceHandlerOptions{
				OnReverseTunnelOpen: func(channel grpctunnel.TunnelChannel) {
					log.Println("New Tunnel Opened")
					// let's ask some stuff
					c := pb.NewTimeTellerClient(channel)

					// Contact the server and print out its response.
					ctx, cancel := context.WithTimeout(context.Background(), time.Second)
					defer cancel()
					r, err := c.WhatTimeIsIt(ctx, &pb.TellRequest{})
					if err != nil {
						log.Fatalf("could not get time: %v", err)
					}
					log.Printf("Time Is: %s", r.GetMessage())
				},
				OnReverseTunnelClose: func(channel grpctunnel.TunnelChannel) {
					log.Println("Tunnel Closed")
				},
			},
		)
		svr := grpc.NewServer()
		tunnelpb.RegisterTunnelServiceServer(svr, handler.Service())

		// Start the gRPC server.
		l, err := net.Listen("tcp", "0.0.0.0:7899")
		if err != nil {
			log.Fatal(err)
		}
		if err := svr.Serve(l); err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
