package adventofcode.java.year2016.day18;

import java.io.IOException;
import java.io.Reader;
import java.io.StringReader;
import java.nio.file.Files;
import java.nio.file.FileSystems;
import java.util.concurrent.TimeUnit;

import org.openjdk.jmh.annotations.Benchmark;
import org.openjdk.jmh.annotations.BenchmarkMode;
import org.openjdk.jmh.annotations.Fork;
import org.openjdk.jmh.annotations.Measurement;
import org.openjdk.jmh.annotations.Mode;
import org.openjdk.jmh.annotations.OutputTimeUnit;
import org.openjdk.jmh.annotations.Scope;
import org.openjdk.jmh.annotations.Setup;
import org.openjdk.jmh.annotations.State;
import org.openjdk.jmh.annotations.Warmup;

@BenchmarkMode(Mode.AverageTime)
@Fork(1)
@Measurement(iterations = 1, time = 1)
@OutputTimeUnit(TimeUnit.MILLISECONDS)
@State(Scope.Benchmark)
@Warmup(iterations = 5, time = 1)
public class Day18Benchmark {
    private Reader input;

    @Setup
    public void setup() throws IOException {
        var path = FileSystems.getDefault().getPath("adventofcode", "inputs", "2016", "18");
        var contents = Files.readString(path);
        this.input = new StringReader(contents);
    }
        
    @Benchmark
    public void part1() throws IOException {
        Day18.part1(this.input);
        this.input.reset();
    }
        
    @Benchmark
    public void part2() throws IOException {
        Day18.part2(this.input);
        this.input.reset();
    }
}
