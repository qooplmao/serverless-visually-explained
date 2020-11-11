<?php

return function($event) {
    return [
        'uppercase' => mb_strtoupper($event['name']),
        'lowercase' => mb_strtolower($event['name']),
    ];
};
