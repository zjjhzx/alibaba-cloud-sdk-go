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

func (client *Client) DescribeRouteTables(request *DescribeRouteTablesRequest) (response *DescribeRouteTablesResponse, err error) {
	response = CreateDescribeRouteTablesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeRouteTablesWithChan(request *DescribeRouteTablesRequest) (<-chan *DescribeRouteTablesResponse, <-chan error) {
	responseChan := make(chan *DescribeRouteTablesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRouteTables(request)
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

func (client *Client) DescribeRouteTablesWithCallback(request *DescribeRouteTablesRequest, callback func(response *DescribeRouteTablesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRouteTablesResponse
		var err error
		defer close(result)
		response, err = client.DescribeRouteTables(request)
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

type DescribeRouteTablesRequest struct {
	*requests.RpcRequest
	PageSize             string `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	VRouterId            string `position:"Query" name:"VRouterId"`
	RouterId             string `position:"Query" name:"RouterId"`
	RouterType           string `position:"Query" name:"RouterType"`
	RouteTableName       string `position:"Query" name:"RouteTableName"`
	RouteTableId         string `position:"Query" name:"RouteTableId"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type DescribeRouteTablesResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	TotalCount  int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber  int    `json:"PageNumber" xml:"PageNumber"`
	PageSize    int    `json:"PageSize" xml:"PageSize"`
	RouteTables struct {
		RouteTable []struct {
			VRouterId      string `json:"VRouterId" xml:"VRouterId"`
			RouteTableId   string `json:"RouteTableId" xml:"RouteTableId"`
			RouteTableType string `json:"RouteTableType" xml:"RouteTableType"`
			CreationTime   string `json:"CreationTime" xml:"CreationTime"`
			RouteEntrys    struct {
				RouteEntry []struct {
					RouteTableId         string `json:"RouteTableId" xml:"RouteTableId"`
					DestinationCidrBlock string `json:"DestinationCidrBlock" xml:"DestinationCidrBlock"`
					Type                 string `json:"Type" xml:"Type"`
					Status               string `json:"Status" xml:"Status"`
					InstanceId           string `json:"InstanceId" xml:"InstanceId"`
					NextHopType          string `json:"NextHopType" xml:"NextHopType"`
					NextHops             struct {
						NextHop []struct {
							NextHopType string `json:"NextHopType" xml:"NextHopType"`
							NextHopId   string `json:"NextHopId" xml:"NextHopId"`
							Enabled     int    `json:"Enabled" xml:"Enabled"`
							Weight      int    `json:"Weight" xml:"Weight"`
						} `json:"NextHop" xml:"NextHop"`
					} `json:"NextHops" xml:"NextHops"`
				} `json:"RouteEntry" xml:"RouteEntry"`
			} `json:"RouteEntrys" xml:"RouteEntrys"`
		} `json:"RouteTable" xml:"RouteTable"`
	} `json:"RouteTables" xml:"RouteTables"`
}

func CreateDescribeRouteTablesRequest() (request *DescribeRouteTablesRequest) {
	request = &DescribeRouteTablesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeRouteTables", "", "")
	return
}

func CreateDescribeRouteTablesResponse() (response *DescribeRouteTablesResponse) {
	response = &DescribeRouteTablesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
