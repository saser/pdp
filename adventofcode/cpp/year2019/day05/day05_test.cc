#include "adventofcode/cpp/year2019/day05/day05.h"

#include <fstream>

#include "gtest/gtest.h"

TEST(Year2019Day05, Part1Actual) {
  std::ifstream input("adventofcode/data/year2019/day05/actual.in");
  std::string output = "12896948";
  adventofcode::answer_t a = day05::part1(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}

TEST(Year2019Day05, Part2Actual) {
  std::ifstream input("adventofcode/data/year2019/day05/actual.in");
  std::string output = "7704130";
  adventofcode::answer_t a = day05::part2(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}
