package product

import (
	"os"
	"testing"
)

func init() {
	os.Chdir("../../../test/testdata")
	// put this teardown maybe? os.Chdir("..")
}

func TestLoadProduct(t *testing.T) {
	t.Fatalf("Expected implementation")
}

func TestSaveProduct(t *testing.T) {
	t.Fatalf("Expected implementation")
}

func TestAddApplication(t *testing.T) {
	t.Fatalf("Expected implementation")
}
