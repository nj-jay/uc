package play

import (

	"io"
	"os"
    "github.com/hajimehoshi/oto"
	"github.com/qiniu/audio"
)

func Play(filename string) error {

	f, err := os.Open(filename)

		if err != nil {
			return err
		}

		defer f.Close()

		d, _, err := audio.Decode(f)

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

