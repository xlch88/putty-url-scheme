# PuTTY URL Scheme Helper 
Open PuTTY as a url scheme

[中文](README_CN.md)

# Install
1. download release binary file to PuTTY path.
2. run `putty-url-scheme.exe --register` to add registry
3. Open a new tab of your browser and enter `ssh://username:password@hostname:port` in the address bar

# URL Scheme
ssh://username:password@hostname:port

## Use Proxy
ssh://username:password@hostname:port/?proxyHost=`[HOSTNAME]`&proxyPort=`[PORT]`&proxyUsername=`[USERNAME]`&proxyPassword=`[PASSWORD]`&proxyMethod=`[METHOD]`

Proxy Method:
1. SOCKS 4
2. SOCKS 5
3. HTTP
4. Telnet
5. Local