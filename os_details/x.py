import timeit
import math
import random


def f(x, y):
    return math.sin(x) + math.cos(y)


def simulated_annealing(max_iterations, start_temperature, cooling_rate):
    x = random.uniform(-math.pi, math.pi)
    y = random.uniform(-math.pi, math.pi)
    best_x, best_y = x, y
    best_f = f(x, y)
    temperature = start_temperature
    for i in range(max_iterations):
        dx = random.uniform(-1, 1)
        dy = random.uniform(-1, 1)
        new_x, new_y = x+dx, y+dy
        new_f = f(new_x, new_y)
        delta_f = new_f - best_f
        if delta_f > 0 or random.uniform(0, 1) < math.exp(delta_f/temperature):
            best_x, best_y = new_x, new_y
            best_f = new_f
        temperature *= cooling_rate
    return best_f, best_x, best_y


def hill_climbing(max_iterations, step_size):
    x = random.uniform(-math.pi, math.pi)
    y = random.uniform(-math.pi, math.pi)
    best_x, best_y = x, y
    best_f = f(x, y)
    for i in range(max_iterations):
        dx = random.uniform(-step_size, step_size)
        dy = random.uniform(-step_size, step_size)
        if f(x+dx, y+dy) > best_f:
            best_f = f(x+dx, y+dy)
            best_x, best_y = x+dx, y+dy
        x, y = best_x, best_y
    return best_f, best_x, best_y



def genetic_algorithm(population_size, num_generations):
    def crossover(p1, p2):
        c1 = (p1[0], p2[1])
        c2 = (p2[0], p1[1])
        return c1, c2

    def mutate(p, mutation_rate):
        x, y = p
        if random.uniform(0, 1) < mutation_rate:
            x += random.uniform(-1, 1)
        if random.uniform(0, 1) < mutation_rate:
            y += random.uniform(-1, 1)
        return x, y

    population = [(random.uniform(-math.pi, math.pi),
                   random.uniform(-math.pi, math.pi)) for i in range(population_size)]
    for generation in range(num_generations):
        population = sorted(
            population, key=lambda p: f(p[0], p[1]), reverse=True)
        elite = population[:int(population_size/4)]
        offspring = []
        for i in range(population_size - len(elite)):
            p1, p2 = random.choices(population[:int(population_size/2)], k=2)
            c1, c2 = crossover(p1, p2)
            c1 = mutate(c1, 0.1)
            c2 = mutate(c2, 0.1)
            offspring.append(c1)
            offspring.append(c2)
        population = elite + offspring
    best_x, best_y = population[0]
    best_f = f(best_x, best_y)
    return best_f, best_x, best_y


num_runs = 10
max_iterations = 10000
step_size = 0.1
start_temperature = 10
cooling_rate = 0.999
population_size = 100
num_generations = 100

hill_climbing_times = []
hill_climbing_results = []
for i in range(num_runs):
    start_time = timeit.default_timer()
    result = hill_climbing(max_iterations, step_size)
    elapsed_time = timeit.default_timer() - start_time
    hill_climbing_times.append(elapsed_time)
    hill_climbing_results.append(result)

simulated_annealing_times = []
simulated_annealing_results = []
for i in range(num_runs):
    start_time = timeit.default_timer()
    result = simulated_annealing(
        max_iterations, start_temperature, cooling_rate)
    elapsed_time = timeit.default_timer() - start_time
    simulated_annealing_times.append(elapsed_time)
    simulated_annealing_results.append(result)

genetic_algorithm_times = []
genetic_algorithm_results = []
for i in range(num_runs):
    start_time = timeit.default_timer()
    result = genetic_algorithm(population_size, num_generations)
    elapsed_time = timeit.default_timer() - start_time
    genetic_algorithm_times.append(elapsed_time)
    genetic_algorithm_results.append(result)


print('Hill Climbing Average time: {:.5f} seconds'.format(sum(hill_climbing_times)/num_runs))
print('Hill Climbing Average optimum value of x: {:.5f}'.format(sum(r[1] for r in hill_climbing_results)/num_runs))
print('Hill Climbing Average optimum value of y: {:.5f}'.format(sum(r[2] for r in hill_climbing_results)/num_runs))
print('Hill Climbing Average optimum result: {:.5f}'.format(sum(r[0] for r in hill_climbing_results)/num_runs))
print('Simulated Annealing Average time: {:.5f} seconds'.format(sum(simulated_annealing_times)/num_runs))
print('Simulated Annealing Average optimum value of x: {:.5f}'.format(sum(r[1] for r in simulated_annealing_results)/num_runs))
print('Simulated Annealing Average optimum value of y: {:.5f}'.format(sum(r[2] for r in simulated_annealing_results)/num_runs))
print('Simulated Annealing Average optimum result: {:.5f}'.format(sum(r[0] for r in simulated_annealing_results)/num_runs))
print('Genetic Algorithm Average time: {:.5f} seconds'.format(sum(genetic_algorithm_times)/num_runs))
print('Genetic Algorithm Average optimum value of x: {:.5f}'.format(sum(r[1] for r in genetic_algorithm_results)/num_runs))
print('Genetic Algorithm Average optimum value of y: {:.5f}'.format(sum(r[2] for r in genetic_algorithm_results)/num_runs))
print('Genetic Algorithm Average optimum result: {:.5f}'.format(sum(r[0] for r in genetic_algorithm_results)/num_runs))
