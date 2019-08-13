package configuration

import (
	"github.com/NYTimes/video-transcoding-api/db"
	"github.com/NYTimes/video-transcoding-api/provider/bitmovinnewsdk/internal/configuration/codec"
	"github.com/NYTimes/video-transcoding-api/provider/bitmovinnewsdk/internal/types"
	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/pkg/errors"
)

// VP8Vorbis is a configuration service for content in this codec pair
type VP8Vorbis struct {
	api *bitmovin.BitmovinApi
}

// NewVP8Vorbis returns a service for managing VP8 / Vorbis configurations
func NewVP8Vorbis(api *bitmovin.BitmovinApi) *VP8Vorbis {
	return &VP8Vorbis{api: api}
}

// Create will create a new VP8 configuration based on a preset
func (c *VP8Vorbis) Create(preset db.Preset) (string, error) {
	audCfgID, err := codec.NewVorbis(c.api, preset.Audio.Bitrate)
	if err != nil {
		return "", err
	}

	vidCfgID, err := codec.NewVP8(c.api, preset, customDataWith(audCfgID, preset.Container))
	if err != nil {
		return "", err
	}

	return vidCfgID, nil
}

// Get retrieves audio / video configuration with a presetID
func (c *VP8Vorbis) Get(cfgID string) (bool, Details, error) {
	vidCfg, customData, err := c.vidConfigWithCustomDataFrom(cfgID)
	if err != nil {
		return false, Details{}, err
	}

	audCfgID, err := AudCfgIDFrom(customData)
	if err != nil {
		return false, Details{}, err
	}

	audCfg, err := c.api.Encoding.Configurations.Audio.Vorbis.Get(audCfgID)
	if err != nil {
		return false, Details{}, errors.Wrapf(err, "getting the audio config with ID %q", audCfgID)
	}

	return true, Details{vidCfg, audCfg, customData}, nil
}

// Delete removes the audio / video configurations
func (c *VP8Vorbis) Delete(cfgID string) (found bool, e error) {
	customData, err := c.vidCustomDataFrom(cfgID)
	if err != nil {
		return found, err
	}

	audCfgID, err := AudCfgIDFrom(customData)
	if err != nil {
		return found, err
	}

	audCfg, err := c.api.Encoding.Configurations.Audio.Vorbis.Get(audCfgID)
	if err != nil {
		return found, errors.Wrap(err, "retrieving audio configuration")
	}
	found = true

	_, err = c.api.Encoding.Configurations.Audio.Vorbis.Delete(audCfg.Id)
	if err != nil {
		return found, errors.Wrap(err, "removing the audio config")
	}

	_, err = c.api.Encoding.Configurations.Video.Vp8.Delete(cfgID)
	if err != nil {
		return found, errors.Wrap(err, "removing the video config")
	}

	return found, nil
}

func (c *VP8Vorbis) vidConfigWithCustomDataFrom(cfgID string) (*model.Vp8VideoConfiguration, types.CustomData, error) {
	vidCfg, err := c.api.Encoding.Configurations.Video.Vp8.Get(cfgID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "retrieving configuration with config ID")
	}

	data, err := c.vidCustomDataFrom(vidCfg.Id)
	if err != nil {
		return nil, nil, err
	}

	return vidCfg, data, nil
}

func (c *VP8Vorbis) vidCustomDataFrom(cfgID string) (types.CustomData, error) {
	data, err := c.api.Encoding.Configurations.Video.Vp8.Customdata.Get(cfgID)
	if err != nil {
		return nil, errors.Wrap(err, "retrieving custom data with config ID")
	}

	return data.CustomData, nil
}
