#include "adventofcode/cpp/year2019/day22/day22.h"

#include <fstream>
#include <string>

#include "gtest/gtest.h"

#include "adventofcode/cpp/adventofcode.h"

TEST(Year2019Day22, Part1Actual) {
  std::ifstream input("adventofcode/data/year2019/day22/actual.in");
  std::string output = "4485";
  adventofcode::answer_t a = day22::part1(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}

TEST(Year2019Day22, Part2Actual) {
  std::ifstream input("adventofcode/data/year2019/day22/actual.in");
  std::string output = "91967327971097";
  adventofcode::answer_t a = day22::part2(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}
