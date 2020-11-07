def httpIntro(event, context):
    body = 'Hello World'

    if 'queryStringParameters' in event and 'name' in event['queryStringParameters']:
        body = 'Hi ' + event['queryStringParameters']['name']

    return {
        'statusCode': 200,
        'headers': {},
        'body': body
    }
