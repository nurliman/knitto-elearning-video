package transcode

import (
	"fmt"
	"log"

	ffmpeg "github.com/floostack/transcoder/ffmpeg"
	"github.com/rs/xid"
)

var (
	format       = "mp4"
	frameRate    = 24
	overwrite    = true
	videoCodec   = "libx264"
	videoFilter  = "scale=-2:1080"
	videoBitRate = "4300k"
	videoProfile = "main"
	movFlags     = "faststart"
	preset       = "fast"
	skipAudio    = true
	bufferSize   = 8600000
	maxRate      = 4300000
)

// Video transcode video
func Video(videoPath *string) {

	opts := ffmpeg.Options{
		FrameRate:       &frameRate,
		OutputFormat:    &format,
		Overwrite:       &overwrite,
		VideoCodec:      &videoCodec,
		VideoFilter:     &videoFilter,
		VideoBitRate:    &videoBitRate,
		MovFlags:        &movFlags,
		Preset:          &preset,
		VideoProfile:    &videoProfile,
		SkipAudio:       &skipAudio,
		BufferSize:      &bufferSize,
		VideoMaxBitRate: &maxRate,
		ExtraArgs:       map[string]interface{}{"-x264opts": "'keyint=48:min-keyint=48:no-scenecut'"},
	}

	ffmpegConf := &ffmpeg.Config{
		FfmpegBinPath:   "/usr/bin/ffmpeg",
		FfprobeBinPath:  "/usr/bin/ffprobe",
		ProgressEnabled: true,
		Verbose:         true,
	}

	progress, err := ffmpeg.
		New(ffmpegConf).
		Input(*videoPath).
		Output(fmt.Sprintf("./output/%s.mp4", xid.New().String())).
		WithOptions(opts).
		Start(opts)

	if err != nil {
		log.Fatal(err)
	}

	for msg := range progress {
		log.Printf("%+v", msg)
	}
}
