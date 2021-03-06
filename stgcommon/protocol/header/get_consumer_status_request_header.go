package header

// getConsumerStatus 获得消费者状态的请求头
// Author rongzhihong
// Since 2017/9/19
type GetConsumerStatusRequestHeader struct {
	Topic      string `json:"topic"`
	Group      string `json:"group"`
	ClientAddr string `json:"clientAddr"`
}

func (header *GetConsumerStatusRequestHeader) CheckFields() error {
	return nil
}

// NewGetConsumerStatusRequestHeader 初始化
// Author: tianyuliang
// Since: 2017/11/6
func NewGetConsumerStatusRequestHeader(topic, group, clientAddr string) *GetConsumerStatusRequestHeader {
	consumerStatusRequestHeader := &GetConsumerStatusRequestHeader{
		Topic:      topic,
		Group:      group,
		ClientAddr: clientAddr,
	}
	return consumerStatusRequestHeader
}
