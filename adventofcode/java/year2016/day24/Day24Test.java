package adventofcode.java.year2016.day24;

import java.io.FileReader;
import java.io.IOException;

import org.junit.Test;
import org.junit.Assert;

public class Day24Test {
    @Test
    public void part1Example() throws IOException {
        try (var input = new FileReader("adventofcode/java/testdata/year2016/day24/example")) {
            var output = "14";
            var result = Day24.part1(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }

    @Test
    public void part1Actual() throws IOException {
        try (var input = new FileReader("adventofcode/data/year2016/day24/actual.in")) {
            var output = "470";
            var result = Day24.part1(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }

     @Test
     public void part2Actual() throws IOException {
         try (var input = new FileReader("adventofcode/data/year2016/day24/actual.in")) {
             var output = "720";
             var result = Day24.part2(input);
             Assert.assertEquals("no error", "", result.error);
             Assert.assertEquals("correct output", output, result.answer);
         }
     }
}
