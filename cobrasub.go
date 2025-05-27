package cobrasub

import "github.com/spf13/cobra"

func MyCommand(token string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "mycommand",
		Short: "A command imported from another repository",
		Long: "This command is a part of a larger application and is imported from another repository.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("Running mycommand with token:", token)
		},
	}

	return cmd
}
