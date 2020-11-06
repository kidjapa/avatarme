package avatarme

import (
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "kidjapa-avatarme/user"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
)

// create the avatarme

var (
    backGroundColor color.Color = color.RGBA{R: 255, G: 255, B: 255, A: 255}
)

type Avatar struct {
    *user.User              // store the user input string and hash
    Ink         color.Color // store the ink avatar color
    Dir         string      // Store the directory to save image name
    Filename    string      // Store the image file name
    ImageWidth  int         // Store the image width
    ImageHeight int         // Store the image height
    Columns     uint        // Store the qty of block Columns in avatar
    Lines       uint        // Store the qty of block Lines in avatar
    Border      int         // Store the border for start draw rects
    Pixels      [][]string  // Store the hexadecimal value of each "pixel" in the image
}

func New(str string, dir string, fileName string, width uint, height uint, columns uint, lines uint, border uint) *Avatar {
    avatar := &Avatar{
        User:        user.New(str),
        Dir:         dir,
        Filename:    fileName,
        ImageWidth:  int(width),
        ImageHeight: int(height),
        Columns:     columns,
        Lines:       lines,
        Border:      int(border),
    }
    avatar._getAvatarColor()
    avatar._getAvatarPixels()
    return avatar
}

/**
Define avatar color using the last 6 digits of md5 hash created
*/
func (a *Avatar) _getAvatarColor() {
    source := a.Hash[26:] // last 6 characteres (Hexcolor)
    var pixels [3]uint8
    for i := 0; i < 3; i++ {
        pixel, _ := strconv.ParseInt(source[i*2:i*2+2], 16, 64)
        pixels[i] = uint8(pixel)
    }
    a.Ink = color.RGBA{R: pixels[0], G: pixels[1], B: pixels[2], A: 255}
}

/**
Get all avatar pixels
*/
func (a *Avatar) _getAvatarPixels() {
    for i := 0; i < int(a.Lines*a.Columns); i++ {
        rIdx := a._getRandomIndex()
        if rIdx >= 31 {
            rIdx = 30
        }
        a.Pixels = append(a.Pixels, []string{a.Hash[rIdx:(rIdx + 1)], a.Hash[(rIdx + 1):(rIdx + 2)]})
    }
}

/**
Get a random pixel index
*/
func (a *Avatar) _getRandomIndex() uint {
    rand.Seed(time.Now().UnixNano())
    min := 0
    max := len(a.Hash) - 1
    r := uint(0)
    for {
        r = uint(rand.Intn(max-min+1) + min)
        if r <= uint(len(a.Hash)-1) {
            break
        }
    }
    return r
}

func (a *Avatar) GenerateImage() {

    canvas := image.NewRGBA(image.Rect(0, 0, a.ImageWidth, a.ImageHeight))                          // create new "canvas"
    draw.Draw(canvas, canvas.Bounds(), &image.Uniform{C: backGroundColor}, image.Point{}, draw.Src) // fill canvas with background
    position := 0                                                                                   // start position at 0

    rectWidth := (a.ImageWidth - (a.Border * 2)) / int(a.Columns)
    rectHeight := (a.ImageHeight - (a.Border * 2)) / int(a.Lines)

    // Create a 4x4 blocks with defined width and height pixels each
    for yLine := 0; yLine < int(a.Lines); yLine++ {
        for xColumn := 0; xColumn < int(a.Columns); xColumn++ {

            hexValue := a.Pixels[position]
            value, _ := strconv.ParseInt(strings.Join(hexValue, ""), 16, 64)
            if value%2 == 0 {
                rect := image.Rect(
                    (xColumn*rectWidth)+a.Border,             // X0
                    (yLine*rectHeight)+(a.Border),            // Y0
                    (xColumn*rectWidth)+(a.Border+rectWidth), // X1
                    (yLine*rectHeight)+(a.Border+rectHeight)) // Y1
                draw.Draw(canvas, rect, &image.Uniform{C: a.Ink}, image.Point{}, draw.Src)
            }
            position++
        }
    }
    fileImage, _ := os.Create(a._getFileName())
    defer fileImage.Close()
    _ = png.Encode(fileImage, canvas)
}

func (a *Avatar) _getFileName() string {
    fileName := a.Filename + ".png"
    if a.Dir != "." {
        if a.Dir[:len(a.Dir)-1] != "/" {
            fileName = a.Dir + "/" + fileName
        } else {
            fileName = a.Dir + fileName
        }
    }
    return fileName
}
