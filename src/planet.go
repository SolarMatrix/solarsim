package solarsim

type State struct {
	position Vector
	velocity Vector
	acceleration Vector
	force Vector
	mass float32
}

type Planet struct {
	State
	name string
	radius float32
}

func (s *State) update(){
	s.force.div(s.mass)
	s.velocity.add(s.force)
	s.position.add(s.velocity)
}