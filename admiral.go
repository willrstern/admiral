package main

import (
  "github.com/spf13/cobra"
)

var destroy bool
var discovery bool
var appName string
var scale int

func main() {
  var cmdDeploy = &cobra.Command{
    Use:   "deploy [app to deploy]",
    Short: "Deploy multiple instances of a fleet service",
    Long: `Deploy multiple instances of a Fleet service`,
    Run: func(cmd *cobra.Command, args []string) {
      appName = args[0]
      deployApp(appName, scale, discovery, destroy)
    },
  }

  var cmdScale = &cobra.Command{
    Use:   "scale [app to deploy]",
    Short: "Change running instances of a Fleet service",
    Long: `Change running instances of a Fleet service
        `,
    Run: func(cmd *cobra.Command, args []string) {
      appName = args[0]
      scaleApp(appName, scale, discovery, destroy)
    },
  }

  var rootCmd = &cobra.Command{Use: "admiral"}
  rootCmd.PersistentFlags().IntVarP(&scale, "scale", "s", 1, "number of instances to run")
  rootCmd.PersistentFlags().BoolVarP(&destroy, "destroy", "d", false, "destroy instances and submit/create new lines")
  rootCmd.PersistentFlags().BoolVarP(&discovery, "discovery", "v", false, "launch <appname>-discovery@.service files for each instance")

  rootCmd.AddCommand(cmdDeploy, cmdScale)
  rootCmd.Execute()
}
