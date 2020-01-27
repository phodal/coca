package com.phodal.coca.analysis;

import com.phodal.coca.java.JavaParserBaseVisitor;
import com.phodal.coca.analysis.calls.JavaCallVisitor;
import com.phodal.coca.analysis.calls.JavaDaoParser;
import com.phodal.coca.analysis.calls.model.JMethodCall;
import com.phodal.coca.analysis.identifier.JavaFileParser;
import com.phodal.coca.analysis.utils.ProcessFiles;
import org.antlr.v4.runtime.tree.ParseTree;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;
import java.util.function.Consumer;

public class JavaCallApp {
    private JavaDaoParser daoParser;

    public JavaCallApp(JavaDaoParser daoParser) {
        this.daoParser = daoParser;
    }

    public void analysisDir(String dir, List<String> clzs) throws IOException, InterruptedException, ExecutionException {
        Path startingDir = Paths.get(dir);
        Consumer<Path> fileAnalysis = parse(clzs);
        int poolSize = 8;
        ExecutorService pool = Executors.newFixedThreadPool(poolSize);
        List<Future> futures = new ArrayList<>();
        ProcessFiles pf = new ProcessFiles(fileAnalysis, pool, futures);
        Files.walkFileTree(startingDir, pf);
        for (Future f : futures) {
            f.get();
        }
        pool.shutdown();
    }

    private Consumer<Path> parse(List<String> clzs) {
        return (Path path) -> {
            try {
                boolean projectJavaFile = !path.toString().endsWith("Tests.java") && !path.toString().endsWith("Test.java") && path.toString().endsWith(".java");
                if (projectJavaFile) {
                    parse(path, clzs);
                }
            } catch (IOException e) {
                e.printStackTrace();
            }
        };
    }

    public void parse(Path path, List<String> clzs) throws IOException {
        System.out.println("Start parse java call: " + path.getFileName());
        ParseTree tree = JavaFileParser.parse(path);
        List<JMethodCall> calls = new ArrayList<>();
        JavaParserBaseVisitor visitor = new JavaCallVisitor(calls, clzs, daoParser);
        visitor.visit(tree);
    }
}