package adventofcode.java.year2016.day21;

import java.io.FileReader;
import java.io.IOException;
import java.io.StringReader;

import org.junit.Test;
import org.junit.Assert;

public class Day21Test {
    @Test
    public void part1Actual() throws IOException {
        try (var input = new FileReader("adventofcode/data/year2016/day21/actual.in")) {
            var output = "bfheacgd";
            var result = Day21.part1(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }

     @Test
     public void part2Actual() throws IOException {
         try (var input = new FileReader("adventofcode/data/year2016/day21/actual.in")) {
             var output = "gcehdbfa";
             var result = Day21.part2(input);
             Assert.assertEquals("no error", "", result.error);
             Assert.assertEquals("correct output", output, result.answer);
         }
     }
}
