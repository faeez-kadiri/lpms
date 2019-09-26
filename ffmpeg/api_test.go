package ffmpeg

import (
	"testing"
)

func TestTranscoderAPI_Stopped(t *testing.T) {

	// Test stopped transcoder
	tc := NewTranscoder()
	tc.StopTranscoder()
	in := &TranscodeOptionsIn{}
	_, err := tc.Transcode(in, nil)
	if err != ErrTranscoderStp {
		t.Errorf("Unexpected error; wanted %v but got %v", ErrTranscoderStp, err)
	}

	// test somehow munged transcoder handle
	tc2 := NewTranscoder()
	tc2.handle = nil // technically this leaks memory ... OK for test
	_, err = tc2.Transcode(in, nil)
	if err != ErrTranscoderStp {
		t.Errorf("Unexpected error; wanted %v but got %v", ErrTranscoderStp, err)
	}
}

func TestTranscoderAPI_TooManyOutputs(t *testing.T) {

	out := make([]TranscodeOptions, 11)
	for i, _ := range out {
		out[i].VideoEncoder = ComponentOptions{Name: "drop"}
	}
	in := &TranscodeOptionsIn{}
	tc := NewTranscoder()
	_, err := tc.Transcode(in, out)
	if err == nil || err.Error() != "Too many outputs" {
		t.Error("Expected 'Too many outputs', got ", err)
	}
}
