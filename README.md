# brainfuck-runtime
Simple brainfuck runtime written in Go, I've programmed this .. for no particular reason, to be honest.

Following the [wikipedia specification](http://en.wikipedia.org/wiki/Brainfuck#Commands) of brainfuck.

## Example usage

This will print out "Hello, world." to the console. Blatantly obvious, right?
```
brain.Execute("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.")
```
