package organization

import (
	"github.com/spf13/cobra"

	"github.com/naiba/cloudssh/cmd/client/cmd/organization/server"
)

// ServerCmd ..
var ServerCmd *cobra.Command

func init() {
	ServerCmd = &cobra.Command{
		Use:   "server",
		Short: "manage organization servers",
	}
	ServerCmd.AddCommand(server.CreateCmd)
	ServerCmd.AddCommand(server.ListCmd)
	ServerCmd.AddCommand(server.DeleteCmd)
	ServerCmd.AddCommand(server.EditCmd)
	ServerCmd.AddCommand(server.DialCmd)
}
