package pso

import (
	"fmt"
	"math/rand"
)

type Particle struct {
	position     []float64
	velocity     []float64
	personalBest []float64
}

func (p *Particle) updatePosition(learning_rate float64) {
	for i := range p.position {
		p.position[i] += learning_rate * p.velocity[i]
	}
}

func (p *Particle) updateVelocity(
	momentum float64,
	personal_weight float64,
	global_weight float64,
	global_best []float64) {
	for dim, velocity := range p.velocity {
		random_weight_personal := rand.Float64()
		random_weight_global := rand.Float64()
		velocity = velocity*momentum +
			random_weight_personal*(p.personalBest[dim]-p.position[dim]) +
			random_weight_global*(global_best[dim]-p.position[dim])
	}
}

/*
Initialize:
For each particle:
	initialize particle's position with uniformly distributed random vector
	initialize particle's best known position to initial position (personal-best) pb <- x
	if f(pb) better than f(global-best), set global-best to pb
	Initialize velocity with some bounds on the search space
Loop (until termination criteria met):
	For each particle:
		For each dimension, d:
			Pick random weights: rand_personal, rand_global ~ U(0, 1)
			Update particle's velocity at dimension:
				velocity[d] <- momentum * velocity[d] +
					personal_weight * rand_personal * (pb[d] - x[d]) +
					global_weight * rand_global * (global_best[d] - x[d])
		Update particle's position: x <- x + learning_rate * velocity
		If f(x) better than f(pb) then pb = x
		If f(pb) better than f(global-best) then global_best = x
*/
func pso() {
	fmt.Println("PSO function")
}
