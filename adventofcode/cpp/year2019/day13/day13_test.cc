#include "adventofcode/cpp/year2019/day13/day13.h"

#include <fstream>
#include <string>

#include "gtest/gtest.h"

#include "adventofcode/cpp/adventofcode.h"

TEST(Year2019Day13, Part1Actual) {
  std::ifstream input("adventofcode/data/year2019/day13/actual.in");
  std::string output = "216";
  adventofcode::answer_t a = day13::part1(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}

TEST(Year2019Day13, Part2Actual) {
  std::ifstream input("adventofcode/data/year2019/day13/actual.in");
  std::string output = "10025";
  adventofcode::answer_t a = day13::part2(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}
