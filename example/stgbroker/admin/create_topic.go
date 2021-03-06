package main

import (
	"fmt"
	"os"

	"github.com/ttstringiot/golangiot/stgbroker"
	"github.com/ttstringiot/golangiot/stgcommon"
	"github.com/ttstringiot/golangiot/stgcommon/logger"
	code "github.com/ttstringiot/golangiot/stgcommon/protocol"
	"github.com/ttstringiot/golangiot/stgcommon/protocol/header"
	"github.com/ttstringiot/golangiot/stgnet/netm"
	"github.com/ttstringiot/golangiot/stgnet/protocol"
)

var (
	request     *protocol.RemotingCommand
	response    *protocol.RemotingCommand
	err         error
	newTopic    = "cloudzone123"
	ctx         netm.Context
	localPort   = 10925
	localIp     = "127.0.0.1"
	namesrvAddr = "127.0.0.1:9876"
)

func main() {

	os.Setenv(stgcommon.NAMESRV_ADDR_ENV, namesrvAddr)

	requestHeader := initRequestHeader(newTopic)
	request = protocol.CreateRequestCommand(code.UPDATE_AND_CREATE_TOPIC, requestHeader)
	request.EncodeHeader() // 将requestHeader的值写入到ExtFields中
	logger.Infof("request ---> %s", request.ToString())

	ctx = CreateDefaultContext()
	controller := stgbroker.CreateBrokerController()
	controller.Initialize()
	admin := stgbroker.NewAdminBrokerProcessor(controller)
	logger.Infof("admin processor ready success")

	response, err = admin.ProcessRequest(ctx, request)
	logger.Infof("admin processor send ok")

	if err != nil {
		logger.Errorf("sync response UPDATE_AND_CREATE_TOPIC failed. err: %s", err.Error())
		return
	}
	if response == nil {
		logger.Errorf("sync response UPDATE_AND_CREATE_TOPIC failed. err: response is nil")
		return
	}
	logger.Infof("response --> %s", response.ToString())

	if response.Code == code.SUCCESS {
		format := "sync response UPDATE_AND_CREATE_TOPIC success. topic=%s"
		logger.Infof(format, newTopic)
		return
	}
	format := "sync handle UPDATE_AND_CREATE_TOPIC failed. code=%d, remark=%s"
	logger.Infof(format, response.Code, response.Remark)
}

func initRequestHeader(topic string) *header.CreateTopicRequestHeader {
	requestHeader := &header.CreateTopicRequestHeader{}
	requestHeader.Topic = topic
	requestHeader.DefaultTopic = fmt.Sprintf("%s%d", topic, stgcommon.GetCurrentTimeMillis())
	requestHeader.ReadQueueNums = 8
	requestHeader.WriteQueueNums = 6
	requestHeader.Perm = 4
	requestHeader.TopicFilterType = stgcommon.SINGLE_TAG
	requestHeader.Order = false
	requestHeader.TopicSysFlag = 0
	return requestHeader
}

func CreateDefaultContext() netm.Context {
	var remoteContext netm.Context

	bootstrap := netm.NewBootstrap()
	bootstrap.Bind(localIp, localPort)
	bootstrap.RegisterHandler(func(buffer []byte, ctx netm.Context) {
		remoteContext = ctx
	})
	go func() {
		bootstrap.Sync()
	}()

	clientBootstrap := netm.NewBootstrap()
	err := clientBootstrap.Connect(localIp, localPort)
	if err != nil {
		logger.Errorf("clientBootstrap.Connect() err: %s", err.Error())
	}
	return clientBootstrap.Contexts()[0]
}
