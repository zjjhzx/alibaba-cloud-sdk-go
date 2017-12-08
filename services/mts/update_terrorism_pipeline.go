package mts

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

func (client *Client) UpdateTerrorismPipeline(request *UpdateTerrorismPipelineRequest) (response *UpdateTerrorismPipelineResponse, err error) {
	response = CreateUpdateTerrorismPipelineResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) UpdateTerrorismPipelineWithChan(request *UpdateTerrorismPipelineRequest) (<-chan *UpdateTerrorismPipelineResponse, <-chan error) {
	responseChan := make(chan *UpdateTerrorismPipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTerrorismPipeline(request)
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

func (client *Client) UpdateTerrorismPipelineWithCallback(request *UpdateTerrorismPipelineRequest, callback func(response *UpdateTerrorismPipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTerrorismPipelineResponse
		var err error
		defer close(result)
		response, err = client.UpdateTerrorismPipeline(request)
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

type UpdateTerrorismPipelineRequest struct {
	*requests.RpcRequest
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Priority             string `position:"Query" name:"Priority"`
	Name                 string `position:"Query" name:"Name"`
	State                string `position:"Query" name:"State"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type UpdateTerrorismPipelineResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Pipeline  struct {
		Id           string `json:"Id" xml:"Id"`
		Name         string `json:"Name" xml:"Name"`
		State        string `json:"State" xml:"State"`
		Priority     int    `json:"Priority" xml:"Priority"`
		NotifyConfig struct {
			Topic string `json:"Topic" xml:"Topic"`
			Queue string `json:"Queue" xml:"Queue"`
		} `json:"NotifyConfig" xml:"NotifyConfig"`
	} `json:"Pipeline" xml:"Pipeline"`
}

func CreateUpdateTerrorismPipelineRequest() (request *UpdateTerrorismPipelineRequest) {
	request = &UpdateTerrorismPipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "UpdateTerrorismPipeline", "", "")
	return
}

func CreateUpdateTerrorismPipelineResponse() (response *UpdateTerrorismPipelineResponse) {
	response = &UpdateTerrorismPipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
