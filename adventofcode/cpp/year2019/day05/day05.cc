#include "adventofcode/cpp/year2019/day05/day05.h"

#include <istream>
#include <string>

#include "absl/strings/str_format.h"

#include "adventofcode/cpp/adventofcode.h"
#include "adventofcode/cpp/year2019/intcode/intcode.h"

adventofcode::answer_t solve(std::istream& is, int part);

namespace day05 {
  adventofcode::answer_t part1(std::istream& is) {
    return solve(is, 1);
  }
  adventofcode::answer_t part2(std::istream& is) {
    return solve(is, 2);
  }
}

adventofcode::answer_t solve(std::istream& is, int part) {
  std::string line;
  std::getline(is, line);
  intcode::memory program = intcode::parse(line);
  int input = part == 1 ? 1 : 5;
  auto result = intcode::run(program, {input});
  auto output = result.second;
  for (size_t i = 0; i < output.size() - 1; i++) {
    if (output[i] != 0) {
      return adventofcode::err(absl::StrFormat("failure in test %d: output = %d", i, output[i]));
    }
  }
  int diagnostic_code = output[output.size() - 1];
  return adventofcode::ok(std::to_string(diagnostic_code));
}
