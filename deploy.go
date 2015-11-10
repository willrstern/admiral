package main

import (
  "fmt"
  "os"
)

func deployApp(app string, scale int, discovery bool, destroy bool) {
  var restartCount int

  if destroy {
    destroyBase(app, discovery);
  }

  current, err := getRunningInstances(app)

  // app is not running
  if err != nil || current < 1 {
    startInstances(app, scale, discovery)
    os.Exit(0)
  }


  // scaling down
  if current > scale {
    scaleDown(app, scale, current, discovery)
    restartCount = scale
  } else {
    restartCount = current
  }

  //rolling restart
  for i := 1; i <= restartCount; i++ {
    if destroy {
      destroySvc(fmt.Sprintf("%s@%v", app, i))
      if discovery {
        destroySvc(fmt.Sprintf("%s-discovery@%v", app, i))
      }
    }

    stopSvc(fmt.Sprintf("%s@%v", app, i))
    if discovery {
      stopSvc(fmt.Sprintf("%s-discovery@%v", app, i))
    }

    startSvc(fmt.Sprintf("%s@%v", app, i), true)
    if discovery {
      startSvc(fmt.Sprintf("%s-discovery@%v", app, i), false)
    }

  }

  //scaling up
  if scale > current {
    scaleUp(app, scale, current, discovery)
  }
}

func destroyBase(app string, discovery bool) {
  destroySvc(fmt.Sprintf("%s@", app))
  if discovery {
    destroySvc(fmt.Sprintf("%s-discovery@", app))
  }

  submitSvc(fmt.Sprintf("%s@", app))

  if discovery {
    submitSvc(fmt.Sprintf("%s-discovery@", app))
  }
}

func startInstances(app string, scale int, discovery bool) {
  println("App not running, starting all instances")
  for i := 1; i <= scale; i++ {
    startSvc(fmt.Sprintf("%s@%v", app, i), false)
    if discovery {
      startSvc(fmt.Sprintf("%s-discovery@%v", app, i), false)
    }
  }
}
