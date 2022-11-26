package ops

import (
	_ "embed"
	"testing"

	"barglvojtech.net/govm/pkg/internal/resources"
	"barglvojtech.net/govm/pkg/internal/testutil"
)

func TestListOp(t *testing.T) {
	op := NewListOp(
		setHttpClientToListOp(&testutil.HttpClient{
			T:           t,
			ExpectedUrl: listVersionsUrl.String(),
			Content:     resources.GoDevDlContent,
		}),
	)

	if err := op.Process(); err != nil {
		t.Fatalf("processed with error: %v", err)
	}

	if l := len(op.Versions()); l != 261 {
		t.Errorf("expected number of version = %d, got %d", 261, l)
	}
}
