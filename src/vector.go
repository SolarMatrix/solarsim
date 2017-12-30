package solarsim

type Vector struct {
	x, y float32
}

func (vector *Vector) add(other Vector) {
	vector.x += other.x
	vector.y += other.y
}

func (vector *Vector) sub(other Vector) {
	vector.x -= other.x
	vector.y -= other.y
}

func (vector *Vector) mult(value float32) {
	vector.x *= value
	vector.y *= value
}

func (vector *Vector) div(value float32) {
	vector.x /= value
	vector.y /= value
}