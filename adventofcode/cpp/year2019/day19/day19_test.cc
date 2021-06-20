#include "adventofcode/cpp/year2019/day19/day19.h"

#include <fstream>
#include <sstream>
#include <string>

#include "gtest/gtest.h"

#include "adventofcode/cpp/adventofcode.h"

TEST(Year2019Day19, Part1Actual) {
  std::ifstream input("adventofcode/data/year2019/day19/actual.in");
  std::string output = "197";
  adventofcode::answer_t a = day19::part1(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}

TEST(Year2019Day19, Part2Actual) {
  std::ifstream input("adventofcode/data/year2019/day19/actual.in");
  std::string output = "9181022";
  adventofcode::answer_t a = day19::part2(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}
