#include "adventofcode/cpp/year2019/day22/day22.h"

#include <fstream>

#include "benchmark/benchmark.h"

static void Year2019Day22Part1(benchmark::State& state) {
  std::ifstream input("adventofcode/data/year2019/day22/actual.in");
  for (auto _ : state) {
    day22::part1(input);
    input.clear();
    input.seekg(0);
  }
}
BENCHMARK(Year2019Day22Part1);

static void Year2019Day22Part2(benchmark::State& state) {
  std::ifstream input("adventofcode/data/year2019/day22/actual.in");
  for (auto _ : state) {
    day22::part2(input);
    input.clear();
    input.seekg(0);
  }
}
BENCHMARK(Year2019Day22Part2);
