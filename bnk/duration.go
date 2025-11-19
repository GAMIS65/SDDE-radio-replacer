package bnk

import (
	"encoding/binary"
	"math"
)

// UpdateTrackDurationBySourceID finds a music track with the given source ID and updates its duration.
// The offset parameter should be the offset from the start of the track descriptor to the duration value.
// Duration should be in milliseconds.
func (sec *ObjectHierarchySection) UpdateTrackDurationBySourceID(sourceID uint32, duration uint32, offset int) bool {
	for _, track := range sec.musicTracks {
		// Read the source ID from the track data at the specified offset
		if offset+4 > len(track.Data) {
			continue
		}
		trackSourceID := binary.LittleEndian.Uint32(track.Data[offset:])
		if trackSourceID == sourceID {
			// Found the matching track, update its duration
			// The duration offset should be relative to the track data start
			durationOffset := TRACK_SRC_DURATION_OFFSET - OBJECT_DESCRIPTOR_BYTES
			if durationOffset+8 <= len(track.Data) {
				// Duration is stored as a 64-bit float (double)
				durationDouble := float64(duration)
				bytes := make([]byte, 8)
				binary.LittleEndian.PutUint64(bytes, math.Float64bits(durationDouble))
				copy(track.Data[durationOffset:durationOffset+8], bytes)
				return true
			}
		}
	}
	return false
}

// UpdateSegmentDurationByChildID finds a music segment with the given child ID and updates both
// its duration and end position.
func (sec *ObjectHierarchySection) UpdateSegmentDurationByChildID(childID uint32, duration uint32) bool {
	for _, segment := range sec.musicSegments {
		// Read the child ID from the segment data at the specified offset
		childIDOffset := SEGMENT_CHILD_ID_OFFSET - OBJECT_DESCRIPTOR_BYTES
		if childIDOffset+4 > len(segment.Data) {
			continue
		}
		segmentChildID := binary.LittleEndian.Uint32(segment.Data[childIDOffset:])
		if segmentChildID == childID {
			// Found the matching segment, update duration and end position
			durationOffset := SEGMENT_DURATION_OFFSET - OBJECT_DESCRIPTOR_BYTES
			endPositionOffset := SEGMENT_END_POSITION_OFFSET - OBJECT_DESCRIPTOR_BYTES

			if durationOffset+8 <= len(segment.Data) && endPositionOffset+8 <= len(segment.Data) {
				// Both duration and end position are 64-bit floats (doubles)
				durationDouble := float64(duration)
				bytes := make([]byte, 8)
				binary.LittleEndian.PutUint64(bytes, math.Float64bits(durationDouble))
				copy(segment.Data[durationOffset:durationOffset+8], bytes)
				copy(segment.Data[endPositionOffset:endPositionOffset+8], bytes)
				return true
			}
		}
	}
	return false
}

// FindTrackIDBySourceID finds the track object ID that has the given source ID.
func (sec *ObjectHierarchySection) FindTrackIDBySourceID(sourceID uint32) (uint32, bool) {
	for trackID, track := range sec.musicTracks {
		sourceIDOffset := TRACK_SOURCEID_OFFSET - OBJECT_DESCRIPTOR_BYTES
		if sourceIDOffset+4 <= len(track.Data) {
			trackSourceID := binary.LittleEndian.Uint32(track.Data[sourceIDOffset:])
			if trackSourceID == sourceID {
				return trackID, true
			}
		}
	}
	return 0, false
}
