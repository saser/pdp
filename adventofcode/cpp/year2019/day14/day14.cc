#include "adventofcode/cpp/year2019/day14/day14.h"

#include <istream>
#include <regex>
#include <string>
#include <unordered_map>
#include <utility>
#include <vector>

#include "adventofcode/cpp/adventofcode.h"

struct reagent {
  unsigned long amount;
  std::string chemical;
};

using productions_t = std::unordered_map<std::string, std::pair<unsigned long, std::vector<reagent>>>;
using produce_t = std::unordered_map<std::string, unsigned long>;

adventofcode::answer_t solve(std::istream& is, int part);
productions_t parse(std::istream& is);
std::pair<produce_t, produce_t> produce(const reagent& r, const productions_t& productions);
unsigned long ore_for_fuel(const unsigned long& amount, const productions_t& productions);

namespace day14 {
  adventofcode::answer_t part1(std::istream& is) {
    return solve(is, 1);
  }

  adventofcode::answer_t part2(std::istream& is) {
    return solve(is, 2);
  }
}

adventofcode::answer_t solve(std::istream& is, int part) {
  auto productions = parse(is);
  unsigned long ore = ore_for_fuel(1, productions);
  if (part == 1) {
    return adventofcode::ok(std::to_string(ore));
  }
  // This is a bit of cleverness. It is based on that if `n` ore can produce `m`
  // fuel, then `target` ore makes _at least_ `m * target / n` fuel. This makes
  // `m * target / n` a good guess for `fuel` for the next iteration.
  // We use `fuel + 1` to not "overshoot" the `target`. Due to how we update
  // `fuel`, we know that the ore for `fuel` will never overshoot the target.
  // We convert everything to doubles before performing the division. This helps
  // with problems that arise due to integer division. We then floor everything
  // by converting it back to a long again.
  //
  // This method was not found by me. I got stuck and looked for hints on the
  // Advent of Code subreddit, and found this solution by user
  // /u/hotzenplotz6. The comment is available here:
  // https://www.reddit.com/r/adventofcode/comments/eafj32/2019_day_14_solutions/faqkkwv/
  auto fuel = 1ul;
  auto target = 1'000'000'000'000ul;
  while ((ore = ore_for_fuel(fuel + 1, productions)) < target) {
    fuel = (unsigned long) ((double) (fuel + 1) * (double) target / (double) ore);
  }
  return adventofcode::ok(std::to_string(fuel));
}

productions_t parse(std::istream& is) {
  std::string line;
  const std::regex re(R"((\d+) (\w+))");
  productions_t productions;
  productions["ORE"] = {1, {}};
  while (std::getline(is, line)) {
    std::vector<reagent> requirements;
    auto matches_it = std::sregex_iterator(line.begin(), line.end(), re);
    auto matches_end = std::sregex_iterator();
    while (matches_it != matches_end) {
      auto match = *matches_it;
      auto amount = std::stoul(match[1]);
      auto chemical = match[2];
      requirements.push_back(reagent {amount, chemical});
      matches_it++;
    }
    auto result = requirements.back();
    requirements.pop_back();
    productions[result.chemical] = {result.amount, requirements};
  }
  return productions;
}

void produce_aux(const reagent& r, const productions_t& productions, produce_t& produced_chemicals, produce_t& available_chemicals) {
  auto to_produce = r.amount;
  auto& available = available_chemicals[r.chemical];
  if (available >= to_produce) {
    return;
  }
  to_produce -= available;
  auto [result_amount, requirements] = productions.at(r.chemical);
  auto times = to_produce / result_amount;
  if (to_produce % result_amount != 0) {
    times++;
  }
  for (auto requirement : requirements) {
    requirement.amount *= times;
    produce_aux(requirement, productions, produced_chemicals, available_chemicals);
    available_chemicals[requirement.chemical] -= requirement.amount;
  }
  auto total = result_amount * times;
  produced_chemicals[r.chemical] += total;
  available += total;
}

std::pair<produce_t, produce_t> produce(const reagent& r, const productions_t& productions) {
  produce_t produced_chemicals;
  produce_t available_chemicals;
  produce_aux(r, productions, produced_chemicals, available_chemicals);
  return {produced_chemicals, available_chemicals};
}

unsigned long ore_for_fuel(const unsigned long& amount, const productions_t& productions) {
  auto [produced, _] = produce(reagent {amount, "FUEL"}, productions);
  return produced.at("ORE");
}
