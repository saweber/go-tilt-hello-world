load('ext://restart_process', 'docker_build_with_restart')

# Build the binary locally, and outside docker for maximum resource utilization
local_resource('go-local-build', cmd='CGO_ENABLED=0 GOOS=linux go build -o main main.go', deps=['main.go'])

# Build container, and update whenever the binary changes
docker_build_with_restart(
  'example-go-image',
  '.',
  #entrypoint='/app/main', # use this entrypoint for no debugging
  #entrypoint='/go/bin/dlv exec --listen=:40000 --api-version=2 --headless=true --accept-multiclient --log /app/main', # use this entrypoint for debugging on startup
  entrypoint='/go/bin/dlv exec --listen=:40000 --api-version=2 --headless=true --accept-multiclient --log --continue /app/main', # use this entrypoint for debugging after startup
  dockerfile='Dockerfile',
  only=[],
  live_update=[
    sync('main', '/app/main'),
  ],
)

# Deploy container with helm, update values whenever the chart, values.yaml, or container image changes
k8s_yaml(local('helm template -f ./helm-tilt-local/values.yaml ./helm-tilt-local'))
watch_file('./helm-tilt-local')
watch_file('./helm-tilt-local/values.yaml')
k8s_resource('release-name-helm-tilt-local',port_forwards=['8001:8000','40001:40000'] )

# Deploy container with kubernetes.yaml
#k8s_yaml('kubernetes.yaml')
#k8s_resource('example-go', port_forwards=['8000:8000','40000:40000'])
