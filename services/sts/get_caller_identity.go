package sts

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

func (client *Client) GetCallerIdentity(request *GetCallerIdentityRequest) (response *GetCallerIdentityResponse, err error) {
	response = CreateGetCallerIdentityResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetCallerIdentityWithChan(request *GetCallerIdentityRequest) (<-chan *GetCallerIdentityResponse, <-chan error) {
	responseChan := make(chan *GetCallerIdentityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCallerIdentity(request)
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

func (client *Client) GetCallerIdentityWithCallback(request *GetCallerIdentityRequest, callback func(response *GetCallerIdentityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCallerIdentityResponse
		var err error
		defer close(result)
		response, err = client.GetCallerIdentity(request)
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

type GetCallerIdentityRequest struct {
	*requests.RpcRequest
}

type GetCallerIdentityResponse struct {
	*responses.BaseResponse
	AccountId string `json:"AccountId" xml:"AccountId"`
	UserId    string `json:"UserId" xml:"UserId"`
	Arn       string `json:"Arn" xml:"Arn"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateGetCallerIdentityRequest() (request *GetCallerIdentityRequest) {
	request = &GetCallerIdentityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sts", "2015-04-01", "GetCallerIdentity", "", "")
	return
}

func CreateGetCallerIdentityResponse() (response *GetCallerIdentityResponse) {
	response = &GetCallerIdentityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
