#include "adventofcode/cpp/year2019/day21/day21.h"

#include <iostream>
#include <string>

#include "adventofcode/cpp/adventofcode.h"
#include "adventofcode/cpp/year2019/intcode/intcode.h"

adventofcode::answer_t solve(std::istream& is, int part);

namespace day21 {
  adventofcode::answer_t part1(std::istream& is) {
    return solve(is, 1);
  }

  adventofcode::answer_t part2(std::istream& is) {
    return solve(is, 2);
  }
}

adventofcode::answer_t solve(std::istream& is, int part) {
  std::string input;
  std::getline(is, input);
  intcode::memory program = intcode::parse(input);
  intcode::execution e {program};
  std::string springscript;
  if (part == 1) {
    springscript =
      "NOT A J\n"
      "NOT C T\n"
      "AND D T\n"
      "OR T J\n"
      "WALK\n";
  } else {
    springscript =
      "NOT A J\n"
      "NOT C T\n"
      "AND D T\n"
      "AND H T\n"
      "OR T J\n"
      "NOT B T\n"
      "AND D T\n"
      "AND H T\n"
      "OR T J\n"
      "RUN\n";
  }
  for (auto c : springscript) {
    e.write(c);
  }
  e.run();
  auto output = e.read_all();
  if (auto damage = output.back(); damage > 127) {
      return adventofcode::ok(std::to_string(damage));
  }
  for (auto c : output) {
    std::cout << (char) c;
  }
  std::cout << std::endl;
  return adventofcode::err("error");
}
