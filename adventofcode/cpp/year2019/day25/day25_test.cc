#include "adventofcode/cpp/year2019/day25/day25.h"

#include <fstream>
#include <string>

#include "gtest/gtest.h"

#include "adventofcode/cpp/adventofcode.h"

TEST(Year2019Day25, Part1Actual) {
  std::ifstream input("adventofcode/data/year2019/day25/actual.in");
  std::string output = "2622472";
  adventofcode::answer_t a = day25::part1(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}
