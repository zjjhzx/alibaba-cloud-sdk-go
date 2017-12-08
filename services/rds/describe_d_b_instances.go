package rds

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

func (client *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
	response = CreateDescribeDBInstancesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDBInstancesWithChan(request *DescribeDBInstancesRequest) (<-chan *DescribeDBInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstances(request)
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

func (client *Client) DescribeDBInstancesWithCallback(request *DescribeDBInstancesRequest, callback func(response *DescribeDBInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstances(request)
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

type DescribeDBInstancesRequest struct {
	*requests.RpcRequest
	PageSize             string `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	ProxyId              string `position:"Query" name:"proxyId"`
	SearchKey            string `position:"Query" name:"SearchKey"`
	Tag5Value            string `position:"Query" name:"Tag.5.value"`
	DBInstanceStatus     string `position:"Query" name:"DBInstanceStatus"`
	Tag3Key              string `position:"Query" name:"Tag.3.key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Tag1Key              string `position:"Query" name:"Tag.1.key"`
	Tag1Value            string `position:"Query" name:"Tag.1.value"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Tag4Value            string `position:"Query" name:"Tag.4.value"`
	Tags                 string `position:"Query" name:"Tags"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	Engine               string `position:"Query" name:"Engine"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	DBInstanceType       string `position:"Query" name:"DBInstanceType"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	Tag5Key              string `position:"Query" name:"Tag.5.key"`
	Tag2Key              string `position:"Query" name:"Tag.2.key"`
	VpcId                string `position:"Query" name:"VpcId"`
	ConnectionMode       string `position:"Query" name:"ConnectionMode"`
	Tag3Value            string `position:"Query" name:"Tag.3.value"`
	Tag2Value            string `position:"Query" name:"Tag.2.value"`
	InstanceNetworkType  string `position:"Query" name:"InstanceNetworkType"`
	Tag4Key              string `position:"Query" name:"Tag.4.key"`
}

type DescribeDBInstancesResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	PageNumber       int    `json:"PageNumber" xml:"PageNumber"`
	TotalRecordCount int    `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageRecordCount  int    `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            struct {
		DBInstance []struct {
			InsId                 int    `json:"InsId" xml:"InsId"`
			DBInstanceId          string `json:"DBInstanceId" xml:"DBInstanceId"`
			DBInstanceDescription string `json:"DBInstanceDescription" xml:"DBInstanceDescription"`
			PayType               string `json:"PayType" xml:"PayType"`
			DBInstanceType        string `json:"DBInstanceType" xml:"DBInstanceType"`
			RegionId              string `json:"RegionId" xml:"RegionId"`
			ExpireTime            string `json:"ExpireTime" xml:"ExpireTime"`
			DBInstanceStatus      string `json:"DBInstanceStatus" xml:"DBInstanceStatus"`
			Engine                string `json:"Engine" xml:"Engine"`
			DBInstanceNetType     string `json:"DBInstanceNetType" xml:"DBInstanceNetType"`
			ConnectionMode        string `json:"ConnectionMode" xml:"ConnectionMode"`
			LockMode              string `json:"LockMode" xml:"LockMode"`
			DBInstanceClass       string `json:"DBInstanceClass" xml:"DBInstanceClass"`
			InstanceNetworkType   string `json:"InstanceNetworkType" xml:"InstanceNetworkType"`
			LockReason            string `json:"LockReason" xml:"LockReason"`
			ZoneId                string `json:"ZoneId" xml:"ZoneId"`
			MutriORsignle         bool   `json:"MutriORsignle" xml:"MutriORsignle"`
			CreateTime            string `json:"CreateTime" xml:"CreateTime"`
			EngineVersion         string `json:"EngineVersion" xml:"EngineVersion"`
			GuardDBInstanceId     string `json:"GuardDBInstanceId" xml:"GuardDBInstanceId"`
			TempDBInstanceId      string `json:"TempDBInstanceId" xml:"TempDBInstanceId"`
			MasterInstanceId      string `json:"MasterInstanceId" xml:"MasterInstanceId"`
			VpcId                 string `json:"VpcId" xml:"VpcId"`
			VSwitchId             string `json:"VSwitchId" xml:"VSwitchId"`
			ResourceGroupId       string `json:"ResourceGroupId" xml:"ResourceGroupId"`
			ReadOnlyDBInstanceIds struct {
				ReadOnlyDBInstanceId []struct {
					DBInstanceId string `json:"DBInstanceId" xml:"DBInstanceId"`
				} `json:"ReadOnlyDBInstanceId" xml:"ReadOnlyDBInstanceId"`
			} `json:"ReadOnlyDBInstanceIds" xml:"ReadOnlyDBInstanceIds"`
		} `json:"DBInstance" xml:"DBInstance"`
	} `json:"Items" xml:"Items"`
}

func CreateDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
	request = &DescribeDBInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstances", "", "")
	return
}

func CreateDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
	response = &DescribeDBInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
