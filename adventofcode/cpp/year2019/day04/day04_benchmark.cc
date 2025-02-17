#include "adventofcode/cpp/year2019/day04/day04.h"

#include <fstream>

#include "benchmark/benchmark.h"

static void Year2019Day04Part1(benchmark::State& state) {
  std::ifstream input("adventofcode/data/year2019/day04/actual.in");
  for (auto _ : state) {
    day04::part1(input);
    input.clear();
    input.seekg(0);
  }
}
BENCHMARK(Year2019Day04Part1);

static void Year2019Day04Part2(benchmark::State& state) {
  std::ifstream input("adventofcode/data/year2019/day04/actual.in");
  for (auto _ : state) {
    day04::part2(input);
    input.clear();
    input.seekg(0);
  }
}
BENCHMARK(Year2019Day04Part2);
