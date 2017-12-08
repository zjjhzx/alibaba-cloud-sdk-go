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

func (client *Client) AddVServerGroupBackendServers(request *AddVServerGroupBackendServersRequest) (response *AddVServerGroupBackendServersResponse, err error) {
	response = CreateAddVServerGroupBackendServersResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AddVServerGroupBackendServersWithChan(request *AddVServerGroupBackendServersRequest) (<-chan *AddVServerGroupBackendServersResponse, <-chan error) {
	responseChan := make(chan *AddVServerGroupBackendServersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddVServerGroupBackendServers(request)
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

func (client *Client) AddVServerGroupBackendServersWithCallback(request *AddVServerGroupBackendServersRequest, callback func(response *AddVServerGroupBackendServersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddVServerGroupBackendServersResponse
		var err error
		defer close(result)
		response, err = client.AddVServerGroupBackendServers(request)
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

type AddVServerGroupBackendServersRequest struct {
	*requests.RpcRequest
	VServerGroupId       string `position:"Query" name:"VServerGroupId"`
	Tags                 string `position:"Query" name:"Tags"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AccessKeyId          string `position:"Query" name:"access_key_id"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	BackendServers       string `position:"Query" name:"BackendServers"`
}

type AddVServerGroupBackendServersResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	VServerGroupId string `json:"VServerGroupId" xml:"VServerGroupId"`
	BackendServers struct {
		BackendServer []struct {
			ServerId string `json:"ServerId" xml:"ServerId"`
			Port     int    `json:"Port" xml:"Port"`
			Weight   int    `json:"Weight" xml:"Weight"`
		} `json:"BackendServer" xml:"BackendServer"`
	} `json:"BackendServers" xml:"BackendServers"`
}

func CreateAddVServerGroupBackendServersRequest() (request *AddVServerGroupBackendServersRequest) {
	request = &AddVServerGroupBackendServersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "AddVServerGroupBackendServers", "", "")
	return
}

func CreateAddVServerGroupBackendServersResponse() (response *AddVServerGroupBackendServersResponse) {
	response = &AddVServerGroupBackendServersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
