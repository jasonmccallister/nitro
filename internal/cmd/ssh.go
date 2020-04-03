package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/craftcms/nitro/config"
	"github.com/craftcms/nitro/internal/nitro"
)

var sshCommand = &cobra.Command{
	Use:   "ssh",
	Short: "SSH into machine",
	Run: func(cmd *cobra.Command, args []string) {
		name := config.GetString("machine", flagMachineName)

		if err := nitro.Run(
			nitro.NewMultipassRunner("multipass"),
			nitro.SSH(name),
		); err != nil {
			log.Fatal(err)
		}
	},
}
