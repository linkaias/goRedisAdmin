package exoprt_utils

import (
	"os"
	"testing"
)

// TestExportUtils_ExportFile verifies that SaveFile creates a JSON export file
// from in-memory snapshot data.
func TestExportUtils_ExportFile(t *testing.T) {
	cont := &ExportUtils{
		data: []*ExportRedisDataModel{
			{Key: "k1", Type: "string", Value: "v1"},
		},
	}

	filePath, err := cont.SaveFile()
	if err != nil {
		t.Fatalf("SaveFile failed: %v", err)
	}

	if _, err := os.Stat(filePath); err != nil {
		t.Fatalf("export file not created: %v", err)
	}

	// Cleanup generated test artifact.
	_ = os.Remove(filePath)
}
