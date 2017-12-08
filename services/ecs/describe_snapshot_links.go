package ecs

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

func (client *Client) DescribeSnapshotLinks(request *DescribeSnapshotLinksRequest) (response *DescribeSnapshotLinksResponse, err error) {
	response = CreateDescribeSnapshotLinksResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeSnapshotLinksWithChan(request *DescribeSnapshotLinksRequest) (<-chan *DescribeSnapshotLinksResponse, <-chan error) {
	responseChan := make(chan *DescribeSnapshotLinksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSnapshotLinks(request)
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

func (client *Client) DescribeSnapshotLinksWithCallback(request *DescribeSnapshotLinksRequest, callback func(response *DescribeSnapshotLinksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSnapshotLinksResponse
		var err error
		defer close(result)
		response, err = client.DescribeSnapshotLinks(request)
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

type DescribeSnapshotLinksRequest struct {
	*requests.RpcRequest
	PageSize             string `position:"Query" name:"PageSize"`
	DiskIds              string `position:"Query" name:"DiskIds"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	SnapshotLinkIds      string `position:"Query" name:"SnapshotLinkIds"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
}

type DescribeSnapshotLinksResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	TotalCount    int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber    int    `json:"PageNumber" xml:"PageNumber"`
	PageSize      int    `json:"PageSize" xml:"PageSize"`
	SnapshotLinks struct {
		SnapshotLink []struct {
			SnapshotLinkId string `json:"SnapshotLinkId" xml:"SnapshotLinkId"`
			RegionId       string `json:"RegionId" xml:"RegionId"`
			InstanceId     string `json:"InstanceId" xml:"InstanceId"`
			InstanceName   string `json:"InstanceName" xml:"InstanceName"`
			SourceDiskId   string `json:"SourceDiskId" xml:"SourceDiskId"`
			SourceDiskSize int    `json:"SourceDiskSize" xml:"SourceDiskSize"`
			SourceDiskType string `json:"SourceDiskType" xml:"SourceDiskType"`
			TotalSize      int    `json:"TotalSize" xml:"TotalSize"`
			TotalCount     int    `json:"TotalCount" xml:"TotalCount"`
		} `json:"SnapshotLink" xml:"SnapshotLink"`
	} `json:"SnapshotLinks" xml:"SnapshotLinks"`
}

func CreateDescribeSnapshotLinksRequest() (request *DescribeSnapshotLinksRequest) {
	request = &DescribeSnapshotLinksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSnapshotLinks", "", "")
	return
}

func CreateDescribeSnapshotLinksResponse() (response *DescribeSnapshotLinksResponse) {
	response = &DescribeSnapshotLinksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
