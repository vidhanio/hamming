package hamming

import "testing"

func TestNewCode(t *testing.T) {
	code := New()
	if code.IsCorrupt() {
		t.Error("New code is corrupt")
	}
}

func TestCorruptedCode(t *testing.T) {
	code := New()
	code.Corrupt()
	if !code.IsCorrupt() {
		t.Error("Corrupted code is not corrupt")
	}
}

func TestFixCorruption(t *testing.T) {
	code := New()
	code.Corrupt()
	code.FixCorruption()
	if code.IsCorrupt() {
		t.Error("Fixed code is corrupt")
	}
}
