package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	validator "gopkg.in/validator.v2"

	"stormcatcher/lib/common"
	"stormcatcher/vnode/mcclient"

	"crypto/md5"

	"stormcatcher/Chain3Go"

	"stormcatcher/vnode/accounts"
	"stormcatcher/vnode/accounts/keystore"
	//	"stormcatcher/vnode/accounts/abi/bind"
	//	"stormcatcher/lib/types"
)

var (
	mcClient *mcclient.Client = nil
	scsId    string           = ""
)

func getMcClient(vnodeIP string) (*Main, error) {

	if mcClient == nil {
		var err error
		mcClient, err = mcclient.Dial(vnodeIP)
		if err != nil {
			return nil, err
		}
	}
	return NewMain(common.HexToAddress(contractAddr), mcClient)
}

var accountsManager *accounts.Manager = nil

/*
 * 创建操作管理
 */
func makeAccountsManager() (*accounts.Manager, error) {

	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP

	var (
		keydir string = "./scskeystore"
	)

	if err := os.MkdirAll(keydir, 0700); err != nil {
		return nil, err
	}
	// Assemble the accounts manager and supported backends
	backends := []accounts.Backend{
		keystore.NewKeyStore(keydir, scryptN, scryptP),
	}

	return accounts.NewManager(backends...), nil
}

/*
 * 获取操作
 */
func fetchKeystore() *keystore.KeyStore {
	if accountsManager == nil {
		accountsManager, _ = makeAccountsManager()
	}
	return accountsManager.Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
}

func getScsId(scsIdTmp string) string {

	if scsIdTmp == "" {
		accs := fetchKeystore().Accounts()
		if len(accs) >= 1 {
			scsIdTmp = accs[0].Address.Hex()
			return scsIdTmp
		}
	}
	return scsIdTmp
}

func getScsJsonStr(pw string) ([]byte, error) {

	fk := fetchKeystore()

	accs := fk.Accounts()

	return fk.Export(accs[0], pw, pw)
}

var q2cMapping map[string](chan string)
var handlerMapping map[string]func(...string) error

/*
func handleIPFSRead(originalFileHash string) error {
	log.Info("inside read gorouting")

	alteredTmpfile, errCheckout := checkoutIPFSFile(originalFileHash)
	if errCheckout != nil {
		return errCheckout
	}
	log.Info("Checkout altered file hash", originalFileHash, "with file", alteredTmpfile.Name())
	defer os.Remove(alteredTmpfile.Name())

	// step 3
	restoredTmpFile := createRestoredTmpFile(alteredTmpfile)
	defer os.Remove(restoredTmpFile.Name())

	// step 4, originalFileHash set to "" means it's not to checkin new written file
	if errCheckInIpfs := checkInIpfsFile(restoredTmpFile, ""); errCheckInIpfs != nil {
		return errCheckInIpfs
	}

	// step 5, remember to unpin restored file
	tFuture := time.Now().Unix() + RestoredFileUnpinInterval
	unpinFileHashQueueAdd(originalFileHash, tFuture)

	log.Info("Ipfs read complete:", originalFileHash)

	return nil
}*/

//链接对方网络，然后再断开
func handleIPFSRead(args ...string) error {

	strArray := strings.Split(args[0], "@#$")

	//	fileHash := strArray[0]
	conAddress := strArray[1]
	shardId := strArray[2]

	log.Info("shardId:", shardId)

	ipfsConnect(conAddress)
	defer ipfsDisconnect(conAddress)
	return nil
}

func HandleIPFSWrite(args ...string) error {
	log.Info("Inside write gorouting")

	strArray := strings.Split(args[0], "@#$")

	fileHash := strArray[0]
	conAddress := strArray[1]
	shardId := strArray[2]
	log.Info("shardId:", shardId)

	// step 1, get the file from ipfs network
	log.Info("Downloading", fileHash, "...")
	tmpFile, errCheckout := checkoutIPFSFile(conAddress, fileHash)
	if errCheckout != nil {
		return errCheckout
	}
	f, _ := tmpFile.Stat()
	stat := FileHashStat{Size: f.Size()}
	updateFileHashStat(fileHash, stat)
	log.Info("Downloaded file hash ", fileHash, "into file", tmpFile.Name())
	defer os.Remove(tmpFile.Name())

	//	// step 2, alter the file content for scs node security
	//	alteredTmpfile := createAlteredTmpFile(tmpFile)
	//	log.Info("Created alterd tmp file to", alteredTmpfile.Name())
	//	defer os.Remove(alteredTmpfile.Name())

	// step 3, check in the altered file into ipfs network
	if errCheckInIpfs := checkInIpfsFile(tmpFile, fileHash); errCheckInIpfs != nil {
		return errCheckInIpfs
	}
	log.Info("Ipfs write complete:", fileHash)

	return nil
}

func HandleIPFSDelete(args ...string) error {
	log.Info("inside delete gorouting")

	strArray := strings.Split(args[0], "@#$")

	fileHash := strArray[0]
	//	conAddress := strArray[1]
	shardId := strArray[2]

	log.Info("shardId:", shardId)

	// false means we actually want to delete a file instead of just clear a cached checkout file
	if err := deleteIpfsFile(fileHash, false); err != nil {
		log.Info("Ipfs delete interrupted:", fileHash)
		return err
	}
	log.Info("Ipfs delete complete:", fileHash)

	if isIpfsFilePined(fileHash) {
		// if there is a cached file, make it unpin now
		tFuture := time.Now().Unix()
		unpinFileHashQueueAdd(fileHash, tFuture)
	}

	if err := deleteFileHashStat(fileHash); err != nil {
		log.Info("Can not remove file hash from redis stat set")
		return err
	}
	log.Info("Removed file hash from redis stat set")

	if err := deleteFileHashMappings(fileHash); err != nil {
		log.Info("Can not remove file hash from redis hash mappings")
		return err
	}
	log.Info("Removed file hash from redis hash mappings")

	return nil
}

func getVerifyOffsetAndLength(fileSize int64, offset int64) ([]int64, bool) {
	if fileSize == 0 {
		return []int64{0, 0, -1, -1}, false
	}

	// determine read start and end
	start := offset % fileSize
	offset = start
	end := (offset + ipfsVerifyReadLength) % fileSize
	// read until the end of the file if end is too large
	if end < start {
		end = fileSize - 1
	}

	startChunkIndex := start / (IpfsChunkSize - 8)
	startChunkOffset := start % (IpfsChunkSize - 8)
	endChunkIndex := end / (IpfsChunkSize - 8)

	ret := []int64{-1, -1, -1, -1}
	if startChunkIndex == endChunkIndex {
		// if read content within one chunk
		// call go file read once with
		mOffset := startChunkIndex*IpfsChunkSize + 8 + startChunkOffset
		mLength := ipfsVerifyReadLength
		if mLength > (fileSize - offset) {
			mLength = fileSize - offset
		}
		ret[0] = mOffset
		ret[1] = mLength
		return ret, false
	} else {
		// if read content across two chunks
		// need to call go file read twice

		// read until end of the first chunk
		mOffset := startChunkIndex*IpfsChunkSize + 8 + startChunkOffset
		mLength := (startChunkIndex+1)*IpfsChunkSize - mOffset
		ret[0] = mOffset
		ret[1] = mLength
		// read content from second chunk
		nOffset := endChunkIndex*IpfsChunkSize + 8
		nLength := ipfsVerifyReadLength - mLength
		ret[2] = nOffset
		ret[3] = nLength
		return ret, true
	}
}

func getFileBytes(originalFileHash string, offset int64) (error, []byte) {

	err, stat := getFileHashStat(originalFileHash)
	if err != nil {
		return err, []byte{}
	}

	//	// deal with file with size 0
	if stat.Size == 0 {
		return nil, []byte{}
	}

	ret, isStraddle := getVerifyOffsetAndLength(stat.Size, offset)
	if isStraddle {
		mOffset := ret[0]
		mLength := ret[1]
		mBytes := readIpfsFileChunk(originalFileHash, mOffset, mLength)
		log.Info(fmt.Sprintf("verify read 1 %s %d %d", originalFileHash, mOffset, mLength), len(mBytes))
		nOffset := ret[2]
		nLength := ret[3]
		nBytes := readIpfsFileChunk(originalFileHash, nOffset, nLength)
		log.Info(fmt.Sprintf("verify read 2 %s %d %d", originalFileHash, nOffset, nLength), len(nBytes))
		return nil, append(mBytes, nBytes...)
	} else {
		mOffset := ret[0]
		mLength := ret[1]
		mBytes := readIpfsFileChunk(originalFileHash, mOffset, mLength)
		log.Info(fmt.Sprintf("verify read 1 %s %d %d", originalFileHash, mOffset, mLength), len(mBytes))
		return nil, mBytes
	}
}

func hashHandler(input1, input2 []byte) (string, error) {

	//MD5算法.
	hash := md5.New()
	_, err := hash.Write(input1)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	result := hash.Sum(input2)
	return fmt.Sprintf("%x", result), nil
}

func getVerifyStr(token *Main, r1, r2 int64) (string, error) {
	fileBig, rErr := token.FileCount(nil)
	if rErr != nil {
		return "", errors.New("vnode request fail")
	}
	fileCount := fileBig.Int64()
	index1 := r1 % fileCount
	index2 := r2 % fileCount

	fileIdBig1, err1 := token.FileList(nil, big.NewInt(index1))
	if err1 != nil {
		return "", errors.New("vnode request fail")
	}
	fileIdBig2, err2 := token.FileList(nil, big.NewInt(index2))
	if err2 != nil {
		return "", errors.New("vnode request fail")
	}

	fileInfo1, err3 := token.FileMapping(nil, fileIdBig1)
	if err3 != nil {
		return "", errors.New("vnode request fail")
	}
	fileInfo2, err4 := token.FileMapping(nil, fileIdBig2)
	if err4 != nil {
		return "", errors.New("vnode request fail")
	}

	err5, bytes1 := getFileBytes(fileInfo1.FileHash, index1)
	if err5 != nil {
		return "", errors.New("vnode request fail")
	}
	err6, bytes2 := getFileBytes(fileInfo2.FileHash, index2)
	if err6 != nil {
		return "", errors.New("vnode request fail")
	}

	return hashHandler(bytes1, bytes2)
}

var HandleIPFSVerify = func(blockH, r1, r2 int64) error {
	// read ipfsVerifyReadLength bytes from the file
	// if file length is less then length byte
	// read the bytes needed in rotate
	log.Info("inside verify gorouting")

	token, err := getMcClient(VnodeIp)
	if err != nil {
		return errors.New("vnode connect fail")
	}

	scsId = getScsId(scsId)
	if scsId == "" {
		return errors.New("scsId get fail")
	}
	nodeInfo, nodeErr := token.NodeMapping(nil, common.HexToAddress(scsId))
	if nodeErr == nil {
		return errors.New("node info get fail")
	}
	//fmt.Println(scsId, nodeInfo)
	if nodeInfo.LastVerifiedBlock.Int64() == blockH {

		str, strType, strErr := getHardDiskInfo()
		if strErr != nil {
			log.Info("getHardDiskInfo fail")
			return strErr
		}

		driveInfo, _ := token.HardDriveMapping(nil, common.HexToAddress(str))
		tmpScsId := driveInfo.ScsId.String()
		if tmpScsId[:2] != "0x" {
			tmpScsId = "0x" + tmpScsId
		}
		if strings.ToLower(tmpScsId) != strings.ToLower(scsId) || strType < driveInfo.Weight.Int64() {
			log.Info("Device information does not match")
			return errors.New("Device information does not match")
		}

		var (
			verifyGroupId, verifyNodeId string
			blockNumber                 int64
			fileHash                    string
			shardId                     int64
		)
		verifyGroupId = common.HexToMoac(nodeInfo.ScsId.Hex() + nodeInfo.LastVerifiedBlock.String()).Hex()
		verifyNodeId = nodeInfo.ScsId.Hex()
		blockNumber = blockH
		var hashErr error
		fileHash, hashErr = getVerifyStr(token, r1, r2)
		if hashErr != nil {
			return errors.New("param splice fail")
		}
		shardId = nodeInfo.ShardId.Int64()
		bytes, bytesErr := submitVerifyTransactionAssembleByte(verifyGroupId, verifyNodeId, blockNumber, fileHash, shardId)
		if bytesErr != nil {
			return errors.New("param splice fail")
		}
		sendTxOperate(bytes)
	} else {
		//给别人投票
		var (
			//			verifyGroupId string
			verifyNodeId string
			votingNodeId string
			blockNumber  int64
			fileHash     string
			shardId      int64
		)

		votingNodeId = scsId
		blockNumber = blockH
		var hashErr error
		fileHash, hashErr = getVerifyStr(token, r1, r2)
		if hashErr != nil {
			return errors.New("param splice fail")
		}
		shardId = nodeInfo.ShardId.Int64()

		var shardSize int = 10
		for i := 0; i < shardSize; i++ {
			comAddr, comErr := token.ShardNodeList(nil, big.NewInt(shardId), big.NewInt(int64(i)))
			if comErr != nil {
				return errors.New("param splice fail")
			}
			if comAddr.Hex() == "0x0000000000000000000000000000000000000000" {
				break
			} else {
				var flag bool = true
				vGroupId := common.HexToMoac(comAddr.Hex() + nodeInfo.LastVerifiedBlock.String()).Hex()
				for j := 0; j < shardSize; j++ {
					vs, vsErr := token.VerifyGroupMapping(nil, common.HexToAddress(vGroupId), big.NewInt(int64(j)))
					if vsErr != nil {
						flag = false
						break
					}
					if vs.ScsId.Hex() == scsId || vs.VerifyNodeId.Hex() == scsId {
						flag = false
						break
					}
				}
				if flag {
					//发送投票交易
					bytes, bytesErr := voteVerifyTransactionAssembleByte(vGroupId, verifyNodeId, votingNodeId, blockNumber, fileHash, shardId)
					if bytesErr != nil {
						return errors.New("param splice fail")
					}
					sendTxOperate(bytes)
				}
			}
		}
	}

	return nil
}

//IPFS管理子链操作
func sendTxOperate(data []byte) (string, error) {

	var (
		netType              int    = 101
		via                  string = "0xD814F2ac2c4cA49b33066582E4e97EBae02F2aB9"
		formAddress          string = getScsId(scsId)
		formKeystore         string = ""
		formKeystorePassword string = "moacscsofflineaccountpwd"
	)

	jsonBytes, pwErr := getScsJsonStr(formKeystorePassword)
	if pwErr != nil {
		return "", errors.New("pw error")
	}
	formKeystore = string(jsonBytes)

	bytes := data

	var tmpRpcClient *Chain3Go.RpcClient
	tmpRpcClient = Chain3Go.NewRpcClient("http://"+SubChainIp, netType) //子链
	//	nonce, nonErr := tmpRpcClient.Mc().SCS_getNonce(to, formAddress)
	nonce, nonErr := tmpRpcClient.Mc().ScsRPCMethod_GetNonce(contractAddr, formAddress)
	if nonErr == nil {
		var rpcClient *Chain3Go.RpcClient
		rpcClient = Chain3Go.NewRpcClient(VnodeIp, netType) //子链
		signStr, err := Chain3Go.SubChainTxSign(
			via,
			netType,
			formKeystore,
			formKeystorePassword,
			formAddress,
			contractAddr,
			big.NewInt(0),
			big.NewInt(0),
			big.NewInt(0),
			1,
			bytes,
			uint64(nonce),
		)
		if err == nil {
			return rpcClient.Mc().MC_sendRawTransaction(signStr)
		}
		return "", err
	}
	return "", nonErr
}

func handleIPFSProxyRead(args ...string) error {

	mProxyReadRequest := args[0]

	proxyReadRequest := new(ProxyReadRequest)
	json.Unmarshal([]byte(mProxyReadRequest), &proxyReadRequest)

	originalFileHash := proxyReadRequest.Filehash
	proxyAddress := proxyReadRequest.ProxyAddress
	log.Info("In handle ipfs proxy read: ", originalFileHash, proxyAddress)

	// call proxy to restore altered file locally on proxy
	restoreToLocalURL := fmt.Sprintf(
		"%s/%s?%s",
		fmt.Sprintf("http://%s", proxyAddress),
		"vnode/restoreToLocal",
		fmt.Sprintf("file_hash=%s", originalFileHash),
	)

	resp, err := http.Get(restoreToLocalURL)
	if err != nil || resp.StatusCode != 200 {
		log.Error("Can not restore file, original hash ", originalFileHash)
	}

	return nil
}

func handleIPFSProxyWrite(args ...string) error {

	mProxyWriteRequest := args[0]

	proxyWriteRequest := new(ProxyWriteRequest)
	json.Unmarshal([]byte(mProxyWriteRequest), &proxyWriteRequest)

	originalFileHash := proxyWriteRequest.Filehash
	proxyAddress := proxyWriteRequest.ProxyAddress
	log.Info("In handle ipfs proxy read: ", originalFileHash, proxyAddress)

	// step 1, direct download file from proxy
	directDownloadURL := fmt.Sprintf(
		"%s/%s?%s",
		fmt.Sprintf("http://%s", proxyAddress),
		"vnode/directDownload",
		fmt.Sprintf("file_hash=%s", originalFileHash),
	)

	log.Info("Downloading", originalFileHash, "...")
	resp, _ := http.Get(directDownloadURL)
	tmpFile, _ := ioutil.TempFile("", IpfsPrefix)
	io.Copy(tmpFile, resp.Body)
	f, _ := tmpFile.Stat()
	stat := FileHashStat{Size: f.Size()}
	updateFileHashStat(originalFileHash, stat)
	log.Info("Downloaded file hash ", originalFileHash, "into file", tmpFile.Name())
	defer os.Remove(tmpFile.Name())

	// step 2, alter the file content for scs node security
	alteredTmpfile := createAlteredTmpFile(tmpFile)
	log.Info("Created alterd tmp file to", alteredTmpfile.Name())
	defer os.Remove(alteredTmpfile.Name())

	// step 3, check in the altered file into ipfs network
	if errCheckInIpfs := checkInIpfsFile(alteredTmpfile, originalFileHash); errCheckInIpfs != nil {
		return errCheckInIpfs
	}
	log.Info("Ipfs write complete:", originalFileHash)

	return nil
}

// Handle fs proxy read
func restoreToLocalHandler(w http.ResponseWriter, r *http.Request) {
	// checkout the altered file, restore it and add the restored file to ipfs
	originalFileHash := r.URL.Query().Get("file_hash")
	restoreToLocalRequest := RestoreToLocalRequest{Filehash: originalFileHash}
	if errs := validator.Validate(restoreToLocalRequest); errs != nil {
		http.Error(w, fmt.Sprintf("Invalid file_hash parameter: %v", errs), 400)
		return
	}

	go handleRestoreToLocal(restoreToLocalRequest.Filehash)
}

// Handle fs proxy write
func directDownloadHandler(w http.ResponseWriter, r *http.Request) {
	// checkout the altered file, restore it and add the restored file to ipfs
	originalFileHash := r.URL.Query().Get("file_hash")
	directDownloadRequest := DirectDownloadRequest{Filehash: originalFileHash}
	if errs := validator.Validate(directDownloadRequest); errs != nil {
		http.Error(w, fmt.Sprintf("Invalid file_hash parameter: %v", errs), 400)
		return
	}

	url := fmt.Sprintf(
		"%s/%s?%s",
		fmt.Sprintf("http://%s", ipfsHostPort),
		"api/v0/cat",
		fmt.Sprintf("arg=/ipfs/%s", originalFileHash),
	)
	log.Infof("Direct download file %s with ipfs call: %s", originalFileHash, url)
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("%v", err)
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
	return
}

func handleRestoreToLocal(alteredFileHash string) {

	// step 1
	alteredTmpfile, errCheckout := checkoutIPFSFile("", alteredFileHash)
	if errCheckout != nil {
		log.Errorf("Check out ipfs file error: %v, %s", errCheckout, alteredFileHash)
	}
	log.Infof("Checkout altered file hash %s with tmp file %s", alteredFileHash, alteredTmpfile.Name())
	defer os.Remove(alteredTmpfile.Name())

	// step 2
	restoredTmpFile := createRestoredTmpFile(alteredTmpfile)
	defer os.Remove(restoredTmpFile.Name())

	// step 3, originalFileHash set to "" means it's not to checkin new written file
	if _, errAddIpfsFile := AddIpfsFile(restoredTmpFile); errAddIpfsFile != nil {
		log.Errorf("Add restored file to ipfs error: %v, %s", errAddIpfsFile, restoredTmpFile.Name())
	}
}

func initThrottledQueues() {
	q2cMapping = make(map[string](chan string))
	queueNames := []string{
		IpfsReadQueueName,
		IpfsWriteQueueName,
		IpfsDeleteQueueName,
		IpfsProxyWriteQueueName,
		IpfsProxyReadQueueName,
	}

	// this create multiple goroutine constantly push new tasks into queue
	// push rate is throttled by queueConcurrency(=10)
	for _, queueName := range queueNames {
		c := make(chan string, queueConcurrency)
		q2cMapping[queueName] = c
		go func(queueName string, _c chan string) {
			log.Info("q2c mapping", queueName, "->", _c)
			for {
				// timeout is zero, so this will block on new tasks
				newTask, err := getTaskFromQueueBlock(queueName)
				if err != nil {
					log.Error("Can't connect to redis", redisHostPort, err)
					time.Sleep(time.Duration(1) * time.Second)
				} else {
					log.Info("Enqueue new task in", queueName, newTask)
					_c <- newTask
				}
			}
		}(queueName, c)
	}
}

func initThrottledQueueHandlers() {
	// init handler mapping
	handlerMapping = make(map[string]func(...string) error)
	handlerMapping[IpfsReadQueueName] = handleIPFSRead
	handlerMapping[IpfsWriteQueueName] = HandleIPFSWrite
	handlerMapping[IpfsDeleteQueueName] = HandleIPFSDelete
	handlerMapping[IpfsProxyReadQueueName] = handleIPFSProxyRead
	handlerMapping[IpfsProxyWriteQueueName] = handleIPFSProxyWrite

	queueNames := []string{
		IpfsReadQueueName,
		IpfsWriteQueueName,
		IpfsDeleteQueueName,
		IpfsProxyWriteQueueName,
		IpfsProxyReadQueueName,
	}
	for _, queueName := range queueNames {
		go func(queueName string) {
			c := q2cMapping[queueName]
			log.Info("c from q:", queueName, c)
			for {
				// block on next task
				task := <-c
				log.Info(fmt.Sprintf("call %s handler with %v", queueName, task))
				err := handlerMapping[queueName](task)
				if err != nil {
					log.Info("Can't handl task, error:", err)
				}
			}
		}(queueName)
	}
}

func runIpfsGC() {
	ipfsRepoGc()
}

//func initGCWorker() {
//	// run ipfs repo gc periodically
//	go func() {
//		for {
//			runIpfsGC()
//			time.Sleep(time.Duration(ipfsGCInterval) * time.Second)
//		}
//	}()
//}

func runIpfsUnpin() {
	tNow := time.Now().Unix()
	fileHashes, err := unpinFileHashQueueRange(int64(0), tNow)
	if err == nil {
		for _, fileHash := range fileHashes {
			// true means it's a delete to clear restored file
			err := deleteIpfsFile(fileHash, true)
			if err != nil {
				log.Info("Failed to clear cache for file", fileHash)
			} else {
				log.Info("Unpined cache for file", fileHash)
				// remove the file hash from unpin queue
				unpinFileHashQueueRemove(fileHash)
			}
		}
	} else {
		log.Info("Failed getting unpin queue.")
	}
}

func initUnpinWorker() {
	// run checkout file clearance periodically
	go func() {
		for {
			runIpfsUnpin()
			time.Sleep(time.Duration(ipfsUnpinInterval) * time.Second)
		}
	}()
}

//提交硬盘信息
func subHardDiskInfo() {

	hardDiskId, hardDiskSize, err := getHardDiskInfo()
	if err == nil { //提交交易
		dBytes, dErr := updateHardDriveAssembleByte(common.HexToMoac(hardDiskId).Hex(), scsId, hardDiskSize)
		if dErr == nil {
			_, hashErr := sendTxOperate(dBytes)
			if hashErr == nil {
				log.Info("硬盘信息已经提交！")
			} else {
				log.Error("hashErr:", hashErr)
			}
		} else {
			log.Error("dErr:", dErr)
		}
	} else {
		log.Error("err:", err)
	}
}

//获取硬盘信息
func getHardDiskInfo() (string, int64, error) {
	cmdStr1 := `ls -l /dev/disk/by-uuid/ | grep ` + memDeviceDir + ` | awk '{print $9}'` //获取硬盘ID
	var cmdStr2 string = ""
	cmdStr3 := `fdisk -l | grep ` + memDeviceDir + ` | awk '{print $2}'` //获取硬盘大小

	outByte1, err1 := execLinuxCommand("/bin/sh", "-c", cmdStr1)
	if err1 == nil {
		outByte3, err3 := execLinuxCommand("/bin/sh", "-c", cmdStr3)
		if err3 != nil {
			return "", 0, err3
		}
		if string(outByte3) == "*" {
			cmdStr2 = `fdisk -l | grep ` + memDeviceDir + ` | awk '{print $5}'`
		} else {
			cmdStr2 = `fdisk -l | grep ` + memDeviceDir + ` | awk '{print $4}'`
		}
		outByte2, err2 := execLinuxCommand("/bin/sh", "-c", cmdStr2)
		if err2 == nil {
			if (string(outByte2))[:7] != "WARNING" {
				hardDiskId := strings.Replace(strings.Replace(string(outByte1), "\n", "", -1), "-", "", -1)
				hardDiskSize, err := strconv.ParseInt(strings.Replace(strings.Replace(string(outByte2), "\n", "", -1), "+", "", -1), 10, 64)
				if err == nil {
					return hardDiskId, int64(float64(hardDiskSize) / 1000 / 1000 / 1000), nil
				} else {
					return "", 0, err
				}
			} else {
				return "", 0, errors.New(string(outByte2))
			}
		} else {
			return "", 0, err2
		}
	} else {
		return "", 0, err1
	}
}

//执行Linux命令
func execLinuxCommand(a, b, cmdStr string) ([]byte, error) {

	cmd := exec.Command(a, b, cmdStr)
	return cmd.Output()
}
