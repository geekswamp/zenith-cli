package command

import (
	"github.com/geekswap/zenith-cli/pkg/printer"
	"github.com/geekswap/zenith-cli/pkg/utils"
	"github.com/spf13/cobra"
)

func PostRunE(_ *cobra.Command, _ []string) error {
	if err := utils.GoFmt(); err != nil {
		printer.Red("🚫 Error: %v\n", err)
		return err
	}

	return nil
}
