/*
 * Copyright (c) 2022 XXIV
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
package main

/*
typedef struct {
  char* buffer;
  char* error;
} wallpaper;

typedef enum {
	Center,
	Crop,
	Fit,
	Span,
	Stretch,
	Tile
} wallpaper_mode;
*/
import "C"
import (
	"github.com/reujab/wallpaper"
)

//export wallpaper_get
func wallpaper_get() C.wallpaper {
  var self C.wallpaper
  background, err := wallpaper.Get()
  if err != nil {
    self.buffer = nil
    self.error = C.CString(err.Error())
    return self 
  }
    self.buffer = C.CString(background)
    self.error = nil
    return self 
}

//export wallpaper_set_from_file
func wallpaper_set_from_file(file *C.char) *C.char {
  err := wallpaper.SetFromFile(C.GoString(file))
  if err != nil {
    return C.CString(err.Error())
  }
  return nil
}

//export wallpaper_set_from_url
func wallpaper_set_from_url(url *C.char) *C.char {
  err := wallpaper.SetFromURL(C.GoString(url))
  if err != nil {
    return C.CString(err.Error())
  }
  return nil
}

//export wallpaper_set_mode
func wallpaper_set_mode(mode C.wallpaper_mode) *C.char {
  if mode == C.Center {
    err := wallpaper.SetMode(wallpaper.Center)
    if err != nil {
      return C.CString(err.Error())
    }
    return nil
  } else if mode == C.Crop {
    err := wallpaper.SetMode(wallpaper.Crop)
    if err != nil {
      return C.CString(err.Error())
    }
    return nil
  } else if mode == C.Fit {
    err := wallpaper.SetMode(wallpaper.Fit)
    if err != nil {
      return C.CString(err.Error())
    }
    return nil
  } else if mode == C.Span {
    err := wallpaper.SetMode(wallpaper.Span)
    if err != nil {
      return C.CString(err.Error())
    }
    return nil
  } else if mode == C.Stretch {
    err := wallpaper.SetMode(wallpaper.Stretch)
    if err != nil {
      return C.CString(err.Error())
    }
    return nil
  } else if mode == C.Tile {
    err := wallpaper.SetMode(wallpaper.Stretch)
    if err != nil {
      return C.CString(err.Error())
    }
    return nil
  } else {
    return C.CString("invalid wallpaper mode")
  }
}

func main() {}
