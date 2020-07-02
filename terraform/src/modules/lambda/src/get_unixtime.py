import base64
import json
import time

import os
import sys
import pytz
from datetime import datetime, timedelta

def lambda_handler(event: dict, context):
    print(datetime.now().isoformat())

    records_count = len(event['Records'])
    print(f'kinesis recourds count: {records_count}')

    for record in event['Records']:
        b64_data = record['kinesis']['data']
        data = base64.b64decode(b64_data).decode('utf-8')
        print(data)

    return {
        'statusCode': 200,
        'body': json.dumps(
            {'message': 'hello world'}
        ),
    }
