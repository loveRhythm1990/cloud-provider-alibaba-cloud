package ens

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

// DescribeEpnBandwitdhByInternetChargeType invokes the ens.DescribeEpnBandwitdhByInternetChargeType API synchronously
func (client *Client) DescribeEpnBandwitdhByInternetChargeType(request *DescribeEpnBandwitdhByInternetChargeTypeRequest) (response *DescribeEpnBandwitdhByInternetChargeTypeResponse, err error) {
	response = CreateDescribeEpnBandwitdhByInternetChargeTypeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEpnBandwitdhByInternetChargeTypeWithChan invokes the ens.DescribeEpnBandwitdhByInternetChargeType API asynchronously
func (client *Client) DescribeEpnBandwitdhByInternetChargeTypeWithChan(request *DescribeEpnBandwitdhByInternetChargeTypeRequest) (<-chan *DescribeEpnBandwitdhByInternetChargeTypeResponse, <-chan error) {
	responseChan := make(chan *DescribeEpnBandwitdhByInternetChargeTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEpnBandwitdhByInternetChargeType(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeEpnBandwitdhByInternetChargeTypeWithCallback invokes the ens.DescribeEpnBandwitdhByInternetChargeType API asynchronously
func (client *Client) DescribeEpnBandwitdhByInternetChargeTypeWithCallback(request *DescribeEpnBandwitdhByInternetChargeTypeRequest, callback func(response *DescribeEpnBandwitdhByInternetChargeTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEpnBandwitdhByInternetChargeTypeResponse
		var err error
		defer close(result)
		response, err = client.DescribeEpnBandwitdhByInternetChargeType(request)
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

// DescribeEpnBandwitdhByInternetChargeTypeRequest is the request struct for api DescribeEpnBandwitdhByInternetChargeType
type DescribeEpnBandwitdhByInternetChargeTypeRequest struct {
	*requests.RpcRequest
	NetworkingModel string `position:"Query" name:"NetworkingModel"`
	Isp             string `position:"Query" name:"Isp"`
	StartTime       string `position:"Query" name:"StartTime"`
	EnsRegionId     string `position:"Query" name:"EnsRegionId"`
	EndTime         string `position:"Query" name:"EndTime"`
}

// DescribeEpnBandwitdhByInternetChargeTypeResponse is the response struct for api DescribeEpnBandwitdhByInternetChargeType
type DescribeEpnBandwitdhByInternetChargeTypeResponse struct {
	*responses.BaseResponse
	BandwidthValue     int64  `json:"BandwidthValue" xml:"BandwidthValue"`
	InternetChargeType string `json:"InternetChargeType" xml:"InternetChargeType"`
	RequestId          string `json:"RequestId" xml:"RequestId"`
	TimeStamp          string `json:"TimeStamp" xml:"TimeStamp"`
}

// CreateDescribeEpnBandwitdhByInternetChargeTypeRequest creates a request to invoke DescribeEpnBandwitdhByInternetChargeType API
func CreateDescribeEpnBandwitdhByInternetChargeTypeRequest() (request *DescribeEpnBandwitdhByInternetChargeTypeRequest) {
	request = &DescribeEpnBandwitdhByInternetChargeTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeEpnBandwitdhByInternetChargeType", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeEpnBandwitdhByInternetChargeTypeResponse creates a response to parse from DescribeEpnBandwitdhByInternetChargeType response
func CreateDescribeEpnBandwitdhByInternetChargeTypeResponse() (response *DescribeEpnBandwitdhByInternetChargeTypeResponse) {
	response = &DescribeEpnBandwitdhByInternetChargeTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
