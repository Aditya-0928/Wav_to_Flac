package tests

import (
	"testing"
	"wav-to-flac-converter/audio"
)

// Sample WAV header for a 1 second silent mono WAV file (44 bytes header + 44100 bytes of silence)
var sampleWAV = []byte{
	'R', 'I', 'F', 'F', // RIFF
	0x00, 0x00, 0x00, 0x00, // file size (placeholder)
	'W', 'A', 'V', 'E', // WAVE
	'f', 'm', 't', ' ', // fmt
	0x10, 0x00, 0x00, 0x00, // subchunk1 size (16 for PCM)
	0x01, 0x00, // audio format (1 for PCM)
	0x01, 0x00, // number of channels (1 for mono)
	0x44, 0xac, 0x00, 0x00, // sample rate (44100)
	0x88, 0x58, 0x01, 0x00, // byte rate (44100 * 1 * 16/8)
	0x02, 0x00, // block align (1 * 16/8)
	0x10, 0x00, // bits per sample (16)
	'd', 'a', 't', 'a', // data
	0x00, 0x00, 0x00, 0x00, // data size (placeholder)
}

func TestConvertWAVToFLAC(t *testing.T) {
	// Prepare sample WAV data with a proper data size
	dataSize := len(sampleWAV) - 44 // 44 bytes header
	sampleWAV[40] = byte(dataSize)
	sampleWAV[41] = byte(dataSize >> 8)
	sampleWAV[42] = byte(dataSize >> 16)
	sampleWAV[43] = byte(dataSize >> 24)

	flacData, err := audio.ConvertWAVToFLAC(sampleWAV)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(flacData) == 0 {
		t.Fatal("Expected FLAC data, got empty slice")
	}
}

func TestInvalidWAV(t *testing.T) {
	invalidWAV := []byte("not a wav file")
	_, err := audio.ConvertWAVToFLAC(invalidWAV)
	if err == nil {
		t.Fatal("Expected error for invalid WAV data, got nil")
	}
}
