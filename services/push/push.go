package push

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) Push(request *PushRequest) (response *PushResponse, err error) {
	response = CreatePushResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) PushWithChan(request *PushRequest) (<-chan *PushResponse, <-chan error) {
	responseChan := make(chan *PushResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Push(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) PushWithCallback(request *PushRequest, callback func(response *PushResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PushResponse
		var err error
		defer close(result)
		response, err = client.Push(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type PushRequest struct {
	*requests.RpcRequest
	AndroidXiaoMiNotifyTitle       string `position:"Query" name:"AndroidXiaoMiNotifyTitle"`
	AndroidNotificationBarPriority string `position:"Query" name:"AndroidNotificationBarPriority"`
	Body                           string `position:"Query" name:"Body"`
	IOSBadgeAutoIncrement          string `position:"Query" name:"iOSBadgeAutoIncrement"`
	AndroidOpenUrl                 string `position:"Query" name:"AndroidOpenUrl"`
	TargetValue                    string `position:"Query" name:"TargetValue"`
	AndroidPopupBody               string `position:"Query" name:"AndroidPopupBody"`
	SendSpeed                      string `position:"Query" name:"SendSpeed"`
	AndroidPopupTitle              string `position:"Query" name:"AndroidPopupTitle"`
	AndroidExtParameters           string `position:"Query" name:"AndroidExtParameters"`
	AppKey                         string `position:"Query" name:"AppKey"`
	AndroidXiaoMiActivity          string `position:"Query" name:"AndroidXiaoMiActivity"`
	SmsDelaySecs                   string `position:"Query" name:"SmsDelaySecs"`
	IOSBadge                       string `position:"Query" name:"iOSBadge"`
	IOSExtParameters               string `position:"Query" name:"iOSExtParameters"`
	IOSRemind                      string `position:"Query" name:"iOSRemind"`
	PushType                       string `position:"Query" name:"PushType"`
	IOSMusic                       string `position:"Query" name:"iOSMusic"`
	IOSRemindBody                  string `position:"Query" name:"iOSRemindBody"`
	PushTime                       string `position:"Query" name:"PushTime"`
	SmsSignName                    string `position:"Query" name:"SmsSignName"`
	Title                          string `position:"Query" name:"Title"`
	SmsTemplateName                string `position:"Query" name:"SmsTemplateName"`
	IOSSubtitle                    string `position:"Query" name:"iOSSubtitle"`
	JobKey                         string `position:"Query" name:"JobKey"`
	AndroidActivity                string `position:"Query" name:"AndroidActivity"`
	ExpireTime                     string `position:"Query" name:"ExpireTime"`
	SmsSendPolicy                  string `position:"Query" name:"SmsSendPolicy"`
	IOSSilentNotification          string `position:"Query" name:"iOSSilentNotification"`
	AndroidPopupActivity           string `position:"Query" name:"AndroidPopupActivity"`
	AndroidNotificationBarType     string `position:"Query" name:"AndroidNotificationBarType"`
	AndroidOpenType                string `position:"Query" name:"AndroidOpenType"`
	AndroidNotifyType              string `position:"Query" name:"AndroidNotifyType"`
	BatchNumber                    string `position:"Query" name:"BatchNumber"`
	IOSNotificationCategory        string `position:"Query" name:"iOSNotificationCategory"`
	IOSApnsEnv                     string `position:"Query" name:"iOSApnsEnv"`
	SmsParams                      string `position:"Query" name:"SmsParams"`
	StoreOffline                   string `position:"Query" name:"StoreOffline"`
	IOSMutableContent              string `position:"Query" name:"iOSMutableContent"`
	AndroidXiaoMiNotifyBody        string `position:"Query" name:"AndroidXiaoMiNotifyBody"`
	AndroidRemind                  string `position:"Query" name:"AndroidRemind"`
	AndroidMusic                   string `position:"Query" name:"AndroidMusic"`
	Target                         string `position:"Query" name:"Target"`
	DeviceType                     string `position:"Query" name:"DeviceType"`
}

type PushResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	MessageId string `json:"MessageId" xml:"MessageId"`
}

func CreatePushRequest() (request *PushRequest) {
	request = &PushRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "Push", "", "")
	return
}

func CreatePushResponse() (response *PushResponse) {
	response = &PushResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
