# essentials

There are some simple things that should be built-in to Go. As I find such things, I'll add them to this package. See the [GoDoc](https://godoc.org/github.com/unixpickle/essentials) for more. I will try to document some of the package's functionality here in the README, but not everything.

# Slice deletion

Deleting elements from slices in Go is annoying and error prone. With `essentials`, the hard work is done for you:

```go
mySlice := []int{1, 2, 3, 4}
essentials.OrderedDelete(&mySlice, 2)
fmt.Println(mySlice) // [1 2 4]
```

If you don't care about the order of your slice, then `essentials.UnorderedDelete` is faster and will also do the job.

# The Die API

This API is useful for CLI apps where you want to exit with an error code in several places. Take this for example:

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
