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

func (client *Client) DescribeSQLLogRecords(request *DescribeSQLLogRecordsRequest) (response *DescribeSQLLogRecordsResponse, err error) {
	response = CreateDescribeSQLLogRecordsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeSQLLogRecordsWithChan(request *DescribeSQLLogRecordsRequest) (<-chan *DescribeSQLLogRecordsResponse, <-chan error) {
	responseChan := make(chan *DescribeSQLLogRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSQLLogRecords(request)
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

func (client *Client) DescribeSQLLogRecordsWithCallback(request *DescribeSQLLogRecordsRequest, callback func(response *DescribeSQLLogRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSQLLogRecordsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSQLLogRecords(request)
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

type DescribeSQLLogRecordsRequest struct {
	*requests.RpcRequest
	EndTime              string `position:"Query" name:"EndTime"`
	PageSize             string `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	Form                 string `position:"Query" name:"Form"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	Database             string `position:"Query" name:"Database"`
	SQLId                string `position:"Query" name:"SQLId"`
	QueryKeywords        string `position:"Query" name:"QueryKeywords"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	User                 string `position:"Query" name:"User"`
}

type DescribeSQLLogRecordsResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	TotalRecordCount int    `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       int    `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int    `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            struct {
		SQLRecord []struct {
			DBName              string `json:"DBName" xml:"DBName"`
			AccountName         string `json:"AccountName" xml:"AccountName"`
			HostAddress         string `json:"HostAddress" xml:"HostAddress"`
			SQLText             string `json:"SQLText" xml:"SQLText"`
			TotalExecutionTimes int64  `json:"TotalExecutionTimes" xml:"TotalExecutionTimes"`
			ReturnRowCounts     int64  `json:"ReturnRowCounts" xml:"ReturnRowCounts"`
			ExecuteTime         string `json:"ExecuteTime" xml:"ExecuteTime"`
			ThreadID            string `json:"ThreadID" xml:"ThreadID"`
		} `json:"SQLRecord" xml:"SQLRecord"`
	} `json:"Items" xml:"Items"`
}

func CreateDescribeSQLLogRecordsRequest() (request *DescribeSQLLogRecordsRequest) {
	request = &DescribeSQLLogRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLLogRecords", "", "")
	return
}

func CreateDescribeSQLLogRecordsResponse() (response *DescribeSQLLogRecordsResponse) {
	response = &DescribeSQLLogRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
