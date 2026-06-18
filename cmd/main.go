package main

import (
	"log"

	"github.com/spf13/cobra"

	"mac-powermetrics-exporter/internal/config"
	"mac-powermetrics-exporter/internal/server"
)

func main() {
	cfg := config.New()

	cmd := &cobra.Command{
		Use:     "mac-powermetrics-exporter",
		Aliases: []string{"mac-exporter"},
		Short:   "A Prometheus exporter for macOS system metrics using powermetrics and vm_stat commands.",
		RunE: func(cmd *cobra.Command, args []string) error {
			srv := server.New(cfg)
			return srv.Start()
		},
	}

	cmd.Flags().StringVar(&cfg.Port, "port", cfg.Port, "Port to listen on")

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
