#include "adventofcode/cpp/year2019/day25/day25.h"

#include <fstream>

#include "benchmark/benchmark.h"

static void Year2019Day25Part1(benchmark::State& state) {
  std::ifstream input("adventofcode/data/year2019/day25/actual.in");
  for (auto _ : state) {
    day25::part1(input);
    input.clear();
    input.seekg(0);
  }
}
BENCHMARK(Year2019Day25Part1);
