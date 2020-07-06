#/bin/bash
#用于生成测试用jar
javac test/Test.java
jar cvfm server.jar manifest.txt test