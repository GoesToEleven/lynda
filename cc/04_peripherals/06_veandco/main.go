package main

// typedef unsigned char Uint8;
// void SineWave(void *userdata, Uint8 *stream, int len);
import "C"
import (
	"math"
	"reflect"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

var toneHz float64 = 440
var volume float64 = 1

const sampleHz = 48000

func dPhase() float64 {
	return 2 * math.Pi * toneHz / sampleHz
}

const (
	winTitle            = "Go-SDL2 Events"
	winWidth, winHeight = 800, 600
	pitchRange          = 800
)

var phase float64

//export SineWave
func SineWave(_ unsafe.Pointer, stream *C.Uint8, length C.int) {
	n := int(length)
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(stream)), Len: n, Cap: n}
	buf := *(*[]C.Uint8)(unsafe.Pointer(&hdr))

	for i := 0; i < n; i++ {
		phase += dPhase()
		sample := C.Uint8((math.Sin(phase) + 0.999999) * 128 * volume)
		buf[i] = sample
	}
}

func main() {
	err := sdl.Init(sdl.INIT_AUDIO)
	if err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	spec := &sdl.AudioSpec{
		Freq:     sampleHz,
		Format:   sdl.AUDIO_U8,
		Channels: 1,
		Samples:  2048,
		Callback: sdl.AudioCallback(C.SineWave),
	}
	err = sdl.OpenAudio(spec, nil)
	if err != nil {
		panic(err)
	}
	sdl.PauseAudio(true)

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				sdl.LockAudio()
				toneHz = (float64(t.X)/winWidth)*pitchRange + 100
				volume = (float64(t.Y) / winHeight)
				sdl.UnlockAudio()

			case *sdl.MouseButtonEvent:
				if t.Button == sdl.BUTTON_LEFT {
					sdl.PauseAudio(t.State == sdl.RELEASED)
//					if t.State == sdl.PRESSED {
//						sdl.PauseAudio(false)
//					} else {
//						sdl.PauseAudio(true)
//					}
				}
			}
		}
		sdl.Delay(100)
	}
}