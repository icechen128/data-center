package temp

import "testing"

func TestParseRss(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test",
			args:    args{url: "https://rsshub.app/pcr/news"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ParseRss(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("ParseRss() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
