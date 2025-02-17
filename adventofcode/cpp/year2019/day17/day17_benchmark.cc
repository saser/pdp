#include "adventofcode/cpp/year2019/day17/day17.h"

#include <fstream>

#include "benchmark/benchmark.h"

static void Year2019Day17Part1(benchmark::State& state) {
  std::ifstream input("adventofcode/data/year2019/day17/actual.in");
  for (auto _ : state) {
    day17::part1(input);
    input.clear();
    input.seekg(0);
  }
}
BENCHMARK(Year2019Day17Part1);

static void Year2019Day17Part2(benchmark::State& state) {
  std::ifstream input("adventofcode/data/year2019/day17/actual.in");
  for (auto _ : state) {
    day17::part2(input);
    input.clear();
    input.seekg(0);
  }
}
BENCHMARK(Year2019Day17Part2);
