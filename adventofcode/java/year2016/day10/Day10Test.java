package adventofcode.java.year2016.day10;

import java.io.FileReader;
import java.io.IOException;

import org.junit.Test;
import org.junit.Assert;

public class Day10Test {
    @Test
    public void part1Actual() throws IOException {
        try (var input = new FileReader("adventofcode/data/year2016/day10/actual.in")) {
            var output = "118";
            var result = Day10.part1(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }

     @Test
     public void part2Actual() throws IOException {
         try (var input = new FileReader("adventofcode/data/year2016/day10/actual.in")) {
             var output = "143153";
             var result = Day10.part2(input);
             Assert.assertEquals("no error", "", result.error);
             Assert.assertEquals("correct output", output, result.answer);
         }
     }
}
