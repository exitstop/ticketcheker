package lCommon

import (
	"io"
	"os"
	//"os/exec"
	//"runtime"
	"github.com/hajimehoshi/oto"
	"github.com/hajimehoshi/go-mp3"
	"time"
	// "strconv"
	//"sync/atomic"
    "fmt"
)




func PlayMusic(path string, div int64) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}
	defer d.Close()

	p, err := oto.NewPlayer(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}	
	defer p.Close()

	if _, err := io.CopyN(p, d, div * 8192 * 22  ); err != nil {
		return err
	}	

	return nil
}

func CheckTimeoit(c1 chan time.Time){


    for {
      select {
        case msg1 := <- c1:
           fmt.Println("\nmain            time: ", msg1.Format("2006/01/02 15:04:05"))
        case <- time.After(time.Second*35):
          fmt.Println("\ntimeout")

          //go LogPrint()
          //LogStop()

          if err := PlayMusic("./sound/soundBelt.mp3", 3 ) ; err != nil {
          }
      }
    }
}

