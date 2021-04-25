package pso

import (
	"fmt"
	"math"
	"math/rand"
)

type Particle struct {
	position            []float64
	velocity            []float64
	personal_best_pos   []float64
	personal_best_value float64
}

func MakeParticle(lower_bounds []float64, upper_bounds []float64, eval_fn evaluationFunction) Particle {
	particle_position := make([]float64, len(lower_bounds))
	particle_velocity := make([]float64, len(lower_bounds))
	for i := range lower_bounds {
		particle_position[i] = randomFromBounds(lower_bounds[i], upper_bounds[i])
		particle_velocity[i] = randomFromBounds(
			-math.Abs(upper_bounds[i]-lower_bounds[i]),
			math.Abs(upper_bounds[i]-lower_bounds[i]),
		)
	}
	return Particle{
		position:            particle_position,
		velocity:            particle_velocity,
		personal_best_pos:   particle_position,
		personal_best_value: eval_fn(particle_position),
	}
}

func (p *Particle) updatePosition(learning_rate float64) {
	fmt.Printf("Before position change: %v\n", p.position)
	for i := range p.position {
		p.position[i] += learning_rate * p.velocity[i]
	}
	fmt.Printf("After position change: %v\n", p.position)
}

func (p *Particle) updatePersonalBest(eval_fn evaluationFunction, optimize_min bool) {
	fmt.Println("Checking personal best for update")
	updateBest(p, &p.personal_best_pos, &p.personal_best_value, eval_fn, optimize_min)
}

func (p *Particle) updateVelocity(
	momentum float64,
	personal_weight float64,
	global_weight float64,
	global_best_pos []float64) {
	for dim, velocity := range p.velocity {
		random_weight_personal := rand.Float64()
		random_weight_global := rand.Float64()
		p.velocity[dim] = velocity*momentum +
			random_weight_personal*(p.personal_best_pos[dim]-p.position[dim]) +
			random_weight_global*(global_best_pos[dim]-p.position[dim])
	}
}

func updateBest(
	particle *Particle,
	best_pos *[]float64,
	best_value *float64,
	eval_fn evaluationFunction,
	optimize_minimum bool) {

	particle_value := eval_fn(particle.position)
	if optimize_minimum {
		if particle_value < *best_value {
			fmt.Printf("Best value updated from %v to %v \n", *best_value, particle_value)
			*best_value = particle_value
			*best_pos = particle.position
		}
	} else {
		if particle_value > *best_value {
			*best_value = particle_value
			*best_pos = particle.position
		}
	}
}
