package adventofcode.java.year2016.day17;

import java.io.FileReader;
import java.io.IOException;
import java.io.StringReader;

import org.junit.Test;
import org.junit.Assert;

public class Day17Test {
    @Test
    public void part1Example1() {
        var input = new StringReader("ihgpwlah");
        var output = "DDRRRD";
        var result = Day17.part1(input);
        Assert.assertEquals("no error", "", result.error);
        Assert.assertEquals("correct output", output, result.answer);
    }

    @Test
    public void part1Example2() {
        var input = new StringReader("kglvqrro");
        var output = "DDUDRLRRUDRD";
        var result = Day17.part1(input);
        Assert.assertEquals("no error", "", result.error);
        Assert.assertEquals("correct output", output, result.answer);
    }

    @Test
    public void part1Example3() {
        var input = new StringReader("ulqzkmiv");
        var output = "DRURDRUDDLLDLUURRDULRLDUUDDDRR";
        var result = Day17.part1(input);
        Assert.assertEquals("no error", "", result.error);
        Assert.assertEquals("correct output", output, result.answer);
    }

    @Test
    public void part1Actual() throws IOException {
        try (var input = new FileReader("adventofcode/data/year2016/day17/actual.in")) {
            var output = "DDRLRRUDDR";
            var result = Day17.part1(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }

    @Test
    public void part2Example1() {
        var input = new StringReader("ihgpwlah");
        var output = "370";
        var result = Day17.part2(input);
        Assert.assertEquals("no error", "", result.error);
        Assert.assertEquals("correct output", output, result.answer);
    }

    @Test
    public void part2Example2() {
        var input = new StringReader("kglvqrro");
        var output = "492";
        var result = Day17.part2(input);
        Assert.assertEquals("no error", "", result.error);
        Assert.assertEquals("correct output", output, result.answer);
    }

    @Test
    public void part2Example3() {
        var input = new StringReader("ulqzkmiv");
        var output = "830";
        var result = Day17.part2(input);
        Assert.assertEquals("no error", "", result.error);
        Assert.assertEquals("correct output", output, result.answer);
    }

    @Test
    public void part2Actual() throws IOException {
        try (var input = new FileReader("adventofcode/data/year2016/day17/actual.in")) {
            var output = "556";
            var result = Day17.part2(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }
}
