# Proxypac Server

proxy.pac ファイルの配信用ローカルサーバー

[Chrome 70 macOS での変更](https://chromium.googlesource.com/chromium/src/+/5b4a3d893cfd6a42a2d1f685fa6a828d34c9c1de%5E%21/#F0)により、Proxy pacファイルの取得URLに `file://` スキームが使えなくなったので、ローカルにproxy.pacファイル配信専用のHTTPサーバを立てて、そこから取得できるようにする。
Safariでも少し前から同様の動作をしているので、使えると思う。

```
proxypac_server [-p port] pacfile
```
