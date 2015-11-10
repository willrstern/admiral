package main

import (
  "bytes"
  "fmt"
  "os"
  "os/exec"
  "time"
)

func startSvc(name string, waitForStart bool) {
  fmt.Printf("starting %v.service", name)
  _, err := exec.Command("fleetctl", "start", name).Output()


  if err != nil {
    fmt.Printf("error starting %v.service\n", name);
    os.Exit(1)
  }

  for !instanceIsRunning(name) {
    print(".")
    time.Sleep(500 * time.Millisecond)
  }

  print("\n")
}

func instanceIsRunning(name string)(status bool) {
  var out bytes.Buffer
  list := exec.Command("fleetctl", "list-units")
  grep := exec.Command("grep", "^" + name)
  grep2  := exec.Command("grep", "active")
  grep.Stdin, _ = list.StdoutPipe()
  grep2.Stdin, _ = grep.StdoutPipe()
  grep2.Stdout = &out
  _ = grep2.Start()
  _ = grep.Start()
  _ = list.Run()
  err := grep2.Wait()

  if err != nil {
    return false
  }
  return true
}
