package adventofcode.java.year2016.day05;

import java.io.FileReader;
import java.io.IOException;
import java.io.StringReader;

import org.junit.Test;
import org.junit.Assert;

public class Day05Test {
    @Test
    public void part1Example() {
        var input = new StringReader("abc");
        var output = "18f47a30";
        var result = Day05.part1(input);
        Assert.assertEquals("no error", "", result.error);
        Assert.assertEquals("correct output", output, result.answer);
    }

    @Test
    public void part1Actual() throws IOException {
        try (var input = new FileReader("adventofcode/data/year2016/day05/actual.in")) {
            var output = "c6697b55";
            var result = Day05.part1(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }

    @Test
    public void part2Example() {
        var input = new StringReader("abc");
        var output = "05ace8e3";
        var result = Day05.part2(input);
        Assert.assertEquals("no error", "", result.error);
        Assert.assertEquals("correct output", output, result.answer);
    }

    @Test
    public void part2Actual() throws IOException {
        try (var input = new FileReader("adventofcode/data/year2016/day05/actual.in")) {
            var output = "8c35d1ab";
            var result = Day05.part2(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }
}
