package main

import (
  "fmt"
)

func scaleDown(app string, scale int, current int, discovery bool) {
  for i := (scale + 1) ; i <= current; i++ {
    destroySvc(fmt.Sprintf("%s@%v", app, i))
    if discovery {
      destroySvc(fmt.Sprintf("%s-discovery@%v", app, i))
    }
  }
}

func scaleUp(app string, scale int, current int, discovery bool) {
  for i := (current + 1) ; i <= scale; i++ {
    startSvc(fmt.Sprintf("%s@%v", app, i), false)
    if discovery {
      startSvc(fmt.Sprintf("%s-discovery@%v", app, i), false)
    }
  }
}
