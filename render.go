package main

import (
	"errors"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"runtime"
	"sync"
)

func render(width, height, samples, superSamples int) error {
	if width <= 0 || height <= 0 || samples <= 0 || superSamples <= 0 {
		return errors.New("parameters must be greater than zero")
	}

	cameraPos := Vec{X: 50.0, Y: 52.0, Z: 220.0}
	cameraDir := Vec{X: 0.0, Y: -0.04, Z: -1.0}.Normalize()
	cameraUp := Vec{X: 0.0, Y: 1.0, Z: 0.0}

	screenHeight := 30.0
	screenWidth := screenHeight * float64(width) / float64(height)

	screenDist := 40.0

	screenX := cameraDir.Cross(cameraUp).Normalize().Mul(screenWidth)
	screenY := screenX.Cross(cameraDir).Normalize().Mul(screenHeight)
	screenCenter := cameraPos.AddVec(cameraDir.Mul(screenDist))

	invWidth := 1.0 / float64(width)
	invHeight := 1.0 / float64(height)
	invSamples := 1.0 / float64(samples*superSamples*superSamples)
	rate := 1.0 / float64(superSamples)

	img := make([]color.Color, width*height)

	log.Println("start")
	wg := sync.WaitGroup{}
	jobNum := runtime.NumCPU()
	finishedCounter := 0
	for job := 0; job < jobNum; job++ {
		wg.Add(1)
		go func(job int) {
			for y := job * (height / jobNum); y*jobNum < (job+1)*height; y++ {
				rnd := rand.New(rand.NewSource(int64(y)))
				for x := 0; x < width; x++ {
					accumulatedRadiance := Color{}
					for sy := 0; sy < superSamples; sy++ {
						for sx := 0; sx < superSamples; sx++ {
							for s := 0; s < samples; s++ {
								screenPosition := screenCenter.
									AddVec(screenX.Mul((rate*(float64(sx)+0.5)+float64(x))*invWidth - 0.5)).
									AddVec(screenY.Mul((rate*(float64(sy)+0.5)+float64(y))*invHeight - 0.5))
								dir := screenPosition.SubVec(cameraPos).Normalize()
								accumulatedRadiance = accumulatedRadiance.AddColor(radiance(Ray{cameraPos, dir}, rnd))
							}
						}
					}
					img[(height-y-1)*width+x] = accumulatedRadiance.Mul(invSamples)
				}

				finishedCounter++
				fmt.Printf("\r%d%%", 100*finishedCounter/height)
			}
			wg.Done()
		}(job)
	}
	wg.Wait()
	fmt.Println()
	log.Println("end")

	SavePNG("go-edupt.png", img, width, height)

	return nil
}
