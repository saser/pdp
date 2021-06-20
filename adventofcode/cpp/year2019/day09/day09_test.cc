#include "adventofcode/cpp/year2019/day09/day09.h"

#include <fstream>
#include <sstream>
#include <string>

#include "gtest/gtest.h"

#include "adventofcode/cpp/adventofcode.h"

TEST(Year2019Day09, Part1Actual) {
  std::ifstream input("adventofcode/data/year2019/day09/actual.in");
  std::string output = "2453265701";
  adventofcode::answer_t a = day09::part1(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}

TEST(Year2019Day09, Part2Actual) {
  std::ifstream input("adventofcode/data/year2019/day09/actual.in");
  std::string output = "80805";
  adventofcode::answer_t a = day09::part2(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}
