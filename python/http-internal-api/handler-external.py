import boto3
import json

lambdaClient = boto3.client('lambda', region_name = 'eu-west-2')

def handler(event, context):
    name = 'None Set'

    if 'queryStringParameters' in event and 'name' in event['queryStringParameters']:
        name = event['queryStringParameters']['name']

    response = lambdaClient.invoke(
        FunctionName = 'HttpInternalApiPython-dev-internalApi',
        InvocationType = 'RequestResponse',
        LogType = 'None',
        Payload = json.dumps({'name': name })
    )

    print(response)

    return {
        'statusCode': 200,
        'headers': {},
        'body': response['Payload'].read()
    }
