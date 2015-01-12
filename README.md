## Docker Who

The names of (Docker)[https://github.com/docker/docker]. 
When you create a Docker container, if you did not specify a name yourself,
Docker will automatically generate a name for you. This generator lives here:
https://github.com/docker/docker/blob/master/pkg/namesgenerator/names-generator.go

What many may not realize is the names are a tribute to renowned scientists from
around the world past and present.

This tool is just a way to further pay tribute and provide a way for accessing
the information about these scientists.

To use, take a name that was generated by Docker, e.g. `romantic_wozniak`, and
run it as an argument to this program.

```bash
docker-who romantic_wozniak
```

or also

```bash
docker-who wozniak
```


### Credits

Thanks to @thaJeztah for the idea here: docker/docker#9042
