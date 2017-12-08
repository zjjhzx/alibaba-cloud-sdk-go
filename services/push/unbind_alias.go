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

func (client *Client) UnbindAlias(request *UnbindAliasRequest) (response *UnbindAliasResponse, err error) {
	response = CreateUnbindAliasResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) UnbindAliasWithChan(request *UnbindAliasRequest) (<-chan *UnbindAliasResponse, <-chan error) {
	responseChan := make(chan *UnbindAliasResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindAlias(request)
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

func (client *Client) UnbindAliasWithCallback(request *UnbindAliasRequest, callback func(response *UnbindAliasResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindAliasResponse
		var err error
		defer close(result)
		response, err = client.UnbindAlias(request)
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

type UnbindAliasRequest struct {
	*requests.RpcRequest
	UnbindAll string `position:"Query" name:"UnbindAll"`
	AliasName string `position:"Query" name:"AliasName"`
	AppKey    string `position:"Query" name:"AppKey"`
	DeviceId  string `position:"Query" name:"DeviceId"`
}

type UnbindAliasResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateUnbindAliasRequest() (request *UnbindAliasRequest) {
	request = &UnbindAliasRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "UnbindAlias", "", "")
	return
}

func CreateUnbindAliasResponse() (response *UnbindAliasResponse) {
	response = &UnbindAliasResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
