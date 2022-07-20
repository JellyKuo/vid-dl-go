package dl

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
)

const basedir = "downloads/"
const ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36"

func Download(url string) uuid.UUID {
	id := uuid.New()
	os.MkdirAll(basedir+"/"+id.String(), 0755)
	go startDownload(id, url)
	return id
}

func startDownload(id uuid.UUID, url string) {
	downloadVideo(id, url)
}

// func startDownload(id uuid.UUID, url string) {
// 	m3u, err := downloadM3u(url)
// 	if err != nil {
// 		println("[DL] Download M3U Error. ID: " + id.String() + " Error: " + err.Error())
// 		return
// 	}
// 	origM3uFile, err := os.Create(basedir + "/" + id.String() + "/original.m3u")
// 	if err != nil {
// 		println("[DL] Writing original M3U Error. ID: " + id.String() + " Error: " + err.Error())
// 		return
// 	}
// 	origM3uFile.WriteString(m3u)
// 	origM3uFile.Close()

// 	lines := strings.Split(m3u, "\n")
// 	wg := &sync.WaitGroup{}
// 	for idx, line := range lines {
// 		if strings.HasPrefix(line, "#") {
// 			continue
// 		}
// 		wg.Add(1)
// 		go downloadVideo(id, line, idx, wg)

// 	}
// }

// func downloadM3u(url string) (string, error) {
// 	client := http.Client{}
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return "", err
// 	}
// 	req.Header.Set("User-Agent", ua)
// 	res, err := client.Do(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	if res.StatusCode != 200 {
// 		return "", errors.New("Status code " + fmt.Sprint(res.StatusCode))
// 	}
// 	defer res.Body.Close()
// 	m3u, err := ioutil.ReadAll(res.Body)
// 	return string(m3u), err
// }

func downloadVideo(id uuid.UUID, url string) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", ua)
	if err != nil {
		println("[DL] Download Video Error. ID: " + id.String() + " Error: " + err.Error())
		return
	}
	res, err := client.Do(req)
	if err != nil {
		println("[DL] Download Video Error. ID: " + id.String() + " Error: " + err.Error())
		return
	}
	if res.StatusCode != 200 {
		println("[DL] Download Video Error. ID: " + id.String() + " Error: Status code " + fmt.Sprint(res.StatusCode))
		return
	}
	defer res.Body.Close()
	f, err := os.Create(basedir + "/" + id.String() + "/" + "downloading.mp4")
	if err != nil {
		println("[DL] Download Video Error. ID: " + id.String() + " Error: " + err.Error())
		return
	}
	io.Copy(f, res.Body)
	f.Close()
	os.Rename(basedir+"/"+id.String()+"/"+"downloading.mp4", basedir+"/"+id.String()+"/"+"video.mp4")
}
