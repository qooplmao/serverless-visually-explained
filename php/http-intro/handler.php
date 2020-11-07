<?php

declare(strict_types=1);

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
        $body = 'Hello World';

        if (null !== $name = ($request->getQueryParams()['name'] ?? null)) {
            $body = sprintf('Hi %s', $name);
        }

        return new Response(200, [], $body);
    }
}

return new HttpHandler();