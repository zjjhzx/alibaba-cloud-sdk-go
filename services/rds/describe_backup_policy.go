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

func (client *Client) DescribeBackupPolicy(request *DescribeBackupPolicyRequest) (response *DescribeBackupPolicyResponse, err error) {
	response = CreateDescribeBackupPolicyResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeBackupPolicyWithChan(request *DescribeBackupPolicyRequest) (<-chan *DescribeBackupPolicyResponse, <-chan error) {
	responseChan := make(chan *DescribeBackupPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBackupPolicy(request)
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

func (client *Client) DescribeBackupPolicyWithCallback(request *DescribeBackupPolicyRequest, callback func(response *DescribeBackupPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBackupPolicyResponse
		var err error
		defer close(result)
		response, err = client.DescribeBackupPolicy(request)
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

type DescribeBackupPolicyRequest struct {
	*requests.RpcRequest
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

type DescribeBackupPolicyResponse struct {
	*responses.BaseResponse
	RequestId                string `json:"RequestId" xml:"RequestId"`
	BackupRetentionPeriod    int    `json:"BackupRetentionPeriod" xml:"BackupRetentionPeriod"`
	PreferredNextBackupTime  string `json:"PreferredNextBackupTime" xml:"PreferredNextBackupTime"`
	PreferredBackupTime      string `json:"PreferredBackupTime" xml:"PreferredBackupTime"`
	PreferredBackupPeriod    string `json:"PreferredBackupPeriod" xml:"PreferredBackupPeriod"`
	BackupLog                string `json:"BackupLog" xml:"BackupLog"`
	LogBackupRetentionPeriod int    `json:"LogBackupRetentionPeriod" xml:"LogBackupRetentionPeriod"`
}

func CreateDescribeBackupPolicyRequest() (request *DescribeBackupPolicyRequest) {
	request = &DescribeBackupPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeBackupPolicy", "", "")
	return
}

func CreateDescribeBackupPolicyResponse() (response *DescribeBackupPolicyResponse) {
	response = &DescribeBackupPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
