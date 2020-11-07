exports.httpIntro = async function (event, context) {
  let body = 'Hello World';

  if (event.queryStringParameters && event.queryStringParameters.name) {
    body = `Hi ${event.queryStringParameters.name}`;
  }

  return {
    statusCode: 200,
    headers: {},
    body
  }
}
