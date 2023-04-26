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

// MeasurementData is a nested struct in ens response
type MeasurementData struct {
	CostCycle              string                                     `json:"CostCycle" xml:"CostCycle"`
	CostEndTime            string                                     `json:"CostEndTime" xml:"CostEndTime"`
	ChargeModel            string                                     `json:"ChargeModel" xml:"ChargeModel"`
	CostStartTime          string                                     `json:"CostStartTime" xml:"CostStartTime"`
	ResourceFeeData        ResourceFeeData                            `json:"ResourceFeeData" xml:"ResourceFeeData"`
	BandWidthFeeDatas      BandWidthFeeDatasInDescribeMeasurementData `json:"BandWidthFeeDatas" xml:"BandWidthFeeDatas"`
	ResourceFeeDataDetails ResourceFeeDataDetails                     `json:"ResourceFeeDataDetails" xml:"ResourceFeeDataDetails"`
}
