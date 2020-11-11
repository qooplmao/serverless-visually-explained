const AWS = require('aws-sdk');

const lambda = new AWS.Lambda({
  apiVersion: 'latest',
  region: 'eu-west-2',
});

const params = {
  FunctionName: 'HttpInternalApiJs-dev-internalApi',
  InvocationType: 'RequestResponse',
  LogType: 'None',
  Payload: JSON.stringify({}),
};

exports.handler = (event, context, callback) => {
  const name = event.queryStringParameters && event.queryStringParameters.name
    ? event.queryStringParameters.name
    : 'None Set';

  params.Payload = JSON.stringify({
    name
  });

  console.log(params.Payload);

  const request = lambda.invoke(params);

  request
    .on('success', ({ data }) => {
      console.log({ name: 'success', data });
      callback(null, data.Payload);
    })
    .on('error', error => {
      console.log({ name: 'error', error });
      callback(error);
    })
    .on('complete', () => {
      console.log('DONE!!!');
    })
    .send()
  ;
}
