package mts

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

type Audio struct {
	Remove     string `json:"Remove" xml:"Remove"`
	Samplerate string `json:"Samplerate" xml:"Samplerate"`
	Channels   string `json:"Channels" xml:"Channels"`
	Qscale     string `json:"Qscale" xml:"Qscale"`
	Profile    string `json:"Profile" xml:"Profile"`
	Codec      string `json:"Codec" xml:"Codec"`
	Bitrate    string `json:"Bitrate" xml:"Bitrate"`
	Volume     Volume `json:"Volume" xml:"Volume"`
}