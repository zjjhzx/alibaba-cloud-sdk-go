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

func (client *Client) PreCheckBeforeImportData(request *PreCheckBeforeImportDataRequest) (response *PreCheckBeforeImportDataResponse, err error) {
	response = CreatePreCheckBeforeImportDataResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) PreCheckBeforeImportDataWithChan(request *PreCheckBeforeImportDataRequest) (<-chan *PreCheckBeforeImportDataResponse, <-chan error) {
	responseChan := make(chan *PreCheckBeforeImportDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PreCheckBeforeImportData(request)
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

func (client *Client) PreCheckBeforeImportDataWithCallback(request *PreCheckBeforeImportDataRequest, callback func(response *PreCheckBeforeImportDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PreCheckBeforeImportDataResponse
		var err error
		defer close(result)
		response, err = client.PreCheckBeforeImportData(request)
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

type PreCheckBeforeImportDataRequest struct {
	*requests.RpcRequest
	DBInstanceId           string `position:"Query" name:"DBInstanceId"`
	ImportDataType         string `position:"Query" name:"ImportDataType"`
	ClientToken            string `position:"Query" name:"ClientToken"`
	SourceDatabaseIp       string `position:"Query" name:"SourceDatabaseIp"`
	OwnerId                string `position:"Query" name:"OwnerId"`
	SourceDatabaseUserName string `position:"Query" name:"SourceDatabaseUserName"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	SourceDatabaseDBNames  string `position:"Query" name:"SourceDatabaseDBNames"`
	ResourceOwnerId        string `position:"Query" name:"ResourceOwnerId"`
	SourceDatabasePassword string `position:"Query" name:"SourceDatabasePassword"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	SourceDatabasePort     string `position:"Query" name:"SourceDatabasePort"`
}

type PreCheckBeforeImportDataResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	PreCheckId string `json:"PreCheckId" xml:"PreCheckId"`
}

func CreatePreCheckBeforeImportDataRequest() (request *PreCheckBeforeImportDataRequest) {
	request = &PreCheckBeforeImportDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "PreCheckBeforeImportData", "", "")
	return
}

func CreatePreCheckBeforeImportDataResponse() (response *PreCheckBeforeImportDataResponse) {
	response = &PreCheckBeforeImportDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
