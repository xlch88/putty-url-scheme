# PuTTY URL Scheme 辅助程序
使用 url scheme 打开 putty。

# 对于中国用户的 *温 馨 提 示*
如果 -register 无效，请关闭傻宝 360 等国产傻宝软件。

# 安装
1. 在 release 下载二进制文件放到 PuTTY 同目录下。
2. 运行 `putty-url-scheme.exe --register` 进行注册表注册。
3. 打开你的浏览器，在地址栏输入 `ssh://username:password@hostname:port` 后按下回车进行测试。

# URL Scheme
ssh://username:password@hostname:port

## 使用代理
ssh://username:password@hostname:port/?proxyHost=`[HOSTNAME]`&proxyPort=`[PORT]`&proxyUsername=`[USERNAME]`&proxyPassword=`[PASSWORD]`&proxyMethod=`[METHOD]`

代理类型：
1. SOCKS 4
2. SOCKS 5
3. HTTP
4. Telnet
5. Local