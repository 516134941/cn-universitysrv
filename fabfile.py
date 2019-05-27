# -*- coding: utf-8 -*-
"""fabric 部署脚本"""

import os
import datetime
from fabric.api import env, local, run, put, hosts

GOPATH = os.environ.get("GOPATH", None)
if GOPATH is None:
    raise Exception("GOPATH unset!")

# 项目名称
PROJECT_NAME = "cn-universitysrv"
LOCAL_FILE_NAME = "/tmp/%s" % PROJECT_NAME
# 当天日期
CURRENT_DAY_STRING = datetime.datetime.today().strftime("%Y%m%d%H%M")
# 远程服务器登陆使用的用户名及IP地址
env.key_filename = '/Users/hcf/pem/macpro.pem'
env.user = 'root'

def build():
    """build"""
    local("rm -rf %s" % LOCAL_FILE_NAME)
    command = "GOOS=linux go build -o %s" % LOCAL_FILE_NAME
    local(command, capture=False)

@hosts("47.106.232.79:22")
def test():
    """测试环境"""

    # 远程路径：/home/tonnn/connectsrv/connetsrv_201609190306
    remote_run_name = os.path.join("/home", PROJECT_NAME, PROJECT_NAME)
    remote_run_backup_name = os.path.join("/home", PROJECT_NAME, "%s.backup" % PROJECT_NAME)
    remote_backup_name = os.path.join("/home", PROJECT_NAME, "backup", \
        "%s_%s" % (PROJECT_NAME, CURRENT_DAY_STRING))

    # 创建远程文件目录
    back_up = os.path.join("/home", PROJECT_NAME, "backup")
    run('mkdir -p %s' % back_up)
    # 将代码归档上传到服务器当中的临时文件夹内
    put(LOCAL_FILE_NAME, remote_backup_name)
    run("chmod +x %s" % remote_backup_name)

    # 备份当前代码
    run("cp -rf %s %s" % (remote_run_name, remote_run_backup_name))
    run("cp -rf %s %s" % (remote_backup_name, remote_run_name))

    # 重启服务
    run("sudo supervisorctl restart cn-universitysrv")


