// tokenTx.go
package main

import (
	"errors"
	"math/big"
	"stormcatcher/lib/common"
	"stormcatcher/vnode/accounts/abi"
	"strings"
)

const (
	//	tokenAbi string = `[{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"shardMapping","outputs":[{"name":"shardId","type":"uint256"},{"name":"nodeCount","type":"uint256"},{"name":"weight","type":"uint256"},{"name":"size","type":"uint256"},{"name":"availableSize","type":"uint256"},{"name":"percentage","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"recentlyUsedList","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"shardCount","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"addrs","type":"address[]"},{"name":"addr","type":"address"}],"name":"have","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"pos","type":"uint256"},{"name":"tosend","type":"address[]"},{"name":"amount","type":"uint256[]"},{"name":"times","type":"uint256[]"}],"name":"postFlush","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"admin","type":"address"}],"name":"removeAdmin","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"value","type":"uint256"}],"name":"recentlyUsed","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"fileMapping","outputs":[{"name":"fileId","type":"uint256"},{"name":"fileHash","type":"string"},{"name":"fileName","type":"string"},{"name":"fileSize","type":"uint256"},{"name":"fileOwner","type":"address"},{"name":"createTime","type":"uint256"},{"name":"verifiedCount","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"fileList","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint256"},{"name":"","type":"uint256"}],"name":"shardNodeList","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"newlist","type":"address[]"}],"name":"updateNodeList","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getAllShards","outputs":[{"name":"","type":"uint256[]"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"num","type":"uint256"}],"name":"setBlockVerificationInterval","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"admins","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"fileCount","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"userAddr","type":"address"},{"name":"pos","type":"uint256"}],"name":"getRedeemMapping","outputs":[{"name":"redeemingAddr","type":"address[]"},{"name":"redeemingAmt","type":"uint256[]"},{"name":"redeemingtime","type":"uint256[]"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"index","type":"uint256"}],"name":"removeFromAddressArray","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getCurNodeList","outputs":[{"name":"nodeList","type":"address[]"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"curNodeList","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"userAddr","type":"address"}],"name":"getEnterRecords","outputs":[{"name":"enterAmt","type":"uint256[]"},{"name":"entertime","type":"uint256[]"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"shardList","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"array","type":"uint256[]"},{"name":"index","type":"uint256"}],"name":"removeFromArray","outputs":[{"name":"","type":"uint256[]"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"verifyGroupId","type":"address"},{"name":"verifyNodeId","type":"address"},{"name":"votingNodeId","type":"address"},{"name":"blockNumber","type":"uint256"},{"name":"fileHash","type":"string"},{"name":"shardId","type":"uint256"}],"name":"voteVerifyTransaction","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"fileId","type":"uint256"},{"name":"ipfsId","type":"string"}],"name":"readFile","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"verifyGroupId","type":"address"},{"name":"verifyNodeId","type":"address"},{"name":"blockNumber","type":"uint256"},{"name":"fileHash","type":"string"},{"name":"shardId","type":"uint256"}],"name":"submitVerifyTransaction","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"admin","type":"address"}],"name":"addAdmin","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"val1","type":"uint256"},{"name":"val2","type":"uint256"}],"name":"min","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"fileHash","type":"string"},{"name":"fileName","type":"string"},{"name":"fileSize","type":"uint256"},{"name":"createTime","type":"uint256"}],"name":"addFile","outputs":[{"name":"","type":"uint256"},{"name":"","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"scsId","type":"address"},{"name":"beneficiary","type":"address"},{"name":"weight","type":"uint256"}],"name":"addNode","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"redeemFromMicroChain","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":false,"inputs":[{"name":"fileHash","type":"string"},{"name":"fileName","type":"string"},{"name":"fileSize","type":"uint256"},{"name":"createTime","type":"uint256"},{"name":"ipfsId","type":"string"}],"name":"addFile","outputs":[{"name":"","type":"uint256"},{"name":"","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"weight","type":"uint256"}],"name":"addShard","outputs":[{"name":"shardId","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"fileId","type":"uint256"}],"name":"removeFile","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"myAddr","type":"address"}],"name":"getMyFileHashes","outputs":[{"name":"","type":"uint256[]"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"shardId","type":"uint256"}],"name":"getAllFilesByShard","outputs":[{"name":"","type":"uint256[]"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"scsId","type":"address"}],"name":"removeNode","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"unassignedNoteList","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"amount","type":"uint256"}],"name":"setAwardAmount","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"size","type":"uint256"}],"name":"setShardSize","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"weight","type":"uint256"},{"name":"size","type":"uint256"}],"name":"setCapacity","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"fileId","type":"uint256"}],"name":"getFileById","outputs":[{"name":"","type":"string"},{"name":"","type":"string"},{"name":"","type":"uint256"},{"name":"","type":"address"},{"name":"","type":"uint256"},{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"capacityMapping","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"enterPos","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"nodeShardMapping","outputs":[{"name":"shardId","type":"uint256"},{"name":"nodeCount","type":"uint256"},{"name":"weight","type":"uint256"},{"name":"size","type":"uint256"},{"name":"availableSize","type":"uint256"},{"name":"percentage","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"nodeMapping","outputs":[{"name":"scsId","type":"address"},{"name":"beneficiary","type":"address"},{"name":"size","type":"uint256"},{"name":"lastVerifiedBlock","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"beneficiary","type":"address"}],"name":"award","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"},{"name":"","type":"uint256"}],"name":"verifyGroupMapping","outputs":[{"name":"scsId","type":"address"},{"name":"verifyNodeId","type":"address"},{"name":"blockNumber","type":"uint256"},{"name":"fileHash","type":"string"},{"name":"totalCount","type":"uint256"},{"name":"votedCount","type":"uint256"},{"name":"affirmCount","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":true,"stateMutability":"payable","type":"constructor"}]`
	tokenAbi string = `[ { "constant": true, "inputs": [ { "name": "", "type": "uint256" } ], "name": "shardMapping", "outputs": [ { "name": "shardId", "type": "uint256" }, { "name": "nodeCount", "type": "uint256" }, { "name": "weight", "type": "uint256" }, { "name": "size", "type": "uint256" }, { "name": "availableSize", "type": "uint256" }, { "name": "percentage", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "uint256" } ], "name": "recentlyUsedList", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "shardCount", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "addrs", "type": "address[]" }, { "name": "addr", "type": "address" } ], "name": "have", "outputs": [ { "name": "", "type": "bool" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "pos", "type": "uint256" }, { "name": "tosend", "type": "address[]" }, { "name": "amount", "type": "uint256[]" }, { "name": "times", "type": "uint256[]" } ], "name": "postFlush", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "fileHash", "type": "string" }, { "name": "fileName", "type": "string" }, { "name": "fileSize", "type": "uint256" }, { "name": "createTime", "type": "uint256" }, { "name": "shardId", "type": "uint256" } ], "name": "addFile", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "admin", "type": "address" } ], "name": "removeAdmin", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "address" } ], "name": "hardDriveMapping", "outputs": [ { "name": "scsId", "type": "address" }, { "name": "weight", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "value", "type": "uint256" } ], "name": "recentlyUsed", "outputs": [ { "name": "", "type": "bool" } ], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "uint256" } ], "name": "fileMapping", "outputs": [ { "name": "fileId", "type": "uint256" }, { "name": "fileHash", "type": "string" }, { "name": "fileName", "type": "string" }, { "name": "fileSize", "type": "uint256" }, { "name": "fileOwner", "type": "address" }, { "name": "createTime", "type": "uint256" }, { "name": "verifiedCount", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "driveId", "type": "address" }, { "name": "scsId", "type": "address" }, { "name": "weight", "type": "uint256" } ], "name": "updateHardDrive", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "uint256" } ], "name": "nodeBeneficiaryList", "outputs": [ { "name": "", "type": "address" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "uint256" } ], "name": "fileList", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "uint256" }, { "name": "", "type": "uint256" } ], "name": "shardNodeList", "outputs": [ { "name": "", "type": "address" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "address" } ], "name": "nodeBeneficiaryMapping", "outputs": [ { "name": "scsId", "type": "address" }, { "name": "beneficiary", "type": "address" }, { "name": "value", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "newlist", "type": "address[]" } ], "name": "updateNodeList", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [], "name": "getAllShards", "outputs": [ { "name": "", "type": "uint256[]" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "num", "type": "uint256" } ], "name": "setBlockVerificationInterval", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "address" } ], "name": "admins", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "fileCount", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "userAddr", "type": "address" }, { "name": "pos", "type": "uint256" } ], "name": "getRedeemMapping", "outputs": [ { "name": "redeemingAddr", "type": "address[]" }, { "name": "redeemingAmt", "type": "uint256[]" }, { "name": "redeemingtime", "type": "uint256[]" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "index", "type": "uint256" } ], "name": "removeFromAddressArray", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "weight", "type": "uint256" } ], "name": "pubAddShard", "outputs": [ { "name": "shardId", "type": "uint256" } ], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "index", "type": "uint256" } ], "name": "removeFromAddressArray2", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [], "name": "getCurNodeList", "outputs": [ { "name": "nodeList", "type": "address[]" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "uint256" } ], "name": "curNodeList", "outputs": [ { "name": "", "type": "address" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "userAddr", "type": "address" } ], "name": "getEnterRecords", "outputs": [ { "name": "enterAmt", "type": "uint256[]" }, { "name": "entertime", "type": "uint256[]" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "scsId", "type": "address" }, { "name": "status", "type": "uint256" } ], "name": "updateScsMsg", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "uint256" } ], "name": "shardList", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "verifyGroupId", "type": "address" }, { "name": "verifyNodeId", "type": "address" }, { "name": "votingNodeId", "type": "address" }, { "name": "blockNumber", "type": "uint256" }, { "name": "fileHash", "type": "string" }, { "name": "shardId", "type": "uint256" } ], "name": "voteVerifyTransaction", "outputs": [ { "name": "", "type": "bool" } ], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "fileId", "type": "uint256" }, { "name": "ipfsId", "type": "string" } ], "name": "readFile", "outputs": [ { "name": "", "type": "bool" } ], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "verifyGroupId", "type": "address" }, { "name": "verifyNodeId", "type": "address" }, { "name": "blockNumber", "type": "uint256" }, { "name": "fileHash", "type": "string" }, { "name": "shardId", "type": "uint256" } ], "name": "submitVerifyTransaction", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "admin", "type": "address" } ], "name": "addAdmin", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "val1", "type": "uint256" }, { "name": "val2", "type": "uint256" } ], "name": "min", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "scsId", "type": "address" }, { "name": "beneficiary", "type": "address" }, { "name": "weight", "type": "uint256" } ], "name": "addNode", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [], "name": "redeemFromMicroChain", "outputs": [], "payable": true, "stateMutability": "payable", "type": "function" }, { "constant": false, "inputs": [], "name": "doAward", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [], "name": "shardSize", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "fileId", "type": "uint256" } ], "name": "removeFile", "outputs": [ { "name": "", "type": "bool" } ], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [ { "name": "myAddr", "type": "address" } ], "name": "getMyFileHashes", "outputs": [ { "name": "", "type": "uint256[]" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "shardId", "type": "uint256" } ], "name": "getAllFilesByShard", "outputs": [ { "name": "", "type": "uint256[]" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "fileHash", "type": "string" }, { "name": "fileName", "type": "string" }, { "name": "fileSize", "type": "uint256" }, { "name": "createTime", "type": "uint256" }, { "name": "ipfsId", "type": "string" }, { "name": "shardId", "type": "uint256" } ], "name": "addFile", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "scsId", "type": "address" } ], "name": "removeNode", "outputs": [ { "name": "", "type": "bool" } ], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "beneficiary", "type": "address" }, { "name": "value", "type": "uint256" } ], "name": "pay", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "uint256" } ], "name": "unassignedNoteList", "outputs": [ { "name": "", "type": "address" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "amount", "type": "uint256" } ], "name": "setAwardAmount", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "size", "type": "uint256" } ], "name": "setShardSize", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [ { "name": "weight", "type": "uint256" }, { "name": "size", "type": "uint256" } ], "name": "setCapacity", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [ { "name": "fileId", "type": "uint256" } ], "name": "getFileById", "outputs": [ { "name": "", "type": "string" }, { "name": "", "type": "string" }, { "name": "", "type": "uint256" }, { "name": "", "type": "address" }, { "name": "", "type": "uint256" }, { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "uint256" } ], "name": "capacityMapping", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "index", "type": "uint256" } ], "name": "removeFromArray", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [], "name": "enterPos", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "address" } ], "name": "nodeShardMapping", "outputs": [ { "name": "shardId", "type": "uint256" }, { "name": "nodeCount", "type": "uint256" }, { "name": "weight", "type": "uint256" }, { "name": "size", "type": "uint256" }, { "name": "availableSize", "type": "uint256" }, { "name": "percentage", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "address" } ], "name": "nodeMapping", "outputs": [ { "name": "shardId", "type": "uint256" }, { "name": "scsId", "type": "address" }, { "name": "beneficiary", "type": "address" }, { "name": "size", "type": "uint256" }, { "name": "lastVerifiedBlock", "type": "uint256" }, { "name": "balance", "type": "uint256" }, { "name": "status", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "address" }, { "name": "", "type": "uint256" } ], "name": "verifyGroupMapping", "outputs": [ { "name": "scsId", "type": "address" }, { "name": "verifyNodeId", "type": "address" }, { "name": "blockNumber", "type": "uint256" }, { "name": "fileHash", "type": "string" }, { "name": "totalCount", "type": "uint256" }, { "name": "votedCount", "type": "uint256" }, { "name": "affirmCount", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "inputs": [], "payable": true, "stateMutability": "payable", "type": "constructor" } ]`
)

func assembleInput(abiStr, name string, args ...interface{}) ([]byte, error) {

	parsed, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		return nil, errors.New("abiStr error")
	}

	return parsed.Pack(name, args...)
}

//func addFileAssembleByte(fileHash, fileName, ipfsId string, fileSize, createTime int64) ([]byte, error) {
//	return assembleInput(tokenAbi, "addFile", fileHash, fileName, big.NewInt(fileSize), big.NewInt(createTime), ipfsId)
//}

//func readFileAssembleByte(fileId int64, ipfsId string) ([]byte, error) {
//	return assembleInput(tokenAbi, "readFile", big.NewInt(fileId), ipfsId)
//}

//func removeFileAssembleByte(fileId int64) ([]byte, error) {
//	return assembleInput(tokenAbi, "removeFile", big.NewInt(fileId))
//}

//func addNodeAssembleByte(scsId, beneficiary string, weight int64) ([]byte, error) {
//	return assembleInput(tokenAbi, "addNode", common.HexToAddress(scsId), common.HexToAddress(beneficiary), big.NewInt(weight))
//}

//func removeNodeAssembleByte(scsId string) ([]byte, error) {
//	return assembleInput(tokenAbi, "removeNode", common.HexToAddress(scsId))
//}

func submitVerifyTransactionAssembleByte(verifyGroupId, verifyNodeId string, blockNumber int64, fileHash string, shardId int64) ([]byte, error) {
	return assembleInput(tokenAbi, "submitVerifyTransaction", common.HexToAddress(verifyGroupId), common.HexToAddress(verifyNodeId), big.NewInt(blockNumber), fileHash, big.NewInt(shardId))
}

func voteVerifyTransactionAssembleByte(verifyGroupId, verifyNodeId, votingNodeId string, blockNumber int64, fileHash string, shardId int64) ([]byte, error) {
	return assembleInput(tokenAbi, "voteVerifyTransaction", common.HexToAddress(verifyGroupId), common.HexToAddress(verifyNodeId), common.HexToAddress(votingNodeId), big.NewInt(blockNumber), fileHash, big.NewInt(shardId))
}

func updateHardDriveAssembleByte(driveId, scsId string, weight int64) ([]byte, error) {
	return assembleInput(tokenAbi, "updateHardDrive", common.HexToAddress(driveId), common.HexToAddress(scsId), big.NewInt(weight))
}
