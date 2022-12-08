package metrics

import "testing"

func TestPriority(t *testing.T) {
	cases := []struct {
		labels []string
		wantP  string
	}{
		{
			labels: []string{},
			wantP:  "P?",
		},
		{
			labels: []string{"S-Critical", "O-Frequent"},
			wantP:  "P1",
		},
		{
			labels: []string{"S-Critical", "O-Occasional"},
			wantP:  "P1",
		},
		{
			labels: []string{"S-Major", "O-Frequent"},
			wantP:  "P1",
		},
		{
			labels: []string{"S-Major", "O-Frequent"},
			wantP:  "P1",
		},
		{
			labels: []string{"S-Minor", "O-Frequent"},
			wantP:  "P2",
		},
		{
			labels: []string{"S-Major", "O-Occasional"},
			wantP:  "P2",
		},
		{
			labels: []string{"S-Critical", "O-Uncommon"},
			wantP:  "P2",
		},
		{
			labels: []string{"S-Major", "O-Uncommon"},
			wantP:  "P3",
		},
		{
			labels: []string{"S-Minor", "O-Occasional"},
			wantP:  "P3",
		},
		{
			labels: []string{"S-Tolerable", "O-Frequent"},
			wantP:  "P3",
		},
		{
			labels: []string{"S-Tolerable", "O-Occasional"},
			wantP:  "P4",
		},
		{
			labels: []string{"S-Minor", "O-Uncommon"},
			wantP:  "P4",
		},
		{
			labels: []string{"S-Tolerable", "O-Uncommon"},
			wantP:  "P4",
		},
		// overrides
		{
			labels: []string{"S-Tolerable", "O-Uncommon", "P2"},
			wantP:  "P2",
		},
		// missing fields
		{
			labels: []string{"O-Frequent"},
			wantP:  "P?",
		},
		{
			labels: []string{"S-Minor"},
			wantP:  "P?",
		},
	}
	for _, c := range cases {
		gotP := Priority(c.labels)
		if gotP != c.wantP {
			t.Errorf("got %v want %v with labels %v", gotP, c.wantP, c.labels)
		}
	}
}
