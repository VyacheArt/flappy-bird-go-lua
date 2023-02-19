package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	lua "github.com/yuin/gopher-lua"
	"image"
	"image/color"
	luar "layeh.com/gopher-luar"
)

type LuaModule struct {
	state  *lua.LState
	draw   *lua.LFunction
	update *lua.LFunction
}

func loadLuaModule(path string) (*LuaModule, error) {
	L := lua.NewState()

	//exported types
	L.SetGlobal("ColorRGBA", luar.NewType(L, color.RGBA{}))

	//exported functions
	L.SetGlobal("GetWindowSize", luar.New(L, func() image.Point {
		width, height := ebiten.WindowSize()
		return image.Point{X: width, Y: height}
	}))
	L.SetGlobal("ColorWhite", luar.New(L, color.White))
	L.SetGlobal("BoundString", luar.New(L, text.BoundString))
	L.SetGlobal("DrawText", luar.New(L, text.Draw))
	L.SetGlobal("DefaultFont", luar.New(L, arcadeFont))

	if err := L.DoFile(path); err != nil {
		return nil, err
	}

	drawFn := L.GetGlobal("draw").(*lua.LFunction)
	updateFn := L.GetGlobal("update").(*lua.LFunction)

	return &LuaModule{
		state:  L,
		draw:   drawFn,
		update: updateFn,
	}, nil
}
