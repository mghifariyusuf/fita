package appointment

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "variant-2",
	Short: "Run variant 2",
	RunE:  run,
}

func run(cmd *cobra.Command, args []string) error {
	fmt.Println(LoadData())
	return nil
}
