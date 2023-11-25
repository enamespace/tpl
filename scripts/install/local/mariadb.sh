#!/usr/bin/env bash

TPL_ROOT=$(dirname "${BASH_SOURCE[0]}")/../..

source ${TPL_ROOT}/scripts/install/common.sh

# 安装后打印必要的信息
function mariadb::info() {
cat << EOF
MariaDB Login: mysql -h127.0.0.1 -u${MARIADB_ADMIN_USERNAME} -p'${MARIADB_ADMIN_PASSWORD}'
EOF
}

function mariadb::install()  {
    # 1. 配置 MariaDB 10.5 apt 源
    common::sudo apt-get install software-properties-common dirmngr apt-transport-https
    common::sudo apt-key adv --fetch-keys 'https://mariadb.org/mariadb_release_signing_key.asc'
    common::sudo add-apt-repository 'deb [arch=amd64,arm64,ppc64el,s390x] https://mirrors.aliyun.com/mariadb/repo/10.5/ubuntu focal main'
    
    # 2. 安装 MariaDB 和 MariaDB 客户端
    common::sudo apt update
    common::sudo apt -y install mariadb-server

    # common::sudo systemctl enable mariadb
    # common::sudo systemctl start mariadb

    common::sudo mysqladmin -u${MARIADB_ADMIN_USERNAME} password ${MARIADB_ADMIN_PASSWORD}


    mariadb::status || return 1
    mariadb::info
    echo "install MariaDB successfully"
}

function mariadb::uninstall()
{
  set +o errexit
  common::sudo "systemctl stop mariadb"
  common::sudo "systemctl disable mariadb"
  common::sudo "apt-get -y remove mariadb-server"
  common::sudo "rm -rf /var/lib/mysql"
  set -o errexit
  echo "uninstall MariaDB successfully"
}

function mariadb::status() {
      # 查看 mariadb 运行状态，如果输出中包含 active (running) 字样说明 mariadb 成功启动。
  systemctl status mariadb |grep -q 'active' || {
    echo "mariadb failed to start, maybe not installed properly"
    return 1
  }

  mysql -u${MARIADB_ADMIN_USERNAME} -p${MARIADB_ADMIN_PASSWORD} -e quit &>/dev/null || {
    echo "can not login with root, mariadb maybe not initialized properly"
    return 1
  }
  echo "MariaDB status active"
}



if [[ "$*" =~ "mariadb::" ]]; then 
    eval $*
fi