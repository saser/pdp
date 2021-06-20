#include "adventofcode/cpp/year2019/day23/day23.h"

#include <fstream>

#include "benchmark/benchmark.h"

static void Year2019Day23Part1(benchmark::State& state) {
  std::ifstream input("adventofcode/data/year2019/day23/actual.in");
  for (auto _ : state) {
    day23::part1(input);
    input.clear();
    input.seekg(0);
  }
}
BENCHMARK(Year2019Day23Part1);

static void Year2019Day23Part2(benchmark::State& state) {
  std::ifstream input("adventofcode/data/year2019/day23/actual.in");
  for (auto _ : state) {
    day23::part2(input);
    input.clear();
    input.seekg(0);
  }
}
BENCHMARK(Year2019Day23Part2);
