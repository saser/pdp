#include "adventofcode/cpp/year2019/day15/day15.h"

#include <fstream>
#include <string>

#include "gtest/gtest.h"

#include "adventofcode/cpp/adventofcode.h"

TEST(Year2019Day15, Part1Actual) {
  std::ifstream input("adventofcode/data/year2019/day15/actual.in");
  std::string output = "232";
  adventofcode::answer_t a = day15::part1(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}

TEST(Year2019Day15, Part2Actual) {
  std::ifstream input("adventofcode/data/year2019/day15/actual.in");
  std::string output = "320";
  adventofcode::answer_t a = day15::part2(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}
