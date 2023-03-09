package Zip

import (
	
	"io"

	"testing"
	"url/Downloader/File"
	"url/Downloader/Web"
)

func TestZip_Archive(t *testing.T) {
	z := &Zip{}
	r1,_:= Web.NewWebDownloader().UrlDownload("http://www.yahoo.com/image_to_read.jpg")
	r2, _ := File.NewFileDownloader().UrlDownload("/Users/konjeti.harika/Downloads/Screenshot.jpeg")
	type args struct {
		names   []string
		readers []io.Reader
	}
	tests := []struct {
		name     string
		
		args     args
	
		wantErr  bool
	}{
		{
			name: "Test1",
			args: args{
				names: []string{
					"image_to_read.jpg",
					"",
				},
				readers: []io.Reader{
					r1,
					r2,
				},
			},
			wantErr: false,
		},
		
	}
	for _, tt := range tests {
		_, err := z.Archive(tt.args.names, tt.args.readers...)
		if (err != nil) != tt.wantErr {
			t.Errorf("Zip.Archive() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
	}
}

