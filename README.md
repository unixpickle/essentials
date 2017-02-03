# essentials

There are some simple things that should be built-in to Go. As I find such things, I'll add them to this package.

Currently, the only thing in here is a "die" routine. Here's an example of the problem:

```go
if dataFile == "" {
	fmt.Fprintln(os.Stderr, "Missing -data flag. See -help for more info.")
	os.Exit(1)
}

log.Println("Loading encoder...")
var enc tweetenc.Encoder
if err := serializer.LoadAny(encFile, &enc.Block); err != nil {
	fmt.Fprintln(os.Stderr, "Load encoder:", err)
	os.Exit(1)
}

dataReader, err := os.Open(dataFile)
if err != nil {
	fmt.Fprintln(os.Stderr, "Open data:", err)
	os.Exit(1)
}
```

In three different places, I print to standard error and then exit with a status code. It would be so much less typing to do this:

```go
if dataFile == "" {
	essentials.Die("Missing -data flag. See -help for more info.")
}

log.Println("Loading encoder...")
var enc tweetenc.Encoder
if err := serializer.LoadAny(encFile, &enc.Block); err != nil {
	essentials.Die("Load encoder:", err)
}

dataReader, err := os.Open(dataFile)
if err != nil {
	essentials.Die("Open data:", err)
}
```