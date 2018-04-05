//GoLang project
//For this project you need to take a picture and rotate it then split the picture into 4 pieces and rotate each of these. It should do this recursively down to single pixels. It should display the image at each iteration and save the file out to disk.


package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/draw"
	"os"
	"math"
	//"log"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
	"code.google.com/p/graphics-go/graphics"
)

func init(){
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func main(){
	gl.StartDriver(appMain)
}

func appMain(driver gxui.Driver){

	fmt.Println("This is my Program in Go")
	
	//import a picture

	width, height := 1000, 1000
	imgfile, err:= os.Open("doggy.jpg")
	fmt.Println("Here are the properties of your image file: ")
	if err != nil{
		fmt.Println("doggy.jpg file not found!")
		os.Exit(1)
	}

	defer imgfile.Close()

	img, err := jpeg.Decode(imgfile)
	if(err != nil) {panic (err)}
	editableImage := image.NewRGBA(img.Bounds())
	draw.Draw(editableImage, editableImage.Bounds(), img, image.Point{0,0}, draw.Src)

	//properties of file
	fmt.Println(img.At(10,10))
	bounds := img.Bounds()
	fmt.Println(bounds)
	canvas := image.NewAlpha(bounds)
	op := canvas.Opaque()
	fmt.Println(op)
	

	
	//display picture
	theme := dark.CreateTheme(driver)
	imgViewer := theme.CreateImage()
	window := theme.CreateWindow(width, height, "Image Viewer for Image Rotator")
	window.AddChild(imgViewer)
	
	//var texture[40]gxui.Texture;
	texture := driver.CreateTexture(editableImage, 1.0)
	imgViewer.SetTexture(texture)
	window.OnClose(driver.Terminate)

	graphics.Rotate(editableImage, img, &graphics.RotateOptions{math.Pi/2.0})
	draw.Draw(editableImage, editableImage.Bounds(), img, image.Point{0,0}, draw.Src)

	
	//done := false;
	
	//rotate func 45 degrees
	/*i := 0;	
	timer := time.AfterFunc(pause, func(){
		i++;
		if i >= 40{
			i = -5;
			done = true;
		}
		if (!done){
			m := image.NewRGBA(image.Rect(0,0, 220,220));
			draw.Draw(m, m.Bounds(), img, image.Point{0,0}, draw.Src)
			rotate(m, 0, 0, 220, 220, i);
			rotate(m, 0, 0, 220, 220, i-10);
			rotate(m, 0, 0, 220, 220, i-20);
			rotate(m, 0, 0, 220, 220, i-30);
			texture[i] = driver.CreateTexture(m, 1.0)
		}
	});*/
	
	/*driver.Call(func(){
		if (i >=0){
			//print
			img.SetTexture(texture[i])
		}
		timer.Reset(pause);
	});*/

	//split into 4 pieces

	
	//save file to disk
} 
