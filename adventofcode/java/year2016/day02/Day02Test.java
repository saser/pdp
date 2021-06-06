package adventofcode.java.year2016.day02;

import java.io.FileReader;
import java.io.IOException;

import org.junit.Test;
import org.junit.Assert;

public class Day02Test {
    @Test
    public void part1Example() throws IOException {
        try (var input = new FileReader("adventofcode/java/testdata/year2016/day02/example")) {
            var output = "1985";
            var result = Day02.part1(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }

    @Test
    public void part1Actual() throws IOException {
        try (var input = new FileReader("adventofcode/inputs/2016/02")) {
            var output = "99332";
            var result = Day02.part1(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }

    @Test
    public void part2Example() throws IOException {
        try (var input = new FileReader("adventofcode/java/testdata/year2016/day02/example")) {
            var output = "5DB3";
            var result = Day02.part2(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }

    @Test
    public void part2Actual() throws IOException {
        try (var input = new FileReader("adventofcode/inputs/2016/02")) {
            var output = "DD483";
            var result = Day02.part2(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }
}
