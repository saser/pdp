#include "adventofcode/cpp/year2019/day02/day02.h"

#include <fstream>
#include <sstream>

#include "gtest/gtest.h"

#include "adventofcode/cpp/adventofcode.h"

TEST(Year2019Day02, Part1Actual) {
  std::ifstream input("adventofcode/data/year2019/day02/actual.in");
  std::string output = "4090689";
  adventofcode::answer_t a = day02::part1(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}

TEST(Year2019Day02, Part2Actual) {
  std::ifstream input("adventofcode/data/year2019/day02/actual.in");
  std::string output = "7733";
  adventofcode::answer_t a = day02::part2(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}
