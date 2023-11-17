# Master Class : Docker

```
docker images
```

> See available images.

```
docker images rm _images_name:_tags_name
```

> See remove images.

```
docker pull _images_name:_tags_name
```

> Download images with specific version.

```
docker build --tag _tag_name:_tag_version .
```

> Build images.

```
docker container ls --all
```

> See available container. If --all not included it will only show the list of running container.

```
docker container create --name _container_name  -p _outside_port:_internal_port -e _env_var_name=_env_var_value --env-file _file_name _images_name:_tags_name
```

> Create container. Random container name will be created if parameter name is not set.

```
docker container start _container_name
```

> Start container by name.

```
docker container stop _container_name
```

> Stop container by name.

```
docker container rm _container_name
```

> Delete container by name.

```
docker push _repository_registry_name/_images_name:_tags_name
```

> Upload to registry.

```
docker tag _images_name:_tags_name _repository_registry_name/_images_name:_tags_name
```

> Adjust tags name.

```
docker login
```

> Login to docker hub.

```
docker container inspect _container_name
```

> See container configuration.
