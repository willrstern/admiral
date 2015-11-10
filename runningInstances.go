package main

import (
  "bytes"
  "os/exec"
  "regexp"
  "strconv"
)

func getRunningInstances(app string) (count int,er error) {
  var out bytes.Buffer
  list := exec.Command("fleetctl", "list-units")
  grep := exec.Command("grep", "^" + app + "@")
  wc  := exec.Command("wc", "-l")
  grep.Stdin, _ = list.StdoutPipe()
  wc.Stdin, _ = grep.StdoutPipe()
  wc.Stdout = &out
  _ = wc.Start()
  _ = grep.Start()
  _ = list.Run()
  err := wc.Wait()

  d := regexp.MustCompile(`\d+`).FindString(out.String())
  current64, _ := strconv.ParseInt(d, 10, 64)

  return int(current64), err
}
