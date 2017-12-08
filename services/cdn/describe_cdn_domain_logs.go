package cdn

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

func (client *Client) DescribeCdnDomainLogs(request *DescribeCdnDomainLogsRequest) (response *DescribeCdnDomainLogsResponse, err error) {
	response = CreateDescribeCdnDomainLogsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeCdnDomainLogsWithChan(request *DescribeCdnDomainLogsRequest) (<-chan *DescribeCdnDomainLogsResponse, <-chan error) {
	responseChan := make(chan *DescribeCdnDomainLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCdnDomainLogs(request)
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

func (client *Client) DescribeCdnDomainLogsWithCallback(request *DescribeCdnDomainLogsRequest, callback func(response *DescribeCdnDomainLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCdnDomainLogsResponse
		var err error
		defer close(result)
		response, err = client.DescribeCdnDomainLogs(request)
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

type DescribeCdnDomainLogsRequest struct {
	*requests.RpcRequest
	EndTime       string `position:"Query" name:"EndTime"`
	PageSize      string `position:"Query" name:"PageSize"`
	StartTime     string `position:"Query" name:"StartTime"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageNumber    string `position:"Query" name:"PageNumber"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	LogDay        string `position:"Query" name:"LogDay"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

type DescribeCdnDomainLogsResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	PageNumber     int64  `json:"PageNumber" xml:"PageNumber"`
	PageSize       int64  `json:"PageSize" xml:"PageSize"`
	TotalCount     int64  `json:"TotalCount" xml:"TotalCount"`
	DomainLogModel struct {
		DomainName       string `json:"DomainName" xml:"DomainName"`
		DomainLogDetails struct {
			DomainLogDetail []struct {
				LogName   string `json:"LogName" xml:"LogName"`
				LogPath   string `json:"LogPath" xml:"LogPath"`
				LogSize   int64  `json:"LogSize" xml:"LogSize"`
				StartTime string `json:"StartTime" xml:"StartTime"`
				EndTime   string `json:"EndTime" xml:"EndTime"`
			} `json:"DomainLogDetail" xml:"DomainLogDetail"`
		} `json:"DomainLogDetails" xml:"DomainLogDetails"`
	} `json:"DomainLogModel" xml:"DomainLogModel"`
}

func CreateDescribeCdnDomainLogsRequest() (request *DescribeCdnDomainLogsRequest) {
	request = &DescribeCdnDomainLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeCdnDomainLogs", "", "")
	return
}

func CreateDescribeCdnDomainLogsResponse() (response *DescribeCdnDomainLogsResponse) {
	response = &DescribeCdnDomainLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
