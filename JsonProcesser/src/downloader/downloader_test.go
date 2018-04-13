package downloader

import (
	"fmt"
	"testing"
)

func TestDownLoad(t *testing.T) {
	url := "https://contestimg.wish.com/api/webimage/568e16a72dfd0133cb3f7a79-large"
	filename := fmt.Sprintf("%s/%s.png", "/Volumes/CPSC585DATA/downloadtest/", "1")
	if err := download(url, filename); err != nil {
		t.Error(err)
	}
}

func TestDownloader(t *testing.T) {
	urls := []string{
		"https://contestimg.wish.com/api/webimage/568e16a72dfd0133cb3f7a79-large",
		"https://contestimg.wish.com/api/webimage/5452f9925f313f502bf119ff-large",
		"https://contestimg.wish.com/api/webimage/540584051d2d435c5a300a82-large",
		"https://contestimg.wish.com/api/webimage/540c6760f4bdc40bcc10296d-large",
		"https://contestimg.wish.com/api/webimage/5447777d4ad3ab71267befe2-large",
		"https://contestimg.wish.com/api/webimage/53d9bdaa302e550f1889f7a4-large",
		"https://contestimg.wish.com/api/webimage/579a47cf58c0db798a399b4a-large",
		"https://contestimg.wish.com/api/webimage/5671317a30a4d06e41fecd7e-large",
		"https://contestimg.wish.com/api/webimage/54450c9866ccaa0fae0060f5-large",
		"https://contestimg.wish.com/api/webimage/561b7ffe388ed1104c4aedb5-large",
		"https://contestimg.wish.com/api/webimage/552cd3c7e508977e333f2c09-large",
		"https://contestimg.wish.com/api/webimage/542127cef8abc86673909d6a-large",
		"https://contestimg.wish.com/api/webimage/54850f5b7379236a3f76394e-large",
		"https://contestimg.wish.com/api/webimage/5459f02b5cf7460f37f6ff88-large",
		"https://contestimg.wish.com/api/webimage/57771d3676038f3c3521ae58-large",
		"https://contestimg.wish.com/api/webimage/569da0226d57220d8ab1556a-large",
		"https://contestimg.wish.com/api/webimage/570ddf0661d6cf20c0ac4f13-large",
		"https://contestimg.wish.com/api/webimage/53dedbe56c402e0ee3889fba-large",
		"https://contestimg.wish.com/api/webimage/55569a72a9c3af165e1ba7d0-large",
		"https://contestimg.wish.com/api/webimage/549d0e8100e90a102575ef6f-large",
		"https://contestimg.wish.com/api/webimage/573d9842ae875e5d389b95e6-large",
		"https://contestimg.wish.com/api/webimage/553777dad39b0c14521fa871-large",
		"https://contestimg.wish.com/api/webimage/53fc30c71f50635c619ac289-large",
		"https://contestimg.wish.com/api/webimage/547d7b579d9e6f2661ae5d8a-large",
		"https://contestimg.wish.com/api/webimage/544e441fe268791092057d06-large",
		"https://contestimg.wish.com/api/webimage/53e096c8d91139446ced7e4a-large",
		"https://contestimg.wish.com/api/webimage/559a3ccf9cb07c404f7bb10a-large",
		"https://contestimg.wish.com/api/webimage/5485b2c25d7e297a2c24562d-large",
		"https://contestimg.wish.com/api/webimage/56e79d670642785c59603385-large",
		"https://contestimg.wish.com/api/webimage/53d1ef9c4fd4990ee4f32fa5-large",
		"https://contestimg.wish.com/api/webimage/55c964729ace8f477c964b61-large",
		"https://contestimg.wish.com/api/webimage/54588e4e3dabbe0eda30f6ce-large",
		"https://contestimg.wish.com/api/webimage/5453533c4096a6167f7683ef-large",
		"https://contestimg.wish.com/api/webimage/54430e1e4ad3ab5af39c4dd5-large",
		"https://contestimg.wish.com/api/webimage/545a1625bf3ff010c7dd96cd-large",
		"https://contestimg.wish.com/api/webimage/543fe93a232fae36509bda91-large",
		"https://contestimg.wish.com/api/webimage/5370d3c0017cff0c8e1240fb-large",
		"https://contestimg.wish.com/api/webimage/5479d6ae90c77622b9a4f4e3-large",
		"https://contestimg.wish.com/api/webimage/53c7dbf1858d6d0f0beada11-large",
		"https://contestimg.wish.com/api/webimage/5399011ba9fa507566dfe85a-large",
		"https://contestimg.wish.com/api/webimage/5763370cd96e8612f4cd4543-large",
		"https://contestimg.wish.com/api/webimage/564697c7f7793242bbb27a65-large",
		"https://contestimg.wish.com/api/webimage/542c1f227541ce6a01460be6-large",
	}

	downloader := New(64)
	defer downloader.Close()
	for k, url := range urls {
		filename := fmt.Sprintf("/Volumes/CPSC585DATA/downloadtest/%d.png", k)
		downloader.AddJob(url, filename)
	}
}
