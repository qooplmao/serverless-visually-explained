<?php

declare(strict_types=1);

use Aws\Lambda\LambdaClient;
use Nyholm\Psr7\Response;
use Psr\Http\Message\ResponseInterface;
use Psr\Http\Message\ServerRequestInterface;
use Psr\Http\Server\RequestHandlerInterface;

class HttpHandler implements RequestHandlerInterface
{
    /**
     * @param ServerRequestInterface $request
     * @return ResponseInterface
     */
    public function handle(ServerRequestInterface $request): ResponseInterface
    {
        $lambda = new LambdaClient([
            'version'   => 'latest',
            'region'    => 'eu-west-2',
        ]);

        $result = $lambda->invoke([
            'FunctionName'      => 'HttpInternalApiPhp-dev-internalApi',
            'InvocationType'    => 'RequestResponse',
            'LogType'           => 'None',
            'Payload'           => json_encode([
                'name'          => $request->getQueryParams()['name'] ?? 'None Set',
            ]),
        ]);

        return new Response(200, [], $result->get('Payload')->getContents());
    }
}

return new HttpHandler();