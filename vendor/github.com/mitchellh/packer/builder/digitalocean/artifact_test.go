package digitalocean

import (
	"testing"

	"github.com/hashicorp/packer/packer"
)

func TestArtifact_Impl(t *testing.T) {
	var raw interface{}
	raw = &Artifact{}
	if _, ok := raw.(packer.Artifact); !ok {
		t.Fatalf("Artifact should be artifact")
	}
}

func TestArtifactId(t *testing.T) {
	a := &Artifact{"packer-foobar", 42, []string{"sfo", "tor1"}, nil}
	expected := "sfo,tor1:42"

	if a.Id() != expected {
		t.Fatalf("artifact ID should match: %v", expected)
	}
}

func TestArtifactIdWithoutMultipleRegions(t *testing.T) {
	a := &Artifact{"packer-foobar", 42, []string{"sfo"}, nil}
	expected := "sfo:42"

	if a.Id() != expected {
		t.Fatalf("artifact ID should match: %v", expected)
	}
}

func TestArtifactString(t *testing.T) {
	a := &Artifact{"packer-foobar", 42, []string{"sfo", "tor1"}, nil}
	expected := "A snapshot was created: 'packer-foobar' (ID: 42) in regions 'sfo,tor1'"

	if a.String() != expected {
		t.Fatalf("artifact string should match: %v", expected)
	}
}

func TestArtifactStringWithoutMultipleRegions(t *testing.T) {
	a := &Artifact{"packer-foobar", 42, []string{"sfo"}, nil}
	expected := "A snapshot was created: 'packer-foobar' (ID: 42) in regions 'sfo'"

	if a.String() != expected {
		t.Fatalf("artifact string should match: %v", expected)
	}
}
