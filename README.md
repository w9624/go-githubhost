# go-githubhost
## 国内github访问过慢问题，项目使用Go实现，支持获取Github相关域名最快的IP，提供dns刷新相关信息
## 修改hosts文件需要sudo权限，更新hosts同时前会将hosts文件内容复制一份到hosts_tmp文件

# 运行项目
## 下载output中二进制文件运行
```shell
# darwin/linux
chmod a+x ./output/githubhost_xxx
sudo ./output/githubhost_xxx

# windows 直接运行
```

## 下载项目编译执行
```shell
git clone https://github.com/w9624/go-githubhost

make run
```

## 项目源文件编译后运行
```shell
make run
```

# 如果有用，亲，记得给个star🌟哈～