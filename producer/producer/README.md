### 概要
- USBからデータを受けてそれをAWS IoTCoreのエンドポイントに送信する
- ラズパイ等のIoT機器として動作することを想定

### 実行環境
- 実行環境にあわせてビルドしなおせばOK
- ラズパイをとりあえず想定

### 参考サイト
- https://mikan.github.io/2018/10/22/accessing-aws-iot-mqtt-through-port-443-from-go/

### how to use
- 鍵、証明書を配置(`cert`フォルダ)する
  - 以下のシェルスクリプト実行
```bash
$ sh MakeKey.sh
```

### データPOSTの例
#### 通常データ
```
$ curl -X POST -H "Content-Type: application/json" -d '{"kind":"start","ver":"ver1"}' localhost:8000/stats
```

#### エラー発生
```
$ curl -X POST -H "Content-Type: application/json" -d '{"kind":"error","ver":"ver1"}' localhost:8000/stats
```
