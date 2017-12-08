package dm

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

func (client *Client) GetSenderAddressList(request *GetSenderAddressListRequest) (response *GetSenderAddressListResponse, err error) {
	response = CreateGetSenderAddressListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetSenderAddressListWithChan(request *GetSenderAddressListRequest) (<-chan *GetSenderAddressListResponse, <-chan error) {
	responseChan := make(chan *GetSenderAddressListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSenderAddressList(request)
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

func (client *Client) GetSenderAddressListWithCallback(request *GetSenderAddressListRequest, callback func(response *GetSenderAddressListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSenderAddressListResponse
		var err error
		defer close(result)
		response, err = client.GetSenderAddressList(request)
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

type GetSenderAddressListRequest struct {
	*requests.RpcRequest
	Total                string `position:"Query" name:"Total"`
	PageSize             string `position:"Query" name:"PageSize"`
	NotifyUrl            string `position:"Query" name:"NotifyUrl"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Keyword              string `position:"Query" name:"Keyword"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	PageNo               string `position:"Query" name:"PageNo"`
	Offset               string `position:"Query" name:"Offset"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type GetSenderAddressListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Total     int    `json:"Total" xml:"Total"`
	PageNo    int    `json:"PageNo" xml:"PageNo"`
	PageSize  int    `json:"PageSize" xml:"PageSize"`
	Data      struct {
		SenderAddressNotificationInfo []struct {
			Region          string `json:"Region" xml:"Region"`
			UpdateTime      string `json:"UpdateTime" xml:"UpdateTime"`
			Status          string `json:"Status" xml:"Status"`
			SenderAddressId string `json:"SenderAddressId" xml:"SenderAddressId"`
			SenderAddress   string `json:"SenderAddress" xml:"SenderAddress"`
		} `json:"senderAddressNotificationInfo" xml:"senderAddressNotificationInfo"`
	} `json:"data" xml:"data"`
}

func CreateGetSenderAddressListRequest() (request *GetSenderAddressListRequest) {
	request = &GetSenderAddressListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "GetSenderAddressList", "", "")
	return
}

func CreateGetSenderAddressListResponse() (response *GetSenderAddressListResponse) {
	response = &GetSenderAddressListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
