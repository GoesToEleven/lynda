package main

// typedef unsigned char Uint8;
// void SineWave(void *userdata, Uint8 *stream, int len);
import "C"
import (
	"math"
	"reflect"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
	"time"
)

// IMPORTANT:
// run this at terminal to import exif package:
// go get github.com/veandco/go-sdl2/sdl


// tonehz, aka, pitch, oscillations per second
// more oscillations / second = higher pitch/tone
// less oscillations / second = lower pitch/tone
var toneHz float64 = 440
var volume float64 = 1

// number of entries in array that makes a second
const sampleHz = 48000

// d stands for delta which is change
func dPhase() float64 {
	// toneHz / sampleHz
	// 440 / 48000 is oscillations per second / entries per second is oscillations / entry
	// 2 * pi is one oscillation
	// this is in radians which is a unitless type (not in programming sense, but math sense)
	// there's no concept attached to the number like with time
	// not like time which is seconds, or minutes
	// it's just a plain number, doesn't
	return 2 * math.Pi * toneHz / sampleHz
}

const (
	winTitle            = "Go-SDL2 Events"
	winWidth, winHeight = 800, 600
	pitchRange          = 800
)

var phase float64

// export SineWave
// this all communicates with C code
// the parameters in the signature are all C types
// the first argument is user data - but we're not using it
// "stream" argument is a pointer to a single uint8 which is the first item in the array
// it is pointing towards the very first empty spot in the array
// the first memory address location
// and then the THIRD ARGUMENT is the length
// the amount of memory we can fill out
// SDL will pass in these arguments
// so SDL will determine how much memory we have to fill out
func SineWave(_ unsafe.Pointer, stream *C.Uint8, length C.int) {
	// convert to Go type int
	n := int(length)
	// create a header to a Go slice: pointer to underlying array, len, cap
	// stream is the pointer to the beginning of the array
	// the rest of the pointer to the underlying array is a bunch of conversion
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(stream)), Len: n, Cap: n}
	// casting / conversion to a pointer to a slice of uint8
	// dereference that with the outer * which gives us the value
	// so type buf is just a slice of uint8
	buf := *(*[]C.Uint8)(unsafe.Pointer(&hdr))

	for i := 0; i < n; i++ {
		// we have an infinite sine wave
		// and we're moving along the waves
		// every 2pi the wave repeats
		// if we graph the sine wave, phase is the x
		// and to get the Y, we call the sine on it, as in next line below
		phase += dPhase()
		// we're not changing how much the sin wave is condensed / expanded
		// we're changing how fast / slow we move through the sine waves
		// to get the Y, we call sin on phase
		// the Y will be between -1 and 1
		// this is the volume
		// we add .9999999, eg, 1 to it b/c we have an uint
		// and therefore our volume will be between 0 and 2
		// multiply by 128 gets us between 0 and 256
		// b/c 256 values (0-255) is the highest we can get in a uint8
		// we then multiply by volume
		// volume is a number between 0 and 1
		sample := C.Uint8((math.Sin(phase) + 0.999999) * 128 * volume)
		buf[i] = sample
	}
}

func main() {
	// init scans the system
	// figures out what hardware is available
	// init takes in an optional argument
	// we're also initializing the audio system
	// if we had put in 0 only the basics would be initialized in SDL (window, input, etc)
	// often the video is initialized
	// also an everything flag
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
		Freq: sampleHz,
		// size of data unsigned 8 bit integers
		Format: sdl.AUDIO_U8,
		// mono channel
		Channels: 1,
		// buffer, entries in the future it can store before it calls for more data
		// must be a power of 2 sdl defines
		// 2048 / 48000 = 0.0426
		// only storing 4/100th second
		// so when we change tone or volume, at worst, the buffer is already full, so it will take 4/100th of a second until the new data that has flowed in plays
		// that's why we want "Samples" to be relatively small
		// if this was an MP3 you'd want several seconds of data
		Samples:  2048,
		Callback: sdl.AudioCallback(C.SineWave),
	}

	// second argument would be a pointer to an empty sdl.AudioSpec
	// sdl.OpenAudio would then fill that out with what is actually found on the audio system
	// if for instance, the audio system being played on didn't support frequency or samples
	// sdl would find the fit that does work on that system and then fill out that empty sdl.AudioSpec and use that
	err = sdl.OpenAudio(spec, nil)
	if err != nil {
		panic(err)
	}
	// audio is paused until mouse button is held
	sdl.PauseAudio(true)

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			// the top left red x has been pushed so quit
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				// sdl runs audio in separate thread
				// race condition possibility: callback is in middle of being called
				// and you're trying to change callback
				// because audio is running on one thread
				// and where the callback would/is being run is on another thread
				sdl.LockAudio()
				// change tone / volume based on position of mouse
				toneHz = (float64(t.X)/winWidth)*pitchRange + 100
				volume = (float64(t.Y) / winHeight)
				sdl.UnlockAudio()

			case *sdl.MouseButtonEvent:
				if t.Button == sdl.BUTTON_LEFT {
					sdl.PauseAudio(t.State == sdl.RELEASED)
					// instead of above line, could do this - it was like this before:
					//					if t.State == sdl.PRESSED {
					//						sdl.PauseAudio(false)
					//					} else {
					//						sdl.PauseAudio(true)
					//					}
				}
			}
		}
		// sdl.Delay(100)
		// use Go for sleep
		// because in SDL the delay pauses the entire thread
		// Go only pauses the goroutine, not the thread
		time.Sleep(100 * time.Millisecond)
	}
}
