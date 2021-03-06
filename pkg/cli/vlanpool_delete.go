package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// vpoolDelCmd deletes a vlan pool.
var (
	vpoolDelCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a vlan pool",
		Long:  `Delete a vlan pool`,
		Run:   runVpoolDelCmd,
	}
)

func init() {
	vpoolCmd.AddCommand(vpoolDelCmd)
	vpoolDelCmd.Flags().StringVar(&aciFlags.name, "name", "", "Name of vlan pool")
	vpoolDelCmd.Flags().StringVar(&aciFlags.vlanMode, "mode", "", "Mode of vlan pool")
	vpoolDelCmd.MarkFlagRequired("name")
	vpoolDelCmd.MarkFlagRequired("mode")
}

func runVpoolDelCmd(cmd *cobra.Command, args []string) {
	if len(aciFlags.name) == 0 {
		cmd.Help()
		return
	}
	if len(aciFlags.vlanMode) == 0 {
		cmd.Help()
		return
	}
	if err := validateArgs(cmd, args); err != nil {
		return
	}

	client := mustClientFromCmd(cmd)

	err := client.Login()
	if err != nil {
		exitWithError(ExitError, err)
	}

	err = client.VlanPoolDel(aciFlags.name, aciFlags.vlanMode)
	if err != nil {
		exitWithError(ExitError, err)
	}

	fmt.Printf("vlan pool %s deleted.\n", aciFlags.name)

	err = client.Logout()
	if err != nil {
		exitWithError(ExitError, err)
	}
}
