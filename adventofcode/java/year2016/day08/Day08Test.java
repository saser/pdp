package adventofcode.java.year2016.day08;

import java.io.BufferedInputStream;
import java.io.FileInputStream;
import java.io.FileReader;
import java.io.IOException;

import org.junit.Test;
import org.junit.Assert;

public class Day08Test {
    @Test
    public void part1Actual() throws IOException {
        try (var input = new FileReader("adventofcode/inputs/2016/08")) {
            var output = "116";
            var result = Day08.part1(input);
            Assert.assertEquals("no error", "", result.error);
            Assert.assertEquals("correct output", output, result.answer);
        }
    }

     @Test
     public void part2Actual() throws IOException {
         try (var input = new FileReader("adventofcode/inputs/2016/08");
              var outputStream = new FileInputStream("adventofcode/java/testdata/year2016/day08/output")) {
             var outputBytes = new BufferedInputStream(outputStream).readAllBytes();
             var output = new String(outputBytes);
             var result = Day08.part2(input);
             Assert.assertEquals("no error", "", result.error);
             Assert.assertEquals("correct output", output, result.answer);
         }
     }
}
