#include "adventofcode/cpp/year2019/day23/day23.h"

#include <fstream>
#include <string>

#include "gtest/gtest.h"

#include "adventofcode/cpp/adventofcode.h"

TEST(Year2019Day23, Part1Actual) {
  std::ifstream input("adventofcode/data/year2019/day23/actual.in");
  std::string output = "18966";
  adventofcode::answer_t a = day23::part1(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}

TEST(Year2019Day23, Part2Actual) {
  std::ifstream input("adventofcode/data/year2019/day23/actual.in");
  std::string output = "14370";
  adventofcode::answer_t a = day23::part2(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}
