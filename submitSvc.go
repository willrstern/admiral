package main

import (
  "fmt"
  "os"
  "os/exec"
)

func submitSvc(name string) {
  fmt.Printf("submitting %v.service\n", name)
  _, err := exec.Command("fleetctl", "submit", name).Output()

  if err == nil {
    fmt.Printf("submitted %v.service\n", name)
  } else {
    fmt.Printf("error submitting %v.service\n", name);
    os.Exit(1)
  }
}
