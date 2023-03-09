package File

import (
	
	"testing"
)

func Test_filedownload_UrlDownload(t *testing.T) {
	fd := &filedownload{}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		
		args    args
		
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				url: "http://www.yahoo.com/image_to_read.jpg",
			},
			wantErr: true,
		},
		{
			name: "Test2",
			args: args{
				url: "/Users/konjeti.harika/Downloads/Screenshot.jpeg",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		
			
			_, err := fd.UrlDownload(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("filedownload.UrlDownload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			
		}
	}

