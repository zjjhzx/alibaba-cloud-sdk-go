package ess

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

func (client *Client) DescribeScheduledTasks(request *DescribeScheduledTasksRequest) (response *DescribeScheduledTasksResponse, err error) {
	response = CreateDescribeScheduledTasksResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeScheduledTasksWithChan(request *DescribeScheduledTasksRequest) (<-chan *DescribeScheduledTasksResponse, <-chan error) {
	responseChan := make(chan *DescribeScheduledTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScheduledTasks(request)
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

func (client *Client) DescribeScheduledTasksWithCallback(request *DescribeScheduledTasksRequest, callback func(response *DescribeScheduledTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScheduledTasksResponse
		var err error
		defer close(result)
		response, err = client.DescribeScheduledTasks(request)
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

type DescribeScheduledTasksRequest struct {
	*requests.RpcRequest
	ScheduledTaskName6   string `position:"Query" name:"ScheduledTaskName.6"`
	ScheduledTaskName20  string `position:"Query" name:"ScheduledTaskName.20"`
	ScheduledTaskName5   string `position:"Query" name:"ScheduledTaskName.5"`
	ScheduledTaskName8   string `position:"Query" name:"ScheduledTaskName.8"`
	ScheduledTaskName7   string `position:"Query" name:"ScheduledTaskName.7"`
	ScheduledTaskName2   string `position:"Query" name:"ScheduledTaskName.2"`
	ScheduledTaskName1   string `position:"Query" name:"ScheduledTaskName.1"`
	ScheduledTaskName4   string `position:"Query" name:"ScheduledTaskName.4"`
	ScheduledTaskName3   string `position:"Query" name:"ScheduledTaskName.3"`
	ScheduledAction16    string `position:"Query" name:"ScheduledAction.16"`
	ScheduledAction17    string `position:"Query" name:"ScheduledAction.17"`
	ScheduledAction18    string `position:"Query" name:"ScheduledAction.18"`
	ScheduledAction19    string `position:"Query" name:"ScheduledAction.19"`
	ScheduledTaskName9   string `position:"Query" name:"ScheduledTaskName.9"`
	ScheduledAction11    string `position:"Query" name:"ScheduledAction.11"`
	ScheduledAction10    string `position:"Query" name:"ScheduledAction.10"`
	ScheduledAction13    string `position:"Query" name:"ScheduledAction.13"`
	ScheduledAction12    string `position:"Query" name:"ScheduledAction.12"`
	ScheduledAction15    string `position:"Query" name:"ScheduledAction.15"`
	ScheduledAction14    string `position:"Query" name:"ScheduledAction.14"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ScheduledTaskId7     string `position:"Query" name:"ScheduledTaskId.7"`
	ScheduledTaskId8     string `position:"Query" name:"ScheduledTaskId.8"`
	ScheduledTaskId5     string `position:"Query" name:"ScheduledTaskId.5"`
	ScheduledTaskId6     string `position:"Query" name:"ScheduledTaskId.6"`
	ScheduledTaskId3     string `position:"Query" name:"ScheduledTaskId.3"`
	ScheduledTaskId4     string `position:"Query" name:"ScheduledTaskId.4"`
	ScheduledTaskId1     string `position:"Query" name:"ScheduledTaskId.1"`
	ScheduledTaskId2     string `position:"Query" name:"ScheduledTaskId.2"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	ScheduledTaskId9     string `position:"Query" name:"ScheduledTaskId.9"`
	ScheduledAction20    string `position:"Query" name:"ScheduledAction.20"`
	PageSize             string `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScheduledTaskId16    string `position:"Query" name:"ScheduledTaskId.16"`
	ScheduledTaskId17    string `position:"Query" name:"ScheduledTaskId.17"`
	ScheduledTaskId14    string `position:"Query" name:"ScheduledTaskId.14"`
	ScheduledTaskId15    string `position:"Query" name:"ScheduledTaskId.15"`
	ScheduledTaskId12    string `position:"Query" name:"ScheduledTaskId.12"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	ScheduledTaskId13    string `position:"Query" name:"ScheduledTaskId.13"`
	ScheduledTaskId10    string `position:"Query" name:"ScheduledTaskId.10"`
	ScheduledTaskId11    string `position:"Query" name:"ScheduledTaskId.11"`
	ScheduledTaskId18    string `position:"Query" name:"ScheduledTaskId.18"`
	ScheduledTaskId19    string `position:"Query" name:"ScheduledTaskId.19"`
	ScheduledTaskName10  string `position:"Query" name:"ScheduledTaskName.10"`
	ScheduledAction3     string `position:"Query" name:"ScheduledAction.3"`
	ScheduledTaskName11  string `position:"Query" name:"ScheduledTaskName.11"`
	ScheduledAction4     string `position:"Query" name:"ScheduledAction.4"`
	ScheduledTaskName12  string `position:"Query" name:"ScheduledTaskName.12"`
	ScheduledAction5     string `position:"Query" name:"ScheduledAction.5"`
	ScheduledTaskName13  string `position:"Query" name:"ScheduledTaskName.13"`
	ScheduledAction6     string `position:"Query" name:"ScheduledAction.6"`
	ScheduledTaskName14  string `position:"Query" name:"ScheduledTaskName.14"`
	ScheduledAction7     string `position:"Query" name:"ScheduledAction.7"`
	ScheduledTaskName15  string `position:"Query" name:"ScheduledTaskName.15"`
	ScheduledAction8     string `position:"Query" name:"ScheduledAction.8"`
	ScheduledTaskName16  string `position:"Query" name:"ScheduledTaskName.16"`
	ScheduledAction9     string `position:"Query" name:"ScheduledAction.9"`
	ScheduledTaskName17  string `position:"Query" name:"ScheduledTaskName.17"`
	ScheduledTaskName18  string `position:"Query" name:"ScheduledTaskName.18"`
	ScheduledTaskId20    string `position:"Query" name:"ScheduledTaskId.20"`
	ScheduledTaskName19  string `position:"Query" name:"ScheduledTaskName.19"`
	ScheduledAction1     string `position:"Query" name:"ScheduledAction.1"`
	ScheduledAction2     string `position:"Query" name:"ScheduledAction.2"`
}

type DescribeScheduledTasksResponse struct {
	*responses.BaseResponse
	TotalCount     int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber     int    `json:"PageNumber" xml:"PageNumber"`
	PageSize       int    `json:"PageSize" xml:"PageSize"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ScheduledTasks struct {
		ScheduledTask []struct {
			ScheduledTaskId      string `json:"ScheduledTaskId" xml:"ScheduledTaskId"`
			ScheduledTaskName    string `json:"ScheduledTaskName" xml:"ScheduledTaskName"`
			Description          string `json:"Description" xml:"Description"`
			ScheduledAction      string `json:"ScheduledAction" xml:"ScheduledAction"`
			RecurrenceEndTime    string `json:"RecurrenceEndTime" xml:"RecurrenceEndTime"`
			LaunchTime           string `json:"LaunchTime" xml:"LaunchTime"`
			RecurrenceType       string `json:"RecurrenceType" xml:"RecurrenceType"`
			RecurrenceValue      string `json:"RecurrenceValue" xml:"RecurrenceValue"`
			LaunchExpirationTime int    `json:"LaunchExpirationTime" xml:"LaunchExpirationTime"`
			TaskEnabled          bool   `json:"TaskEnabled" xml:"TaskEnabled"`
		} `json:"ScheduledTask" xml:"ScheduledTask"`
	} `json:"ScheduledTasks" xml:"ScheduledTasks"`
}

func CreateDescribeScheduledTasksRequest() (request *DescribeScheduledTasksRequest) {
	request = &DescribeScheduledTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DescribeScheduledTasks", "", "")
	return
}

func CreateDescribeScheduledTasksResponse() (response *DescribeScheduledTasksResponse) {
	response = &DescribeScheduledTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
