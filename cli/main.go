package main

import (
	"log"

	"github.com/mghifariyusuf/fita/cli/appointment"
	"github.com/mghifariyusuf/fita/cli/checkout"
	"github.com/spf13/cobra"
)

var (
	mainCmd = &cobra.Command{
		Use: "run",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	mainCmd.AddCommand(checkout.Cmd, appointment.Cmd)
	if err := mainCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
