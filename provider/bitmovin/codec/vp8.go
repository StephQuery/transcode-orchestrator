package codec

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/db"
)

type CodecVP8 struct {
	codec
	cfg model.Vp8VideoConfiguration
}

func (c *CodecVP8) set(preset db.Preset) (ok bool) {
	return c.setVideo(VideoPTR{
		Name:         &c.cfg.Name,
		Width:        &c.cfg.Width,
		Height:       &c.cfg.Height,
		Bitrate:      &c.cfg.Bitrate,
		EncodingMode: &c.cfg.EncodingMode,
	}, preset)
}