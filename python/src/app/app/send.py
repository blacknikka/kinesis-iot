# -*- coding:utf8 -*-
import os
from AWSIoTPythonSDK.MQTTLib import AWSIoTMQTTClient
import time
import json
from dotenv import load_dotenv

load_dotenv(verbose=True)

def customCallback(client, userdata, message):
    print('Received a new message: ')
    print(message.payload)
    print('from topic: ')
    print(message.topic)
    print('--------------\n\n')

# Read environment variables from .env file
clientId = os.environ.get("CLIENT_ID")
iotEndpoint = os.environ.get("AWS_IOT_ENDPOINT")
privateKeyFileName = os.environ.get("AWS_IOT_PRIVATE_PEM_KEY")
certFileName = os.environ.get("AWS_IOT_CERT_PEM")

# For certificate based connection
myMQTTClient = AWSIoTMQTTClient(clientId)
myMQTTClient.configureEndpoint(iotEndpoint, 8883)
myMQTTClient.configureCredentials('rootCA.pem', privateKeyFileName, certFileName)
myMQTTClient.configureOfflinePublishQueueing(-1) # Infinite offline Publish queueing
myMQTTClient.configureDrainingFrequency(2) # Draining: 2 Hz
myMQTTClient.configureConnectDisconnectTimeout(10) # 10 sec
myMQTTClient.configureMQTTOperationTimeout(5) # 5 sec
myMQTTClient.connect()


myMQTTClient.subscribe("myTopic", 1, customCallback)
time.sleep(2)

loopCount = 0
while True:
    message = {}
    message['message'] = 'from python!'
    message['sequence'] = loopCount
    messageJson = json.dumps(message)
    myMQTTClient.publish("myTopic", messageJson, 1)
    print("send")
    loopCount += 1
    time.sleep(1)
