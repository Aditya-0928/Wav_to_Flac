package audio

import (
	"os"
	"os/exec"
)

// ConvertWAVToFLAC uses ffmpeg to convert WAV data to FLAC format.
func ConvertWAVToFLAC(wavData []byte) ([]byte, error) {
	// Create a temporary file for the WAV data
	wavFile, err := os.CreateTemp("", "*.wav")
	if err != nil {
		return nil, err
	}
	defer os.Remove(wavFile.Name()) // Clean up

	// Write the WAV data to the temporary file
	if _, err := wavFile.Write(wavData); err != nil {
		return nil, err
	}
	wavFile.Close() // Close the file to flush the data

	// Create a temporary file for the FLAC output
	flacFile, err := os.CreateTemp("", "*.flac")
	if err != nil {
		return nil, err
	}
	defer os.Remove(flacFile.Name()) // Clean up

	// Run ffmpeg to convert WAV to FLAC
	cmd := exec.Command("ffmpeg", "-i", wavFile.Name(), flacFile.Name())
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	// Read the FLAC data from the temporary file
	flacData, err := os.ReadFile(flacFile.Name())
	if err != nil {
		return nil, err
	}

	return flacData, nil
}
