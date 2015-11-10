package main

import (
  "fmt"
  "os/exec"
)

func destroySvc(name string) {
  fmt.Printf("destroying %v.service\n", name)
  _, err := exec.Command("fleetctl", "destroy", name).Output()

  if err == nil {
    fmt.Printf("destroyed %v.service\n", name)
  }
}
