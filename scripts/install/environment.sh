#!/usr/bin/env bash

readonly DATABASE="tpl"
readonly USERNAME="tpl"
readonly PASSWORD="1234"


# Mariadb 配置
readonly MARIADB_ADMIN_USERNAME=${MARIADB_ADMIN_USERNAME:-root} 
readonly MARIADB_ADMIN_PASSWORD=${MARIADB_ADMIN_PASSWORD:-${PASSWORD}} 
readonly MARIADB_HOST=${MARIADB_HOST:-127.0.0.1:3306}
readonly MARIADB_DATABASE=${MARIADB_DATABASE:-${DATABASE}}
readonly MARIADB_USERNAME=${MARIADB_USERNAME:-${USERNAME}} 
readonly MARIADB_PASSWORD=${MARIADB_PASSWORD:-${PASSWORD}} 


# Redis 配置信息
readonly REDIS_HOST=${REDIS_HOST:-127.0.0.1} # Redis 主机地址
readonly REDIS_PORT=${REDIS_PORT:-6379} # Redis 监听端口
readonly REDIS_USERNAME=${REDIS_USERNAME:-''} # Redis 用户名
readonly REDIS_PASSWORD=${REDIS_PASSWORD:-${PASSWORD}} # Redis 密码

# MongoDB 配置
readonly MONGO_ADMIN_USERNAME=${MONGO_ADMIN_USERNAME:-root} # MongoDB root 用户
readonly MONGO_ADMIN_PASSWORD=${MONGO_ADMIN_PASSWORD:-${PASSWORD}} # MongoDB root 用户密码
readonly MONGO_HOST=${MONGO_HOST:-127.0.0.1} # MongoDB 地址
readonly MONGO_PORT=${MONGO_PORT:-27017} # MongoDB 端口
readonly MONGO_USERNAME=${MONGO_USERNAME:-${USERNAME}} # MongoDB 用户名
readonly MONGO_PASSWORD=${MONGO_PASSWORD:-${PASSWORD}} # MongoDB 密码