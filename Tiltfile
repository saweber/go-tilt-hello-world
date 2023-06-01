load('ext://restart_process', 'docker_build_with_restart')

local_resource('go-local-build', cmd='CGO_ENABLED=0 GOOS=linux go build -o main main.go ', deps=['main.go'])

docker_build_with_restart(
  'example-go-image',
  '.',
  entrypoint=['/app/main'],
  dockerfile='Dockerfile',
  only=[],
  live_update=[
    sync('main', '/app/main'),
  ],
)

k8s_yaml('kubernetes.yaml')
k8s_resource('example-go', port_forwards=8000)
