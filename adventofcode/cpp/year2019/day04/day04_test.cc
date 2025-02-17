#include "adventofcode/cpp/year2019/day04/day04.h"

#include <fstream>

#include "gtest/gtest.h"

#include "adventofcode/cpp/adventofcode.h"

TEST(Year2019Day04, Part1Actual) {
  std::ifstream input("adventofcode/data/year2019/day04/actual.in");
  std::string output = "1246";
  adventofcode::answer_t a = day04::part1(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}

TEST(Year2019Day04, Part2Actual) {
  std::ifstream input("adventofcode/data/year2019/day04/actual.in");
  std::string output = "814";
  adventofcode::answer_t a = day04::part2(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}
