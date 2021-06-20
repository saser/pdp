package adventofcode.java.year2016.day07;

import java.io.FileReader;
import java.io.IOException;

import org.junit.Test;
import org.junit.Assert;

public class Day07Test {
    @Test
    public void part1Example() throws IOException {
        try (var input = new FileReader("adventofcode/java/testdata/year2016/day07/p1example")) {
            var output = "2";
            var result = Day07.part1(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }

    @Test
    public void part1Actual() throws IOException {
        try (var input = new FileReader("adventofcode/data/year2016/day07/actual.in")) {
            var output = "105";
            var result = Day07.part1(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }

    @Test
    public void part2Example() throws IOException {
        try (var input = new FileReader("adventofcode/java/testdata/year2016/day07/p2example")) {
            var output = "3";
            var result = Day07.part2(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }

    @Test
    public void part2Actual() throws IOException {
        try (var input = new FileReader("adventofcode/data/year2016/day07/actual.in")) {
            var output = "258";
            var result = Day07.part2(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }
}
