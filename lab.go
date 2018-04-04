//GoLang project
//For this project you need to take a picture and rotate it then split the picture into 4 pieces and rotate each of these. It should do this recursively down to single pixels. It should display the image at each iteration and save the file out to disk.


package main

import (
	"fmt"
	"image"
	_"image/jpeg"
	"os"
	"log"
	"golang.org/exp/shiny/driver"
	"golang.org/exp/shiny/screen"
	"golang.org/mobile/lifecycle"
	"golang.org/exp/shiny/widget"
)

func init(){
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func main(){

	fmt.Println("This is my Program in Go")
	
	//import a picture
	imgfile, err:= os.Open("./goLab/doggy.jpg")
	fmt.PrintLn("Here are the properties of your image file: ")
	if err != nil{
		fmt.Println("doggy.jpg file not found!")
		os.Exit(1)
	}

	defer imgfile.Close()

	img, _, err := image.Decode(imgfile)
	fmt.Println(img.At(10,10))
	bounds := img.Bounds()
	fmt.Println(bounds)
	canvas := image.NewAlpha(bounds)
	op := canvas.Opaque()
	fmt.Println(op)
	
	//rotate picture 45 degrees
	
	//split into 4 pieces
	
	//display picture
	log.SetFlags(0)
        driver.Main(func(s screen.Screen){
                w := widget.NewSheet(widget.NewImage(img, img.Bounds()))
		if err := widget.RunWindows(s, w, &widget.RunWindowOptions{
			NewWindowOptions: screen.NewWindowOptions{
				Title: "Rotate and chop Doggy Image",
			}
		}); err != nil{
			log.Fatal(err)
		}
	})

	//save file to disk
} 
