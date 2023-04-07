#!/bin/bash
set -euo pipefail

app="admin-api"
build_time=$(date +"%Y-%m-%d %H:%M:%S")
current_path=$(
  cd "$(dirname "$0")" || exit
  pwd
)
docker_image_name="github.com.account.test/${app}"
docker_image_tag="${docker_image_name}:$(date +%Y%m%d%H%M)"

function go_build() {
  green "==> build ${app}..."

  export GOPROXY="https://goproxy.cn,direct"

  flags="-w -s"

  set -x
  go build -v -o ${app} -ldflags "${flags}"
  set +x
}

function docker_build() {
  green "==> build docker image..."

  cd "${current_path}" || exit
  docker build -t "${docker_image_tag}" .
}

function docker_push() {
  green "==> push docker image..."

  cd "${current_path}" || exit
  docker login -p password -u username repository_name # TODO replace password、username、repository_name
  docker push "${docker_image_tag}"
}

function build_success() {
  green "==> build success"
  echo "${docker_image_tag}"
}

function green() {
  echo -e "\033[32m${1}\033[0m"
}

function red() {
  echo -e "\033[31m${1}\033[0m"
}

function main() {
  go_build
  docker_build
  docker_push
  build_success
}

main
