package slb

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

func (client *Client) SetCACertificateName(request *SetCACertificateNameRequest) (response *SetCACertificateNameResponse, err error) {
	response = CreateSetCACertificateNameResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SetCACertificateNameWithChan(request *SetCACertificateNameRequest) (<-chan *SetCACertificateNameResponse, <-chan error) {
	responseChan := make(chan *SetCACertificateNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetCACertificateName(request)
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

func (client *Client) SetCACertificateNameWithCallback(request *SetCACertificateNameRequest, callback func(response *SetCACertificateNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetCACertificateNameResponse
		var err error
		defer close(result)
		response, err = client.SetCACertificateName(request)
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

type SetCACertificateNameRequest struct {
	*requests.RpcRequest
	CACertificateId      string `position:"Query" name:"CACertificateId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AccessKeyId          string `position:"Query" name:"access_key_id"`
	CACertificateName    string `position:"Query" name:"CACertificateName"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type SetCACertificateNameResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateSetCACertificateNameRequest() (request *SetCACertificateNameRequest) {
	request = &SetCACertificateNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "SetCACertificateName", "", "")
	return
}

func CreateSetCACertificateNameResponse() (response *SetCACertificateNameResponse) {
	response = &SetCACertificateNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
