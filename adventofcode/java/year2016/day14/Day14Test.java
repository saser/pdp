package adventofcode.java.year2016.day14;

import java.io.FileReader;
import java.io.IOException;
import java.io.StringReader;

import org.junit.Test;
import org.junit.Assert;

public class Day14Test {
    @Test
    public void part1Example() {
        var input = new StringReader("abc");
        var output = "22728";
        var result = Day14.part1(input);
        Assert.assertEquals("no error", "", result.error);
        Assert.assertEquals("correct output", output, result.answer);
    }

    @Test
    public void part1Actual() throws IOException {
        try (var input = new FileReader("adventofcode/data/year2016/day14/actual.in")) {
            var output = "16106";
            var result = Day14.part1(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }

     @Test
     public void part2Example() {
         var input = new StringReader("abc");
         var output = "22551";
         var result = Day14.part2(input);
         Assert.assertEquals("no error", "", result.error);
         Assert.assertEquals("correct output", output, result.answer);
     }

     @Test
     public void part2Actual() throws IOException {
         try (var input = new FileReader("adventofcode/data/year2016/day14/actual.in")) {
             var output = "22423";
             var result = Day14.part2(input);
             Assert.assertEquals("no error", "", result.error);
             Assert.assertEquals("correct output", output, result.answer);
         }
     }
}
