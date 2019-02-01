package main

import (
	//	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"strconv"

	validator "gopkg.in/validator.v2"

	"stormcatcher/lib/common/hexutil"
)

var log = setupLogging(true)

func init() {

	flag.StringVar(&listenAddressAndPort, "listen-host-port", "127.0.0.1:18080", "host:port, e.g. 127.0.0.1:18080")
	flag.StringVar(&redisHostPort, "redis-host-port", "localhost:6379", "host:port, e.g. 127.0.0.1:6379")
	flag.StringVar(&ipfsHostPort, "ipfs-host-port", "localhost:5001", "host:port, e.g. 127.0.0.1:5001")
	flag.StringVar(&contractAddr, "subChain-contract-address", "", "contractAddress, e.g. 0x")
	flag.StringVar(&SubChainIp, "subChain-address", "localhost:50068/rpc", "subChainAddress, e.g. 127.0.0.1:50068")
}

// AccessType is the ipfs file access type
type AccessType int

// Read=0, Write=1, Delete=2, Verify=3
const (
	Read   AccessType = 0
	Write  AccessType = 1
	Remove AccessType = 2
	Verify AccessType = 3
)

// FileHash is the ipfs hash and the type
type FileHash struct {
	hashString string
	hashType   string
}

type ReadRequest struct {
	Filehash string `validate:"len=46,regexp=^[a-zA-Z0-9]*$"`
	IpfsId   string `validate:"regexp=^[a-zA-Z0-9]*$"`
	ShardId  string `validate:"regexp=^[0-9]*$"`
}

type ProxyReadRequest struct {
	Filehash     string `validate:"len=46,regexp=^[a-zA-Z0-9]*$"`
	ProxyAddress string
}

type WriteRequest struct {
	Filehash string `validate:"len=46,regexp=^[a-zA-Z0-9]*$"`
	IpfsId   string `validate:"regexp=^[a-zA-Z0-9]*$"`
	ShardId  string `validate:"regexp=^[0-9]*$"`
}

type ProxyWriteRequest struct {
	Filehash     string `validate:"len=46,regexp=^[a-zA-Z0-9]*$"`
	ProxyAddress string
}

type DeleteRequest struct {
	Filehash string `validate:"len=46,regexp=^[a-zA-Z0-9]*$"`
	IpfsId   string `validate:"regexp=^[a-zA-Z0-9]*$"`
	ShardId  string `validate:"regexp=^[0-9]*$"`
}

type VerifyRequest struct {
	//	Filehash string `validate:"len=46,regexp=^[a-zA-Z0-9]*$"`
	//	Offset   string `validate:"regexp=^[0-9]*$"`

	Block_int string `validate:"regexp=^[0-9]*$"`
	Random_1  string `validate:"regexp=^[0-9]*$"`
	Random_2  string `validate:"regexp=^[0-9]*$"`
}

type RestoreToLocalRequest struct {
	Filehash string `validate:"len=46,regexp=^[a-zA-Z0-9]*$"`
}

type DirectDownloadRequest struct {
	Filehash string `validate:"len=46,regexp=^[a-zA-Z0-9]*$"`
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	// sample query:
	// curl "http://127.0.0.1:18080/ipfs/read?file_hash=QmTor1GsqZQwJdFoTYjAdEEjXDZgYDm1oc3Lj8waHUKRFN"
	fileHash := r.URL.Query().Get("file_hash")
	readRequest := ReadRequest{Filehash: fileHash}
	if errs := validator.Validate(readRequest); errs != nil {
		http.Error(w, fmt.Sprintf("Invalid file_hash parameter: %v", errs), 400)
		return
	}

	if err := AddTaskToQueue(IpfsReadQueueName, fileHash); err != nil {
		http.Error(w, "Can not enqueue read request.", 500)
		return
	}
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	// sample query:
	// curl "http://127.0.0.1:18080/ipfs/write?file_hash=QmTor1GsqZQwJdFoTYjAdEEjXDZgYDm1oc3Lj8waHUKRFN"
	fileHash := r.URL.Query().Get("file_hash")
	conAddress := r.URL.Query().Get("connect_address")
	writeRequest := WriteRequest{Filehash: fileHash}
	if errs := validator.Validate(writeRequest); errs != nil {
		http.Error(w, fmt.Sprintf("Invalid file_hash parameter: %v", errs), 400)
		return
	}

	if err := AddTaskToQueue(IpfsWriteQueueName, fileHash, conAddress); err != nil {
		http.Error(w, "Can not enqueue write request.", 500)
		return
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	// sample query:
	// curl "http://127.0.0.1:18080/ipfs/delete?file_hash=QmTor1GsqZQwJdFoTYjAdEEjXDZgYDm1oc3Lj8waHUKRFN"
	fileHash := r.URL.Query().Get("file_hash")
	deleteRequest := DeleteRequest{Filehash: fileHash}
	if errs := validator.Validate(deleteRequest); errs != nil {
		http.Error(w, fmt.Sprintf("Invalid file_hash parameter: %v", errs), 400)
		return
	}

	if err := AddTaskToQueue(IpfsDeleteQueueName, fileHash); err != nil {
		http.Error(w, "Can not enqueue delete request.", 500)
		return
	}
}

func verifyHandler(w http.ResponseWriter, r *http.Request) {
	// sample query:
	// curl "http://127.0.0.1:18080/verify?file_hash=QmTor1GsqZQwJdFoTYjAdEEjXDZgYDm1oc3Lj8waHUKRFN&offset=0"
	q := r.URL.Query()
	block_int := q.Get("block_int")
	random_1 := q.Get("random_1")
	random_2 := q.Get("random_2")

	verifyRequest := VerifyRequest{Block_int: block_int, Random_1: random_1, Random_2: random_2}
	if errs := validator.Validate(verifyRequest); errs != nil {
		http.Error(w, fmt.Sprintf("Invalid verify parameter: %v", errs), 400)
		return
	}

	blockH, err1 := strconv.ParseInt(verifyRequest.Block_int, 10, 64)
	if err1 != nil {
		http.Error(w, "Can not parse offset value.", 400)
		return
	}

	r1, err2 := strconv.ParseInt(verifyRequest.Random_1, 10, 64)
	if err2 != nil {
		http.Error(w, "Can not parse offset value.", 400)
		return
	}

	r2, err3 := strconv.ParseInt(verifyRequest.Random_2, 10, 64)
	if err3 != nil {
		http.Error(w, "Can not parse offset value.", 400)
		return
	}

	err := HandleIPFSVerify(blockH, r1, r2)
	if err != nil {
		http.Error(w, err.Error(), 404)
	}
	w.Write([]byte(""))
}

func catchHandler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")

	//文件hash，shardId，IPFS_ID
	switch fileHash, ipfsId, aType, shardId, _ := decodeInput(input); AccessType(aType) {
	case Remove:
		shardIdStr := strconv.FormatInt(shardId, 10)
		deleteRequest := DeleteRequest{Filehash: fileHash, IpfsId: ipfsId, ShardId: shardIdStr}
		if errs := validator.Validate(deleteRequest); errs != nil {
			http.Error(w, fmt.Sprintf("Invalid file_hash parameter: %v", errs), 400)
			return
		}

		if err := AddTaskToQueue(IpfsDeleteQueueName, fileHash, ipfsId, shardIdStr); err != nil {
			http.Error(w, "Can not enqueue delete request.", 500)
			return
		}
	case Write:
		shardIdStr := strconv.FormatInt(shardId, 10)
		writeRequest := WriteRequest{Filehash: fileHash, IpfsId: ipfsId, ShardId: shardIdStr}
		if errs := validator.Validate(writeRequest); errs != nil {
			http.Error(w, fmt.Sprintf("Invalid file_hash parameter: %v", errs), 400)
			return
		}

		if err := AddTaskToQueue(IpfsWriteQueueName, fileHash, ipfsId, shardIdStr); err != nil {
			http.Error(w, "Can not enqueue write request.", 500)
			return
		}
	case Read:
		shardIdStr := strconv.FormatInt(shardId, 10)
		readRequest := ReadRequest{Filehash: fileHash, IpfsId: ipfsId, ShardId: shardIdStr}
		if errs := validator.Validate(readRequest); errs != nil {
			http.Error(w, fmt.Sprintf("Invalid file_hash parameter: %v", errs), 400)
			return
		}

		if err := AddTaskToQueue(IpfsReadQueueName, fileHash, ipfsId, shardIdStr); err != nil {
			http.Error(w, "Can not enqueue read request.", 500)
			return
		}
	}
}

// DecodeIpfsParams decodes input byte array and returns the value of ipfsFile parameters.
func decodeInput(hexString string) (string, string, int64, int64, error) {
	invalidIpfsHasErr := errors.New("invalid ipfs hash format")
	log.Debugf("hexString, |%v|", hexString)
	hexLen := len(hexString)
	if hexLen > 11*64 && hexLen < 12*64 {
		log.Errorf("Ipfs Input Code is too short. %v", hexString)
		return "", "", -1, -1, invalidIpfsHasErr
	}
	log.Debugf("hexString, %v", hexString)

	funcCode := hexString[:8] //get function code
	if funcCode != "633fb659" {
		log.Errorf("Invalid ipfs function call. %v", funcCode)
		return "", "", -1, -1, errors.New("invalid function call.")
	}
	aType, aTypeErr := strconv.ParseInt(hexString[10+64:10+64*2], 16, 64)

	shardId, shardIdErr := strconv.ParseInt(hexString[10+64*2:10+64*3], 16, 64)

	fileHashBytes, fileHashErr := hexutil.Decode("0x" + hexString[10+5*64:10+7*64])

	ipfsIdBytes, ipfsIdErr := hexutil.Decode("0x" + hexString[10+8*64:10+11*64])

	if aTypeErr != nil && shardIdErr != nil && fileHashErr != nil && ipfsIdErr != nil {
		return "", "", -1, -1, errors.New("input parameter error")
	}

	fileHash := string(fileHashBytes)
	ipfsId := string(ipfsIdBytes)

	log.Debugf("fileHash, %v", fileHash)
	if fileHash[0:2] != "Qm" {
		log.Debugf("hex |%v|", hexString[0:2])
		log.Errorf("Ipfs Hash is in wrong format. %v", fileHash)
		return "", "", -1, -1, invalidIpfsHasErr
	}
	return fileHash, ipfsId, aType, shardId, nil
}

/*func proxyReadHandler(w http.ResponseWriter, r *http.Request) {
	// read through proxy(vnode) server
	originalFileHash := r.URL.Query().Get("file_hash")
	proxyAddress := r.URL.Query().Get("proxy_address")
	proxyReadRequest := ProxyReadRequest{
		Filehash:     originalFileHash,
		ProxyAddress: proxyAddress,
	}

	if errs := validator.Validate(proxyReadRequest); errs != nil {
		http.Error(w, fmt.Sprintf("Invalid parameter: %v", errs), 400)
		return
	}

	mProxyReadRequest, _ := json.Marshal(proxyReadRequest)
	if err := AddTaskToQueue(IpfsProxyReadQueueName, string(mProxyReadRequest)); err != nil {
		http.Error(w, "Can not enqueue proxy read request.", 500)
		return
	}
}*/

/*func proxyWriteHandler(w http.ResponseWriter, r *http.Request) {
	// write through proxy(vnode) server
	originalFileHash := r.URL.Query().Get("file_hash")
	proxyAddress := r.URL.Query().Get("proxy_address")
	proxyWriteRequest := ProxyWriteRequest{
		Filehash:     originalFileHash,
		ProxyAddress: proxyAddress,
	}

	if errs := validator.Validate(proxyWriteRequest); errs != nil {
		http.Error(w, fmt.Sprintf("Invalid parameter: %v", errs), 400)
		return
	}

	mProxyWriteRequest, _ := json.Marshal(proxyWriteRequest)
	if err := AddTaskToQueue(IpfsProxyWriteQueueName, string(mProxyWriteRequest)); err != nil {
		http.Error(w, "Can not enqueue proxy write request.", 500)
		return
	}
}*/

func main() {

	flag.Parse()

	// verify does not use queue and should return immediately the result
	http.HandleFunc("/verify", verifyHandler)
	// catch will determine what routing it should do.
	http.HandleFunc("/catch", catchHandler)

	// routing
	//	http.HandleFunc("/ipfs/read", readHandler)
	//	http.HandleFunc("/ipfs/write", writeHandler)
	//	http.HandleFunc("/ipfs/delete", deleteHandler)

	// proxy read/write
	//	http.HandleFunc("/proxy/read", proxyReadHandler)
	//	http.HandleFunc("/proxy/write", proxyWriteHandler)

	// vnode proxy endpoints
	http.HandleFunc("/vnode/restoreToLocal", restoreToLocalHandler)
	http.HandleFunc("/vnode/directDownload", directDownloadHandler)

	// print config
	log.Info("StormCatcher:", listenAddressAndPort)
	log.Info("Redis:", redisHostPort)
	log.Info("Ipfs:", ipfsHostPort)
	log.Info("contractAddr:", contractAddr)

	// init redis
	initRedisClient()
	// init queue and queue handler
	initThrottledQueues()
	log.Info("Channel setup done. Storm Catcher is ready!")
	initThrottledQueueHandlers()

	//init gc routine
	initGCWorker()

	//init unpin worker
	initUnpinWorker()

	// start server
	log.Critical(http.ListenAndServe(listenAddressAndPort, nil))
}
