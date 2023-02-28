package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/mv-kan/nust/console"
	"github.com/mv-kan/nust/core"
	"github.com/spf13/cobra"
)

var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "undo a nust task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		makeargs, _ := cmd.Flags().GetString("makeargs")
		nocolor, _ := cmd.Flags().GetBool("no-color")
		nomoretries, _ := cmd.Flags().GetBool("no-more-tries")

		if nocolor {
			color.NoColor = true // disables colorized output
		}

		i := 0
		for {
			var err error
			if force {
				err = core.UndoTaskForce(args[0], makeargs)
			} else {
				err = core.UndoTask(args[0], makeargs)
			}
			if err != nil {
				console.Danger(fmt.Sprintf("(nust try number %d): %v\n", i, err))
				i++
				if i >= 10 || nomoretries {
					os.Exit(1)
					break
				}
			} else {
				break
			}
		}
	},
}

func init() {
	undoCmd.PersistentFlags().StringP("makeargs", "m", "", "pass flags and values to makefile")
	undoCmd.PersistentFlags().BoolP("force", "f", false, "run without checks in the exec info json file")
	undoCmd.PersistentFlags().Bool("no-more-tries", false, "no more tries for this undo task command")

	rootCmd.AddCommand(undoCmd)
}
