package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"text/template"
	"os"
	"log"
	"fmt"
)

type Netprobe struct {
	Managedentity string
	Probename string
}

func netprobe(c echo.Context) error{
	name := c.Param("netprobe")
	ProbeFile := fmt.Sprintf("assets/netprobe/%s.xml", name)
	netprobe := Netprobe{name,name}
	tmpl := template.New("probe template")
	tmpl, err := tmpl.ParseFiles("assets/netprobe/probe.xml.tmpl")
	if err != nil {
		log.Fatal("Parse: ", err)
		return err
	}

	f,err := os.Create(ProbeFile)
	if err != nil {
		log.Fatal("Create file: ", err)
		return err
	}

	err = tmpl.ExecuteTemplate(f,"body",netprobe)
	if err != nil {
		log.Fatal("Execute: ", err)
		return err
	}

	f.Close()

	return c.File(ProbeFile)
}

func test(c echo.Context) error{
	return c.File("assets/test/probe.xml")
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/activeconsole", "assets/activeconsole")
	e.GET("/netprobe/:netprobe", netprobe)
	e.GET("/test/*", test)
	e.Logger.Fatal(e.Start(":8000"))
}
