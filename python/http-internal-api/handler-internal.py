import json

def handler(event, context):
    return {
        'uppercase': event['name'].upper(),
        'lowercase': event['name'].lower(),
    }
