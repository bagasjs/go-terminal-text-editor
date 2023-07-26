package core

type Vec2 struct {
    X int32
    Y int32
}

func Vec2Zeros() Vec2 {
    return Vec2{ X: 0, Y: 0 }
}
