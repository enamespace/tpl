#!/usr/bin/env bash
# 脚本执行返回非0值退出
set -o errexit 
# 如果尝试使用未设置的变量，脚本会立即退出
set +o nounset
# 如果管道中的任何一个命令失败，整个管道的返回值将是失败，而不是默认情况下只取最后一个命令的返回值
set -o pipefail


TPL_ROOT=$(dirname "${BASH_SOURCE[0]}")/../..

ENABLE_DOCKER=true

source "${TPL_ROOT}/scripts/install/environment.sh"

function common::sudo {
    local command=("$@")
 
    "${command[@]}"
    # or echo ${LINUX_PASSWORD} | sudo -S "${command[@]}"
}
