package adventofcode.java.year2016.day15;

import java.io.FileReader;
import java.io.IOException;

import org.junit.Test;
import org.junit.Assert;

public class Day15Test {
    @Test
    public void part1Example() throws IOException {
        try (var input = new FileReader("adventofcode/java/testdata/year2016/day15/example")) {
            var output = "5";
            var result = Day15.part1(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }

    @Test
    public void part1Actual() throws IOException {
        try (var input = new FileReader("adventofcode/data/year2016/day15/actual.in")) {
            var output = "148737";
            var result = Day15.part1(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }

     @Test
     public void part2Actual() throws IOException {
         try (var input = new FileReader("adventofcode/data/year2016/day15/actual.in")) {
             var output = "2353212";
             var result = Day15.part2(input);
             Assert.assertEquals("no error", "", result.error);
             Assert.assertEquals("correct output", output, result.answer);
         }
     }
}
