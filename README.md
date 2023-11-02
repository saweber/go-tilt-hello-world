# go-tilt-hello-world

## Setup
- Install tilt - `brew install tilt` or [follow install document](https://docs.tilt.dev/install.html)
- Install kind - `brew install kind` or [follow install document](https://kind.sigs.k8s.io/docs/user/quick-start)
- Configure kubectl to point at kind (or any other local cluster)
- Run `tilt up`
- Go to `http://localhost:8001/` in a browser to reach the service
- Use port `40001` to attach a debugger to the application

## Local Debugging with IntelliJ or GoLand
Two configurations are in the .run folder.
- `Remote Debug` is for attaching a debugger to the `kubernetes.yaml` deployment
- `Remote Debug Helm` is for attaching a debugger to the helm deployment