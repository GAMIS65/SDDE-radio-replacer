package util

import (
	"encoding/binary"
	"fmt"
	"io"
)

// WAVDuration extracts the duration in milliseconds from a PCM WAV file.
// It reads the headers to determine the sample rate and data size.
func WAVDuration(r io.ReadSeeker) (uint32, error) {
	// --- Read and validate RIFF header ---
	var riffHeader [4]byte
	if _, err := r.Read(riffHeader[:]); err != nil {
		return 0, fmt.Errorf("failed to read RIFF header: %w", err)
	}
	if string(riffHeader[:]) != "RIFF" {
		return 0, fmt.Errorf("invalid RIFF header (expected 'RIFF')")
	}

	// Skip file size (4 bytes)
	if _, err := r.Seek(4, io.SeekCurrent); err != nil {
		return 0, fmt.Errorf("failed to skip file size: %w", err)
	}

	// Read and validate WAVE header
	var waveHeader [4]byte
	if _, err := r.Read(waveHeader[:]); err != nil {
		return 0, fmt.Errorf("failed to read WAVE header: %w", err)
	}
	if string(waveHeader[:]) != "WAVE" {
		return 0, fmt.Errorf("invalid WAVE header (expected 'WAVE')")
	}

	// --- Parse chunks ---
	var (
		sampleRate    uint32
		byteRate      uint32
		blockAlign    uint16
		bitsPerSample uint16
		numChannels   uint16
	)

	// Find "fmt " chunk
	for {
		var chunkID [4]byte
		if _, err := r.Read(chunkID[:]); err != nil {
			return 0, fmt.Errorf("failed to read chunk ID: %w", err)
		}

		var chunkSize uint32
		if err := binary.Read(r, binary.LittleEndian, &chunkSize); err != nil {
			return 0, fmt.Errorf("failed to read chunk size: %w", err)
		}

		if string(chunkID[:]) == "fmt " {
			// Parse the "fmt " chunk
			var audioFormat uint16
			if err := binary.Read(r, binary.LittleEndian, &audioFormat); err != nil {
				return 0, fmt.Errorf("failed to read audio format: %w", err)
			}

			if err := binary.Read(r, binary.LittleEndian, &numChannels); err != nil {
				return 0, fmt.Errorf("failed to read numChannels: %w", err)
			}
			if err := binary.Read(r, binary.LittleEndian, &sampleRate); err != nil {
				return 0, fmt.Errorf("failed to read sampleRate: %w", err)
			}
			if err := binary.Read(r, binary.LittleEndian, &byteRate); err != nil {
				return 0, fmt.Errorf("failed to read byteRate: %w", err)
			}
			if err := binary.Read(r, binary.LittleEndian, &blockAlign); err != nil {
				return 0, fmt.Errorf("failed to read blockAlign: %w", err)
			}
			if err := binary.Read(r, binary.LittleEndian, &bitsPerSample); err != nil {
				return 0, fmt.Errorf("failed to read bitsPerSample: %w", err)
			}

			// Skip any remaining bytes in fmt chunk (e.g., for non-PCM formats)
			remaining := int64(chunkSize) - 16
			if remaining > 0 {
				if _, err := r.Seek(remaining, io.SeekCurrent); err != nil {
					return 0, fmt.Errorf("failed to skip extra fmt data: %w", err)
				}
			}

			// If chunk size is odd, skip padding byte
			if chunkSize%2 == 1 {
				_, _ = r.Seek(1, io.SeekCurrent)
			}

			break
		}

		// Skip this chunk and pad byte if needed
		if _, err := r.Seek(int64(chunkSize), io.SeekCurrent); err != nil {
			return 0, fmt.Errorf("failed to skip non-fmt chunk: %w", err)
		}
		if chunkSize%2 == 1 {
			_, _ = r.Seek(1, io.SeekCurrent)
		}
	}

	if sampleRate == 0 || byteRate == 0 {
		return 0, fmt.Errorf("invalid or missing sample rate/byte rate")
	}

	// --- Find "data" chunk ---
	for {
		var chunkID [4]byte
		if _, err := r.Read(chunkID[:]); err != nil {
			return 0, fmt.Errorf("failed to read chunk ID: %w", err)
		}

		var chunkSize uint32
		if err := binary.Read(r, binary.LittleEndian, &chunkSize); err != nil {
			return 0, fmt.Errorf("failed to read chunk size: %w", err)
		}

		if string(chunkID[:]) == "data" {
			// Duration (ms) = (dataBytes / byteRate) * 1000
			durationMs := (float64(chunkSize) / float64(byteRate)) * 1000
			return uint32(durationMs), nil
		}

		// Skip non-data chunk
		if _, err := r.Seek(int64(chunkSize), io.SeekCurrent); err != nil {
			return 0, fmt.Errorf("failed to skip non-data chunk: %w", err)
		}
		if chunkSize%2 == 1 {
			_, _ = r.Seek(1, io.SeekCurrent)
		}
	}
}
