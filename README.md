# asm-package

asm-package 是服务网格独立部署的命令行打包工具.
```bash
Usage:
  asm-package [command]

Available Commands:
  bundle      打包operator镜像
  chart       打包chart-global-asm及镜像
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  images      打包chart-global-asm及镜像

Flags:
      --dockerCmd string   docker binary (default "docker")
      --helmCmd string     helm v3 binary (default "helm")
  -h, --help               help for asm-package
      --output string      output directory (default "/tmp")
      --registry string    registry (default "build-harbor.alauda.cn")
```
---
## 构建
### 本地测试
make build 
### linux-amd64
make build-linux
### linux-arm64
make build-linux-arm

---
## 样例
### bundle 打包
```bash
bin/asm-package bundle --asmBundleVersion=v3.7-13-ge53b7de --flaggerBundleVersion=v3.7-3-ga0a14d5
```

### chart打包
```bash
bin/asm-package chart --chartVersion=v3.7.0-alpha.681
```

### images打包
```bash
bin/asm-package images --chartFolder=/tmp/global-asm
```
默认会在 /tmp 目录下生成 部署包.
```bash
drwxr-xr-x  6 yuan          wheel   192B 11 18 10:42 global-asm
-rw-------  1 yuan          wheel   2.6G 11 18 11:05 asm-images.tar
-rw-------  1 yuan          wheel    72K 11 18 11:07 asm-bundle.tar
```

**注意:** 
1. 程序会调用docker和helm 命令（helm需要v3版本），如果命令没有在PATH中设置，需要通过参数明确指定。
    ```bash
    bin/asm-package chart --chartVersion=v3.7.0-alpha.681 --dockerCmd=/usr/local/bin/docker --helmCmd=/usr/local/bin/helm
    ```
2. registry 需要口令登陆，请联系ASM Group管理人员获取.
    ```bash
    docker login <registry> -u <user> -p <passwd>
    ```
3. images 打包是对chart中镜像的拉取和打包，所以执行前需要先完成chart下载。