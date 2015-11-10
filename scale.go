package main

func scaleApp(app string, scale int, discovery bool, destroy bool) {

  current, err := getRunningInstances(app)

  if err != nil {
    println(err)
  } else {
    // scaling down
    if current > scale {
      scaleDown(app, scale, current, discovery)
    }

    //rolling restart

    //scaling up
    if scale > current {
      scaleUp(app, scale, current, discovery)
    }
  }
}
