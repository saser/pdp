#include "adventofcode/cpp/year2019/day09/day09.h"

#include <fstream>

#include "benchmark/benchmark.h"

static void Year2019Day09Part1(benchmark::State& state) {
  std::ifstream input("adventofcode/data/year2019/day09/actual.in");
  for (auto _ : state) {
    day09::part1(input);
    input.clear();
    input.seekg(0);
  }
}
BENCHMARK(Year2019Day09Part1);

static void Year2019Day09Part2(benchmark::State& state) {
  std::ifstream input("adventofcode/data/year2019/day09/actual.in");
  for (auto _ : state) {
    day09::part2(input);
    input.clear();
    input.seekg(0);
  }
}
BENCHMARK(Year2019Day09Part2);
