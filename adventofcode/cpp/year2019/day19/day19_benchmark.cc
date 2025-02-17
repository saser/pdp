#include "adventofcode/cpp/year2019/day19/day19.h"

#include <fstream>

#include "benchmark/benchmark.h"

static void Year2019Day19Part1(benchmark::State& state) {
  std::ifstream input("adventofcode/data/year2019/day19/actual.in");
  for (auto _ : state) {
    day19::part1(input);
    input.clear();
    input.seekg(0);
  }
}
BENCHMARK(Year2019Day19Part1);

static void Year2019Day19Part2(benchmark::State& state) {
  std::ifstream input("adventofcode/data/year2019/day19/actual.in");
  for (auto _ : state) {
    day19::part2(input);
    input.clear();
    input.seekg(0);
  }
}
BENCHMARK(Year2019Day19Part2);
