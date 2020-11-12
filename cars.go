/*
MIT License
Copyright (c) 2020 Jean-Jacques François Reibel
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import "fmt"

type car struct{
        wheels int
        doors int
        cylinders int
}

func new(wheels int, doors int, cylinders int) car{
  myCar := car{wheels, doors, cylinders}
  return myCar
}

func addWheel(myCar *car, wheelsToAdd int) int{
        fmt.Printf("Adding wheels. \n")
        myCar.wheels += wheelsToAdd
        return myCar.wheels
}

func addDoor(myCar *car, doorsToAdd int) int{
        fmt.Printf("Adding doors. \n")
        myCar.doors += doorsToAdd
        return myCar.doors
}

func addCylinder(myCar *car, cylindersToAdd int) int{
        fmt.Printf("Adding cylinders. \n")
        myCar.cylinders += cylindersToAdd
        return myCar.cylinders
}

func deleteWheel(myCar *car, wheelsToRipOff int) int{
        fmt.Printf("Ripping off wheels. \n")
        myCar.wheels -= wheelsToRipOff
        return myCar.wheels
}

func deleteDoor(myCar *car, doorsToRipOff int) int{
        fmt.Printf("Ripping off doors. \n")
        myCar.doors -= doorsToRipOff
        return myCar.doors
}

func deleteCylinder(myCar *car, cylindersToRipOff int) int{
        fmt.Printf("Ripping off cylinders \n")
        myCar.cylinders -= cylindersToRipOff
        return myCar.cylinders
}

func main() {
        fmt.Printf("Creating car: \n")
        subaru := new(4, 4, 4)
        fmt.Printf("Wheels check: %d \n", subaru.wheels)
        fmt.Printf("Cylinders check: %d \n", subaru.cylinders)
        fmt.Printf("Doors check: %d \n", subaru.doors)
        fmt.Printf("\n")

        subaru.wheels = 5
        fmt.Printf("Added wheels directly to struct element: \n")
        fmt.Printf("Wheels check: %d \n", subaru.wheels)
        fmt.Printf("Cylinders check: %d \n", subaru.cylinders)
        fmt.Printf("Doors check: %d \n", subaru.doors)
        fmt.Printf("\n")

        deleteWheel(&subaru, 1)
        fmt.Printf("Remove wheels from struct element: \n")
        fmt.Printf("Wheels check: %d \n", subaru.wheels)
        fmt.Printf("Cylinders check: %d \n", subaru.cylinders)
        fmt.Printf("Doors check: %d \n", subaru.doors)
} 
