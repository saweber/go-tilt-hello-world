load('ext://restart_process', 'docker_build_with_restart')

local_resource('go-local-build', cmd='CGO_ENABLED=0 GOOS=linux go build -o main main.go', deps=['main.go'])

docker_build_with_restart(
  'example-go-image',
  '.',
  entrypoint='/go/bin/dlv exec --listen=:40000 --api-version=2 --headless=true --accept-multiclient --log /app/main',
  dockerfile='Dockerfile',
  only=[],
  live_update=[
    sync('main', '/app/main'),
  ],
)

k8s_yaml('kubernetes.yaml')
k8s_resource('example-go', port_forwards=['8000:8000','40000:40000'])
