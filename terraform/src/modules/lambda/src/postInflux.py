import json
import base64
import pytz
import requests
import os
from datetime import datetime, timedelta

# endpointを取得
influxEndpoint = os.environ['INFLUX_ENDPOINT']

def lambda_handler(event: dict, context):
    influx_raw_query = ''
    for record in event['Records']:
        # base64されたデータを取得
        content64 = record['kinesis']['data']

        # utf-8でエンコードしてからbase64デコードする
        content = base64.decodebytes(content64.encode('utf-8'))

        # 取得した文字列を辞書データに変換
        content_dict = json.loads(content)

        # timestamp取得
        if 'timestamp' in content_dict:
            timestamp = content_dict.pop('timestamp')
        else:
            UTC = pytz.timezone('UTC')
            now = datetime.now(UTC)
            now = int(now.timestamp())
            timestamp = str(now * pow(10, 9))

        # influxのクエリの記載
        raw_query = 'sens '
        raw_query += ",".join(["{0}={1}".format(key, value) for (key, value) in content_dict.items()]) + ' '

        # timestampの記載
        raw_query += timestamp

        # 追記する
        influx_raw_query += raw_query + '\n'

    url = f'http://{influxEndpoint}:8086/write?db=iot'
    response = requests.post(url, data=influx_raw_query)

    if response.status_code >= 200 and response.status_code < 300:
        return {
            'statusCode': 200,
        }
    else:
        return {
            'statusCode': response.status_code,
        }
