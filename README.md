### 起動方法
```bash
cp .env.example .env
# .envを記載する（AWSアカウント）
docker-compose up -d
docker-compose exec tf bash

# コンテナ内でterraform
terraform init
terraform apply
```

### kinesisへのデータのPOST方法
- 参考
  - https://docs.aws.amazon.com/ja_jp/iot/latest/developerguide/http.html

```bash
curl --tlsv1.2 --cert iot-cert.cert.pem --key iot-moor.private.key -X POST -d '{"sens": 1.0, "sens2": 2.5}' https://<iot-core-endpoint>:8443/topics/iot/something
```

### Lambdaのビルド方法
```bash
make
```
