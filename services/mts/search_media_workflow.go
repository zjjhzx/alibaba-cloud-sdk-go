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

func (client *Client) SearchMediaWorkflow(request *SearchMediaWorkflowRequest) (response *SearchMediaWorkflowResponse, err error) {
	response = CreateSearchMediaWorkflowResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SearchMediaWorkflowWithChan(request *SearchMediaWorkflowRequest) (<-chan *SearchMediaWorkflowResponse, <-chan error) {
	responseChan := make(chan *SearchMediaWorkflowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchMediaWorkflow(request)
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

func (client *Client) SearchMediaWorkflowWithCallback(request *SearchMediaWorkflowRequest, callback func(response *SearchMediaWorkflowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchMediaWorkflowResponse
		var err error
		defer close(result)
		response, err = client.SearchMediaWorkflow(request)
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

type SearchMediaWorkflowRequest struct {
	*requests.RpcRequest
	PageSize             string `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	StateList            string `position:"Query" name:"StateList"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type SearchMediaWorkflowResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	TotalCount        int64  `json:"TotalCount" xml:"TotalCount"`
	PageNumber        int64  `json:"PageNumber" xml:"PageNumber"`
	PageSize          int64  `json:"PageSize" xml:"PageSize"`
	MediaWorkflowList struct {
		MediaWorkflow []struct {
			MediaWorkflowId string `json:"MediaWorkflowId" xml:"MediaWorkflowId"`
			Name            string `json:"Name" xml:"Name"`
			Topology        string `json:"Topology" xml:"Topology"`
			State           string `json:"State" xml:"State"`
			CreationTime    string `json:"CreationTime" xml:"CreationTime"`
		} `json:"MediaWorkflow" xml:"MediaWorkflow"`
	} `json:"MediaWorkflowList" xml:"MediaWorkflowList"`
}

func CreateSearchMediaWorkflowRequest() (request *SearchMediaWorkflowRequest) {
	request = &SearchMediaWorkflowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SearchMediaWorkflow", "", "")
	return
}

func CreateSearchMediaWorkflowResponse() (response *SearchMediaWorkflowResponse) {
	response = &SearchMediaWorkflowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
