package main

// will be set in flags
var listenAddressAndPort string
var redisHostPort string
var ipfsHostPort string
var contractAddr string
var SubChainIp string
var queueConcurrency = 10
var ipfsGCInterval = 100                         // in seconds
var ipfsUnpinInterval = 100                      // in seconds
var unpinInterval = 60                           // in seconds
var RestoredFileUnpinInterval = int64(3600 * 24) // in seconds

//var vnodeIP string = "http://gateway.moac.io/mainnet"
//var VnodeIp string = "http://gateway.moac.io/testnet"
//var VnodeIp string = "http://120.78.2.59:8547"
var VnodeIp string

//存储设备目录
var memDeviceDir string
