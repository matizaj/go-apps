package streamer

type Encoder interface {
	EncodeToMP4(v *Video, baseFileName string) error
}

type VideoEncoder struct{}

func (ve *VideoEncoder) EncodeToMP4(v *Video, baseFileName string) error {
	return nil
}
