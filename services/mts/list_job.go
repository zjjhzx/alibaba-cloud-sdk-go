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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) ListJob(request *ListJobRequest) (response *ListJobResponse, err error) {
	response = CreateListJobResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListJobWithChan(request *ListJobRequest) (<-chan *ListJobResponse, <-chan error) {
	responseChan := make(chan *ListJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListJob(request)
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

func (client *Client) ListJobWithCallback(request *ListJobRequest, callback func(response *ListJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListJobResponse
		var err error
		defer close(result)
		response, err = client.ListJob(request)
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

type ListJobRequest struct {
	*requests.RpcRequest
	StartOfJobCreatedTimeRange string `position:"Query" name:"StartOfJobCreatedTimeRange"`
	EndOfJobCreatedTimeRange   string `position:"Query" name:"EndOfJobCreatedTimeRange"`
	PipelineId                 string `position:"Query" name:"PipelineId"`
	ResourceOwnerAccount       string `position:"Query" name:"ResourceOwnerAccount"`
	State                      string `position:"Query" name:"State"`
	ResourceOwnerId            string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount               string `position:"Query" name:"OwnerAccount"`
	MaximumPageSize            string `position:"Query" name:"MaximumPageSize"`
	OwnerId                    string `position:"Query" name:"OwnerId"`
	NextPageToken              string `position:"Query" name:"NextPageToken"`
}

type ListJobResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	NextPageToken string `json:"NextPageToken" xml:"NextPageToken"`
	JobList       struct {
		Job []struct {
			JobId        string `json:"JobId" xml:"JobId"`
			State        string `json:"State" xml:"State"`
			Code         string `json:"Code" xml:"Code"`
			Message      string `json:"Message" xml:"Message"`
			Percent      int64  `json:"Percent" xml:"Percent"`
			PipelineId   string `json:"PipelineId" xml:"PipelineId"`
			CreationTime string `json:"CreationTime" xml:"CreationTime"`
			FinishTime   string `json:"FinishTime" xml:"FinishTime"`
			Input        struct {
				Bucket   string `json:"Bucket" xml:"Bucket"`
				Location string `json:"Location" xml:"Location"`
				Object   string `json:"Object" xml:"Object"`
			} `json:"Input" xml:"Input"`
			Output struct {
				TemplateId         string `json:"TemplateId" xml:"TemplateId"`
				UserData           string `json:"UserData" xml:"UserData"`
				Rotate             string `json:"Rotate" xml:"Rotate"`
				VideoStreamMap     string `json:"VideoStreamMap" xml:"VideoStreamMap"`
				AudioStreamMap     string `json:"AudioStreamMap" xml:"AudioStreamMap"`
				DeWatermark        string `json:"DeWatermark" xml:"DeWatermark"`
				Priority           string `json:"Priority" xml:"Priority"`
				WaterMarkConfigUrl string `json:"WaterMarkConfigUrl" xml:"WaterMarkConfigUrl"`
				MergeConfigUrl     string `json:"MergeConfigUrl" xml:"MergeConfigUrl"`
				OutputFile         struct {
					Bucket   string `json:"Bucket" xml:"Bucket"`
					Location string `json:"Location" xml:"Location"`
					Object   string `json:"Object" xml:"Object"`
					RoleArn  string `json:"RoleArn" xml:"RoleArn"`
				} `json:"OutputFile" xml:"OutputFile"`
				M3U8NonStandardSupport struct {
					TS struct {
						Md5Support  bool `json:"Md5Support" xml:"Md5Support"`
						SizeSupport bool `json:"SizeSupport" xml:"SizeSupport"`
					} `json:"TS" xml:"TS"`
				} `json:"M3U8NonStandardSupport" xml:"M3U8NonStandardSupport"`
				Properties struct {
					Width      string `json:"Width" xml:"Width"`
					Height     string `json:"Height" xml:"Height"`
					Bitrate    string `json:"Bitrate" xml:"Bitrate"`
					Duration   string `json:"Duration" xml:"Duration"`
					Fps        string `json:"Fps" xml:"Fps"`
					FileSize   string `json:"FileSize" xml:"FileSize"`
					FileFormat string `json:"FileFormat" xml:"FileFormat"`
					Streams    struct {
						VideoStreamList struct {
							VideoStream []struct {
								Index          string `json:"Index" xml:"Index"`
								CodecName      string `json:"CodecName" xml:"CodecName"`
								CodecLongName  string `json:"CodecLongName" xml:"CodecLongName"`
								Profile        string `json:"Profile" xml:"Profile"`
								CodecTimeBase  string `json:"CodecTimeBase" xml:"CodecTimeBase"`
								CodecTagString string `json:"CodecTagString" xml:"CodecTagString"`
								CodecTag       string `json:"CodecTag" xml:"CodecTag"`
								Width          string `json:"Width" xml:"Width"`
								Height         string `json:"Height" xml:"Height"`
								HasBFrames     string `json:"HasBFrames" xml:"HasBFrames"`
								Sar            string `json:"Sar" xml:"Sar"`
								Dar            string `json:"Dar" xml:"Dar"`
								PixFmt         string `json:"PixFmt" xml:"PixFmt"`
								Level          string `json:"Level" xml:"Level"`
								Fps            string `json:"Fps" xml:"Fps"`
								AvgFPS         string `json:"AvgFPS" xml:"AvgFPS"`
								Timebase       string `json:"Timebase" xml:"Timebase"`
								StartTime      string `json:"StartTime" xml:"StartTime"`
								Duration       string `json:"Duration" xml:"Duration"`
								Bitrate        string `json:"Bitrate" xml:"Bitrate"`
								NumFrames      string `json:"NumFrames" xml:"NumFrames"`
								Lang           string `json:"Lang" xml:"Lang"`
								NetworkCost    struct {
									PreloadTime   string `json:"PreloadTime" xml:"PreloadTime"`
									CostBandwidth string `json:"CostBandwidth" xml:"CostBandwidth"`
									AvgBitrate    string `json:"AvgBitrate" xml:"AvgBitrate"`
								} `json:"NetworkCost" xml:"NetworkCost"`
							} `json:"VideoStream" xml:"VideoStream"`
						} `json:"VideoStreamList" xml:"VideoStreamList"`
						AudioStreamList struct {
							AudioStream []struct {
								Index          string `json:"Index" xml:"Index"`
								CodecName      string `json:"CodecName" xml:"CodecName"`
								CodecTimeBase  string `json:"CodecTimeBase" xml:"CodecTimeBase"`
								CodecLongName  string `json:"CodecLongName" xml:"CodecLongName"`
								CodecTagString string `json:"CodecTagString" xml:"CodecTagString"`
								CodecTag       string `json:"CodecTag" xml:"CodecTag"`
								SampleFmt      string `json:"SampleFmt" xml:"SampleFmt"`
								Samplerate     string `json:"Samplerate" xml:"Samplerate"`
								Channels       string `json:"Channels" xml:"Channels"`
								ChannelLayout  string `json:"ChannelLayout" xml:"ChannelLayout"`
								Timebase       string `json:"Timebase" xml:"Timebase"`
								StartTime      string `json:"StartTime" xml:"StartTime"`
								Duration       string `json:"Duration" xml:"Duration"`
								Bitrate        string `json:"Bitrate" xml:"Bitrate"`
								NumFrames      string `json:"NumFrames" xml:"NumFrames"`
								Lang           string `json:"Lang" xml:"Lang"`
							} `json:"AudioStream" xml:"AudioStream"`
						} `json:"AudioStreamList" xml:"AudioStreamList"`
						SubtitleStreamList struct {
							SubtitleStream []struct {
								Index string `json:"Index" xml:"Index"`
								Lang  string `json:"Lang" xml:"Lang"`
							} `json:"SubtitleStream" xml:"SubtitleStream"`
						} `json:"SubtitleStreamList" xml:"SubtitleStreamList"`
					} `json:"Streams" xml:"Streams"`
					Format struct {
						NumStreams     string `json:"NumStreams" xml:"NumStreams"`
						NumPrograms    string `json:"NumPrograms" xml:"NumPrograms"`
						FormatName     string `json:"FormatName" xml:"FormatName"`
						FormatLongName string `json:"FormatLongName" xml:"FormatLongName"`
						StartTime      string `json:"StartTime" xml:"StartTime"`
						Duration       string `json:"Duration" xml:"Duration"`
						Size           string `json:"Size" xml:"Size"`
						Bitrate        string `json:"Bitrate" xml:"Bitrate"`
					} `json:"Format" xml:"Format"`
				} `json:"Properties" xml:"Properties"`
				Clip struct {
					TimeSpan struct {
						Seek     string `json:"Seek" xml:"Seek"`
						Duration string `json:"Duration" xml:"Duration"`
					} `json:"TimeSpan" xml:"TimeSpan"`
				} `json:"Clip" xml:"Clip"`
				SuperReso struct {
					IsHalfSample string `json:"IsHalfSample" xml:"IsHalfSample"`
				} `json:"SuperReso" xml:"SuperReso"`
				SubtitleConfig struct {
					SubtitleList struct {
						Subtitle []struct {
							Map string `json:"Map" xml:"Map"`
						} `json:"Subtitle" xml:"Subtitle"`
					} `json:"SubtitleList" xml:"SubtitleList"`
					ExtSubtitleList struct {
						ExtSubtitle []struct {
							FontName string `json:"FontName" xml:"FontName"`
							CharEnc  string `json:"CharEnc" xml:"CharEnc"`
							Input3   struct {
								Bucket   string `json:"Bucket" xml:"Bucket"`
								Location string `json:"Location" xml:"Location"`
								Object   string `json:"Object" xml:"Object"`
							} `json:"Input" xml:"Input"`
						} `json:"ExtSubtitle" xml:"ExtSubtitle"`
					} `json:"ExtSubtitleList" xml:"ExtSubtitleList"`
				} `json:"SubtitleConfig" xml:"SubtitleConfig"`
				TransConfig struct {
					TransMode               string `json:"TransMode" xml:"TransMode"`
					IsCheckReso             string `json:"IsCheckReso" xml:"IsCheckReso"`
					IsCheckResoFail         string `json:"IsCheckResoFail" xml:"IsCheckResoFail"`
					IsCheckVideoBitrate     string `json:"IsCheckVideoBitrate" xml:"IsCheckVideoBitrate"`
					IsCheckAudioBitrate     string `json:"IsCheckAudioBitrate" xml:"IsCheckAudioBitrate"`
					AdjDarMethod            string `json:"AdjDarMethod" xml:"AdjDarMethod"`
					IsCheckVideoBitrateFail string `json:"IsCheckVideoBitrateFail" xml:"IsCheckVideoBitrateFail"`
					IsCheckAudioBitrateFail string `json:"IsCheckAudioBitrateFail" xml:"IsCheckAudioBitrateFail"`
				} `json:"TransConfig" xml:"TransConfig"`
				MuxConfig struct {
					Segment struct {
						Duration string `json:"Duration" xml:"Duration"`
					} `json:"Segment" xml:"Segment"`
					Gif struct {
						Loop            string `json:"Loop" xml:"Loop"`
						FinalDelay      string `json:"FinalDelay" xml:"FinalDelay"`
						IsCustomPalette string `json:"IsCustomPalette" xml:"IsCustomPalette"`
						DitherMode      string `json:"DitherMode" xml:"DitherMode"`
					} `json:"Gif" xml:"Gif"`
				} `json:"MuxConfig" xml:"MuxConfig"`
				Audio struct {
					Codec      string `json:"Codec" xml:"Codec"`
					Profile    string `json:"Profile" xml:"Profile"`
					Samplerate string `json:"Samplerate" xml:"Samplerate"`
					Bitrate    string `json:"Bitrate" xml:"Bitrate"`
					Channels   string `json:"Channels" xml:"Channels"`
					Qscale     string `json:"Qscale" xml:"Qscale"`
					Volume     struct {
						Level  string `json:"Level" xml:"Level"`
						Method string `json:"Method" xml:"Method"`
					} `json:"Volume" xml:"Volume"`
				} `json:"Audio" xml:"Audio"`
				Video struct {
					Codec      string `json:"Codec" xml:"Codec"`
					Profile    string `json:"Profile" xml:"Profile"`
					Bitrate    string `json:"Bitrate" xml:"Bitrate"`
					Crf        string `json:"Crf" xml:"Crf"`
					Width      string `json:"Width" xml:"Width"`
					Height     string `json:"Height" xml:"Height"`
					Fps        string `json:"Fps" xml:"Fps"`
					Gop        string `json:"Gop" xml:"Gop"`
					Preset     string `json:"Preset" xml:"Preset"`
					ScanMode   string `json:"ScanMode" xml:"ScanMode"`
					Bufsize    string `json:"Bufsize" xml:"Bufsize"`
					Maxrate    string `json:"Maxrate" xml:"Maxrate"`
					PixFmt     string `json:"PixFmt" xml:"PixFmt"`
					Degrain    string `json:"Degrain" xml:"Degrain"`
					Qscale     string `json:"Qscale" xml:"Qscale"`
					Crop       string `json:"Crop" xml:"Crop"`
					Pad        string `json:"Pad" xml:"Pad"`
					MaxFps     string `json:"MaxFps" xml:"MaxFps"`
					BitrateBnd struct {
						Max string `json:"Max" xml:"Max"`
						Min string `json:"Min" xml:"Min"`
					} `json:"BitrateBnd" xml:"BitrateBnd"`
				} `json:"Video" xml:"Video"`
				Container struct {
					Format string `json:"Format" xml:"Format"`
				} `json:"Container" xml:"Container"`
				Encryption struct {
					Type    string `json:"Type" xml:"Type"`
					Id      string `json:"Id" xml:"Id"`
					Key     string `json:"Key" xml:"Key"`
					KeyUri  string `json:"KeyUri" xml:"KeyUri"`
					KeyType string `json:"KeyType" xml:"KeyType"`
					SkipCnt string `json:"SkipCnt" xml:"SkipCnt"`
				} `json:"Encryption" xml:"Encryption"`
				WaterMarkList struct {
					WaterMark []struct {
						WaterMarkTemplateId string `json:"WaterMarkTemplateId" xml:"WaterMarkTemplateId"`
						Width               string `json:"Width" xml:"Width"`
						Height              string `json:"Height" xml:"Height"`
						Dx                  string `json:"Dx" xml:"Dx"`
						Dy                  string `json:"Dy" xml:"Dy"`
						ReferPos            string `json:"ReferPos" xml:"ReferPos"`
						Type                string `json:"Type" xml:"Type"`
						InputFile           struct {
							Bucket   string `json:"Bucket" xml:"Bucket"`
							Location string `json:"Location" xml:"Location"`
							Object   string `json:"Object" xml:"Object"`
						} `json:"InputFile" xml:"InputFile"`
					} `json:"WaterMark" xml:"WaterMark"`
				} `json:"WaterMarkList" xml:"WaterMarkList"`
				MergeList struct {
					Merge []struct {
						MergeURL string `json:"MergeURL" xml:"MergeURL"`
						Start    string `json:"Start" xml:"Start"`
						Duration string `json:"Duration" xml:"Duration"`
						RoleArn  string `json:"RoleArn" xml:"RoleArn"`
					} `json:"Merge" xml:"Merge"`
				} `json:"MergeList" xml:"MergeList"`
				OpeningList struct {
					Merge []struct {
						OpenUrl string `json:"openUrl" xml:"openUrl"`
						Start   string `json:"Start" xml:"Start"`
						Width   string `json:"Width" xml:"Width"`
						Height  string `json:"Height" xml:"Height"`
					} `json:"Merge" xml:"Merge"`
				} `json:"OpeningList" xml:"OpeningList"`
				TailSlateList struct {
					Merge []struct {
						TailUrl       string `json:"TailUrl" xml:"TailUrl"`
						Start         string `json:"Start" xml:"Start"`
						BlendDuration string `json:"BlendDuration" xml:"BlendDuration"`
						Width         string `json:"Width" xml:"Width"`
						Height        string `json:"Height" xml:"Height"`
						IsMergeAudio  bool   `json:"IsMergeAudio" xml:"IsMergeAudio"`
						BgColor       string `json:"BgColor" xml:"BgColor"`
					} `json:"Merge" xml:"Merge"`
				} `json:"TailSlateList" xml:"TailSlateList"`
			} `json:"Output" xml:"Output"`
			MNSMessageResult struct {
				MessageId    string `json:"MessageId" xml:"MessageId"`
				ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
				ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
			} `json:"MNSMessageResult" xml:"MNSMessageResult"`
		} `json:"Job" xml:"Job"`
	} `json:"JobList" xml:"JobList"`
}

func CreateListJobRequest() (request *ListJobRequest) {
	request = &ListJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ListJob", "", "")
	return
}

func CreateListJobResponse() (response *ListJobResponse) {
	response = &ListJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
