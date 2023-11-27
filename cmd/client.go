/*
Copyright © 2023 Daniel Valdivia <c@dvm.sh>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/dvaldivia/grpctunnel_example/pkg/tunnel_server"
	"log"

	"github.com/jhump/grpctunnel"
	"github.com/spf13/cobra"

	pb "github.com/dvaldivia/grpctunnel_example/stubs"
	"github.com/jhump/grpctunnel/tunnelpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Starts a network client",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Client")

		// Dial the server.
		cc, err := grpc.Dial(
			"127.0.0.1:7899",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatal(err)
		}

		// Register services for reverse tunnels.
		tunnelStub := tunnelpb.NewTunnelServiceClient(cc)
		channelServer := grpctunnel.NewReverseTunnelServer(tunnelStub)

		pb.RegisterTimeTellerServer(channelServer, &tunnel_server.MyTimeTellServer{})

		log.Println("Starting Client")
		// Open the reverse tunnel and serve requests.
		if _, err := channelServer.Serve(context.Background()); err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
