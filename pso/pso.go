package pso

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type evaluationFunction func(position []float64) float64

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
func PSO(
	num_particles int,
	num_iterations int,
	upper_bounds []float64,
	lower_bounds []float64,
	momentum float64,
	learning_rate float64,
	personal_weight float64,
	global_weight float64,
	optimize_minimum bool) {

	rand.Seed(time.Now().UTC().UnixNano())
	var eval_fn evaluationFunction = beale
	global_best_position := make([]float64, len(lower_bounds))
	var global_best_value float64
	if optimize_minimum {
		global_best_value = math.Inf(1)
	} else {
		global_best_value = math.Inf(-1)
	}

	// Initialize particles
	particles := make([]Particle, num_particles)
	for i := range particles {
		particles[i] = MakeParticle(lower_bounds, upper_bounds, eval_fn)
		updateBest(&particles[i], &global_best_position, &global_best_value, eval_fn, optimize_minimum)
	}

	for iter := 0; iter < num_iterations; iter++ {
		fmt.Printf("Starting iteration %v \n", iter)
		for particle_idx := 0; particle_idx < len(particles); particle_idx++ {
			particle := &particles[particle_idx]
			fmt.Printf("Upating particle %v \n", particle_idx)
			particle.updateVelocity(momentum, personal_weight, global_weight, global_best_position)
			fmt.Printf("Eval before position change: %v\n", eval_fn(particle.position))
			particle.updatePosition(learning_rate)
			fmt.Printf("Eval after position change: %v\n", eval_fn(particle.position))

			fmt.Printf("Pre personal best change: %v\n", particle.personal_best_value)
			particle.updatePersonalBest(eval_fn, optimize_minimum)
			fmt.Printf("Post personal best change: %v\n", particle.personal_best_value)

			fmt.Println("Checking global best for update")
			updateBest(particle, &global_best_position, &global_best_value, eval_fn, optimize_minimum)
		}
	}
	fmt.Printf("Best value found: %v \n At position: %v  \n", global_best_value, global_best_position)
}

func randomFromBounds(lower_bound float64, upper_bound float64) float64 {
	spread := upper_bound - lower_bound
	untranslated := rand.Float64() * spread

	// "Translate" the range into the appropriate position
	return untranslated + lower_bound
}
