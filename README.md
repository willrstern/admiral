# Admiral
A wrapper for coreos fleet - allowing scaling and rolling deploys of apps

While Fleetctl is excellent, it doesn't provide the ability to command all instances of an app.  Until (hopefully) that day comes, `admiral` provides the ability to do zero-downtime rolling restarts a.k.a. deployments, scaling, & even container rebalancing with the `--destroy` option

## Installation
```
wget -O /usr/local/bin/admiral https://github.com/willrstern/admiral/releases/download/0.0.3/admiral
chmod +x /usr/local/bin/admiral
```

## Usage
`admiral deploy someapp 4`
will submit `someapp@` service file, and start/rolling-restart 4 instances of it

`admiral scale someapp 5`
will start an additional instance of `someapp@`

`admiral scale someapp 3`
will destroy `someapp@4` and `someapp@5`

options:
- `--discovery`, `-d` also submits `someapp-discovery@` and runs 4 instances of it
- `--destroy`/`-x` destroys each service before restarting it 
  - use for container rebalancing when coreos nodes have been added
  - use to resubmit service/service-discovery systemd file when changes have been made

