package main

import (
  "github.com/spf13/cobra"
  "strconv"
)

var destroy bool
var discovery bool
var appName string

func main() {
  var cmdDeploy = &cobra.Command{
    Use:   "deploy [app to deploy] [# of instances]",
    Short: "Deploy multiple instances of a fleet service",
    Long: `Deploy multiple instances of a Fleet service`,
    Run: func(cmd *cobra.Command, args []string) {
      if len(args) < 2 {
        println("Incorrect amount of arguments:\nadmiral deploy [app to deploy] [# of instances]")
        return
      }
      appName = args[0]
      scale, _ := strconv.ParseInt(args[1], 10, 64)
      deployApp(appName, int(scale), discovery, destroy)
    },
  }

  var cmdScale = &cobra.Command{
    Use:   "scale [app to deploy] [# of instances]",
    Short: "Change running instances of a Fleet service",
    Long: `Change running instances of a Fleet service
        `,
    Run: func(cmd *cobra.Command, args []string) {
      if len(args) < 2 {
        println("Incorrect amount of arguments:\nadmiral scale [app to deploy] [# of instances]")
        return
      }
      appName = args[0]
      scale, _ := strconv.ParseInt(args[1], 10, 64)
      scaleApp(appName, int(scale), discovery, destroy)
    },
  }

  var rootCmd = &cobra.Command{Use: "admiral"}
  rootCmd.PersistentFlags().BoolVarP(&destroy, "destroy", "x", false, "destroy instances and submit/create new lines")
  rootCmd.PersistentFlags().BoolVarP(&discovery, "discovery", "d", false, "launch <appname>-discovery@.service files for each instance")

  rootCmd.AddCommand(cmdDeploy, cmdScale)
  rootCmd.Execute()
}
