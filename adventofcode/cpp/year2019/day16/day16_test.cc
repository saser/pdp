#include "adventofcode/cpp/year2019/day16/day16.h"

#include <fstream>
#include <sstream>
#include <string>

#include "gtest/gtest.h"

#include "adventofcode/cpp/adventofcode.h"

TEST(Year2019Day16, Part1Example1) {
  std::istringstream input("80871224585914546619083218645595");
  std::string output = "24176176";
  adventofcode::answer_t a = day16::part1(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
}

TEST(Year2019Day16, Part1Example2) {
  std::istringstream input("19617804207202209144916044189917");
  std::string output = "73745418";
  adventofcode::answer_t a = day16::part1(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
}

TEST(Year2019Day16, Part1Example3) {
  std::istringstream input("69317163492948606335995924319873");
  std::string output = "52432133";
  adventofcode::answer_t a = day16::part1(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
}

TEST(Year2019Day16, Part1Actual) {
  std::ifstream input("adventofcode/data/year2019/day16/actual.in");
  std::string output = "49254779";
  adventofcode::answer_t a = day16::part1(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}

TEST(Year2019Day16, Part2Example1) {
  std::istringstream input("03036732577212944063491565474664");
  std::string output = "84462026";
  adventofcode::answer_t a = day16::part2(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
}

TEST(Year2019Day16, Part2Example2) {
  std::istringstream input("02935109699940807407585447034323");
  std::string output = "78725270";
  adventofcode::answer_t a = day16::part2(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
}

TEST(Year2019Day16, Part2Example3) {
  std::istringstream input("03081770884921959731165446850517");
  std::string output = "53553731";
  adventofcode::answer_t a = day16::part2(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
}

TEST(Year2019Day16, Part2Actual) {
  std::ifstream input("adventofcode/data/year2019/day16/actual.in");
  std::string output = "55078585";
  adventofcode::answer_t a = day16::part2(input);
  EXPECT_EQ("", a.error);
  EXPECT_EQ(output, a.answer);
  input.close();
}
