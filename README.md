Reproduction of memory issues in [pants build](https://www.pantsbuild.org/docs/reference-golang)

This does seem to be related to included libraries. Builds with fewer requirements succeed and cache correctly, but adding more requirements quickly reaches the threshold of memory granted to pantsd.

# Clean reproduction steps

``` bash
pants check ::
# first run fills the immutable cache, second and onward more quickly demonstrates excess memory use
pants check ::
```

