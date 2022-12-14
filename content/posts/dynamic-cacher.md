+++
date = "2007-10-02T13:38:00Z"
title = "Dynamic Cacher"
description = "My own simple cache service written in PHP.  More a proof of concept than anything else."
aliases = [
  "/snip/2*",
  "/snip/02*",
  "/snip/dynamic-cacher"
]
tags = ["php", "code"]
categories = ["programming", "snip"]
+++

My own simple cache service written in PHP. More a proof of concept than anything else.

```php
<?php

/**
 * My own simple cache service written in PHP.  More a proof of concept than anything else.
 *
 * @author Phill Sparks <me@phills.me.uk>
 * @license http://creativecommons.org/licenses/by-sa/2.0/uk/ Creative Commons Attribution-ShareAlike 2.0 UK: England & Wales
 */

set_time_limit(3600);

global $fp, $headers;
$fp = null; $headers = array();

$url = 'http://milk-hub.net/'.($path = trim($_SERVER['REQUEST_URI'], '/'));
if ($path == '') $path = 'index';
$nfopath = dirname($path).'/._'.basename($path);

if ($path[0] == '.' || $path == 'index.php') {
    header('HTTP/1.0 403 '.statusmessage(403));
    echo '403 '. statusmessage(403);
    exit;
}
else if (is_dir($path)) {
    header('HTTP/1.0 403 '.statusmessage(403));
    echo '403 '. statusmessage(403);
    exit;
}

// Make dir for file
if (!is_dir(dirname($path))) mkdir(dirname($path), 0755, true);

// Do the fetch
$date = (file_exists($path) && file_exists($nfopath)) ? filemtime($path) : false;

if ($date < time() - 60*60*24 || $_GET['nocache']) {
    $date = $date ? date('r', $date) : false;

    if (file_exists($path))    unlink($path);
    if (file_exists($nfopath)) unlink($nfopath);

    if ($fp = fopen($path, 'wb')) {
        ignore_user_abort(true); // Even if the user stops fetching the file we need to finish caching it!

        $info = download($url, $date);
        $info['headers'] = $headers;
        file_put_contents($nfopath, serialize($info));

        fclose($fp);
        chmod($path, 0644);
    }
    else {
        header('HTTP/1.0 500 '.statusmessage(500));
        echo '500 '. statusmessage(500);
        exit;
    }
}
else {
    $info = unserialize(file_get_contents($nfopath));

    header('HTTP/1.0 '.$info['http_code'].' '.statusmessage($info['http_code']));
    if (is_array($info['headers'])) {
        foreach ($info['headers'] as $header) {
            header($header);
        }
    }
    else {
        header('Content-type: '.$info['content_type']);
        header('Content-Length: '.$info['download_content_length']);
        header('Last-Modified: '.date('r', filemtime($path)));
        header('Content-Disposition: inline; filename='.basename($path));
    }
    readfile($path);
}

$fsize = filesize($path);
if ($info['http_code'] < 200 || $info['http_code'] >= 400 || $info['download_content_length'] != $fsize || $fsize == 0) {
    if (file_exists($path))    unlink($path);
    if (file_exists($nfopath)) unlink($nfopath);
}

exit;

function download($url, $date) {
    global $fp;

    $ch = curl_init($url);

    $header = array(
        'Accept: text/xml,application/xml,application/xhtml+xml,text/html;q=0.9,text/plain;q=0.8,image/png,*/*;q=0.5',
        'Cache-Control: max-age=0',
        'Connection: keep-alive',
        'Keep-Alive: 300',
        'Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7',
        'Accept-Language: en-us,en;q=0.5',
        'Pragma: ', // browsers keep this blank.
    );
    if ($date) {
        $header[] = 'If-Modified-Since: '.$date;
    }

    curl_setopt($ch, CURLOPT_BINARYTRANSFER, 1);
    curl_setopt($ch, CURLOPT_ENCODING,    'gzip,deflate');
    curl_setopt($ch, CURLOPT_FILE, $fp);
    curl_setopt($ch, CURLOPT_HTTPHEADER,  $header);
    curl_setopt($ch, CURLOPT_REFERER,     $_SERVER['HTTP_REFERER']);
    // curl_setopt($ch, CURLOPT_TIMEOUT,     10);
    curl_setopt($ch, CURLOPT_USERAGENT,  'CacheBot/1.0 (+http://cache.milk-hub.net)');

    curl_setopt($ch, CURLOPT_HEADERFUNCTION, 'read_header');
    curl_setopt($ch, CURLOPT_WRITEFUNCTION, 'read_body');

    curl_exec($ch);
    $info = curl_getinfo($ch);
    curl_close($ch);

    return $info;
}

function read_header($ch, $data) {
    global $headers;

    list($field, $value) = explode(':', $data, 2);

    switch ($field) {
        case 'Content-type':
        case 'Content-Disposition':
        case 'Content-Length':
        case 'Last-Modified':
            header($data);
            $headers[] = $data;
        break;
    }

    return strlen($data);
}

function read_body($ch, $data) {
    global $fp;

    fwrite($fp, $data);
    echo $data; flush();

    return strlen($data);
}

function statusmessage($statuscode) {
    switch($statuscode) {
        case 100:
            return 'Continue';
        case 101:
            return 'Switching Protocols';
        case 200:
            return 'OK';
        case 201:
            return 'Created';
        case 202:
            return 'Accepted';
        case 203:
            return 'Non-Authoritative Information';
        case 204:
            return 'No Content';
        case 205:
            return 'Reset Content';
        case 206:
            return 'Partial Content';
        case 300:
            return 'Multiple Choices';
        case 301:
            return 'Moved Permanently';
        case 302:
            return 'Found';
        case 303:
            return 'See Other';
        case 304:
            return 'Not Modified';
        case 305:
            return 'Use Proxy';
        case 306:
            return 'No Longer Used';
        case 307:
            return 'Temporary Redirect';
        case 400:
            return 'Bad Request';
        case 401:
            return 'Not Authorised';
        case 402:
            return 'Payment Required';
        case 403:
            return 'Forbidden';
        case 404:
            return 'Not Found';
        case 405:
            return 'Method Not Allowed';
        case 406:
            return 'Not Acceptable';
        case 407:
            return 'Proxy Authentication Required';
        case 408:
            return 'Request Timeout';
        case 409:
            return 'Conflict';
        case 410:
            return 'Gone';
        case 411:
            return 'Length Required';
        case 412:
            return 'Precondition Failed';
        case 413:
            return 'Request Entity Too Large';
        case 414:
            return 'Request URI Too Long';
        case 415:
            return 'Unsupported Media Type';
        case 416:
            return 'Requested Range Not Satisfiable';
        case 417:
            return 'Expectation Failed';
        case 500:
            return 'Internal Server Error';
        case 501:
            return 'Not Implemented';
        case 502:
            return 'Bad Gateway';
        case 503:
            return 'Service Unavailable';
        case 504:
            return 'Gateway Timeout';
        case 505:
            return 'HTTP Version Not Supported ';
        default:
            return 'Unknown';
    }
}
```
