package main

import (
  "fmt"
  "os"
  "os/exec"
)

func stopSvc(name string) {
  fmt.Printf("stopping %v.service\n", name)
  _, err := exec.Command("fleetctl", "stop", name).Output()

  if err != nil {
    fmt.Printf("error stopping %v.service\n", name);
    os.Exit(1)
  }
}
