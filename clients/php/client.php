#!/usr/bin/env php
<?php

error_reporting(E_ALL);

require_once __DIR__ . '/vendor/autoload.php';

// KujiService
require_once __DIR__ . '/kuji/KujiService.php';
require_once __DIR__ . '/kuji/Types.php';


use Thrift\ClassLoader\ThriftClassLoader;

$GEN_DIR_Kuji = realpath(dirname(__FILE__)).'/kuji';

$loader = new ThriftClassLoader();
$loader->registerNamespace('Thrift', __DIR__ . '/vendor/apache/lib/php/lib');
$loader->registerDefinition('kuji', $GEN_DIR_Kuji);
$loader->register();

use Thrift\Protocol\TJSONProtocol;
use Thrift\Transport\TSocket;
use Thrift\Transport\TBufferedTransport;
use Thrift\Exception\TException;
use Thrift\Factory\TJSONProtocolFactory;
use Thrift\Factory\TTransportFactory;
use Thrift\Protocol\TMultiplexedProtocol;

try {
  $hostname = '127.0.0.1';
  $socket = new TSocket($hostname, 50051);
  $transport = new TBufferedTransport($socket, 1024, 1024);
  $protocol = new TJSONProtocol($transport);
  $multiplexKujiService = new TMultiplexedProtocol($protocol, "KujiService");

  $client = new \kuji\KujiServiceClient($multiplexKujiService);

  $transport->open();

  $candidate1 = new \kuji\KujiCandidate([
          'id' => 1,
      'weight' => 200,
  ]);
  $candidate2 = new \kuji\KujiCandidate([
          'id' => 2,
      'weight' => 500,
  ]);
  $candidate3 = new \kuji\KujiCandidate([
          'id' => 3,
      'weight' => 300,
  ]);
  $candidates = new \kuji\KujiCandidates([
          'candidates' => [
              $candidate1,
              $candidate2,
              $candidate3
          ]
  ]);
  $req = new \kuji\ReqCandidates([
          'key' => "php_simple",
          'candidates' => $candidates
  ]);
  $res = $client->ThRegisterCandidatesWithKey($req);

  $result = (array)$res;
  var_dump($result);

  $req2 = new \kuji\ReqPickOneByKey([
          'key' => "php_simple"
  ]);
  $res2 = $client->ThPickOneByKey($req2);
  $result = (array)$res2;
  var_dump($result);

  $transport->close();

} catch (TException $tx) {
  print 'TException: '.$tx->getMessage()."\n";
}
