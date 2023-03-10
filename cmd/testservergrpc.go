// Package cmd contains the cli command definitions for gomicroservicestarter
package cmd

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

import (
	"github.com/clarkezone/gomicroservicestarter/internal"
	"github.com/clarkezone/gomicroservicestarter/pkg/basicserver"
	"github.com/clarkezone/gomicroservicestarter/pkg/config"
	"github.com/clarkezone/gomicroservicestarter/pkg/greetingservice"
	clarkezoneLog "github.com/clarkezone/gomicroservicestarter/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// TestServerGrpcCmd is the command to start a test grpc server
type TestServerGrpcCmd struct {
	bs  *basicserver.Grpc
	mid *basicserver.PromMetricsMiddlewareGrpc
}

func newTestServerGrpcCmd(partent *cobra.Command) (*TestServerGrpcCmd, error) {
	bsGrpc := basicserver.CreateGrpc()
	tsGrpc := &TestServerGrpcCmd{
		bs: bsGrpc,
	}
	cmd := &cobra.Command{
		Use:   "testservergrpc",
		Short: "Starts a test grpc server to test logging and metrics",
		Long: `Starts a listener that will
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clarkezoneLog.Successf("gomicroservicestarter version %v,%v started in testservergrpc mode\n",
				config.VersionString, config.VersionHash)
			clarkezoneLog.Successf("Log level set to %v", internal.LogLevel)

			tsGrpc.mid = basicserver.NewPromMetricsMiddlewareGrpc("gomicroservicestarter_grpc_server")
			bsGrpc.AddUnaryInterceptor(tsGrpc.mid.MetricsUnaryInterceptor)
			clarkezoneLog.Successf("Starting grpc server on port %v", internal.Port)
			bsGrpc.StartMetrics()
			clarkezoneLog.Successf("Starting metrics on port %v", internal.MetricsPort)
			serv := bsGrpc.StartListen("")
			greetingservice.RegisterGreeterServer(serv, &greetingservice.GreetingServer{})
			return bsGrpc.WaitforInterupt()
		},
	}
	err := tsGrpc.configFlags(cmd)
	if err != nil {
		return nil, err
	}
	partent.AddCommand(cmd)
	return tsGrpc, nil
}

func (ts *TestServerGrpcCmd) configFlags(cmd *cobra.Command) error {
	m := modeValue(internal.StartupMode)

	cmd.PersistentFlags().VarP(&m, "startupmode", "", "startup mode (httpserver, grpcserver, grpcclient) (default is httpserver)")
	err := viper.BindPFlag("startupmode", cmd.PersistentFlags().Lookup(internal.StartupMode))

	if err != nil {
		return err
	}
	return nil
}
