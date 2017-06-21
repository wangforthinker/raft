package controller

type Controller interface {
	//put data
	Put(key string, data []byte, opt* DataOpt)
	//get by key and version, if version == 0, means no version
	Get(key string, version int64)
}

type MyController struct {

}
