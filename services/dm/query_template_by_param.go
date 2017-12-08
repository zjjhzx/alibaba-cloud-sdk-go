package dm

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

func (client *Client) QueryTemplateByParam(request *QueryTemplateByParamRequest) (response *QueryTemplateByParamResponse, err error) {
	response = CreateQueryTemplateByParamResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryTemplateByParamWithChan(request *QueryTemplateByParamRequest) (<-chan *QueryTemplateByParamResponse, <-chan error) {
	responseChan := make(chan *QueryTemplateByParamResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTemplateByParam(request)
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

func (client *Client) QueryTemplateByParamWithCallback(request *QueryTemplateByParamRequest, callback func(response *QueryTemplateByParamResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTemplateByParamResponse
		var err error
		defer close(result)
		response, err = client.QueryTemplateByParam(request)
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

type QueryTemplateByParamRequest struct {
	*requests.RpcRequest
	PageSize             string `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Status               string `position:"Query" name:"Status"`
	KeyWord              string `position:"Query" name:"KeyWord"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	PageNo               string `position:"Query" name:"PageNo"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	FromType             string `position:"Query" name:"FromType"`
}

type QueryTemplateByParamResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	Data       struct {
		Template []struct {
			TemplateId      string `json:"TemplateId" xml:"TemplateId"`
			TemplateName    string `json:"TemplateName" xml:"TemplateName"`
			TemplateComment string `json:"TemplateComment" xml:"TemplateComment"`
			TemplateStatus  string `json:"TemplateStatus" xml:"TemplateStatus"`
			CreateTime      string `json:"CreateTime" xml:"CreateTime"`
			UtcCreatetime   int64  `json:"UtcCreatetime" xml:"UtcCreatetime"`
			TemplateType    int    `json:"TemplateType" xml:"TemplateType"`
			SmsTemplateCode int    `json:"SmsTemplateCode" xml:"SmsTemplateCode"`
			Smsrejectinfo   int    `json:"Smsrejectinfo" xml:"Smsrejectinfo"`
			SmsStatus       int    `json:"SmsStatus" xml:"SmsStatus"`
		} `json:"template" xml:"template"`
	} `json:"data" xml:"data"`
}

func CreateQueryTemplateByParamRequest() (request *QueryTemplateByParamRequest) {
	request = &QueryTemplateByParamRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "QueryTemplateByParam", "", "")
	return
}

func CreateQueryTemplateByParamResponse() (response *QueryTemplateByParamResponse) {
	response = &QueryTemplateByParamResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
