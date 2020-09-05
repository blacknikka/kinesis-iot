import json
import base64
import pytz
import requests
import os
import pymongo
from datetime import datetime, timedelta

# endpointを取得(influxDBエンドポイント)
influxEndpoint = os.environ['INFLUX_ENDPOINT']

# documentDB endpoint
docdbClusterEndpoint = os.environ['DOCDB_CLUSTER_ENDPOINT']

# documentDB admin
docdbAdminUser = os.environ['DOCDB_ADMIN_USER']
docdbAdminPassword = os.environ['DOCDB_ADMIN_PASSWORD']

def lambda_handler(event: dict, context):
    influx_raw_query = ''
    for record in event['Records']:
        # partitionによってPOST先を変更する
        if record['kinesis']['partitionKey'] == 'iot/time':
            # timeseris (InfluxDB)
            print('timeseries')

            # base64されたデータを取得
            content64 = record['kinesis']['data']
            # utf-8でエンコードしてからbase64デコードする
            content = base64.decodebytes(content64.encode('utf-8'))

            # create query from string
            raw_query = createfluxQuery(content)

            # 追記する
            influx_raw_query += raw_query + '\n'
        # 'stats'を期待
        else:
            # documentDB
            print('stats')
            print(record)
            postDocumentDB(record['kinesis'])

    # insert to influxDB
    if influx_raw_query != '':
        response = insertToInflux(influx_raw_query)
        if response.status_code >= 200 and response.status_code < 300:
            return {
                'statusCode': 200,
            }
        else:
            return {
                'statusCode': response.status_code,
            }
    else:
        return {
            'statusCode': 200,
        }

# insert to InfluxDB
def insertToInflux(influx_raw_query):
    print("insertToInflux")
    print(influx_raw_query)
    url = f'http://{influxEndpoint}:8086/write?db=iot'
    response = requests.post(url, data=influx_raw_query)

    return response

# create InfluxDB query string
def createfluxQuery(content):
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

    return raw_query

# post data to document DB
def postDocumentDB(data):
    print("postDocumentDB")
    client = pymongo.MongoClient(
        f'mongodb://{docdbAdminUser}:{docdbAdminPassword}@{docdbClusterEndpoint}:27017/?ssl=true&ssl_ca_certs=rds-combined-ca-bundle.pem&replicaSet=rs0&readPreference=secondaryPreferred&retryWrites=false'
    )

    db = client.machines
    col = db.playdata

    # utf-8でエンコードしてからbase64デコードする
    content64 = data['data']
    content = base64.decodebytes(content64.encode('utf-8'))
    json_dict = json.loads(content)

    col.insert_one({'data': json_dict, 'timestamp': data['approximateArrivalTimestamp']})

    client.close()
