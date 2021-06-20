package adventofcode.java.year2016.day04;

import java.io.FileReader;
import java.io.IOException;

import org.junit.Test;
import org.junit.Assert;

public class Day04Test {
    @Test
    public void part1Example() throws IOException {
        try (var input = new FileReader("adventofcode/java/testdata/year2016/day04/example")) {
            var output = "1514";
            var result = Day04.part1(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }

    @Test
    public void part1Actual() throws IOException {
        try (var input = new FileReader("adventofcode/data/year2016/day04/actual.in")) {
            var output = "361724";
            var result = Day04.part1(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }

    @Test
    public void part2Actual() throws IOException {
        try (var input = new FileReader("adventofcode/data/year2016/day04/actual.in")) {
            var output = "482";
            var result = Day04.part2(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }
}
