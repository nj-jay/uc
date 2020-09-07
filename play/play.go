package play

import (
	"io"
	_ "io/ioutil"
	_ "log"

	"github.com/hajimehoshi/oto"
	_ "net/http"
	"os"

	"github.com/qiniu/audio"
	_ "github.com/qiniu/audio/mp3"
	_ "github.com/qiniu/audio/wav"
	_ "github.com/qiniu/audio/wav/adpcm"
)

func PlayAudio(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer f.Close()
	d, _, err := audio.Decode(f)
	//fmt.Println(d)
	if err != nil {
		return err
	}

	defer f.Close()
	c, err := oto.NewContext(d.SampleRate(), d.Channels(), d.BytesPerSample(), 8192)
	if err != nil {
		return err
	}
	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()

	_, err = io.Copy(p, d)
	return err


}
