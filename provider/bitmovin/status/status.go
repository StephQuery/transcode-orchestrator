package status

import (
	"fmt"
	"time"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/query"
	"github.com/cbsinteractive/transcode-orchestrator/provider"
	"github.com/pkg/errors"
)

// ToProviderStatus maps Bitmovin's status to the local provider status
func ToProviderStatus(status model.Status) provider.Status {
	switch status {
	case model.Status_CREATED, model.Status_QUEUED:
		return provider.StatusQueued
	case model.Status_RUNNING:
		return provider.StatusStarted
	case model.Status_FINISHED:
		return provider.StatusFinished
	case model.Status_ERROR:
		return provider.StatusFailed
	case model.Status_CANCELED:
		return provider.StatusCanceled
	default:
		return provider.StatusUnknown
	}
}

// EnrichSourceInfo adds information about the video input source to a job status and returns a new status or an error
func EnrichSourceInfo(api *bitmovin.BitmovinApi, s provider.JobStatus) (provider.JobStatus, error) {
	inStreams, err := api.Encoding.Encodings.Streams.List(s.ProviderJobID, func(params *query.StreamListQueryParams) {
		params.Limit = 1
		params.Offset = 0
	})
	if err != nil {
		return s, errors.Wrap(err, "retrieving input streams from the Bitmovin API")
	}
	if len(inStreams.Items) == 0 {
		return s, fmt.Errorf("no streams found for encodingID %s", s.ProviderJobID)
	}

	inStream := inStreams.Items[0]
	streamInput, err := api.Encoding.Encodings.Streams.Input.Get(s.ProviderJobID, inStream.Id)
	if err != nil {
		return s, errors.Wrap(err, "retrieving stream input details from the Bitmovin API")
	}

	var (
		vidCodec      string
		width, height int64
	)
	if len(streamInput.VideoStreams) > 0 {
		vidStreamInput := streamInput.VideoStreams[0]
		vidCodec = vidStreamInput.Codec
		width = int64(int32Value(vidStreamInput.Width))
		height = int64(int32Value(vidStreamInput.Height))
	}

	s.SourceInfo = provider.SourceInfo{
		Duration:   time.Duration(floatValue(streamInput.Duration) * float64(time.Second)),
		Width:      width,
		Height:     height,
		VideoCodec: vidCodec,
	}

	return s, nil
}

func floatValue(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

func int32Value(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}
