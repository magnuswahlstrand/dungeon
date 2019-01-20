package game

import "github.com/peterhellberg/gfx"

var debug bool = false
var rubberband bool

var hook = gfx.Vec{50, 32}

var playerID = "abc321"
var hookID = "hook321"

var gravity = 0.25

const collisionScaling = 100

const jumpSpeed = 7.0

const accX = 0.2
