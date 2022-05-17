<?php


function tcp() 
{
	$addr = "127.0.0.1";
	$socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
	$connRes = socket_connect($socket, $addr, 9999);
	if (!$connRes) {
		echo ("tcp链接失败\n");
		return;
	}

	$sendData = "hello server tcp" . PHP_EOL;
	$sendLen = socket_send($socket, $sendData, strlen($sendData), MSG_OOB);
	echo "发送成功，长度：{$sendLen}\n";

	$recvData = null;
	$recvData = socket_read($socket, 1024, PHP_BINARY_READ);
	if ($recvData === false) {
		echo "socket_read() failed; reason: " . socket_strerror(socket_last_error($socket)) . "\n";
		return;
	}
	echo ("recv: {$recvData}\n");

	
	socket_close($socket);
}


function udp() 
{
	$addr = "127.0.0.1";
	$socket = socket_create(AF_INET, SOCK_DGRAM, SOL_UDP);
	$connRes = socket_connect($socket, $addr, 9999);
	if (!$connRes) {
		echo ("udp链接失败\n");
		return;
	}

	$sendData = "hello server udp";
	$sendLen = socket_write($socket, $sendData, strlen($sendData));
	echo "发送成功，长度：{$sendLen}\n";

	$recvData = null;
	$recvData = socket_read($socket, 1024, PHP_BINARY_READ);
	if ($recvData === false) {
		echo "socket_read() failed; reason: " . socket_strerror(socket_last_error($socket)) . "\n";
		return;
	}
	echo ("recv: {$recvData}\n");

	
	socket_close($socket);
}


tcp();
udp();






